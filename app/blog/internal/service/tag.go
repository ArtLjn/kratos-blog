package service

import (
	"context"
	"kratos-blog/app/blog/internal/biz"

	pb "kratos-blog/api/tag"
)

type TagService struct {
	pb.UnimplementedTagServer
	uc *biz.TagUseCase
}

func NewTagService(uc *biz.TagUseCase) *TagService {
	return &TagService{uc: uc}
}

func (s *TagService) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.CreateTagReply, error) {
	return s.uc.CreateTag(ctx, req), nil
}
func (s *TagService) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*pb.DeleteTagReply, error) {
	return s.uc.DeleteTag(ctx, req), nil
}
func (s *TagService) ListTag(ctx context.Context, req *pb.ListTagRequest) (*pb.ListTagReply, error) {
	return s.uc.SearchAllTag(ctx, req), nil
}
