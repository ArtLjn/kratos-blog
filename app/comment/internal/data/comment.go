package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-blog/api/v1/comment"
	"kratos-blog/app/comment/internal/biz"
	"kratos-blog/pkg/util"
	"kratos-blog/pkg/vo"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const sensitiveWords = "sensitive_words"

type commRepo struct {
	data          *Data
	log           *log.Helper
	sensitiveList []string
}

type Comment struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	OriginID      string `json:"origin_id"`
	ArticleID     string `json:"article_id"`
	Comment       string `json:"comment"`
	CommentTime   string `json:"comment_time"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	RewardName    string `json:"reward_name"`
	RewardTime    string `json:"reward_time"`
	RewardContent string `json:"reward_content"`
	RewardEmail   string `json:"reward_email"`
	CommentAddr   string `json:"comment_addr"`
	RewardAddr    string `json:"reward_addr"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func NewCommRepo(data *Data, logger log.Logger, sensitiveList []string) biz.CommRepo {
	return &commRepo{
		data:          data,
		log:           log.NewHelper(logger),
		sensitiveList: sensitiveList,
	}
}

func QueryWords(db *gorm.DB) []string {
	var words []string

	// 查询敏感词
	if err := db.Table(sensitiveWords).Select("word").Find(&words).Error; err != nil {
		panic(err)
	}

	return words
}

func (c *commRepo) CheckWords(word string) []string {
	// 构建正则表达式模式
	pattern := strings.Join(c.sensitiveList, "|")
	regex := regexp.MustCompile(pattern)
	// 检查敏感词
	matches := regex.FindAllString(word, -1)
	return matches
}

func (c *commRepo) GetIp(ip string) string {
	uri := c.data.c.Api.GetIp() + ip
	body, err := util.Request(http.MethodGet, uri, nil)
	if err != nil {
		c.log.Log(log.LevelError, err)
		return ""
	}

	jsonStr := util.GetJsonVal(body, "result")
	data, e := json.Marshal(jsonStr)
	m := strings.ReplaceAll(string(data), "\"", "")
	if e != nil {
		c.log.Log(log.LevelError, err)
		return ""
	} else if len(m) == 0 {
		return ""
	}
	addrList := util.GetJsonVal(string(data), "addr")
	strList, _ := json.Marshal(addrList)
	var list []interface{}
	eee, res := c.data.pf.ParseJSONStrToStruct(string(strList), &list)
	if !res {
		c.log.Log(log.LevelError, eee)
	}
	return list[0].(string)
}

func (c *commRepo) AddComment(ctx context.Context, req *comment.CreateCommentRequest) *comment.CreateCommentReply {
	user := c.data.role.QueryUserMsg(ctx)
	name := user.U.Data[0]
	email := user.U.Data[1]
	words := c.CheckWords(req.Comment)
	if len(words) != 0 {
		return &comment.CreateCommentReply{
			Code:   400,
			Result: vo.TALK_ERROR,
		}
	}
	ipAddr := c.GetIp(req.CommentAddr)
	if len(ipAddr) == 0 {
		return &comment.CreateCommentReply{Code: vo.BAD_REQUEST, Result: vo.IPADDR_ERROR}
	}
	commentBean := Comment{
		ArticleID:   req.ArticleId,
		Comment:     req.Comment,
		CommentAddr: ipAddr,
		CommentTime: time.Now().Format("2022-12-22 01:12"),
		Name:        name,
		Email:       email,
	}
	err := c.data.pf.CreateFunc(commentBean, &Comment{})
	if err != nil {
		c.log.Log(log.LevelError, err)
		return &comment.CreateCommentReply{
			Code:   vo.BAD_REQUEST,
			Result: vo.INSERT_ERROR,
		}
	}
	return &comment.CreateCommentReply{Code: vo.SUCCESS_REQUEST, Result: vo.INSERT_SUCCESS}
}
