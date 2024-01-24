package service

import (
	"context"
	pb "kratos-blog/api/v1/comment"
	"kratos-blog/app/comment/internal/biz"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	uc *biz.CommUseCase
}

func NewCommentService(uc *biz.CommUseCase) *CommentService {
	return &CommentService{uc: uc}
}

func (s *CommentService) AddComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	return &pb.CreateCommentReply{}, nil
}
func (s *CommentService) AddReward(ctx context.Context, req *pb.CreateRewardRequest) (*pb.CreateRewardReply, error) {
	return &pb.CreateRewardReply{}, nil
}
func (s *CommentService) ExtractParentComments(ctx context.Context, req *pb.ExtractParentCommentsRequest) (*pb.ExtractParentCommentsReply, error) {
	return &pb.ExtractParentCommentsReply{}, nil
}
