package service

import (
	"context"
	"kratos-blog/app/blog/internal/biz"

	pb "kratos-blog/api/v1/tag"
)

type TagService struct {
	pb.UnimplementedTagServer
	uc *biz.TagUseCase
}

func NewTagService(uc *biz.TagUseCase) *TagService {
	return &TagService{uc: uc}
}

func (s *TagService) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.CreateTagReply, error) {
	return &pb.CreateTagReply{}, nil
}
func (s *TagService) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*pb.DeleteTagReply, error) {
	return &pb.DeleteTagReply{}, nil
}
func (s *TagService) ListTag(ctx context.Context, req *pb.ListTagRequest) (*pb.ListTagReply, error) {
	return &pb.ListTagReply{}, nil
}
