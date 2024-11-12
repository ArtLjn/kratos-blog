package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/friend"
)

type FriendRepo interface {
	CreateFriend(ctx context.Context, request *pb.CreateFriendRequest) *pb.CreateFriendReply
	DeleteFriend(ctx context.Context, request *pb.DeleteFriendRequest) *pb.DeleteFriendReply
	SearchAllFriend(ctx context.Context, request *pb.ListFriendRequest) *pb.ListFriendReply
	UpdateFriend(ctx context.Context, request *pb.UpdateFriendRequest) *pb.UpdateFriendReply
	GetFriendByCond(ctx context.Context, request *pb.GetFriendRequest) *pb.GetFriendReply
}

type FriendUseCase struct {
	repo FriendRepo
	log  *log.Helper
}

func NewFriendUseCase(repo FriendRepo, logger log.Logger) *FriendUseCase {
	return &FriendUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (f *FriendUseCase) CreateFriend(ctx context.Context, request *pb.CreateFriendRequest) *pb.CreateFriendReply {
	return f.repo.CreateFriend(ctx, request)
}
func (f *FriendUseCase) DeleteFriend(ctx context.Context, request *pb.DeleteFriendRequest) *pb.DeleteFriendReply {
	return f.repo.DeleteFriend(ctx, request)
}
func (f *FriendUseCase) SearchAllFriend(ctx context.Context, request *pb.ListFriendRequest) *pb.ListFriendReply {
	return f.repo.SearchAllFriend(ctx, request)
}
func (f *FriendUseCase) UpdateFriend(ctx context.Context, request *pb.UpdateFriendRequest) *pb.UpdateFriendReply {
	return f.repo.UpdateFriend(ctx, request)
}
func (f *FriendUseCase) GetFriendByCond(ctx context.Context, request *pb.GetFriendRequest) *pb.GetFriendReply {
	return f.repo.GetFriendByCond(ctx, request)
}
