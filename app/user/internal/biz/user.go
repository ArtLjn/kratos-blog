package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/api/user"
	"kratos-blog/app/user/internal/conf"
	cf "kratos-blog/pkg/conf"
	"kratos-blog/pkg/util"
	"kratos-blog/pkg/vo"
	"strings"
	"time"
)

type UserRepo interface {
	GetCurrentConfig() (cf.Config, error)
	AddUser(ctx context.Context, request *user.CreateUserRequest) (string, error)
	Login(ctx context.Context, request *user.LoginRequest) (string, string, error)
	SendEmail(body, to, sub string) bool
	CacheCode(email string, code int) bool
	UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest) (string, error)
	SetBlack(ctx context.Context, request *user.SetBlackRequest) (string, error)
	GetUserMsg(request *user.GetUserRequest) []string
	AdminLogin(ctx context.Context, request *user.AdminLoginRequest) *user.AdminLoginReply
	VerifyToken(ctx context.Context, request *user.VerifyTokenRequest) *user.VerifyTokenReply
	LogOut(ctx context.Context, request *user.LogOutRequest) *user.LogOutReply
	QueryAllUser(ctx context.Context, request *user.QueryAllUserRequest) *user.QueryAllUserResponse
	SetAdmin(ctx context.Context, request *user.SetAdminRequest) *user.SetAdminReply
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
	conf *conf.Bootstrap
}

var emailList = make(map[string]int64)

func NewUserUseCase(repo UserRepo, conf *conf.Bootstrap, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, conf: conf, log: log.NewHelper(logger)}
}

func (u *UserUseCase) Login(ctx context.Context, request *user.LoginRequest) *user.LoginReply {
	res, token, err := u.repo.Login(ctx, request)
	if err != nil {
		u.log.Log(log.LevelError, err)
		return &user.LoginReply{Common: &user.CommonReply{Code: 400, Result: res}}
	}
	return &user.LoginReply{Common: &user.CommonReply{Code: 200, Result: res}, Data: []string{token}}
}

func (u *UserUseCase) Register(ctx context.Context, request *user.CreateUserRequest) *user.CreateUserReply {
	res, err := u.repo.AddUser(ctx, request)
	if err != nil {
		u.log.Log(log.LevelError, err)
		return &user.CreateUserReply{Common: &user.CommonReply{Code: 400, Result: res}}
	}
	return &user.CreateUserReply{Common: &user.CommonReply{Code: 200, Result: res}}
}

func (u *UserUseCase) SendEmail(ctx context.Context, request *user.SendEmailRequest) *user.SendEmailReply {
	emailPrefix := strings.Split(request.Email, "@")[1]
	if emailPrefix != "qq.com" {
		return &user.SendEmailReply{Common: &user.CommonReply{Code: 300, Result: vo.SUPPORT_QQ_EMAIL}}
	}
	currentTime := time.Now().Unix()
	if _, ok := emailList[request.Email]; ok {
		coolDown := emailList[request.Email]
		if currentTime < coolDown {
			return &user.SendEmailReply{Common: &user.CommonReply{Code: 301, Result: vo.EMAIL_COOLDOWN}}
		} else {
			delete(emailList, request.Email)
		}
	}
	code := util.RandomInt()
	body := fmt.Sprintf("您好，您的验证码为\n%d\n请妥善保管，本次验证码5分钟内有效", code)
	u.log.Log(log.LevelInfo, code)
	res := u.repo.SendEmail(body, request.Email, "注册邮件")
	if res {
		u.repo.CacheCode(request.Email, code)
		emailList[request.Email] = currentTime + 60
		return &user.SendEmailReply{Common: &user.CommonReply{Code: 200, Result: vo.EMAIL_SUCCESS}}
	}
	return &user.SendEmailReply{Common: &user.CommonReply{Code: 400, Result: vo.EMAIL_FAIL}}
}

func (u *UserUseCase) UpdateUserPassword(ctx context.Context, request *user.UpdatePasswordRequest) *user.UpdatePasswordReply {
	res, err := u.repo.UpdatePassword(ctx, request)
	if err != nil {
		return &user.UpdatePasswordReply{Common: &user.CommonReply{Code: 400, Result: res}}
	}
	return &user.UpdatePasswordReply{Common: &user.CommonReply{Code: 200, Result: res}}
}

func (u *UserUseCase) SetUserBlack(ctx context.Context, request *user.SetBlackRequest) *user.SetBlackReply {
	res, err := u.repo.SetBlack(ctx, request)
	if err != nil {
		return &user.SetBlackReply{Common: &user.CommonReply{Code: 400, Result: res}}
	}
	return &user.SetBlackReply{Common: &user.CommonReply{Code: 200, Result: res}}
}

func (u *UserUseCase) GetUserMsg(ctx context.Context, request *user.GetUserRequest) *user.GetUserReply {
	data := u.repo.GetUserMsg(request)
	if len(data) == 0 {
		return &user.GetUserReply{Common: &user.CommonReply{Code: 300, Result: vo.QUERY_EMPTY}}
	}
	return &user.GetUserReply{
		Common: &user.CommonReply{Code: 200, Result: vo.QUERY_SUCCESS},
		Data:   data,
	}
}

func (u *UserUseCase) AdminLoginRsg(ctx context.Context, request *user.AdminLoginRequest) *user.AdminLoginReply {
	return u.repo.AdminLogin(ctx, request)
}

func (u *UserUseCase) VerifyToken(ctx context.Context, request *user.VerifyTokenRequest) *user.VerifyTokenReply {
	return u.repo.VerifyToken(ctx, request)
}

func (u *UserUseCase) LogOut(ctx context.Context, request *user.LogOutRequest) *user.LogOutReply {
	return u.repo.LogOut(ctx, request)
}

func (u *UserUseCase) SendEmailCommon(ctx context.Context, request *user.SendEmailCommonRequest) *user.SendEmailCommonReply {
	if u.repo.SendEmail(request.GetBody(), request.GetEmail(), request.GetSubject()) {
		return &user.SendEmailCommonReply{Common: &user.CommonReply{Code: 200, Result: vo.EMAIL_SUCCESS}}
	}
	return &user.SendEmailCommonReply{Common: &user.CommonReply{Code: 400, Result: vo.EMAIL_FAIL}}
}

func (u *UserUseCase) QueryAllUser(ctx context.Context, request *user.QueryAllUserRequest) *user.QueryAllUserResponse {
	return u.repo.QueryAllUser(ctx, request)
}

func (u *UserUseCase) SetAdmin(ctx context.Context, request *user.SetAdminRequest) *user.SetAdminReply {
	return u.repo.SetAdmin(ctx, request)
}
