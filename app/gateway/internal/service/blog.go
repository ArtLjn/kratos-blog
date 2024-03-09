package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/blog"
	"kratos-blog/app/gateway/internal/data"
)

type BlogService struct {
	pb.UnimplementedBlogServer
	uc   pb.BlogClient
	log  *log.Helper
	role *data.Role
}

func NewBlogService(logger log.Logger, uc pb.BlogClient, role *data.Role) *BlogService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &BlogService{uc: uc, log: l, role: role}
}

func (s *BlogService) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogReply, error) {
	return s.uc.CreateBlog(ctx, req)
}
func (s *BlogService) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.UpdateBlogReply, error) {
	return s.uc.UpdateBlog(ctx, req)
}
func (s *BlogService) UpdateIndividualFields(ctx context.Context, req *pb.UpdateIndividualFieldsRequest) (*pb.UpdateIndividualFieldsReply, error) {
	return s.uc.UpdateIndividualFields(ctx, req)
}
func (s *BlogService) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.DeleteBlogReply, error) {
	return s.uc.DeleteBlog(ctx, req)
}
func (s *BlogService) GetBlogByTag(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogReply, error) {
	req = &pb.GetBlogRequest{
		Tag: req.Tag,
		Permission: &pb.Permission{
			Admin: s.role.QueryUserMsg(ctx).GetRole().CheckPermission(),
		},
	}
	return s.uc.GetBlogByTag(ctx, req)
}
func (s *BlogService) ListBlog(ctx context.Context, req *pb.ListBlogRequest) (*pb.ListBlogReply, error) {
	req = &pb.ListBlogRequest{
		Permission: &pb.Permission{
			Admin: s.role.QueryUserMsg(ctx).GetRole().CheckPermission(),
		},
	}
	return s.uc.ListBlog(ctx, req)
}
func (s *BlogService) GetBlogByID(ctx context.Context, req *pb.GetBlogIDRequest) (*pb.GetBlogIDReply, error) {
	req = &pb.GetBlogIDRequest{
		Id: req.Id,
		Permission: &pb.Permission{
			Admin: s.role.QueryUserMsg(ctx).GetRole().CheckPermission(),
		},
	}
	return s.uc.GetBlogByID(ctx, req)
}
func (s *BlogService) GetBlogByTitle(ctx context.Context, req *pb.GetBlogByTitleRequest) (*pb.GetBlogByTitleReply, error) {
	return s.uc.GetBlogByTitle(ctx, req)
}
func (s *BlogService) UpdateOnly(ctx context.Context, req *pb.UpdateOnlyRequest) (*pb.UpdateOnlyReply, error) {
	return s.uc.UpdateOnly(ctx, req)
}
func (s *BlogService) CacheBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogReply, error) {
	return s.uc.CacheBlog(ctx, req)
}
func (s *BlogService) GetCacheBlog(ctx context.Context, req *pb.ListBlogRequest) (*pb.ListCacheReply, error) {
	return s.uc.GetCacheBlog(ctx, req)
}
func (s *BlogService) DeleteCacheBlog(ctx context.Context, req *pb.DeleteCacheBlogRequest) (*pb.DeleteCacheBlogReply, error) {
	return s.uc.DeleteCacheBlog(ctx, req)
}
func (s *BlogService) AddSuggestBlog(ctx context.Context, req *pb.SuggestBlogRequest) (*pb.SuggestBlogReply, error) {
	return s.uc.AddSuggestBlog(ctx, req)
}
func (s *BlogService) GetAllSuggest(ctx context.Context, req *pb.SearchAllSuggest) (*pb.SearchAllReply, error) {
	return s.uc.GetAllSuggest(ctx, req)
}
func (s *BlogService) DeleteSuggestBlog(ctx context.Context, req *pb.SuggestBlogRequest) (*pb.SuggestBlogReply, error) {
	return s.uc.DeleteSuggestBlog(ctx, req)
}
