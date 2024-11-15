package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	blog2 "kratos-blog/api/blog"
	comment2 "kratos-blog/api/comment"
	friend2 "kratos-blog/api/friend"
	photo2 "kratos-blog/api/photo"
	tag2 "kratos-blog/api/tag"
	user2 "kratos-blog/api/user"
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
			logging.Server(logger),
		),
		http.Filter(
			filter.DeleteCache(),
			filter.CommentLimiter(),
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
	user2.RegisterUserHTTPServer(srv, user)
	comment2.RegisterCommentHTTPServer(srv, comment)
	blog2.RegisterBlogHTTPServer(srv, blog)
	tag2.RegisterTagHTTPServer(srv, tag)
	friend2.RegisterFriendHTTPServer(srv, friend)
	photo2.RegisterPhotoHTTPServer(srv, photo)
	return srv
}
