package data

import (
	"errors"
	"kratos-blog/api/v1/blog"
	"kratos-blog/pkg/vo"
)

type BlogFactory interface {
	QueryTag() ([]*blog.BlogData, error)
	QueryListBlog() ([]*blog.BlogData, error)
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
	var b []Blog

	if exists, _ := a.r.data.rdb.Exists(CTX, AdminNotes).Result(); exists == 1 {
		return a.r.restoreList(AdminNotes), nil
	}
	a.r.data.db.Model(Blog{}).Order("createTime desc").Find(&b)
	a.r.parse(b, &blogs)
	a.r.setCacheList(AdminNotes, a.r.parseList(blogs))
	if len(blogs) == 0 {
		a.r.log.Info(vo.QUERY_EMPTY)
		return nil, errors.New(vo.QUERY_EMPTY)
	}
	return blogs, nil
}

func (a *UserOrVisitFactory) QueryListBlog() ([]*blog.BlogData, error) {
	var blogs []*blog.BlogData
	var b []Blog
	if exists, _ := a.r.data.rdb.Exists(CTX, Notes).Result(); exists == 1 {
		return a.r.restoreList(Notes), nil
	}
	a.r.data.db.Model(Blog{}).Order("createTime desc").Where("appear = ?", true).Find(&b)
	a.r.parse(b, &blogs)
	a.r.setCacheList(Notes, a.r.parseList(blogs))
	if len(blogs) == 0 {
		a.r.log.Info(vo.QUERY_EMPTY)
		return nil, errors.New(vo.QUERY_EMPTY)
	}
	return blogs, nil
}
