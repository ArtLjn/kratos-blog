package data

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/api/v1/blog"
	"kratos-blog/pkg/server"
	"kratos-blog/pkg/vo"
)

// BlogRoleFactory 接口定义了查询博客的方法，根据用户角色不同实现不同的逻辑
type BlogRoleFactory interface {
	QueryTag() ([]*blog.BlogData, error)
	QueryListBlog() ([]*blog.BlogData, error)
}

// AdminRole 实现管理员的查询逻辑
type AdminRole struct {
	r   *blogRepo
	req *blog.GetBlogRequest
	reb *blog.ListBlogRequest
}

// VipRole 实现 VIP 用户的查询逻辑
type VipRole struct {
	r   *blogRepo
	req *blog.GetBlogRequest
	reb *blog.ListBlogRequest
}

// VisitorRole 实现普通用户的查询逻辑
type VisitorRole struct {
	r   *blogRepo
	req *blog.GetBlogRequest
	reb *blog.ListBlogRequest
}

// QueryTag 实现管理员的按标签查询逻辑
func (a *AdminRole) QueryTag() ([]*blog.BlogData, error) {
	return a.r.queryByCondition(map[string]interface{}{"tag": a.req.Tag})
}

// QueryTag 实现 VIP 用户的按标签查询逻辑
func (v *VipRole) QueryTag() ([]*blog.BlogData, error) {
	return v.r.queryByCondition(map[string]interface{}{"tag": v.req.Tag})
}

// QueryTag 实现普通用户的按标签查询逻辑，过滤掉未开放的博客
func (v *VisitorRole) QueryTag() ([]*blog.BlogData, error) {
	return v.r.queryByCondition(map[string]interface{}{"tag": v.req.Tag, "appear": true})
}

// QueryListBlog 实现管理员的博客列表查询逻辑
func (a *AdminRole) QueryListBlog() ([]*blog.BlogData, error) {
	return a.r.queryListByCondition(nil) // 管理员无条件限制
}

// QueryListBlog 实现 VIP 用户的博客列表查询逻辑
func (v *VipRole) QueryListBlog() ([]*blog.BlogData, error) {
	return v.r.queryListByCondition(nil) // VIP 用户无条件限制
}

// QueryListBlog 实现普通用户的博客列表查询逻辑，过滤掉未开放的博客
func (v *VisitorRole) QueryListBlog() ([]*blog.BlogData, error) {
	return v.r.queryListByCondition(map[string]interface{}{"appear": true})
}

// 通用查询方法，从 AdminRole 中提取出来放到 blogRepo 中
func (r *blogRepo) queryByCondition(condition map[string]interface{}) ([]*blog.BlogData, error) {
	var blogs []*blog.BlogData
	res, err := r.data.pf.QueryFunc(Blog{}, condition, true)
	if err != nil || res == nil {
		r.log.Info(vo.QUERY_EMPTY)
		return nil, errors.New(vo.QUERY_EMPTY)
	}
	if err = r.data.pf.ParseJSONToStruct(res, &blogs); err != nil {
		return nil, fmt.Errorf(vo.QUERY_FAIL)
	}
	if len(blogs) == 0 {
		return nil, errors.New(vo.LIST_EMPTY)
	}
	return blogs, nil
}

// 通用列表查询方法，从 AdminRole 中提取出来放到 blogRepo 中
func (r *blogRepo) queryListByCondition(condition map[string]interface{}) ([]*blog.BlogData, error) {
	var blogs []*blog.BlogData
	var b []Blog
	r.data.db.Model(Blog{}).Order("createTime desc").Where(condition).Find(&b)
	if err := r.data.pf.ParseJSONToStruct(b, &blogs); err != nil {
		return nil, fmt.Errorf(vo.QUERY_FAIL)
	}
	if len(blogs) == 0 {
		r.log.Info(vo.QUERY_EMPTY)
		return nil, errors.New(vo.QUERY_EMPTY)
	}
	return blogs, nil
}

// NewBlogRoleFactory 工厂方法，返回不同角色的实现
func NewBlogRoleFactory(role int32, r *blogRepo, req *blog.GetBlogRequest, reb *blog.ListBlogRequest) BlogRoleFactory {
	switch role {
	case server.SystemManager: // 系统管理员
		return &AdminRole{r: r, req: req, reb: reb}
	case server.Vip: // VIP 用户
		return &VipRole{r: r, req: req, reb: reb}
	case server.VisitorOrUser: // 普通用户/访客
		return &VisitorRole{r: r, req: req, reb: reb}
	default:
		return &VisitorRole{r: r, req: req, reb: reb}
	}
}

// GetByTagName 博客仓库方法，根据角色进行查询
func (r *blogRepo) GetByTagName(ctx context.Context, request *blog.GetBlogRequest) (string, []*blog.BlogData, error) {
	roleFactory := NewBlogRoleFactory(request.GetPermission().Role, r, request, nil)
	blogs, err := roleFactory.QueryTag()
	if err != nil {
		return vo.QUERY_FAIL, nil, err
	}
	return vo.QUERY_SUCCESS, blogs, nil
}

func (r *blogRepo) ListBlog(ctx context.Context, request *blog.ListBlogRequest) (string, []*blog.BlogData, error) {
	roleFactory := NewBlogRoleFactory(request.GetPermission().Role, r, nil, request)
	blogs, err := roleFactory.QueryListBlog()
	if err != nil {
		return vo.QUERY_FAIL, nil, err
	}
	return vo.QUERY_SUCCESS, blogs, nil
}

// QueryBlogById :dev more blog post ID query blog posts
func (r *blogRepo) QueryBlogById(ctx context.Context, request *blog.GetBlogIDRequest) (msg string,
	da *blog.BlogData, e error) {
	var b Blog
	if err := r.data.db.Where("id = ?", request.Id).First(&b).Error; err != nil {
		r.log.Errorf("query error %s", err)
		return vo.QUERY_FAIL, &blog.BlogData{}, fmt.Errorf(vo.QUERY_FAIL)
	}
	if err := r.data.pf.ParseJSONToStruct(b, &da); err != nil {
		return vo.JSON_ERROR, &blog.BlogData{}, fmt.Errorf(vo.JSON_ERROR)
	}
	if !b.Appear && request.GetPermission().Role == server.VisitorOrUser {
		return vo.FORBIDDEN_ACCESS, &blog.BlogData{}, fmt.Errorf(vo.FORBIDDEN_ACCESS)
	}
	return vo.QUERY_SUCCESS, da, nil
}
