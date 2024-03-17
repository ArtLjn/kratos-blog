package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"kratos-blog/app/gateway/internal/conf"
	"kratos-blog/pkg/upload"
	"net/http"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGinRouter)

// NewGinRouter :dev 工具类方法接口
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
	utilGroup.StaticFS("/img", http.Dir("pkg/upload/file/"))

}
