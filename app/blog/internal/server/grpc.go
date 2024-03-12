package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "kratos-blog/api/v1/blog"
	v3 "kratos-blog/api/v1/friend"
	v4 "kratos-blog/api/v1/photo"
	v2 "kratos-blog/api/v1/tag"
	"kratos-blog/app/blog/internal/conf"
	"kratos-blog/app/blog/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, blog *service.BlogService,
	friend *service.FriendService,
	tag *service.TagService,
	photo *service.PhotoService,
	logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterBlogServer(srv, blog)
	v2.RegisterTagServer(srv, tag)
	v3.RegisterFriendServer(srv, friend)
	v4.RegisterPhotoServer(srv, photo)
	return srv
}
