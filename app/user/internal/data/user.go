package data

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
	"kratos-blog/api/v1/user"
	"kratos-blog/app/user/internal/biz"
	"kratos-blog/pkg/util"
	"kratos-blog/pkg/vo"
	"strconv"
	"sync"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
	mu   sync.Mutex
}

// Role
var (
	AdminRole  = "admin"
	CommonRole = "user"
)

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
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// AddUser :dev create user func
func (u *userRepo) AddUser(ctx context.Context, request *user.CreateUserRequest) (string, error) {
	//Enable mutex
	u.mu.Lock()
	defer u.mu.Unlock()
	//Check the registration conditions
	if err := u.validateUser(request); err != nil {
		u.log.Log(log.LevelError, err)
		return fmt.Sprintf("%s", err), err
	}
	userFactory := u.createUserFromRequest(request)
	user := userFactory()
	//Create a user record
	if err := u.data.db.Create(&user).Error; err != nil {
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
		token := fmt.Sprintf("user-token:%s", util.GenerateToken())
		// set up cache
		u.data.rdb.Set(context.Background(), token, request.Name, 7*24*time.Hour)
		return vo.LOGIN_SUCCESS, token, nil
	}
	return vo.LOGIN_FAIL, "", errors.New(vo.LOGIN_FAIL)
}

// SendEmail :dev send email func
func (u *userRepo) SendEmail(body, to, sub string) bool {
	conf := u.data.c.Mail
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", sub)
	m.SetBody("text/plain", body)
	d := gomail.NewDialer(conf.Host, int(conf.Port), conf.Username, conf.Password)
	if err := d.DialAndSend(m); err != nil {
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
	cacheCode, _ := u.data.rdb.Get(context.Background(), request.Email).Result()
	if request.Code != cacheCode {
		return errors.New(vo.KEY_ERROR)
	} else if exs || request.Name == u.data.c.Admin.Username {
		return errors.New(vo.NAME_EXIST)
	} else if ems {
		return errors.New(vo.EMAIL_EXIST)
	}
	return nil
}

// createUserFromRequest :dev create a user record
func (u *userRepo) createUserFromRequest(request *user.CreateUserRequest) func() User {
	return func() User {
		var user User
		body, err := json.Marshal(&request)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &user); err != nil {
			panic(err)
		}
		user.UUID = uuid.New().String()
		user.Black = false
		user.Password = MD5(user.Password)
		user.Role = CommonRole
		return user
	}
}

// update validate validateUpdatePass
func (u *userRepo) validateUpdatePass(request *user.UpdatePasswordRequest) error {
	var user User
	u.QueryMessage(map[string]interface{}{"email": user.Email}, &user)
	if user.Email != request.Email {
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
	if err := u.data.pf.UpdateFunc(User{}, map[string]interface{}{"name": request.Name}, map[string]interface{}{"black": true}, false); err != nil {
		return vo.UPDATE_FAIL, errors.New(vo.UPDATE_FAIL)
	}

	var user User
	u.QueryMessage(map[string]interface{}{"name": request.Name}, &user)
	// Create a coroutine to send a message
	go func() {
		u.mu.Lock()
		conf := u.data.c.Mail
		blackBody := fmt.Sprintf("亲爱的用户,您好！\n系统检测到您的账号\n %s \n有异常行为,你的账号被暂停使用!"+
			",如有疑问请联系管理员 %s ,谢谢您的配合!", user.Email, conf.Username)
		body := fmt.Sprintf("系统检测到该用户有异常操作行为\n %s \n %s \n 现已被封禁如有异常请您进行处理。", user.Name, user.Email)
		u.SendEmail(body, conf.Username, "通知邮件")
		u.SendEmail(blackBody, user.Email, "违规邮件")
		u.mu.Unlock()
	}()
	return vo.UPDATE_SUCCESS, nil
}

// QueryMessage :dev query user information
func (u *userRepo) QueryMessage(cond map[string]interface{}, data *User) {
	res, err := u.data.pf.QueryFunc(User{}, cond, false)
	var user User
	err = u.data.pf.ParseJSONToStruct(res, &user)
	if err != nil {
		u.log.Log(log.LevelError, err)
	}
	*data = user
}

func (u *userRepo) GetUserMsg(request *user.GetUserRequest) []string {
	var data User
	var list []string
	if request.Name == u.data.c.Admin.Username {
		f := make([]string, 3)
		f = append(f, u.data.c.Admin.Username, AdminRole)
		list = append([]string{}, f...)
		return list
	}
	u.QueryMessage(map[string]interface{}{"name": request.Name}, &data)
	list = append(list, data.Name, data.Email, strconv.Itoa(data.ID), data.UUID, data.Role, strconv.FormatBool(data.Black))
	return list
}

// AdminLogin :dev Super admin users are unique
func (u *userRepo) AdminLogin(request *user.AdminLoginRequest) *user.AdminLoginReply {
	status := func(comm *user.CommonReply, list []string) *user.AdminLoginReply {
		return &user.AdminLoginReply{
			Common: comm,
			Data:   list,
		}
	}
	if request.Name == u.data.c.Admin.Username &&
		request.Password == u.data.c.Admin.Password {
		// generate token
		token := fmt.Sprintf("admin-token:%s", util.GenerateToken())
		// set up cache
		u.data.rdb.Set(context.Background(), token, request.Name, 7*24*time.Hour)
		return status(&user.CommonReply{Code: 200, Result: vo.LOGIN_SUCCESS}, []string{token})
	}
	return status(&user.CommonReply{Code: 400, Result: vo.LOGIN_FAIL}, nil)
}
