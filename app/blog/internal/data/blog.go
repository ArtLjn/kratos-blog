package data

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-blog/api/v1/blog"
	"kratos-blog/api/v1/user"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/pkg/vo"
	"strconv"
	"time"
)

const (
	Token      string = "token"
	AdminNotes string = "adminNotes"
	Notes      string = "notes"
	TableName  string = "BlogVisits"
	Comment    string = "comment"
	Appear     string = "appear"
)

var (
	CTX = context.Background()
)

type blogRepo struct {
	data *Data
	log  *log.Helper
}

func NewBlogRepo(data *Data, logger log.Logger) biz.BlogRepo {
	return &blogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type Blog struct {
	ID         int    `json:"id" gorm:"primary_key;auto_increment"`
	Title      string `json:"title"`
	Preface    string `json:"preface"`
	Photo      string `json:"photo"`
	Tag        string `json:"tag"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Visits     uint64 `json:"visits"`
	Content    string `json:"content"`
	Appear     bool   `json:"appear"`
	Comment    bool   `json:"comment"`
}

func (b Blog) TableName() string {
	return "person_table"
}

// CreateBlog :dev
func (r *blogRepo) CreateBlog(ctx context.Context, request *blog.CreateBlogRequest) (string, error) {
	var blog Blog
	re := r.createBlogFromRequest(request)
	blog = re()
	if err := r.data.db.Create(&blog).Error; err != nil {
		r.log.Info(err)
		return vo.INSERT_ERROR, err
	}
	return vo.INSERT_SUCCESS, nil
}

func (r *blogRepo) UpdateBlog(ctx context.Context, request *blog.UpdateBlogRequest) (string, error) {
	var blog Blog
	f, _ := json.Marshal(&request.Data)
	if err := json.Unmarshal(f, &blog); err != nil {
		panic(err)
	}
	blog.UpdateTime = time.Now().Format("2006-01-02")
	if err := r.data.db.Model(&blog).Where("id = ?", request.Id).Updates(blog).Error; err != nil {
		r.log.Log(log.LevelError, "Error updating blog:", err)
		return vo.UPDATE_FAIL, err
	}
	return vo.UPDATE_SUCCESS, nil
}

func (r *blogRepo) DeleteBlog(ctx context.Context, request *blog.DeleteBlogRequest) (string, error) {
	cacheKey, _ := r.data.rdb.Get(context.Background(), "privateKey").Result()
	if request.Key == cacheKey {
		if err := r.data.db.Where("id = ?", request.Id).Delete(&Blog{}).Error; err != nil {
			r.log.Log(log.LevelError, err)
			return vo.DELETE_ERROR, err
		}
	} else {
		return vo.KEY_ERROR, errors.New(vo.KEY_ERROR)
	}
	return vo.DELETE_SUCCESS, nil
}

// UpdateIndividualFields :dev Whether comments are allowed on the blog
func (r *blogRepo) UpdateIndividualFields(ctx context.Context, request *blog.UpdateIndividualFieldsRequest) (string, error) {
	var condName string
	switch request.Raw {
	case 0:
		condName = Comment
	case 1:
		condName = Appear
	}
	if err := r.data.pf.UpdateFunc(Blog{}, nil, map[string]interface{}{condName: request.Status}, true); err != nil {
		err := errors.New(vo.UPDATE_FAIL)
		r.log.Log(log.LevelError, err)
		return vo.UPDATE_FAIL, err
	}
	return vo.UPDATE_SUCCESS, nil
}

// createBlogFromRequest :dev create a blog record
func (r *blogRepo) createBlogFromRequest(request *blog.CreateBlogRequest) func() Blog {
	return func() Blog {
		var blog Blog
		f, _ := json.Marshal(&request.Data)
		if err := json.Unmarshal(f, &blog); err != nil {
			panic(err)
		}
		blog.CreateTime = time.Now().Format("2006-01-02")
		blog.UpdateTime = time.Now().Format("2006-01-02")
		blog.Comment = true
		blog.Visits = 0
		blog.Appear = true
		return blog
	}
}

// GetByTagName :dev search blog posts based on tags
func (r *blogRepo) GetByTagName(ctx context.Context, request *blog.GetBlogRequest) (string, []*blog.BlogData, error) {
	role := r.queryUserMsg(ctx).GetRole().CheckPermission()
	r.log.Log(log.LevelInfo, role)
	var (
		blogs []*blog.BlogData
		err   error
	)
	if role {
		admin := AdminRoleFactory{r: r, req: request}
		blogs, err = admin.QueryTag()
	} else {
		v := UserOrVisitFactory{r: r, req: request}
		blogs, err = v.QueryTag()
	}
	if err != nil {
		return err.Error(), nil, err
	}
	return vo.QUERY_SUCCESS, blogs, nil
}

// queryUserMsg :dev query user information
func (r *blogRepo) queryUserMsg(ctx context.Context) *RolePermission {
	req, ok := http.RequestFromServerContext(ctx)
	if !ok {
		panic(errors.New("ctx false"))
	}
	token := req.Header.Get(Token)

	grantVisitorRole := func() *RolePermission {
		visit := make([]string, 4)
		visit = append(visit, Visitor)
		return &RolePermission{
			u: &user.GetUserReply{
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
	username, _ := r.data.rdb.Get(context.Background(), token).Result()
	r.log.Info("username:", username)
	// call grpc to query the user
	res, err := r.data.uc.GetUser(context.Background(), &user.GetUserRequest{
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
	return &RolePermission{u: res}
}

// ListBlog :dev query all blog posts based on permissions
func (r *blogRepo) ListBlog(ctx context.Context, request *blog.ListBlogRequest) (string, []*blog.BlogData, error) {
	role := r.queryUserMsg(ctx).GetRole().CheckPermission()
	var (
		blogs []*blog.BlogData
		err   error
	)
	if role {
		admin := AdminRoleFactory{r: r, reb: request}
		blogs, err = admin.QueryListBlog()
	} else {
		v := UserOrVisitFactory{r: r, reb: request}
		blogs, err = v.QueryListBlog()
	}
	if err != nil {
		return vo.QUERY_FAIL, blogs, nil
	}
	return vo.QUERY_SUCCESS, blogs, nil
}

// QueryBlogById :dev more blog post ID query blog posts
func (r *blogRepo) QueryBlogById(ctx context.Context, request *blog.GetBlogIDRequest) (msg string, da blog.BlogData, e error) {
	role := r.queryUserMsg(ctx).GetRole().CheckPermission()
	r.log.Log(log.LevelInfo, role)
	var b Blog
	if err := r.data.db.Where("id = ?", request.Id).First(&b).Error; err != nil {
		r.log.Errorf("query error %s", err)
		return vo.QUERY_FAIL, blog.BlogData{}, err
	}
	if err := r.data.pf.ParseJSONToStruct(b, &da); err != nil {
		return vo.JSON_ERROR, blog.BlogData{}, err
	}

	if !b.Appear && !role {
		return vo.FORBIDDEN_ACCESS, blog.BlogData{}, nil
	}
	strID := strconv.Itoa(int(request.Id))
	if r.hasHashField(TableName, strID) {
		currentCount := r.getHashField(TableName, strID)
		r.setHashField(TableName, strID, currentCount+1)
	} else {
		r.setHashField(TableName, strID, 1)
	}
	return vo.QUERY_SUCCESS, da, nil
}

// UpdateBlogVisitsCount :dev update the number of blog post visits
func (r *blogRepo) UpdateBlogVisitsCount() {
	var blogs []Blog
	if err := r.data.db.Table(Blog{}.TableName()).Find(&blogs).Error; err != nil {
		panic(err)
	}

	// traverse the list
	for _, blog := range blogs {
		visitCount := blog.Visits
		var res error
		cacheCount := r.getHashField(TableName, strconv.Itoa(blog.ID))
		if !r.hasHashField(TableName, strconv.Itoa(blog.ID)) {
			res = r.data.pf.UpdateFunc(Blog{}, map[string]interface{}{"id": blog.ID}, map[string]interface{}{"visits": 0}, false)
		} else if cacheCount < visitCount {
			r.setHashField(TableName, strconv.Itoa(blog.ID), visitCount)
		} else {
			res = r.data.pf.UpdateFunc(Blog{}, map[string]interface{}{"id": blog.ID}, map[string]interface{}{"visits": cacheCount}, false)
		}
		if res != nil {
			r.log.Info(blog.ID, "update failed!")
		}
	}
}

// QueryBlogByTitle :dev query for matching blog posts based on the title
func (r *blogRepo) QueryBlogByTitle(ctx context.Context, request *blog.GetBlogByTitleRequest) (string, []*blog.BlogData, error) {
	role := r.queryUserMsg(ctx).GetRole().CheckPermission()
	var (
		data  []*blog.BlogData
		blogs []Blog
	)
	if role {
		keyword := "%" + request.Title + "%"
		if err := r.data.db.Where("title LIKE ?", keyword).Find(&blogs).Error; err != nil {
			r.log.Log(log.LevelError, err)
			return vo.QUERY_FAIL, nil, errors.New(vo.QUERY_FAIL)
		}
		r.data.pf.ParseJSONToStruct(blogs, &data)
		return vo.QUERY_SUCCESS, data, nil
	}
	return vo.PERMISSION_ERROR, nil, errors.New(vo.PERMISSION_ERROR)
}

func (r *blogRepo) UpdateOnly(ctx context.Context, request *blog.UpdateOnlyRequest) *blog.UpdateOnlyReply {
	role := r.queryUserMsg(ctx).GetRole().CheckPermission()
	if role {
		var condName string
		switch request.Raw {
		case 0:
			condName = Comment
		case 1:
			condName = Appear
		}
		if err := r.data.pf.UpdateFunc(Blog{}, map[string]interface{}{"id": request.Id},
			map[string]interface{}{condName: request.Res}, false); err != nil {
			return &blog.UpdateOnlyReply{Common: &blog.CommonReply{Code: 400, Result: vo.UPDATE_FAIL}}
		}
		return &blog.UpdateOnlyReply{Common: &blog.CommonReply{Code: 200, Result: vo.UPDATE_SUCCESS}}
	}
	return &blog.UpdateOnlyReply{
		Common: &blog.CommonReply{
			Code:   401,
			Result: vo.PERMISSION_ERROR,
		},
	}
}

// ***************** Redis Util *********************** //

func (r *blogRepo) setHashField(hashKey, field string, val interface{}) {
	err := r.data.rdb.HSet(CTX, hashKey, field, val).Err()
	if err != nil {
		r.log.Log(log.LevelError, err)
	}
}

func (r *blogRepo) hasHashField(hashKey, field string) bool {
	val, _ := r.data.rdb.HExists(CTX, hashKey, field).Result()
	return val
}

func (r *blogRepo) getHashField(hashKey, field string) uint64 {
	val, err := r.data.rdb.HGet(CTX, hashKey, field).Uint64()
	if err != nil {
		r.log.Log(log.LevelError, err)
	}
	return val
}

func (r *blogRepo) parseList(o []*blog.BlogData) []interface{} {
	m, _ := json.Marshal(&o)
	var newList []interface{}
	if err := json.Unmarshal(m, &newList); err != nil {
		panic(err)
	}
	return newList
}

func (r *blogRepo) restoreList(key string) []*blog.BlogData {
	var newList []*blog.BlogData
	serializedValues, err := r.data.rdb.LRange(CTX, key, 0, -1).Result()
	if err != nil {
		r.log.Log(log.LevelError, "Failed to retrieve values from Redis list:", err)
		panic(err)
	}
	for _, serializedValue := range serializedValues {
		var value blog.BlogData
		if e := json.Unmarshal([]byte(serializedValue), &value); e != nil {
			r.log.Log(log.LevelError, "Failed to deserialize value from JSON:", e)
			continue
		}
		newList = append(newList, &value)
	}
	return newList
}

func (r *blogRepo) setCacheList(key string, o []interface{}) {
	for _, value := range o {
		serializedValue, err := json.Marshal(value)
		if err != nil {
			r.log.Log(log.LevelError, err)
			continue
		}
		i, e := r.data.rdb.RPush(CTX, key, serializedValue).Result()
		if e != nil {
			r.log.Log(log.LevelError, "Failed to push serialized value to Redis list:", e)
		}
		r.log.Log(log.LevelInfo, "Successfully pushed value to Redis list", i)
	}
}
