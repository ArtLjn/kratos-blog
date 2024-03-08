package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/friend"
)

type FriendService struct {
	pb.UnimplementedFriendServer
	uc  pb.FriendClient
	log *log.Helper
}

func NewFriendService(uc pb.FriendClient, logger log.Logger) *FriendService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &FriendService{uc: uc, log: l}
}

func (s *FriendService) CreateFriend(ctx context.Context, req *pb.CreateFriendRequest) (*pb.CreateFriendReply, error) {
	return s.uc.CreateFriend(ctx, req)
}
func (s *FriendService) UpdateFriend(ctx context.Context, req *pb.UpdateFriendRequest) (*pb.UpdateFriendReply, error) {
	return s.uc.UpdateFriend(ctx, req)
}
func (s *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendReply, error) {
	return s.uc.DeleteFriend(ctx, req)
}
func (s *FriendService) GetFriend(ctx context.Context, req *pb.GetFriendRequest) (*pb.GetFriendReply, error) {
	return s.uc.GetFriend(ctx, req)
}
func (s *FriendService) ListFriend(ctx context.Context, req *pb.ListFriendRequest) (*pb.ListFriendReply, error) {
	return s.uc.ListFriend(ctx, req)
}
