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
	UpdateIndividualFields(ctx context.Context, request *pb.UpdateIndividualFieldsRequest) (string, error)
	GetByTagName(ctx context.Context, request *pb.GetBlogRequest) (string, []*pb.BlogData, error)
	ListBlog(ctx context.Context, request *pb.ListBlogRequest) (string, []*pb.BlogData, error)
	QueryBlogById(ctx context.Context, request *pb.GetBlogIDRequest) (string, *pb.BlogData, error)
	SetBlogVisit(id int)
	QueryBlogByTitle(ctx context.Context, request *pb.GetBlogByTitleRequest) (string, []*pb.BlogData, error)
	UpdateOnly(ctx context.Context, request *pb.UpdateOnlyRequest) *pb.UpdateOnlyReply
	CacheBlog(ctx context.Context, request *pb.CreateBlogRequest) *pb.CreateBlogReply
	GetCacheBlog(ctx context.Context) *pb.ListCacheReply
	DeleteCacheBlog(ctx context.Context, request *pb.DeleteCacheBlogRequest) *pb.DeleteCacheBlogReply
	AddSuggestBlog(ctx context.Context, request *pb.SuggestBlogRequest) *pb.SuggestBlogReply
	GetAllSuggestBlog(ctx context.Context, request *pb.SearchAllSuggest) *pb.SearchAllReply
	DeleteSuggestBlog(ctx context.Context, request *pb.SuggestBlogRequest) *pb.SuggestBlogReply
	UpdateBlogVisitsCount()
	GetCommentPower(path string) (bool, error)
}

type BlogUseCase struct {
	repo BlogRepo
	log  *log.Helper
}

func NewBlogUseCase(repo BlogRepo, logger log.Logger) *BlogUseCase {
	return &BlogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func SetStatus(code int64, msg string) *pb.CommonReply {
	return &pb.CommonReply{
		Code:   code,
		Result: msg,
	}
}

func (uc *BlogUseCase) UpdateBlogVisit(ctx context.Context, request *pb.UpdateBlogVisitRq) *pb.UpdateBlogVisitRq {
	uc.repo.UpdateBlogVisitsCount()
	return &pb.UpdateBlogVisitRq{}
}

func (uc *BlogUseCase) AddBlog(ctx context.Context, request *pb.CreateBlogRequest) *pb.CreateBlogReply {
	u, err := uc.repo.CreateBlog(ctx, request)
	status := func(code int64, msg string) *pb.CreateBlogReply {
		return &pb.CreateBlogReply{
			Common: SetStatus(code, msg),
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
			Common: SetStatus(code, msg),
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
			Common: SetStatus(code, msg),
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u)
	}
	return status(200, u)
}

func (uc *BlogUseCase) UpdateIndividualFields(ctx context.Context, request *pb.UpdateIndividualFieldsRequest) *pb.UpdateIndividualFieldsReply {
	u, err := uc.repo.UpdateIndividualFields(ctx, request)
	status := func(code int64, msg string) *pb.UpdateIndividualFieldsReply {
		return &pb.UpdateIndividualFieldsReply{
			Common: SetStatus(code, msg),
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
			Common: SetStatus(code, msg),
			List:   d,
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
			Common: SetStatus(code, msg),
			List:   d,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u, nil)
	}
	return status(200, u, d)
}

func (uc *BlogUseCase) QueryBlogByID(ctx context.Context, request *pb.GetBlogIDRequest) *pb.GetBlogIDReply {
	u, d, err := uc.repo.QueryBlogById(ctx, request)
	status := func(code int64, msg string, data *pb.BlogData) *pb.GetBlogIDReply {
		return &pb.GetBlogIDReply{
			Common: SetStatus(code, msg),
			Data:   data,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u, d)
	}
	uc.repo.SetBlogVisit(int(request.GetId()))
	return status(200, u, d)
}

func (uc *BlogUseCase) QueryBlogByTitle(ctx context.Context, request *pb.GetBlogByTitleRequest) *pb.GetBlogByTitleReply {
	u, d, err := uc.repo.QueryBlogByTitle(ctx, request)
	status := func(code int64, msg string, data []*pb.BlogData) *pb.GetBlogByTitleReply {
		return &pb.GetBlogByTitleReply{
			Common: SetStatus(code, msg),
			Data:   data,
		}
	}
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return status(400, u, d)
	}
	return status(200, u, d)
}

func (uc *BlogUseCase) UpdateOnly(ctx context.Context, request *pb.UpdateOnlyRequest) *pb.UpdateOnlyReply {
	return uc.repo.UpdateOnly(ctx, request)
}

func (uc *BlogUseCase) CacheBlog(ctx context.Context, request *pb.CreateBlogRequest) *pb.CreateBlogReply {
	return uc.repo.CacheBlog(ctx, request)
}

func (uc *BlogUseCase) GetAllCacheBlog(ctx context.Context) *pb.ListCacheReply {
	return uc.repo.GetCacheBlog(ctx)
}

func (uc *BlogUseCase) DeleteCacheBlog(ctx context.Context, request *pb.DeleteCacheBlogRequest) *pb.DeleteCacheBlogReply {
	return uc.repo.DeleteCacheBlog(ctx, request)
}

func (uc *BlogUseCase) SetSuggestBlog(ctx context.Context, request *pb.SuggestBlogRequest) *pb.SuggestBlogReply {
	return uc.repo.AddSuggestBlog(ctx, request)
}

func (uc *BlogUseCase) GetSuggestBlog(ctx context.Context, request *pb.SearchAllSuggest) *pb.SearchAllReply {
	return uc.repo.GetAllSuggestBlog(ctx, request)
}
func (uc *BlogUseCase) DeleteSuggestBlog(ctx context.Context, request *pb.SuggestBlogRequest) *pb.SuggestBlogReply {
	return uc.repo.DeleteSuggestBlog(ctx, request)
}
func (uc *BlogUseCase) GetCommentPower(ctx context.Context, request *pb.GetCommentPowerRq) *pb.GetCommentPowerReply {
	bo, err := uc.repo.GetCommentPower(request.Path)
	if err != nil {
		uc.log.Log(log.LevelError, err)
		return &pb.GetCommentPowerReply{
			Allow: false,
		}
	}
	return &pb.GetCommentPowerReply{
		Allow: bo,
	}
}
