package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-blog/pkg/server"
	h "net/http"
)

var (
	// PathList 当对文章进行修改时,需要清除redis缓存的接口
	PathList = []string{
		"/api/addBlog",
		"/api/updateBlog",
		"/api/updateIndividual",
		"/api/deleteBlog",
		"/api/updateOnly",
		"/api/cacheBlog",
	}
	// BlackList GET请求需要传递Token
	BlackList = []string{
		"/api/searchBlog",
	}
	// WhiteList 不需要传递Token进行验证
	WhiteList = []string{
		"/util/upload",
		"/util/getBingPhoto",
	}
)

type FilterInterface interface {
	DeleteCache() http.FilterFunc
	FilterPermission(whiteList, blackList []string) http.FilterFunc
}

type FilterRepo struct {
	data *Data
	log  *log.Helper
}

func NewFilterRepo(data *Data, l log.Logger) *FilterRepo {
	lg := log.NewHelper(log.With(l, "module", "filter"))
	return &FilterRepo{data: data, log: lg}
}

// DeleteCache :dev 清除redis缓存
func (f *FilterRepo) DeleteCache() http.FilterFunc {
	return func(handler h.Handler) h.Handler {
		return h.HandlerFunc(func(writer h.ResponseWriter, request *h.Request) {
			method := request.Method
			if f.data.role.IsPathInList(request.URL.Path, PathList) {
				if method == h.MethodPost || method == h.MethodDelete || method == h.MethodPut {
					f.log.Info(f.data.rdb.Del(CTX, server.AdminNotes).Result())
					f.log.Info(f.data.rdb.Del(CTX, server.Notes).Result())
				}
			}
			handler.ServeHTTP(writer, request)
		})
	}
}

func (f *FilterRepo) FilterPermission(whiteList, blackList []string) http.FilterFunc {
	return f.data.role.FilterPermission(whiteList, blackList)
}
