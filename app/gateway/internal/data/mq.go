package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"kratos-blog/api/v1/comment"
	"kratos-blog/api/v1/user"
	"kratos-blog/app/gateway/internal/conf"
	"kratos-blog/pkg/mq"
	"kratos-blog/pkg/vo"
	"sync"
)

type CommentReq[T any, R any] func(ctx context.Context, cm *T, opts ...grpc.CallOption) (*R, error)
type CommentFilter func(ctx context.Context, articleId string) error
type SendEmailF func(ctx context.Context, in *user.SendEmailCommonRequest, opts ...grpc.CallOption) (*user.SendEmailCommonReply, error)

// 消费订单消息并处理
var workerPool = make(chan struct{}, 10) // 假设最多同时处理10个消息

type Mq struct {
	cn *mq.MQ
	l  *log.Helper
	cf *conf.Mq
}

func NewCommentMq(c *conf.Mq, l *log.Helper) *Mq {
	cn, err := mq.ConnMq(c.GetUrl())
	if err != nil {
		err = cn.Conn.Close()
		if err != nil {
			return nil
		}
	}
	return &Mq{
		cn: cn,
		cf: c,
		l:  l,
	}
}

// StartMq 处理评论信息
func (c *Mq) StartMq(queue ...string) *mq.MQ {
	// 声明交换机
	err := mq.CreateExchange(c.cn, c.cf.GetExchange()[0], "direct")
	if err != nil {
		err = c.cn.Conn.Close()
		if err != nil {
			return nil
		}
	}
	for _, q := range queue {
		err = mq.CreateQueueAndBind(c.cn, q, q, c.cf.GetExchange()[0])
		if err != nil {
			err = c.cn.Conn.Close()
			if err != nil {
				return nil
			}
		}
	}
	return c.cn
}

func (c *Mq) PushComment(m *comment.CreateCommentRequest) error {
	// 将订单数据序列化
	commentBytes, err := json.Marshal(&m)
	if err != nil {
		return fmt.Errorf("序列化失败")
	}
	err = mq.PublicMsg(c.cn.Channel, c.cf.GetExchange()[0], c.cf.GetQueue()[0], commentBytes)
	if err != nil {
		c.l.Error("发送消息失败")
		return fmt.Errorf("发送消息失败")
	}
	c.l.Info("发送消息成功")
	return nil
}

// PushReward push邮件任务到队列中
func (c *Mq) PushReward(m *comment.CreateRewardRequest) error {
	// 将订单数据序列化
	commentBytes, err := json.Marshal(&m)
	if err != nil {
		return fmt.Errorf("序列化失败")
	}
	err = mq.PublicMsg(c.cn.Channel, c.cf.GetExchange()[0], c.cf.GetQueue()[1], commentBytes)
	if err != nil {
		c.l.Error("发送消息失败")
		return fmt.Errorf("发送消息失败")
	}
	c.l.Info("发送消息成功")
	return nil
}

type MailS struct {
	Email string
}

func (c *Mq) PushMail(mails MailS) error {
	mailsBytes, err := json.Marshal(&mails)
	if err != nil {
		return fmt.Errorf("序列化失败")
	}
	err = mq.PublicMsg(c.cn.Channel, c.cf.GetExchange()[0], c.cf.GetQueue()[2], mailsBytes)
	if err != nil {
		c.l.Error("发送消息失败")
		return fmt.Errorf("发送消息失败")
	}
	c.l.Info("发送消息成功")
	return nil
}

// ReceiveComment 处理评论信息
func (c *Mq) ReceiveComment(fc CommentReq[comment.CreateCommentRequest, comment.CreateCommentReply],
	filter CommentFilter, mails SendEmailF) {
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
		c.l.Error("消费消息失败")
		return
	}

	forever := make(chan bool)
	go func() {
		var wg sync.WaitGroup
		for d := range msg {
			wg.Add(1)
			c.l.Info("收到消息")
			workerPool <- struct{}{} // 获取信号量，控制并发数
			go func(delivery amqp.Delivery) {
				defer wg.Done()
				defer func() {
					<-workerPool
					// 将消息反序列化
					var cm comment.CreateCommentRequest
					err = json.Unmarshal(d.Body, &cm)
					if err != nil {
						c.l.Error("反序列化消息失败")
						return
					}

					// 处理消息
					err = filter(context.Background(), cm.ArticleId)
					if err != nil {
						c.l.Error("过滤消息失败")
					}

					resp, _ := fc(context.Background(), &cm)
					if resp.Code != vo.SUCCESS_REQUEST {
						c.l.Error("处理消息失败")
						if resp.Result == vo.TALK_ERROR {
							go func() {
								res, _ := mails(context.Background(), &user.SendEmailCommonRequest{
									Email:   cm.GetEmail(),
									Subject: "评论异常",
									Body:    "系统检测到你有异常评论行为，请文明评论，否则将会被封禁处理",
								})
								if res.GetCommon().GetCode() != vo.SUCCESS_REQUEST {
									c.l.Error("发送邮件失败")
								}
							}()
						}
						return
					}
				}()
			}(d)
			// 消息确认
			err = d.Ack(false)
			if err != nil {
				c.l.Error("消息确认失败")
			}
			wg.Wait()
		}
	}()
	<-forever
}

// ReceiveReward 处理回复信息
func (c *Mq) ReceiveReward(fc CommentReq[comment.CreateRewardRequest, comment.CreateRewardReply], filter CommentFilter, mails SendEmailF) {
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
		c.l.Error("消费消息失败")
		return
	}

	forever := make(chan bool)
	go func() {
		var wg sync.WaitGroup
		for d := range msg {
			wg.Add(1)
			c.l.Info("收到消息")
			workerPool <- struct{}{} // 获取信号量，控制并发数
			go func(delivery amqp.Delivery) {
				defer wg.Done()
				defer func() {
					<-workerPool
					// 将消息反序列化
					var cm comment.CreateRewardRequest
					err = json.Unmarshal(d.Body, &cm)
					if err != nil {
						c.l.Error("反序列化消息失败")
						return
					}

					// 处理消息
					err = filter(context.Background(), cm.ArticleId)
					if err != nil {
						c.l.Error("过滤消息失败")
					}

					resp, _ := fc(context.Background(), &cm)
					if resp.Code != vo.SUCCESS_REQUEST {
						c.l.Error("处理消息失败")
						if resp.Result == vo.TALK_ERROR {
							go func() {
								res, _ := mails(context.Background(), &user.SendEmailCommonRequest{
									Email:   cm.GetEmail(),
									Subject: "评论异常",
									Body:    "系统检测到你有异常评论行为，请文明评论，否则将会被封禁处理",
								})
								if res.GetCommon().GetCode() != vo.SUCCESS_REQUEST {
									c.l.Error("发送邮件失败")
								}
							}()
						}
						return
					}
				}()
			}(d)
			// 消息确认
			err = d.Ack(false)
			if err != nil {
				c.l.Error("消息确认失败")
			}
			wg.Wait()
		}
	}()
	<-forever
}

func (c *Mq) ReceiveEmailTask(userClient user.UserClient) {
	msg, err := c.cn.Channel.Consume(
		c.cf.GetQueue()[2],
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.l.Error("消费消息失败")
		return
	}

	forever := make(chan bool)
	go func() {
		var wg sync.WaitGroup
		for d := range msg {
			wg.Add(1)
			c.l.Info("收到消息")
			workerPool <- struct{}{} // 获取信号量，控制并发数
			go func(delivery amqp.Delivery) {
				defer wg.Done()
				defer func() {
					<-workerPool // 释放信号量
				}()

				// 将消息反序列化
				var email MailS
				err = json.Unmarshal(delivery.Body, &email)
				if err != nil {
					c.l.Error("反序列化消息失败")
					// 拒绝消息，并将其重新放回队列
					delivery.Nack(false, true)
					return
				}

				// 发送邮件
				res, _ := userClient.SendEmail(context.Background(), &user.SendEmailRequest{
					Email: email.Email,
				})
				if res.Common.Code != 200 {
					// 发送邮件失败，拒绝消息并重新放回队列
					delivery.Nack(false, true)
					c.l.Error("发送邮件失败")
					return
				}

				// 处理成功，确认消息
				err = delivery.Ack(false)
				if err != nil {
					c.l.Error("消息确认失败")
				}
			}(d)
		}
		wg.Wait()
	}()
	<-forever
}
