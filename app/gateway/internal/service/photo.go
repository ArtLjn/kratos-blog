package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-blog/api/v1/photo"
)

type PhotoService struct {
	pb.UnimplementedPhotoServer
	uc  pb.PhotoClient
	log *log.Helper
}

func NewPhotoService(uc pb.PhotoClient, logger log.Logger) *PhotoService {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &PhotoService{uc: uc, log: l}
}

func (s *PhotoService) CreatePhoto(ctx context.Context, req *pb.CreatePhotoRequest) (*pb.CreatePhotoReply, error) {
	return s.uc.CreatePhoto(ctx, req)
}
func (s *PhotoService) DeletePhoto(ctx context.Context, req *pb.DeletePhotoRequest) (*pb.DeletePhotoReply, error) {
	return s.uc.DeletePhoto(ctx, req)
}
func (s *PhotoService) ListPhoto(ctx context.Context, req *pb.ListPhotoRequest) (*pb.ListPhotoReply, error) {
	return s.uc.ListPhoto(ctx, req)
}
