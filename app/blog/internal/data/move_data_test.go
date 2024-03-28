package data

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"testing"
)

const (
	remoteSource = "root:123456@tcp(127.0.0.1:3306)/beifen"
	localSource  = "root:123456@tcp(127.0.0.1:3306)/test"
)

var (
	wg sync.Mutex
	mx sync.WaitGroup
)

func newDB(source string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type RemoteBlog struct {
	ID         int    `json:"id" gorm:"primary_key;auto_increment"`
	Title      string `json:"title"`
	Preface    string `json:"preface"`
	Photo      string `json:"photo"`
	Tag        string `json:"tag"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Comment    string `json:"comment"`
	Visits     string `json:"visits"`
	Content    string `json:"content"`
	Appear     string `json:"appear"`
}

func GetAllBlog() []RemoteBlog {
	db := newDB(remoteSource)
	var blogs []RemoteBlog
	if err := db.Table("person_table").Find(&blogs).Error; err != nil {
		panic(err)
	}
	return blogs
}

func UpdateLocalBlogData() {
	db := newDB(localSource)
	remoteBlog := GetAllBlog()
	mx.Add(1)
	go func() {
		defer mx.Done()
		for _, blog := range remoteBlog {
			wg.Lock()
			var newBlog Blog
			newBlog.ID = blog.ID
			newBlog.Title = blog.Title
			newBlog.Preface = blog.Preface
			newBlog.Photo = blog.Photo
			newBlog.Tag = blog.Tag
			newBlog.CreateTime = blog.CreateTime
			newBlog.UpdateTime = blog.UpdateTime
			newBlog.Visits = 0
			newBlog.Content = blog.Content
			newBlog.Comment = true
			newBlog.Appear = true
			if err := db.Table("person_table").Create(&newBlog).Error; err != nil {
				fmt.Println(err)
				continue
			}
			wg.Unlock()
		}
	}()
	mx.Wait()
}

type BlogPhoto struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Photo    string `json:"photo"`
	Date     string `json:"date"`
	Title    string `json:"title"`
	Position string `json:"position"`
}

type RemoteTag struct {
	ID      int    `json:"id" gorm:"primary_key;auto_increment"`
	TagName string `json:"tag_name" gorm:"column:tag_name"`
}

type LocalTag struct {
	ID      int    `json:"id" gorm:"primary_key;auto_increment"`
	TagName string `json:"tagName" gorm:"column:tagName"`
}

func GetAllTag() []RemoteTag {
	db := newDB(remoteSource)
	var tags []RemoteTag
	if err := db.Table("tag").Find(&tags).Error; err != nil {
		panic(err)
	}
	return tags
}

func UpdateTag() {
	db := newDB(localSource)
	remoteTag := GetAllTag()
	mx.Add(1)
	go func() {
		defer mx.Done()
		for _, tag := range remoteTag {
			wg.Lock()
			var t LocalTag
			t.ID = tag.ID
			t.TagName = tag.TagName
			if err := db.Table("tag").Create(&t).Error; err != nil {
				fmt.Println(err)
				continue
			}
			wg.Unlock()
		}
	}()
	mx.Wait()
}
func TestCreateBlog(t *testing.T) {
	//UpdateLocalBlogData()
	UpdateTag()
}
