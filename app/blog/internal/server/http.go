package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "kratos-blog/api/v1/blog"
	v3 "kratos-blog/api/v1/friend"
	v2 "kratos-blog/api/v1/tag"
	"kratos-blog/app/blog/internal/conf"
	"kratos-blog/app/blog/internal/data"
	"kratos-blog/app/blog/internal/service"
	"kratos-blog/pkg/upload"
)

var (
	blackList = []string{
		"/api/searchBlog",
	}
	whiteList = []string{
		"/util/upload",
		"/util/getBingPhoto",
	}
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cf *conf.Bootstrap, blog *service.BlogService,
	friend *service.FriendService, tag *service.TagService,
	f *data.FilterRepo, logger log.Logger) *http.Server {
	c := cf.Server
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.Filter(
			f.NewFilter(whiteList, blackList),
			f.DeleteCache()),
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
	// 注册工具类接口
	router := gin.Default()
	NewGinRouter(cf, router)
	srv.HandlePrefix("/util", router)

	v1.RegisterBlogHTTPServer(srv, blog)
	v2.RegisterTagHTTPServer(srv, tag)
	v3.RegisterFriendHTTPServer(srv, friend)
	return srv
}

func NewGinRouter(c *conf.Bootstrap, router *gin.Engine) {
	po := upload.UploadRepo{
		UploadPath: c.Upload.Path,
		Domain:     c.Upload.Domain,
		MaxSize:    c.Upload.Maxsize,
		Url:        c.Upload.Uri,
	}
	utilGroup := router.Group("/util")
	utilGroup.POST("/upload", po.GinUploadImg)
	utilGroup.GET("/getBingPhoto", po.GetBingPhoto)
}
