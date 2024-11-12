package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/api/blog"
	pb "kratos-blog/api/comment"
	"kratos-blog/app/gateway/internal/data"
	"kratos-blog/pkg/vo"
	"strconv"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	UC   pb.CommentClient
	bc   blog.BlogClient
	role *data.Role
	log  *log.Helper
	q    *data.Mq
}

func NewCommentService(logger log.Logger, uc pb.CommentClient,
	bc blog.BlogClient, role *data.Role, q *data.Mq) *CommentService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &CommentService{UC: uc, bc: bc, log: l, role: role, q: q}
}

func (s *CommentService) AddComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {

	res := s.role.QueryUserMsg(ctx)

	if err := s.VerifyComment(res); err != nil {
		return &pb.CreateCommentReply{Code: vo.PERMISSION_REQUEST, Result: err.Error()}, nil
	}
	req.Name = &res.U.Data[0]
	req.Email = &res.U.Data[1]
	if err := s.q.PushComment(req); err != nil {
		return &pb.CreateCommentReply{Code: vo.BAD_REQUEST, Result: err.Error()}, nil
	}
	return &pb.CreateCommentReply{Code: vo.SUCCESS_REQUEST, Result: vo.MESSAGE_PROCESSING}, nil
}
func (s *CommentService) AddReward(ctx context.Context, req *pb.CreateRewardRequest) (*pb.CreateRewardReply, error) {
	res := s.role.QueryUserMsg(ctx)

	if err := s.VerifyComment(res); err != nil {
		return &pb.CreateRewardReply{Code: vo.PERMISSION_REQUEST, Result: err.Error()}, nil
	}
	req.Name = &res.U.Data[0]
	req.Email = &res.U.Data[1]
	if err := s.q.PushReward(req); err != nil {
		return &pb.CreateRewardReply{Code: vo.BAD_REQUEST, Result: err.Error()}, nil
	}
	return &pb.CreateRewardReply{Code: vo.SUCCESS_REQUEST, Result: vo.MESSAGE_PROCESSING}, nil
}

// ExtractParentComments 查询评论
func (s *CommentService) ExtractParentComments(ctx context.Context, req *pb.ExtractParentCommentsRequest) (*pb.ExtractParentCommentsReply, error) {
	return s.UC.ExtractParentComments(ctx, req)
}

func (s *CommentService) VerifyComment(res *data.Permission) error {
	// 判断用户是否在黑名单中
	if r, err := strconv.ParseBool(res.U.Data[5]); r || err != nil {
		return fmt.Errorf(vo.PERMISSION_ERROR)
	}
	// 如果用户名和，邮箱都为为空，则返回错误
	if len(res.U.Data[0]) == 0 || len(res.U.Data[1]) == 0 {
		return fmt.Errorf(vo.PERMISSION_ERROR)
	}
	return nil
}

func (s *CommentService) HasAllowComment(ctx context.Context, articleId string) error {
	if len(articleId) != 0 {
		if articleId == "friendlink" || articleId == "message" {
			res, _ := s.bc.GetCommentPower(ctx, &blog.GetCommentPowerRq{
				Path: articleId,
			})
			if !res.GetAllow() {
				return errors.New(vo.TALK_FORBIDDEN)
			}
			return nil
		}
		id, err := strconv.Atoi(articleId)
		if err != nil {
			return errors.New(vo.JSON_ERROR)
		}
		res, err := s.bc.GetBlogByID(ctx, &blog.GetBlogIDRequest{
			Id: int64(id),
		})
		if err != nil {
			return errors.New(vo.QUERY_FAIL)
		}
		if !res.Data.GetComment() {
			return errors.New(vo.TALK_FORBIDDEN)
		}
	}
	return nil
}
