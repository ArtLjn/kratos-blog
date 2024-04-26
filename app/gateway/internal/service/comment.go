package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/api/v1/blog"
	pb "kratos-blog/api/v1/comment"
	"kratos-blog/app/gateway/internal/data"
	"kratos-blog/pkg/vo"
	"strconv"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	uc   pb.CommentClient
	bc   blog.BlogClient
	role *data.Role
	log  *log.Helper
}

func NewCommentService(logger log.Logger, uc pb.CommentClient, bc blog.BlogClient, role *data.Role) *CommentService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &CommentService{uc: uc, bc: bc, log: l, role: role}
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

func (s *CommentService) AddComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	if err := s.HasAllowComment(ctx, req.ArticleId); err != nil {
		return &pb.CreateCommentReply{Code: vo.BAD_REQUEST, Result: err.Error()}, nil
	}
	res := s.role.QueryUserMsg(ctx)
	if r, err := strconv.ParseBool(res.U.Data[5]); !r || err != nil {
		return &pb.CreateCommentReply{Code: vo.PERMISSION_REQUEST, Result: vo.PERMISSION_ERROR}, nil
	}
	if len(res.U.Data[0]) == 0 || len(res.U.Data[1]) == 0 {
		return &pb.CreateCommentReply{Code: vo.PERMISSION_REQUEST, Result: vo.PERMISSION_ERROR}, nil
	}
	req.Name = &res.U.Data[0]
	req.Email = &res.U.Data[1]
	return s.uc.AddComment(ctx, req)
}
func (s *CommentService) AddReward(ctx context.Context, req *pb.CreateRewardRequest) (*pb.CreateRewardReply, error) {
	if err := s.HasAllowComment(ctx, req.ArticleId); err != nil {
		return &pb.CreateRewardReply{Code: vo.BAD_REQUEST, Result: err.Error()}, nil
	}
	res := s.role.QueryUserMsg(ctx)
	if r, err := strconv.ParseBool(res.U.Data[5]); !r || err != nil {
		return &pb.CreateRewardReply{Code: vo.PERMISSION_REQUEST, Result: vo.PERMISSION_ERROR}, nil
	}
	if len(res.U.Data[0]) == 0 || len(res.U.Data[1]) == 0 {
		return &pb.CreateRewardReply{Code: vo.PERMISSION_REQUEST, Result: vo.PERMISSION_ERROR}, nil
	}
	req.Name = &res.U.Data[0]
	req.Email = &res.U.Data[1]
	return s.uc.AddReward(ctx, req)
}
func (s *CommentService) ExtractParentComments(ctx context.Context, req *pb.ExtractParentCommentsRequest) (*pb.ExtractParentCommentsReply, error) {
	return s.uc.ExtractParentComments(ctx, req)
}
