package data

import (
	"errors"
	"kratos-blog/api/v1/blog"
	"kratos-blog/api/v1/user"
	"kratos-blog/pkg/vo"
)

const (
	Admin   = "admin"
	User    = "user"
	Visitor = "visitor"
)

type RolePermissionStrategy interface {
	CheckPermission() bool
	GetRole() RolePermissionStrategy
}
type BlogFactory interface {
	QueryTag() ([]*blog.BlogData, error)
	QueryListBlog() ([]*blog.BlogData, error)
}

type RolePermission struct {
	u    *user.GetUserReply
	role string
}
type AdminRoleFactory struct {
	r   *blogRepo
	req *blog.GetBlogRequest
	reb *blog.ListBlogRequest
}
type UserOrVisitFactory struct {
	r   *blogRepo
	req *blog.GetBlogRequest
	reb *blog.ListBlogRequest
}

func (r *RolePermission) CheckPermission() bool {
	switch r.role {
	case Admin:
		return true
	case User, Visitor:
		return false
	}
	return false
}

func (r *RolePermission) GetRole() RolePermissionStrategy {
	r.role = r.u.Data[4]
	return r
}

func (a *AdminRoleFactory) QueryTag() ([]*blog.BlogData, error) {
	var blogs []*blog.BlogData
	res, err := a.r.data.pf.QueryFunc(Blog{}, map[string]interface{}{
		"tag": a.req.Tag,
	}, true)
	if err != nil || res == nil {
		a.r.log.Info(vo.QUERY_EMPTY)
		return nil, errors.New(vo.QUERY_EMPTY)
	}
	a.r.data.pf.ParseJSONToStruct(res, &blogs)
	if len(blogs) == 0 {
		return nil, errors.New(vo.LIST_EMPTY)
	}
	return blogs, nil
}

func (a *UserOrVisitFactory) QueryTag() ([]*blog.BlogData, error) {
	var blogs []*blog.BlogData
	res, err := a.r.data.pf.QueryFunc(Blog{}, map[string]interface{}{
		"tag":    a.req.Tag,
		"appear": true,
	}, true)
	if err != nil || res == nil {
		a.r.log.Info(vo.QUERY_EMPTY)
		return nil, errors.New(vo.QUERY_EMPTY)
	}
	a.r.data.pf.ParseJSONToStruct(res, &blogs)
	if len(blogs) == 0 {
		return nil, errors.New(vo.LIST_EMPTY)
	}
	return blogs, nil
}

func (r *blogRepo) parse(m interface{}, blogs *[]*blog.BlogData) {
	var data []*blog.BlogData
	e := r.data.pf.ParseJSONToStruct(m, &data)
	if e != nil {
		panic(e)
	}
	*blogs = data
}

func (a *AdminRoleFactory) QueryListBlog() ([]*blog.BlogData, error) {
	var blogs []*blog.BlogData
	if exists, _ := a.r.data.rdb.Exists(CTX, AdminNotes).Result(); exists == 1 {
		return a.r.restoreList(AdminNotes), nil
	}
	res, err := a.r.data.pf.QueryFunc(Blog{}, nil, true)
	a.r.parse(res, &blogs)
	a.r.setCacheList(AdminNotes, a.r.parseList(blogs))
	if err != nil || res == nil {
		a.r.log.Info(vo.QUERY_EMPTY)
		return nil, errors.New(vo.QUERY_EMPTY)
	}
	return blogs, nil
}

func (a *UserOrVisitFactory) QueryListBlog() ([]*blog.BlogData, error) {
	var blogs []*blog.BlogData
	if exists, _ := a.r.data.rdb.Exists(CTX, Notes).Result(); exists == 1 {
		return a.r.restoreList(Notes), nil
	}
	res, err := a.r.data.pf.QueryFunc(Blog{}, map[string]interface{}{"appear": true}, true)
	a.r.parse(res, &blogs)
	a.r.setCacheList(Notes, a.r.parseList(blogs))
	if err != nil || res == nil {
		a.r.log.Info(vo.QUERY_EMPTY)
		return nil, errors.New(vo.QUERY_EMPTY)
	}
	return blogs, nil
}
