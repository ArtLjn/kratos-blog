package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/comment"
	"kratos-blog/app/gateway/internal/data"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	uc   pb.CommentClient
	role *data.Role
	log  *log.Helper
}

func NewCommentService(logger log.Logger, uc pb.CommentClient, role *data.Role) *CommentService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &CommentService{uc: uc, log: l, role: role}
}

func (s *CommentService) AddComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	res := s.role.QueryUserMsg(ctx)
	req.Name = &res.U.Data[0]
	req.Email = &res.U.Data[1]
	return s.uc.AddComment(ctx, req)
}
func (s *CommentService) AddReward(ctx context.Context, req *pb.CreateRewardRequest) (*pb.CreateRewardReply, error) {
	res := s.role.QueryUserMsg(ctx)
	req.Name = &res.U.Data[0]
	req.Email = &res.U.Data[1]
	return s.uc.AddReward(ctx, req)
}
func (s *CommentService) ExtractParentComments(ctx context.Context, req *pb.ExtractParentCommentsRequest) (*pb.ExtractParentCommentsReply, error) {
	return s.uc.ExtractParentComments(ctx, req)
}

// ParseJSONToStruct :dev parse the string into the struct
func ParseJSONToStruct(jsonStr interface{}, resultStruct interface{}) error {
	f, _ := json.Marshal(jsonStr)
	err := json.Unmarshal(f, resultStruct)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %s", err)
	}
	return nil
}
