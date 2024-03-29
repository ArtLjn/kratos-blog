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
