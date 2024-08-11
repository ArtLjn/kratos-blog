package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "kratos-blog/api/v1/user"
	"kratos-blog/app/gateway/internal/data"
	"kratos-blog/pkg/server"
	"kratos-blog/pkg/vo"
)

type UserService struct {
	pb.UnimplementedUserServer
	UC   pb.UserClient
	log  *log.Helper
	q    *data.Mq
	role *data.Role
}

func NewUserService(logger log.Logger, uc pb.UserClient, q *data.Mq, role *data.Role) *UserService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &UserService{UC: uc, log: l, q: q, role: role}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return s.UC.CreateUser(ctx, req)
}
func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	ok, token := s.verifyToken(&ctx)
	if ok && len(token) != 0 {
		return &pb.LoginReply{Common: &pb.CommonReply{Code: vo.SUCCESS_REQUEST, Result: vo.LOGIN_SUCCESS}, Data: []string{token}}, nil
	} else if !ok {
		return &pb.LoginReply{Common: &pb.CommonReply{Code: vo.PERMISSION_REQUEST, Result: token}}, nil
	}
	return s.UC.LoginUser(ctx, req)
}
func (s *UserService) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	err := s.q.PushMail(data.MailS{
		Email: req.Email,
	})
	if err != nil {
		return &pb.SendEmailReply{Common: &pb.CommonReply{Code: vo.BAD_REQUEST, Result: vo.EMAIL_COOLDOWN}}, nil
	}
	return &pb.SendEmailReply{Common: &pb.CommonReply{Code: vo.SUCCESS_REQUEST, Result: vo.EMAIL_SUCCESS}}, nil
}
func (s *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return s.UC.UpdatePassword(ctx, req)
}
func (s *UserService) SetBlack(ctx context.Context, req *pb.SetBlackRequest) (*pb.SetBlackReply, error) {
	return s.UC.SetBlack(ctx, req)
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return s.UC.GetUser(ctx, req)
}
func (s *UserService) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginReply, error) {
	res, ok := http.RequestFromServerContext(ctx)
	if !ok {
		log.Errorf("ctx false")
		return &pb.AdminLoginReply{Common: &pb.CommonReply{Code: vo.PERMISSION_REQUEST, Result: vo.LOGIN_FAIL}}, nil
	}
	token := res.Header.Get(server.Token)
	if len(token) == 0 {
		return s.UC.AdminLogin(ctx, req)
	}
	r := s.role.VerifyManagerToken(token)
	if r {
		return &pb.AdminLoginReply{Common: &pb.CommonReply{Code: vo.SUCCESS_REQUEST, Result: vo.LOGIN_SUCCESS}, Data: []string{token}}, nil
	} else {
		return &pb.AdminLoginReply{Common: &pb.CommonReply{Code: vo.PERMISSION_REQUEST, Result: vo.LOGIN_FAIL}}, nil
	}
}
func (s *UserService) LogOut(ctx context.Context, req *pb.LogOutRequest) (*pb.LogOutReply, error) {
	return s.UC.LogOut(ctx, req)
}
func (s *UserService) verifyToken(ctx *context.Context) (bool, string) {

	res, ok := http.RequestFromServerContext(*ctx)
	if !ok {
		s.log.Log(log.LevelError, "parse context failed")
	}
	token := res.Header.Get(server.Token)
	if len(token) != 0 {
		response, _ := s.UC.VerifyToken(context.Background(), &pb.VerifyTokenRequest{
			Token: token,
		})
		if response.Common.Code == vo.SUCCESS_REQUEST {
			return true, token
		} else {
			return false, response.Common.Result
		}
	}
	return true, ""
}
func (s *UserService) QueryAllUser(ctx context.Context, req *pb.QueryAllUserRequest) (*pb.QueryAllUserResponse, error) {
	return s.UC.QueryAllUser(ctx, req)
}
func (s *UserService) SetAdmin(ctx context.Context, req *pb.SetAdminRequest) (*pb.SetAdminReply, error) {
	return s.UC.SetAdmin(ctx, req)
}
