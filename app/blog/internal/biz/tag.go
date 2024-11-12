package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/tag"
)

type TagRepo interface {
	CreateTag(ctx context.Context, request *pb.CreateTagRequest) *pb.CreateTagReply
	DeleteTag(ctx context.Context, request *pb.DeleteTagRequest) *pb.DeleteTagReply
	SearchAllTag(ctx context.Context, request *pb.ListTagRequest) *pb.ListTagReply
}

type TagUseCase struct {
	repo TagRepo
	log  *log.Helper
}

func NewTagUseCase(repo TagRepo, logger log.Logger) *TagUseCase {
	return &TagUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (t *TagUseCase) CreateTag(ctx context.Context, request *pb.CreateTagRequest) *pb.CreateTagReply {
	return t.repo.CreateTag(ctx, request)
}

func (t *TagUseCase) DeleteTag(ctx context.Context, request *pb.DeleteTagRequest) *pb.DeleteTagReply {
	return t.repo.DeleteTag(ctx, request)
}

func (t *TagUseCase) SearchAllTag(ctx context.Context, request *pb.ListTagRequest) *pb.ListTagReply {
	return t.repo.SearchAllTag(ctx, request)
}
