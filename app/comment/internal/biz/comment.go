package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/api/comment"
)

type CommRepo interface {
	CheckWords(word string) []string
	AddComment(ctx context.Context, req *comment.CreateCommentRequest) *comment.CreateCommentReply
	AddReward(ctx context.Context, req *comment.CreateRewardRequest) *comment.CreateRewardReply
	ExtractParentComments(ctx context.Context,
		req *comment.ExtractParentCommentsRequest) *comment.ExtractParentCommentsReply
}

type CommUseCase struct {
	repo CommRepo
	log  *log.Helper
}

func NewUCommUseCase(repo CommRepo, logger log.Logger) *CommUseCase {
	return &CommUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *CommUseCase) AddComment(ctx context.Context, req *comment.CreateCommentRequest) *comment.CreateCommentReply {
	return c.repo.AddComment(ctx, req)
}

func (c *CommUseCase) AddReward(ctx context.Context, req *comment.CreateRewardRequest) *comment.CreateRewardReply {
	return c.repo.AddReward(ctx, req)
}

func (c *CommUseCase) ExtractParentComments(ctx context.Context, req *comment.ExtractParentCommentsRequest) *comment.ExtractParentCommentsReply {
	return c.repo.ExtractParentComments(ctx, req)
}
