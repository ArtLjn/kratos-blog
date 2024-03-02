package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	h "net/http"
)

type FilterRepo struct {
	data *Data
	log  *log.Helper
}

// 当对文章进行修改时,需要清除redis缓存的接口
var (
	pathList = []string{
		"/api/addBlog",
		"/api/updateBlog",
		"/api/updateIndividual",
		"/api/deleteBlog",
		"/api/updateOnly",
		"/api/cacheBlog",
	}
)

func NewFilterRepo(data *Data, logger log.Logger) *FilterRepo {
	return &FilterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// NewFilter :dev 权限认证
func (f *FilterRepo) NewFilter(whiteList, blackList []string) http.FilterFunc {
	return f.data.role.FilterPermission(whiteList, blackList)
}

// DeleteCache :dev 清除redis缓存
func (f *FilterRepo) DeleteCache() http.FilterFunc {
	return func(handler h.Handler) h.Handler {
		return h.HandlerFunc(func(writer h.ResponseWriter, request *h.Request) {
			method := request.Method
			if f.data.role.IsPathInList(request.URL.Path, pathList) {
				if method == h.MethodPost || method == h.MethodDelete || method == h.MethodPut {
					f.log.Info(f.data.rdb.Del(CTX, AdminNotes).Result())
					f.log.Info(f.data.rdb.Del(CTX, Notes).Result())
				}
			}
			handler.ServeHTTP(writer, request)
		})
	}
}
