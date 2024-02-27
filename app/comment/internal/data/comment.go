package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-blog/api/v1/comment"
	"kratos-blog/app/comment/internal/biz"
	"kratos-blog/app/comment/internal/data/bean"
	"kratos-blog/pkg/util"
	"kratos-blog/pkg/vo"
	"net/http"
	"regexp"
	"strconv"
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
	uri := fmt.Sprintf("%s%s", c.data.c.Api.GetIp(), ip)
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
		CommentTime: time.Now().Format("2006-01-02 15:04"),
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
	return &comment.CreateCommentReply{Code: vo.SUCCESS_REQUEST, Result: vo.TALK_SUCCESS}
}

func (c *commRepo) AddReward(ctx context.Context, req *comment.CreateRewardRequest) *comment.CreateRewardReply {
	if words := c.CheckWords(req.RewardContent); len(words) != 0 {
		return &comment.CreateRewardReply{Code: vo.BAD_REQUEST, Result: vo.TALK_ERROR}
	}
	ipAddr := c.GetIp(req.RewardAddr)
	if len(ipAddr) == 0 {
		return &comment.CreateRewardReply{Code: vo.BAD_REQUEST, Result: vo.IPADDR_ERROR}
	}

	originComment := c.QueryCommentCond(map[string]interface{}{"id": req.RewardId, "article_id": req.ArticleId})
	if len(originComment.ArticleID) == 0 {
		return &comment.CreateRewardReply{Code: vo.BAD_REQUEST, Result: vo.REQUEST_FAIL}
	}

	currentComment := originComment
	if len(currentComment.RewardName) != 0 {
		c.P(&currentComment, ctx, req.RewardContent, ipAddr, req.RewardId)
		currentComment.ID = 0
		err := c.data.pf.CreateFunc(currentComment, &Comment{})
		if err != nil {
			c.log.Log(log.LevelError, err)
			return &comment.CreateRewardReply{Code: vo.BAD_REQUEST, Result: vo.TALK_FAIL}
		}
	} else {
		c.P(&originComment, ctx, req.RewardContent, ipAddr, req.RewardId)
		var data map[string]interface{}
		c.data.pf.ParseJSONToStruct(originComment, &data)
		err := c.data.pf.UpdateFunc(Comment{}, map[string]interface{}{
			"id": req.RewardId,
		}, data, false)
		if err != nil {
			c.log.Log(log.LevelError, err)
			return &comment.CreateRewardReply{Code: vo.BAD_REQUEST, Result: vo.TALK_FAIL}
		}
	}

	return &comment.CreateRewardReply{Code: vo.SUCCESS_REQUEST, Result: vo.TALK_SUCCESS}
}

func (c *commRepo) P(com *Comment, ctx context.Context, data ...string) {
	user := c.data.role.QueryUserMsg(ctx)
	name, email := user.U.Data[0], user.U.Data[1]
	if len(name) == 0 || len(email) == 0 {
		name = "访客"
		email = "example@qq.com"
	}
	com.RewardName = name
	com.RewardContent = data[0]
	com.RewardEmail = email
	com.RewardAddr = data[1]
	com.RewardTime = time.Now().Format("2006-01-02 15:04")
	com.OriginID = data[2]
}

func (c *commRepo) QueryCommentCond(cond map[string]interface{}) Comment {
	originComment, err := c.data.pf.QueryFunc(Comment{}, cond, false)
	comment := Comment{}
	empty := comment
	if err != nil {
		c.log.Log(log.LevelError, err)
		return empty
	}
	c.data.pf.ParseJSONToStruct(originComment, &comment)
	return comment
}

func (c *commRepo) ExtractParentComments(ctx context.Context,

	req *comment.ExtractParentCommentsRequest) *comment.ExtractParentCommentsReply {
	body, err := c.data.pf.QueryFunc(Comment{}, map[string]interface{}{"article_id": req.Id}, true)
	if err != nil {
		c.log.Log(log.LevelError, err)
		return &comment.ExtractParentCommentsReply{Code: vo.BAD_REQUEST, Result: vo.QUERY_FAIL}
	}
	var (
		comments []Comment
		cmo      bean.CommentBean
		child    bean.ChildComment
	)
	c.data.pf.ParseJSONToStruct(body, &comments)

	commentSet := make(map[string][]bean.CommentBean)
	for _, comment := range comments {
		c.data.pf.ParseJSONToStruct(comment, &cmo)
		cmo.ID = strconv.Itoa(comment.ID)
		if len(comment.OriginID) != 0 {
			c.data.pf.ParseJSONToStruct(comment, &child)
			if _, ok := commentSet[comment.OriginID]; !ok {
				cmo.ChildComments = append(cmo.ChildComments, []bean.ChildComment{child}...)
				commentSet[comment.OriginID] = append(commentSet[comment.OriginID], cmo)
			} else {
				for index, parentComment := range commentSet[comment.OriginID] {
					if parentComment.ID == comment.OriginID {
						commentSet[comment.OriginID][index].ChildComments =
							append(commentSet[comment.OriginID][index].ChildComments, child)
						break
					}
				}
			}
		} else {
			commentSet[strconv.Itoa(comment.ID)] = append(commentSet[strconv.Itoa(comment.ID)], cmo)
		}
	}
	byteRes, _ := json.Marshal(commentSet)
	strBody := string(byteRes)
	return &comment.ExtractParentCommentsReply{Code: vo.SUCCESS_REQUEST, Result: vo.QUERY_SUCCESS,
		List: []string{strBody}}
}
