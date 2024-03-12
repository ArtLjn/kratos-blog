package data

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-blog/app/gateway/internal/conf"
	"kratos-blog/pkg/server"
	"kratos-blog/pkg/vo"
	h "net/http"
	"strings"
	"sync"
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
)

type FilterInterface interface {
	DeleteCache() http.FilterFunc
	FilterPermission(whiteList, blackList []string) http.FilterFunc
	AllowDomainsMiddleware(cf conf.Domain) http.FilterFunc
}

type FilterRepo struct {
	data *Data
	log  *log.Helper
	mu   sync.Mutex
	wg   sync.WaitGroup
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

func (f *FilterRepo) AllowDomainsMiddleware(cf conf.Domain) http.FilterFunc {
	return func(handler h.Handler) h.Handler {
		return h.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			origin := req.Referer()
			// 检查请求的 Origin 是否在允许的域名列表中

			if !cf.Open {
				handler.ServeHTTP(w, req)
				return
			} else {
				f.mu.Lock()
				var allow bool
				for _, domain := range cf.GetOrigin() {
					if strings.HasPrefix(origin, domain) {
						allow = true
						break
					}
				}
				f.mu.Unlock()
				if allow {
					handler.ServeHTTP(w, req)
					return
				}
				WritePermissionError(w)
				return
			}
		})
	}
}

func WritePermissionError(w http.ResponseWriter) {
	w.WriteHeader(401)
	rep := map[string]interface{}{
		"code":   401,
		"result": vo.PERMISSION_ERROR,
	}
	byteRep, _ := json.Marshal(&rep)
	w.Write(byteRep)
}
