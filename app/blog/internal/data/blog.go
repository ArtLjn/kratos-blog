package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// UpdateAllCommentStatus :dev Whether comments are allowed on the blog
func (r *blogRepo) UpdateAllCommentStatus(ctx context.Context, request *blog.UpdateAllCommentStatusRequest) (string, error) {
	if err := r.data.db.Model(&Blog{}).Update("comment", request.Status).Error; err != nil {
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
	role := r.queryUserMsg(ctx).Data[4]
	r.log.Log(log.LevelInfo, role)
	var blogs []*blog.BlogData
	switch role {
	case Admin:
		blogs = r.queryAllBlog([]interface{}{"tag", request.Tag})
	case User, Visitor:
		blogs = r.queryBlog(request, true)
	}
	return vo.QUERY_SUCCESS, blogs, nil
}

// queryBlog :dev query blog posts based on tags and permissions
func (r *blogRepo) queryBlog(req *blog.GetBlogRequest, appear bool) []*blog.BlogData {
	var (
		blogs    []Blog
		blogData []*blog.BlogData
	)
	if err := r.data.db.Where("tag = ?", req.Tag).Where("appear = ?", appear).
		Find(&blogs).Error; err != nil {
		panic(err)
	}
	d, _ := json.Marshal(&blogs)
	if err := json.Unmarshal(d, &blogData); err != nil {
		panic(err)
	}
	return blogData
}

// queryAllBlog :dev search all blog posts based on tags
func (r *blogRepo) queryAllBlog(cond []interface{}) []*blog.BlogData {
	var (
		blogs    []Blog
		blogData []*blog.BlogData
		rule     = "create_time DESC"
	)
	if len(cond) != 0 {
		if len(cond)%2 != 0 {
			cond = append(cond, cond[len(cond)])
		}
		// sort in descending order based on criteria
		str := fmt.Sprintf("%s = ?", cond[0])
		if err := r.data.db.Order(rule).Where(str, cond[1]).
			Find(&blogs).Error; err != nil {
			panic(err)
		}
	} else {
		if err := r.data.db.Order(rule).
			Find(&blogs).Error; err != nil {
			panic(err)
		}
	}
	d, _ := json.Marshal(&blogs)
	if err := json.Unmarshal(d, &blogData); err != nil {
		panic(err)
	}
	return blogData
}

// queryUserMsg :dev query user information
func (r *blogRepo) queryUserMsg(ctx context.Context) *user.GetUserReply {
	req, ok := http.RequestFromServerContext(ctx)
	if !ok {
		panic(errors.New("ctx false"))
	}
	token := req.Header.Get(Token)
	// The token is empty,granting the visitor permissions
	if token == "" {
		visit := make([]string, 4)
		visit = append(visit, Visitor)
		return &user.GetUserReply{
			Common: &user.CommonReply{
				Code:   200,
				Result: vo.QUERY_SUCCESS,
			},
			Data: visit,
		}
	}
	username, _ := r.data.rdb.Get(context.Background(), token).Result()
	r.log.Info("username:", username)
	// call grpc to query the user
	res, err := r.data.uc.GetUser(context.Background(), &user.GetUserRequest{
		Name: username,
	})
	if err != nil {
		panic(err)
	}
	if res.Common.Code != 200 {
		panic(errors.New(res.Common.Result))
	}
	return res
}

// ListBlog :dev query all blog posts based on permissions
func (r *blogRepo) ListBlog(ctx context.Context, request *blog.ListBlogRequest) (string, []*blog.BlogData, error) {
	role := r.queryUserMsg(ctx).Data[4]
	r.log.Log(log.LevelInfo, role)
	var data []*blog.BlogData
	switch role {
	case Admin:
		if exists, _ := r.data.rdb.Exists(CTX, AdminNotes).Result(); exists == 1 {
			return vo.QUERY_SUCCESS, r.restoreList(AdminNotes), nil
		}
		data = r.queryAllBlog(nil)
		r.setCacheList(AdminNotes, r.parseList(data))
	case User, Visitor:
		if exists, _ := r.data.rdb.Exists(CTX, Notes).Result(); exists == 1 {
			return vo.QUERY_SUCCESS, r.restoreList(Notes), nil
		}
		data = r.queryAllBlog([]interface{}{"appear", true})
		r.setCacheList(Notes, r.parseList(data))
	}

	return vo.QUERY_SUCCESS, data, nil
}

// QueryBlogById :dev more blog post ID query blog posts
func (r *blogRepo) QueryBlogById(ctx context.Context, request *blog.GetBlogIDRequest) (msg string, da blog.BlogData, e error) {
	role := r.queryUserMsg(ctx).Data[4]
	r.log.Log(log.LevelInfo, role)
	var b Blog
	if err := r.data.db.Where("id = ?", request.Id).First(&b).Error; err != nil {
		return vo.QUERY_FAIL, blog.BlogData{}, err
	}
	f, _ := json.Marshal(&b)
	if err := json.Unmarshal(f, &da); err != nil {
		return vo.JSON_ERROR, blog.BlogData{}, err
	}
	if b.Appear == false {
		if role == User || role == Visitor {
			return vo.FORBIDDEN_ACCESS, blog.BlogData{}, nil
		}
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
		if err := json.Unmarshal([]byte(serializedValue), &value); err != nil {
			r.log.Log(log.LevelError, "Failed to deserialize value from JSON:", err)
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
		i, err := r.data.rdb.RPush(CTX, key, serializedValue).Result()
		if err != nil {
			r.log.Log(log.LevelError, "Failed to push serialized value to Redis list:", err)
		}
		r.log.Log(log.LevelInfo, "Successfully pushed value to Redis list", i)
	}
}
