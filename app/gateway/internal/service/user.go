package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "kratos-blog/api/v1/user"
	"kratos-blog/pkg/server"
	"kratos-blog/pkg/vo"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  pb.UserClient
	log *log.Helper
}

func NewUserService(logger log.Logger, uc pb.UserClient) *UserService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &UserService{uc: uc, log: l}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return s.uc.CreateUser(ctx, req)
}
func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	ok, token := s.verifyToken(&ctx)
	if ok && len(token) != 0 {
		return &pb.LoginReply{Common: &pb.CommonReply{Code: vo.SUCCESS_REQUEST, Result: vo.LOGIN_SUCCESS}, Data: []string{token}}, nil
	} else if !ok {
		return &pb.LoginReply{Common: &pb.CommonReply{Code: vo.PERMISSION_REQUEST, Result: token}}, nil
	}
	return s.uc.LoginUser(ctx, req)
}
func (s *UserService) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	return s.uc.SendEmail(ctx, req)
}
func (s *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return s.uc.UpdatePassword(ctx, req)
}
func (s *UserService) SetBlack(ctx context.Context, req *pb.SetBlackRequest) (*pb.SetBlackReply, error) {
	return s.uc.SetBlack(ctx, req)
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return s.uc.GetUser(ctx, req)
}
func (s *UserService) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginReply, error) {
	ok, token := s.verifyToken(&ctx)
	if ok && len(token) != 0 {
		return &pb.AdminLoginReply{Common: &pb.CommonReply{Code: vo.SUCCESS_REQUEST, Result: vo.LOGIN_SUCCESS}, Data: []string{token}}, nil
	} else if !ok {
		return &pb.AdminLoginReply{Common: &pb.CommonReply{Code: vo.PERMISSION_REQUEST, Result: token}}, nil
	}
	return s.uc.AdminLogin(ctx, req)
}
func (s *UserService) LogOut(ctx context.Context, req *pb.LogOutRequest) (*pb.LogOutReply, error) {
	return s.uc.LogOut(ctx, req)
}
func (s *UserService) verifyToken(ctx *context.Context) (bool, string) {

	res, ok := http.RequestFromServerContext(*ctx)
	if !ok {
		s.log.Log(log.LevelError, "parse context failed")
	}
	token := res.Header.Get(server.Token)
	if len(token) != 0 {
		response, _ := s.uc.VerifyToken(context.Background(), &pb.VerifyTokenRequest{
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
