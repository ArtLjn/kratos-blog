package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/photo"
)

type PhotoRepo interface {
	CreatePhoto(ctx context.Context, request *pb.CreatePhotoRequest) *pb.CreatePhotoReply
	DeletePhoto(ctx context.Context, request *pb.DeletePhotoRequest) *pb.DeletePhotoReply
	SearchAllPhoto(ctx context.Context, request *pb.ListPhotoRequest) *pb.ListPhotoReply
}

type PhotoUseCase struct {
	repo PhotoRepo
	log  *log.Helper
}

func NewPhotoUseCase(repo PhotoRepo, logger log.Logger) *PhotoUseCase {
	return &PhotoUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (t *PhotoUseCase) CreatePhoto(ctx context.Context, request *pb.CreatePhotoRequest) *pb.CreatePhotoReply {
	return t.repo.CreatePhoto(ctx, request)
}

func (t *PhotoUseCase) DeletePhoto(ctx context.Context, request *pb.DeletePhotoRequest) *pb.DeletePhotoReply {
	return t.repo.DeletePhoto(ctx, request)
}

func (t *PhotoUseCase) SearchAllPhoto(ctx context.Context, request *pb.ListPhotoRequest) *pb.ListPhotoReply {
	return t.repo.SearchAllPhoto(ctx, request)
}
