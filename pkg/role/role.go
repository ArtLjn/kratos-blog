package role

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/redis/go-redis/v9"
	"kratos-blog/api/v1/user"
	"kratos-blog/pkg/vo"
	h "net/http"
	"strings"
)

const (
	Token   = "token"
	Admin   = "admin"
	User    = "user"
	Visitor = "visitor"
)

type Role struct {
	rdb *redis.Client
	uc  user.UserClient
	log *log.Helper
}

func NewRole(rdb *redis.Client, uc user.UserClient, logger *log.Helper) *Role {
	return &Role{rdb: rdb, uc: uc, log: logger}
}

type PermissionStrategy interface {
	CheckPermission() bool
	GetRole() PermissionStrategy
}
type Permission struct {
	U    *user.GetUserReply
	Role string
}

// QueryUserMsg :dev query user information
func (r *Role) QueryUserMsg(ctx context.Context) *Permission {
	req, ok := http.RequestFromServerContext(ctx)
	if !ok {
		panic(errors.New("ctx false"))
	}
	token := req.Header.Get(Token)

	grantVisitorRole := func() *Permission {
		visit := make([]string, 4)
		visit = append(visit, Visitor)
		return &Permission{
			U: &user.GetUserReply{
				Common: &user.CommonReply{
					Code:   200,
					Result: vo.QUERY_SUCCESS,
				},
				Data: visit,
			},
		}
	}
	// The token is empty,granting the visitor permissions
	if token == "" {
		return grantVisitorRole()
	}
	username, _ := r.rdb.Get(context.Background(), token).Result()
	r.log.Info("username:", username)
	// call grpc to query the user
	res, err := r.uc.GetUser(context.Background(), &user.GetUserRequest{
		Name: username,
	})
	if res.Data[4] == "" {
		return grantVisitorRole()
	}
	if err != nil {
		panic(err)
	}
	if res.Common.Code != 200 {
		panic(errors.New(res.Common.Result))
	}
	return &Permission{U: res}
}

func (r *Permission) CheckPermission() bool {
	switch r.Role {
	case Admin:
		return true
	case User, Visitor:
		return false
	}
	return false
}

func (r *Permission) GetRole() PermissionStrategy {
	r.Role = r.U.Data[4]
	return r
}

func (r *Role) FilterPermission(whiteList, blackList []string) http.FilterFunc {
	return func(handler h.Handler) h.Handler {
		return h.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			if !r.IsPathInList(req.URL.Path, blackList) {
				if req.Method == h.MethodGet || r.IsPathInList(req.URL.Path, whiteList) {
					handler.ServeHTTP(w, req)
					return
				}
			}

			token := req.Header.Get("token")
			if token == "" {
				r.WritePermissionError(w)
				return
			}

			username, _ := r.rdb.Get(context.Background(), token).Result()
			res, err := r.uc.GetUser(context.Background(), &user.GetUserRequest{
				Name: username,
			})
			if err != nil {
				panic(err)
			}

			m := &Permission{U: res}
			if !m.GetRole().CheckPermission() {
				r.WritePermissionError(w)
				return
			}
			handler.ServeHTTP(w, req)
		})
	}
}

func (r *Role) IsPathInList(path string, List []string) bool {
	for _, wlPath := range List {
		if strings.HasPrefix(path, wlPath) {
			return true
		}
	}
	return false
}

func (r *Role) WritePermissionError(w http.ResponseWriter) {
	w.WriteHeader(401)
	rep := map[string]interface{}{
		"code":   401,
		"result": vo.PERMISSION_ERROR,
	}
	byteRep, _ := json.Marshal(&rep)
	w.Write(byteRep)
}
