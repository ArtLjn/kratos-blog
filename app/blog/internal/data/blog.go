package data

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/api/v1/blog"
	"kratos-blog/app/blog/internal/biz"
	"kratos-blog/pkg/vo"
	"time"
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
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Visits     string `json:"visits"`
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
	cacheKey, _ := r.data.rdb.Get(context.Background(), "privateKey").Result() //TODO 设置为 admin密码验证
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
		f, _ := json.Marshal(&request)
		if err := json.Unmarshal(f, &blog); err != nil {
			panic(err)
		}
		blog.CreateTime = time.Now().Format("2006-01-02")
		blog.UpdateTime = time.Now().Format("2006-01-02")
		blog.Comment = true
		blog.Visits = "0"
		return blog
	}
}

//func (r *blogRepo) GetByTagName(ctx context.Context, request *v1.GetBlogRequest) (string, []string, error) {
//
//}
