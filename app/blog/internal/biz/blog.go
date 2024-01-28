package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/blog"
	"kratos-blog/pkg/vo"
)

type BlogRepo interface {
	CreateBlog(ctx context.Context, request *pb.CreateBlogRequest) (string, error)
	UpdateBlog(ctx context.Context, request *pb.UpdateBlogRequest) (string, error)
	DeleteBlog(ctx context.Context, request *pb.DeleteBlogRequest) (string, error)
	UpdateAllCommentStatus(ctx context.Context, request *pb.UpdateAllCommentStatusRequest) (string, error)
	GetByTagName(ctx context.Context, request *pb.GetBlogRequest) (string, []*pb.BlogData, error)
	ListBlog(ctx context.Context, request *pb.ListBlogRequest) (string, []*pb.BlogData, error)
	QueryBlogById(ctx context.Context, request *pb.GetBlogIDRequest) (msg string, da pb.BlogData, e error)
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
			Common: &pb.CommonReply{
				Code:   code,
				Result: msg,
			},
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
			Common: &pb.CommonReply{
				Code:   code,
				Result: msg,
			},
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
			Common: &pb.CommonReply{
				Code:   code,
				Result: msg,
			},
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u)
	}
	return status(200, u)
}

func (uc *BlogUseCase) UpdateAllCommentStatus(ctx context.Context, request *pb.UpdateAllCommentStatusRequest) *pb.UpdateAllCommentStatusReply {
	u, err := uc.repo.UpdateAllCommentStatus(ctx, request)
	status := func(code int64, msg string) *pb.UpdateAllCommentStatusReply {
		return &pb.UpdateAllCommentStatusReply{
			Common: &pb.CommonReply{
				Code:   code,
				Result: msg,
			},
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u)
	}
	return status(200, u)
}

func (uc *BlogUseCase) GetByTagName(ctx context.Context, request *pb.GetBlogRequest) *pb.GetBlogReply {
	u, d, err := uc.repo.GetByTagName(ctx, request)
	status := func(code int64, msg string, list []*pb.BlogData) *pb.GetBlogReply {
		return &pb.GetBlogReply{
			Common: &pb.CommonReply{
				Code:   code,
				Result: msg,
			},
			List: d,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u, nil)
	}
	return status(200, u, d)
}

func (uc *BlogUseCase) ListBlog(ctx context.Context, request *pb.ListBlogRequest) *pb.ListBlogReply {
	u, d, err := uc.repo.ListBlog(ctx, request)
	status := func(code int64, msg string, list []*pb.BlogData) *pb.ListBlogReply {
		return &pb.ListBlogReply{
			Common: &pb.CommonReply{
				Code:   code,
				Result: msg,
			},
			List: d,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u, nil)
	} else if len(d) == 0 {
		return status(300, vo.QUERY_EMPTY, nil)
	}
	return status(200, u, d)
}

func (uc *BlogUseCase) QueryBlogByID(ctx context.Context, request *pb.GetBlogIDRequest) *pb.GetBlogIDReply {
	u, d, err := uc.repo.QueryBlogById(ctx, request)
	status := func(code int64, msg string, data pb.BlogData) *pb.GetBlogIDReply {
		return &pb.GetBlogIDReply{
			Common: &pb.CommonReply{
				Code:   code,
				Result: msg,
			},
			Data: &data,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u, d)
	}
	return status(200, u, d)
}
