package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "kratos-blog/api/blog"
	v3 "kratos-blog/api/friend"
	v2 "kratos-blog/api/tag"
	"kratos-blog/app/blog/internal/conf"
	"kratos-blog/app/blog/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cf *conf.Bootstrap, blog *service.BlogService,
	friend *service.FriendService,
	tag *service.TagService) *http.Server {
	c := cf.Server
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// 注册服务接口
	v1.RegisterBlogHTTPServer(srv, blog)
	v2.RegisterTagHTTPServer(srv, tag)
	v3.RegisterFriendHTTPServer(srv, friend)
	return srv
}
