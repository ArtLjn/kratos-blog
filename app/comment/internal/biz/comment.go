package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/api/v1/comment"
)

type CommRepo interface {
	CheckWords(word string) []string
	AddComment(ctx context.Context, req *comment.CreateCommentRequest) *comment.CreateCommentReply
}

type CommUseCase struct {
	repo CommRepo
	log  *log.Helper
}

func NewUCommUseCase(repo CommRepo, logger log.Logger) *CommUseCase {
	return &CommUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *CommUseCase) AddComment(ctx context.Context, req *comment.CreateCommentRequest) *comment.CreateCommentReply {
	resp := c.repo.AddComment(ctx, req)
	return resp
}
