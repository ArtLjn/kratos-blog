package data

import (
	"context"
	"encoding/json"
	"fmt"
	"kratos-blog/api/v1/comment"
	"kratos-blog/app/comment/internal/biz"
	"kratos-blog/app/comment/internal/data/bean"
	"kratos-blog/pkg/util"
	"kratos-blog/pkg/vo"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

const sensitiveWords = "sensitive_words"

type commRepo struct {
	data          *Data
	log           *log.Helper
	sensitiveList []string
	mu            sync.Mutex
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
	if len(c.sensitiveList) == 0 {
		return nil
	}
	// 构建正则表达式模式
	pattern := strings.Join(c.sensitiveList, "|")
	regex := regexp.MustCompile(pattern)
	// 检查敏感词
	matches := regex.FindAllString(word, -1)
	return matches
}

func (c *commRepo) GetIp(ip string) (string, error) {
	checkIp := net.ParseIP(ip)
	if checkIp == nil {
		return "", fmt.Errorf("不合法的IP地址")
	}
	uri := fmt.Sprintf("%s%s", c.data.c.Api.GetIp(), ip)
	body, err := util.Request(http.MethodGet, uri, nil)
	if err != nil {
		c.log.Log(log.LevelError, err)
		return "", fmt.Errorf("请求超时")
	}

	jsonStr := util.GetJsonVal(body, "result")
	data, e := json.Marshal(jsonStr)
	m := strings.ReplaceAll(string(data), "\"", "")
	if e != nil {
		c.log.Log(log.LevelError, err)
		return "", fmt.Errorf("解析失败")
	} else if len(m) == 0 {
		return "", fmt.Errorf("解析失败")
	}
	addrList := util.GetJsonVal(string(data), "addr")
	if addrList == nil {
		return "", fmt.Errorf("未知地址")
	}
	strList, _ := json.Marshal(addrList)
	var list []interface{}
	eee, res := c.data.pf.ParseJSONStrToStruct(string(strList), &list)
	if !res {
		c.log.Log(log.LevelError, eee)
	}
	return list[0].(string), nil
}

func (c *commRepo) AddComment(ctx context.Context, req *comment.CreateCommentRequest) *comment.CreateCommentReply {
	words := c.CheckWords(req.Comment)
	if len(words) != 0 {
		return &comment.CreateCommentReply{
			Code:   400,
			Result: vo.TALK_ERROR,
		}
	}
	ipAddr, err := c.GetIp(req.CommentAddr)
	if err != nil {
		return &comment.CreateCommentReply{Code: vo.BAD_REQUEST, Result: err.Error()}
	}

	var commentParent bean.CommentParent
	e := c.data.pf.ParseJSONToStruct(req, &commentParent)
	if e != nil {
		c.log.Log(log.LevelError, e)
		return &comment.CreateCommentReply{
			Code:   vo.BAD_REQUEST,
			Result: vo.INSERT_ERROR,
		}
	}
	commentParent.CommentAddr = ipAddr
	commentParent.CommentTime = time.Now().Format("2006-01-02 15:04")
	err = c.data.pf.CreateFunc(commentParent, &bean.CommentParent{})
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
	if words := c.CheckWords(req.Comment); len(words) != 0 {
		return &comment.CreateRewardReply{Code: vo.BAD_REQUEST, Result: vo.TALK_ERROR}
	}
	ipAddr, err := c.GetIp(req.CommentAddr)
	if err != nil {
		return &comment.CreateRewardReply{Code: vo.BAD_REQUEST, Result: err.Error()}
	}
	var commentSubTwo bean.CommentSubTwo
	e := c.data.pf.ParseJSONToStruct(req, &commentSubTwo)
	if e != nil {
		c.log.Log(log.LevelError, e)
		return &comment.CreateRewardReply{
			Code:   vo.BAD_REQUEST,
			Result: vo.INSERT_ERROR,
		}
	}
	commentSubTwo.CommentTime = time.Now().Format("2006-01-02 15:04")
	commentSubTwo.CommentAddr = ipAddr
	err = c.data.pf.CreateFunc(commentSubTwo, &bean.CommentSubTwo{})
	if err != nil {
		c.log.Log(log.LevelError, err)
		return &comment.CreateRewardReply{
			Code:   vo.BAD_REQUEST,
			Result: vo.INSERT_ERROR,
		}
	}

	return &comment.CreateRewardReply{Code: vo.SUCCESS_REQUEST, Result: vo.TALK_SUCCESS}
}

func (c *commRepo) ExtractParentComments(ctx context.Context, req *comment.ExtractParentCommentsRequest) *comment.ExtractParentCommentsReply {
	commentParent, err := c.data.pf.QueryFunc(bean.CommentParent{}, map[string]interface{}{"article_id": req.Id}, true)
	if err != nil {
		c.log.Log(log.LevelError, err)
		return &comment.ExtractParentCommentsReply{Code: vo.BAD_REQUEST, Result: vo.QUERY_FAIL}
	}

	commentSub, err := c.data.pf.QueryFunc(bean.CommentSubTwo{}, map[string]interface{}{"article_id": req.Id}, true)
	if err != nil {
		c.log.Log(log.LevelError, err)
		return &comment.ExtractParentCommentsReply{Code: vo.BAD_REQUEST, Result: vo.QUERY_FAIL}
	}

	var (
		parents []bean.CommentParent
		subs    []bean.CommentSubTwo
		mapList []*comment.ExtractResult
	)

	if err = c.data.pf.ParseJSONToStruct(commentParent, &parents); err != nil {
		c.log.Log(log.LevelError, err)
		return &comment.ExtractParentCommentsReply{Code: vo.BAD_REQUEST, Result: vo.QUERY_FAIL}
	}

	if err = c.data.pf.ParseJSONToStruct(commentSub, &subs); err != nil {
		c.log.Log(log.LevelError, err)
		return &comment.ExtractParentCommentsReply{Code: vo.BAD_REQUEST, Result: vo.QUERY_FAIL}
	}

	for i := 0; i < len(parents); i++ {
		c.mu.Lock()
		parent := parents[i]
		var mp comment.ExtractResult
		if err := c.data.pf.ParseJSONToStruct(parent, &mp); err != nil {
			c.mu.Unlock()
			c.log.Log(log.LevelError, err)
			return &comment.ExtractParentCommentsReply{Code: vo.BAD_REQUEST, Result: vo.QUERY_FAIL}
		}

		var child []*comment.CommentSubResult
		for j := 0; j < len(subs); j++ {
			if subs[j].ParentID == parent.ID {
				var s comment.CommentSubResult
				if err := c.data.pf.ParseJSONToStruct(subs[j], &s); err != nil {
					c.mu.Unlock()
					c.log.Log(log.LevelError, err)
					return &comment.ExtractParentCommentsReply{Code: vo.BAD_REQUEST, Result: vo.QUERY_FAIL}
				}
				child = append(child, &s)
			}
		}
		mp.Child = child
		mapList = append(mapList, &mp)
		c.mu.Unlock()
	}

	return &comment.ExtractParentCommentsReply{Code: vo.SUCCESS_REQUEST, Result: vo.QUERY_SUCCESS, List: mapList}
}
