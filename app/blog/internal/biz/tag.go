package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/tag"
)

type TagRepo interface {
	CreateTag(ctx context.Context, request *pb.CreateTagRequest) *pb.CreateTagReply
}

type TagUseCase struct {
	repo TagRepo
	log  *log.Helper
}

func NewTagUseCase(repo TagRepo, logger log.Logger) *TagUseCase {
	return &TagUseCase{repo: repo, log: log.NewHelper(logger)}
}
