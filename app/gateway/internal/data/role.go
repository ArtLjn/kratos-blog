package data

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/redis/go-redis/v9"
	"kratos-blog/api/user"
	"kratos-blog/pkg/auth"
	"kratos-blog/pkg/server"
	"kratos-blog/pkg/vo"
	h "net/http"
	"strings"
)

const Token = server.Token

var CTX = context.Background()

// RoleSwitch
const (
	VIP     = server.VIP
	User    = server.User
	Visitor = server.Visitor
)

type Role struct {
	rdb *redis.Client
	uc  user.UserClient
	log *log.Helper
}

func NewRole(rdb *redis.Client, uc user.UserClient, logger log.Logger) *Role {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Role{rdb: rdb, uc: uc, log: l}
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

	adminCacheToken, _ := r.rdb.Get(CTX, server.AdminToken).Result()
	if len(adminCacheToken) != 0 {
		if adminCacheToken == token {
			return &Permission{
				U: &user.GetUserReply{
					Common: &user.CommonReply{
						Code:   200,
						Result: vo.QUERY_SUCCESS,
					},
					Data: []string{server.SystemAdmin},
				},
				Role: server.SystemAdmin,
			}
		}
	}
	// The token is empty,granting the visitor permissions
	if len(token) == 0 {
		return grantVisitorRole()
	}

	username := auth.GetLoginName(token)
	r.log.Info("username:", username)
	// call grpc to query the user
	res, err := r.uc.GetUser(context.Background(), &user.GetUserRequest{
		Name: username,
	})
	if res == nil || len(res.Data[4]) == 0 {
		return grantVisitorRole()
	}
	if err != nil {
		panic(err)
	}
	if res.Common.Code != 200 {
		panic(errors.New(res.Common.Result))
	}
	return &Permission{U: res, Role: res.Data[4]}
}

/**
 * @dev FilterPermission
 */

func (r *Role) FilterPermission(whiteList []string) http.FilterFunc {
	return func(handler h.Handler) h.Handler {
		return h.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			handler.ServeHTTP(w, req)
			return
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

func (r *Role) VerifyManagerToken(token string) bool {
	cacheToken, err := r.rdb.Get(CTX, server.AdminToken).Result()
	if err != nil {
		return false
	}
	return token == cacheToken
}
