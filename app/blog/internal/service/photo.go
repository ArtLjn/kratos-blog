package service

import (
	"context"
	pb "kratos-blog/api/photo"
	"kratos-blog/app/blog/internal/biz"
)

type PhotoService struct {
	pb.UnimplementedPhotoServer
	uc *biz.PhotoUseCase
}

func NewPhotoService(uc *biz.PhotoUseCase) *PhotoService {
	return &PhotoService{uc: uc}
}

func (s *PhotoService) CreatePhoto(ctx context.Context, req *pb.CreatePhotoRequest) (*pb.CreatePhotoReply, error) {
	return s.uc.CreatePhoto(ctx, req), nil
}
func (s *PhotoService) DeletePhoto(ctx context.Context, req *pb.DeletePhotoRequest) (*pb.DeletePhotoReply, error) {
	return s.uc.DeletePhoto(ctx, req), nil
}
func (s *PhotoService) ListPhoto(ctx context.Context, req *pb.ListPhotoRequest) (*pb.ListPhotoReply, error) {
	return s.uc.SearchAllPhoto(ctx, req), nil
}
