package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/tag"
)

type TagService struct {
	pb.UnimplementedTagServer
	uc  pb.TagClient
	log *log.Helper
}

func NewTagService(logger log.Logger, uc pb.TagClient) *TagService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &TagService{uc: uc, log: l}
}

func (s *TagService) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.CreateTagReply, error) {
	return s.uc.CreateTag(ctx, req)
}
func (s *TagService) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*pb.DeleteTagReply, error) {
	return s.uc.DeleteTag(ctx, req)
}
func (s *TagService) ListTag(ctx context.Context, req *pb.ListTagRequest) (*pb.ListTagReply, error) {
	return s.uc.ListTag(ctx, req)
}
