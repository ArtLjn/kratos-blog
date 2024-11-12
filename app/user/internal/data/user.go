package data

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"
	"kratos-blog/api/user"
	"kratos-blog/app/user/internal/biz"
	"kratos-blog/pkg/auth"
	"kratos-blog/pkg/conf"
	"kratos-blog/pkg/model"
	"kratos-blog/pkg/server"
	"kratos-blog/pkg/vo"
	"strconv"
	"sync"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
	mu   sync.Mutex
	model.PublicFunc
	key auth.TokenManager
}

type User struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
	Black    bool   `json:"black"`
	Role     string `json:"role"`
}

func (u User) TableName() string {
	return "user"
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(logger)
	return &userRepo{
		data: data,
		log:  l,
		key:  auth.NewToken(data.rdb, l),
	}
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func (u *userRepo) GetCurrentConfig() (conf.Config, error) {
	var result conf.Config
	err := u.data.collect.FindOne(context.TODO(),
		bson.D{{"open", true}}).Decode(&result)
	if err != nil {
		log.Error(err)
		return conf.Config{}, err
	}
	return result, nil
}

// AddUser :dev create user func
func (u *userRepo) AddUser(ctx context.Context, request *user.CreateUserRequest) (string, error) {
	//Check the registration conditions
	if err := u.validateUser(request); err != nil {
		u.log.Log(log.LevelError, err)
		return fmt.Sprintf("%s", err), err
	}
	userFactory := u.createUserFromRequest(request)
	factory := userFactory()
	//Create a user record
	if err := u.data.db.Create(&factory).Error; err != nil {
		return vo.INSERT_ERROR, errors.New(vo.INSERT_ERROR)
	} else {
		//Delete the mailbox cache
		u.data.rdb.Del(context.Background(), request.Email)
	}
	return vo.REGISTER_SUCCESS, nil
}

// Login :dev Common user login func
func (u *userRepo) Login(ctx context.Context, request *user.LoginRequest) (string, string, error) {
	if u.data.pf.HasExist(User{}, map[string]interface{}{
		"name":     request.Name,
		"password": MD5(request.Pass),
	}) {
		//generate token
		token := u.key.SaveToken(request.Name)
		if len(token) == 0 {
			return vo.GENERATE_TOKEN_FAIL, "", fmt.Errorf(vo.GENERATE_TOKEN_FAIL)
		}
		return vo.LOGIN_SUCCESS, token, nil
	}
	return vo.LOGIN_FAIL, "", errors.New(vo.LOGIN_FAIL)
}

// SendEmail :dev send email func
func (u *userRepo) SendEmail(body, to, sub string) bool {
	c, err := u.GetCurrentConfig()
	if err != nil {
		u.data.log.Log(log.LevelError, err)
		return false
	}
	m := gomail.NewMessage()
	m.SetHeader("From", c.QQEmail.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", sub)
	m.SetBody("text/plain", body)
	d := gomail.NewDialer(c.QQEmail.Host, c.QQEmail.Port, c.QQEmail.Username, c.QQEmail.Password)
	if err = d.DialAndSend(m); err != nil {
		u.data.log.Log(log.LevelError, err)
		return false
	}
	return true
}

func (u *userRepo) CacheCode(email string, code int) bool {
	_, err := u.data.rdb.Set(context.Background(), email, code, 5*time.Minute).Result()
	return err == nil
}

// UpdatePassword :dev update user password func
func (u *userRepo) UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest) (string, error) {
	u.mu.Lock()
	defer u.data.rdb.Del(context.Background(), request.Email)
	defer u.mu.Unlock()
	// check the updateValidate conditions
	err := u.validateUpdatePass(request)
	if err != nil {
		return fmt.Sprintf("%s", err), err
	}
	return vo.UPDATE_SUCCESS, nil
}

// validateUser :dev check the registration
func (u *userRepo) validateUser(request *user.CreateUserRequest) error {
	res := func(cond, value string) bool {
		if err := u.data.db.Where(map[string]interface{}{cond: value}).First(&User{}).Error; err != nil {
			return false
		}
		return true
	}
	exs := res("name", request.Name)
	ems := res("email", request.Email)
	c, err := u.GetCurrentConfig()
	if err != nil {
		u.log.Log(log.LevelError, err)
		return nil
	}
	cacheCode, _ := u.data.rdb.Get(context.Background(), request.Email).Result()
	if request.Code != cacheCode {
		return errors.New(vo.KEY_ERROR)
	} else if exs || request.Name == c.Admin.Username {
		return errors.New(vo.NAME_EXIST)
	} else if ems {
		return errors.New(vo.EMAIL_EXIST)
	}
	return nil
}

// createUserFromRequest :dev create a user record
func (u *userRepo) createUserFromRequest(request *user.CreateUserRequest) func() User {
	return func() User {
		var us User
		body, err := json.Marshal(&request)
		if err != nil {
			panic(err)
		}
		if e := json.Unmarshal(body, &us); e != nil {
			panic(e)
		}
		us.UUID = uuid.New().String()[:8]
		us.Black = false
		us.Password = MD5(us.Password)
		us.Role = server.User
		return us
	}
}

// update validate validateUpdatePass
func (u *userRepo) validateUpdatePass(request *user.UpdatePasswordRequest) error {
	var uS User
	data, _ := u.data.pf.QueryFunc(User{}, map[string]interface{}{"email": request.Email}, false)
	err := u.data.pf.ParseJSONToStruct(data, &uS)
	if err != nil {
		return err
	}
	if uS.Email != request.Email {
		return errors.New(vo.EMAIL_NOT_MATCH)
	} else if cacheCode, _ := u.data.rdb.Get(context.Background(), request.Email).Result(); cacheCode != request.Code {
		return errors.New(vo.KEY_ERROR)
	} else if err := u.data.pf.UpdateFunc(User{}, map[string]interface{}{
		"email": request.Email,
	}, map[string]interface{}{"password": MD5(request.Password)}, false); err != nil {
		return errors.New(vo.UPDATE_FAIL)
	}
	return nil
}

// SetBlack :dev Set a blacklist of users who violate the rules
func (u *userRepo) SetBlack(ctx context.Context, request *user.SetBlackRequest) (string, error) {
	if err := u.data.pf.UpdateFunc(User{}, map[string]interface{}{"name": request.Name},
		map[string]interface{}{"black": request.GetBlack()}, false); err != nil {
		return vo.UPDATE_FAIL, errors.New(vo.UPDATE_FAIL)
	}
	if request.GetBlack() {
		var uS User
		data, _ := u.data.pf.QueryFunc(User{}, map[string]interface{}{"name": request.Name}, false)
		if err := u.data.pf.ParseJSONToStruct(data, &uS); err != nil {
			u.log.Log(log.LevelError, err)
			return vo.UPDATE_FAIL, errors.New(vo.UPDATE_FAIL)
		}
		// Create a coroutine to send a message
		go func() {
			u.mu.Lock()
			c, err := u.GetCurrentConfig()
			if err != nil {
				u.log.Log(log.LevelError, err)
				return
			}
			blackBody := fmt.Sprintf("亲爱的用户,您好！\n系统检测到您的账号\n %s \n有异常行为,你的账号被暂停使用!"+
				",如有疑问请联系管理员 %s ,谢谢您的配合!", uS.Email, c.QQEmail.Username)
			body := fmt.Sprintf("系统检测到该用户有异常操作行为\n %s \n %s \n 现已被封禁如有异常请您进行处理。", uS.Name, uS.Email)
			u.SendEmail(body, c.QQEmail.Username, "通知邮件")
			u.SendEmail(blackBody, uS.Email, "违规邮件")
			u.mu.Unlock()
		}()
	}
	return vo.UPDATE_SUCCESS, nil
}

func (u *userRepo) GetUserMsg(request *user.GetUserRequest) []string {
	u.log.Log(log.LevelDebug, request.Name)
	var data User
	var list []string
	d, err := u.data.pf.QueryFunc(User{}, map[string]interface{}{"name": request.Name}, false)
	if err != nil {
		u.log.Log(log.LevelError, err)
	}
	if err = u.data.pf.ParseJSONToStruct(d, &data); err != nil {
		return nil
	}
	list = append(list, data.Name, data.Email, strconv.Itoa(data.ID), data.UUID, data.Role, strconv.FormatBool(data.Black))
	return list
}

// AdminLogin :dev Super admin users are unique
func (u *userRepo) AdminLogin(ctx context.Context, request *user.AdminLoginRequest) *user.AdminLoginReply {
	status := func(comm *user.CommonReply, list []string) *user.AdminLoginReply {
		return &user.AdminLoginReply{
			Common: comm,
			Data:   list,
		}
	}
	c, err := u.GetCurrentConfig()
	if err != nil {
		u.log.Log(log.LevelError, err)
		return nil
	}
	if request.Name == c.Admin.Username &&
		request.Password == c.Admin.Password {
		// generate token
		token := fmt.Sprintf("admin-token:%s", uuid.New().String())
		u.data.rdb.Set(context.Background(), server.AdminToken, token, 7*time.Hour*24)
		return status(&user.CommonReply{Code: 200, Result: vo.LOGIN_SUCCESS}, []string{token})
	}
	return status(&user.CommonReply{Code: 400, Result: vo.LOGIN_FAIL}, nil)
}

func (u *userRepo) VerifyToken(ctx context.Context, request *user.VerifyTokenRequest) *user.VerifyTokenReply {
	err := u.key.VerifyToken(request.GetToken())
	if err != nil {
		return &user.VerifyTokenReply{
			Common: &user.CommonReply{
				Code:   vo.PERMISSION_REQUEST,
				Result: err.Error(),
			},
		}
	}
	return &user.VerifyTokenReply{
		Common: &user.CommonReply{
			Code:   vo.SUCCESS_REQUEST,
			Result: vo.VERIFY_THORUGH,
		},
	}
}

func (u *userRepo) LogOut(ctx context.Context, request *user.LogOutRequest) *user.LogOutReply {
	u.key.LogOutToken(request.GetName())
	return &user.LogOutReply{
		Common: &user.CommonReply{
			Code:   vo.SUCCESS_REQUEST,
			Result: vo.OPERATE_SUCCESS,
		},
	}
}

func (u *userRepo) QueryAllUser(ctx context.Context, request *user.QueryAllUserRequest) *user.QueryAllUserResponse {
	var data []*user.QueryUser
	if err := u.data.db.Model(&User{}).Find(&data).Error; err != nil {
		log.Log(log.LevelError, err)
	}
	return &user.QueryAllUserResponse{
		Common: &user.CommonReply{
			Code:   vo.SUCCESS_REQUEST,
			Result: vo.QUERY_SUCCESS,
		},
		Data: data,
	}
}

func (u *userRepo) SetAdmin(ctx context.Context, request *user.SetAdminRequest) *user.SetAdminReply {
	cond := u.data.db.Model(&User{}).Where("name = ?", request.GetName())
	var role string
	if request.GetSet() {
		role = server.VIP
	} else {
		role = server.User
	}
	if err := cond.Update("role", role).Error; err != nil {
		return &user.SetAdminReply{
			Common: &user.CommonReply{
				Code:   vo.SUCCESS_REQUEST,
				Result: vo.UPDATE_FAIL,
			},
		}
	}
	return &user.SetAdminReply{
		Common: &user.CommonReply{
			Code:   vo.SUCCESS_REQUEST,
			Result: vo.UPDATE_SUCCESS,
		},
	}
}
