package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	blog2 "kratos-blog/api/v1/blog"
	comment2 "kratos-blog/api/v1/comment"
	friend2 "kratos-blog/api/v1/friend"
	photo2 "kratos-blog/api/v1/photo"
	tag2 "kratos-blog/api/v1/tag"
	user2 "kratos-blog/api/v1/user"
	"kratos-blog/app/gateway/internal/conf"
	"kratos-blog/app/gateway/internal/data"
	"kratos-blog/app/gateway/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cf *conf.Bootstrap,
	logger log.Logger,
	user *service.UserService,
	comment *service.CommentService,
	blog *service.BlogService,
	tag *service.TagService,
	friend *service.FriendService,
	photo *service.PhotoService,
	// 过滤器
	filter *data.FilterRepo,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.Filter(
			//filter.FilterPermission(cf.Path.GetWhite(), cf.Path.GetBlack()),
			filter.DeleteCache(),
			filter.AllowDomainsMiddleware(*cf.GetDomain()),
		),
	}
	c := cf.Server
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
	// 注册工具类接口
	router := gin.Default()
	NewGinRouter(cf, router)
	srv.HandlePrefix("/util", router)

	user2.RegisterUserHTTPServer(srv, user)
	comment2.RegisterCommentHTTPServer(srv, comment)
	blog2.RegisterBlogHTTPServer(srv, blog)
	tag2.RegisterTagHTTPServer(srv, tag)
	friend2.RegisterFriendHTTPServer(srv, friend)
	photo2.RegisterPhotoHTTPServer(srv, photo)
	return srv
}
