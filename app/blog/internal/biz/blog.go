package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/blog"
)

type BlogRepo interface {
	CreateBlog(ctx context.Context, request *pb.CreateBlogRequest) (string, error)
	UpdateBlog(ctx context.Context, request *pb.UpdateBlogRequest) (string, error)
	DeleteBlog(ctx context.Context, request *pb.DeleteBlogRequest) (string, error)
	UpdateAllCommentStatus(ctx context.Context, request *pb.UpdateAllCommentStatusRequest) (string, error)
}

type BlogUseCase struct {
	repo BlogRepo
	log  *log.Helper
}

func NewBlogUseCase(repo BlogRepo, logger log.Logger) *BlogUseCase {
	return &BlogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BlogUseCase) AddBlog(ctx context.Context, request *pb.CreateBlogRequest) *pb.CreateBlogReply {
	u, err := uc.repo.CreateBlog(ctx, request)
	status := func(code int64, msg string) *pb.CreateBlogReply {
		return &pb.CreateBlogReply{
			Code:   code,
			Result: msg,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u)
	}
	return status(200, u)
}

func (uc *BlogUseCase) UpdateBlog(ctx context.Context, request *pb.UpdateBlogRequest) *pb.UpdateBlogReply {
	u, err := uc.repo.UpdateBlog(ctx, request)
	status := func(code int64, msg string) *pb.UpdateBlogReply {
		return &pb.UpdateBlogReply{
			Code:   code,
			Result: msg,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u)
	}
	return status(200, u)
}

func (uc *BlogUseCase) DeleteBlog(ctx context.Context, request *pb.DeleteBlogRequest) *pb.DeleteBlogReply {
	u, err := uc.repo.DeleteBlog(ctx, request)
	status := func(code int64, msg string) *pb.DeleteBlogReply {
		return &pb.DeleteBlogReply{
			Code:   code,
			Result: msg,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u)
	}
	return status(200, u)
}

//
//func (uc *BlogUseCase) UpdateAllCommentStatus(ctx context.Context, request *pb.UpdateAllCommentStatusRequest) *pb.UpdateAllCommentStatusReply {
//	u, err := uc.repo.UpdateAllCommentStatus(ctx, request)
//}
