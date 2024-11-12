package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	grpcx "google.golang.org/grpc"
	"kratos-blog/api/blog"
	"kratos-blog/api/comment"
	"kratos-blog/api/friend"
	"kratos-blog/api/photo"
	"kratos-blog/api/tag"
	"kratos-blog/api/user"
	"kratos-blog/app/gateway/internal/conf"
	"kratos-blog/pkg/server"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	log  *log.Helper
	c    *conf.Bootstrap
	uc   user.UserClient
	rdb  *redis.Client
	role *Role
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger, uc user.UserClient, rdb *redis.Client, role *Role) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, c: c, rdb: rdb, uc: uc, role: role}, nil
}

func NewRDB(conf *conf.Data) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     conf.Redis.Addr,
			DB:       int(conf.Redis.Db),
			Password: conf.Redis.Password,
		},
	)
}

func NewGrpcServiceClient(serviceName string, rr registry.Discovery) *grpcx.ClientConn {
	endpoint := fmt.Sprintf("discovery:///%s", serviceName)
	selector.SetGlobalSelector(wrr.NewBuilder())
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(rr),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
		grpc.WithHealthCheck(true),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return conn
}

func NewUserServiceClient(rr registry.Discovery) user.UserClient {
	return user.NewUserClient(NewGrpcServiceClient(server.UserService, rr))
}
func NewBlogServiceClient(rr registry.Discovery) blog.BlogClient {
	return blog.NewBlogClient(NewGrpcServiceClient(server.BlogService, rr))
}
func NewCommentServiceClient(rr registry.Discovery) comment.CommentClient {
	return comment.NewCommentClient(NewGrpcServiceClient(server.CommentService, rr))
}
func NewTagServiceClient(rr registry.Discovery) tag.TagClient {
	return tag.NewTagClient(NewGrpcServiceClient(server.BlogService, rr))
}
func NewFriendServiceClient(rr registry.Discovery) friend.FriendClient {
	return friend.NewFriendClient(NewGrpcServiceClient(server.BlogService, rr))
}
func NewPhotoServiceClient(rr registry.Discovery) photo.PhotoClient {
	return photo.NewPhotoClient(NewGrpcServiceClient(server.BlogService, rr))
}
