package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/comment"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	uc  pb.CommentClient
	log *log.Helper
}

func NewCommentService(logger log.Logger, uc pb.CommentClient) *CommentService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &CommentService{uc: uc, log: l}
}

func (s *CommentService) AddComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	return s.uc.AddComment(ctx, req)
}
func (s *CommentService) AddReward(ctx context.Context, req *pb.CreateRewardRequest) (*pb.CreateRewardReply, error) {
	return s.uc.AddReward(ctx, req)
}
func (s *CommentService) ExtractParentComments(ctx context.Context, req *pb.ExtractParentCommentsRequest) (*pb.ExtractParentCommentsReply, error) {
	return s.uc.ExtractParentComments(ctx, req)
}
