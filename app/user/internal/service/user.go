package service

import (
	"context"
	pb "kratos-blog/api/v1/user"
	"kratos-blog/app/user/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return s.uc.Register(ctx, req), nil
}
func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return s.uc.Login(ctx, req), nil
}
func (s *UserService) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	return s.uc.SendEmail(ctx, req), nil
}
func (s *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return s.uc.UpdateUserPassword(ctx, req), nil
}
func (s *UserService) SetBlack(ctx context.Context, req *pb.SetBlackRequest) (*pb.SetBlackReply, error) {
	return s.uc.SetUserBlack(ctx, req), nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return s.uc.GetUserMsg(ctx, req), nil
}
func (s *UserService) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginReply, error) {
	return s.uc.AdminLoginRsg(ctx, req), nil
}
func (s *UserService) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenReply, error) {
	return s.uc.VerifyToken(ctx, req), nil
}
func (s *UserService) LogOut(ctx context.Context, req *pb.LogOutRequest) (*pb.LogOutReply, error) {
	return s.uc.LogOut(ctx, req), nil
}
func (s *UserService) SendEmailCommon(ctx context.Context, req *pb.SendEmailCommonRequest) (*pb.SendEmailCommonReply, error) {
	return s.uc.SendEmailCommon(ctx, req), nil
}
func (s *UserService) QueryAllUser(ctx context.Context, req *pb.QueryAllUserRequest) (*pb.QueryAllUserResponse, error) {
	return s.uc.QueryAllUser(ctx, req), nil
}
func (s *UserService) SetAdmin(ctx context.Context, req *pb.SetAdminRequest) (*pb.SetAdminReply, error) {
	return s.uc.SetAdmin(ctx, req), nil
}
