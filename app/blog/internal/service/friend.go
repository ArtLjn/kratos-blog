package service

import (
	"context"
	pb "kratos-blog/api/friend"
	"kratos-blog/app/blog/internal/biz"
)

type FriendService struct {
	pb.UnimplementedFriendServer
	uc *biz.FriendUseCase
}

func NewFriendService(uc *biz.FriendUseCase) *FriendService {
	return &FriendService{uc: uc}
}

func (s *FriendService) CreateFriend(ctx context.Context, req *pb.CreateFriendRequest) (*pb.CreateFriendReply, error) {
	return s.uc.CreateFriend(ctx, req), nil
}
func (s *FriendService) UpdateFriend(ctx context.Context, req *pb.UpdateFriendRequest) (*pb.UpdateFriendReply, error) {
	return s.uc.UpdateFriend(ctx, req), nil
}
func (s *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendReply, error) {
	return s.uc.DeleteFriend(ctx, req), nil
}
func (s *FriendService) GetFriend(ctx context.Context, req *pb.GetFriendRequest) (*pb.GetFriendReply, error) {
	return s.uc.GetFriendByCond(ctx, req), nil
}
func (s *FriendService) ListFriend(ctx context.Context, req *pb.ListFriendRequest) (*pb.ListFriendReply, error) {
	return s.uc.SearchAllFriend(ctx, req), nil
}
