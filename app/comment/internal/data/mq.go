package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/api/v1/comment"
	"kratos-blog/app/comment/internal/conf"
	"kratos-blog/pkg/mq"
	"kratos-blog/pkg/vo"
)

type CommentMq struct {
	cn *mq.MQ
	l  *log.Helper
	cf *conf.Mq
}

func NewCommentMq(c *conf.Mq) *CommentMq {
	cn, err := mq.ConnMq(c.GetUrl())
	if err != nil {
		err = cn.Conn.Close()
		if err != nil {
			return nil
		}
	}
	return &CommentMq{
		cn: cn,
		cf: c,
	}
}

// StartMq 处理评论信息
func (c *CommentMq) StartMq(queue string) *mq.MQ {
	// 声明交换机
	err := mq.CreateExchange(c.cn, c.cf.GetExchange()[0], "direct")
	if err != nil {
		err = c.cn.Conn.Close()
		if err != nil {
			return nil
		}
	}
	err = mq.CreateQueueAndBind(c.cn, queue, queue, c.cf.GetExchange()[0])
	if err != nil {
		err = c.cn.Conn.Close()
		if err != nil {
			return nil
		}
	}
	return c.cn
}

func (c *CommentMq) PublicComment(m *comment.CreateCommentRequest) {
	// 将订单数据序列化
	commentBytes, err := json.Marshal(&m)
	if err != nil {
		c.l.Error("❌ 序列化订单数据失败")
		return
	}
	err = mq.PublicMsg(c.cn.Channel, c.cf.GetExchange()[0], c.cf.GetQueue()[0], commentBytes)
	if err != nil {
		c.l.Error("❌ 发送消息失败")
		return
	}
	c.l.Info("✅ 发送消息成功")
}

func (c *CommentMq) PublicReward(m *comment.CreateCommentRequest) {
	// 将订单数据序列化
	commentBytes, err := json.Marshal(&m)
	if err != nil {
		c.l.Error("❌ 序列化订单数据失败")
		return
	}
	err = mq.PublicMsg(c.cn.Channel, c.cf.GetExchange()[0], c.cf.GetQueue()[1], commentBytes)
	if err != nil {
		c.l.Error("❌ 发送消息失败")
		return
	}
	c.l.Info("✅ 发送消息成功")
}

// ReceiveComment 处理评论信息
func (c *CommentMq) ReceiveComment(fc func(ctx context.Context, cm *comment.CreateCommentRequest) *comment.CreateCommentReply) {
	msg, err := c.cn.Channel.Consume(
		c.cf.GetQueue()[0],
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.l.Error("❌ 消费消息失败")
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msg {
			c.l.Info("✅ 收到消息")
			// 将消息反序列化
			var cm comment.CreateCommentRequest
			err = json.Unmarshal(d.Body, &cm)
			if err != nil {
				c.l.Error("❌ 反序列化消息失败")
				return
			}
			// 处理消息
			c.l.Info("✅ 处理消息")
			resp := fc(context.Background(), &cm)
			if resp.Code != vo.SUCCESS_REQUEST {
				c.l.Error("❌ 处理消息失败")
				// TODO 消息同步客户端
				continue
			}
			// 消息确认
			d.Ack(false)
		}
	}()
	<-forever
}

// ReceiveReward 处理回复信息
func (c *CommentMq) ReceiveReward(fc func(cm *comment.CreateRewardRequest)) {
	msg, err := c.cn.Channel.Consume(
		c.cf.GetQueue()[1],
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.l.Error("❌ 消费消息失败")
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msg {
			c.l.Info("✅ 收到消息")
			// 将消息反序列化
			var cm comment.CreateRewardRequest
			err = json.Unmarshal(d.Body, &cm)
			if err != nil {
				c.l.Error("❌ 反序列化消息失败")
				return
			}
			// 处理消息
			c.l.Info("✅ 处理消息")
			fc(&cm)
			// 消息确认
			d.Ack(false)
		}
	}()
	<-forever
}
