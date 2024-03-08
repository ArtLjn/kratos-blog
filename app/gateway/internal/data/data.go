package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	grpcx "google.golang.org/grpc"
	"kratos-blog/api/v1/blog"
	"kratos-blog/api/v1/comment"
	"kratos-blog/api/v1/friend"
	"kratos-blog/api/v1/tag"
	"kratos-blog/api/v1/user"
	"kratos-blog/app/gateway/internal/conf"
	"kratos-blog/pkg/server"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRegistrar, NewDiscovery)

// Data .
type Data struct {
	log *log.Helper
	c   *conf.Bootstrap
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, c: c}, nil
}

// NewRegistrar add consul
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	c.Namespace = ""
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
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
