// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"

	"kratos-blog/api/v1/blog"
	"kratos-blog/app/gateway/internal/conf"
	"kratos-blog/app/gateway/internal/data"
	"kratos-blog/app/gateway/internal/server"
	"kratos-blog/app/gateway/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confData *conf.Bootstrap, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	r := data.NewRegistrar(registry)
	rdb := data.NewRDB(confData.Data)
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(discovery)
	role := data.NewRole(rdb, userClient, logger)
	commentClient := data.NewCommentServiceClient(discovery)
	blogClient := data.NewBlogServiceClient(discovery)
	q := data.NewCommentMq(confData.Mq, log.NewHelper(logger))
	if q == nil {
		log.Fatalf("conncet mq is failed")
	}
	q.StartMq(confData.Mq.GetQueue()...)
	commentService := service.NewCommentService(logger, commentClient, blogClient, role, q)
	userService := service.NewUserService(logger, userClient)
	// 启动评论和回复的消费者
	go q.ReceiveComment(commentService.UC.AddComment, commentService.HasAllowComment, userService.UC.SendEmailCommon)
	go q.ReceiveReward(commentService.UC.AddReward, commentService.HasAllowComment, userService.UC.SendEmailCommon)
	Cron(blogClient)
	blogService := service.NewBlogService(logger, blogClient, role)
	tagClient := data.NewTagServiceClient(discovery)
	tagService := service.NewTagService(logger, tagClient)
	friendClient := data.NewFriendServiceClient(discovery)
	friendService := service.NewFriendService(friendClient, logger)
	dataData, err := data.NewData(confData, logger, userClient, rdb, role)
	photoClinet := data.NewPhotoServiceClient(discovery)
	photoService := service.NewPhotoService(photoClinet, logger)
	if err != nil {
		return nil, nil, err
	}
	filterRepo := data.NewFilterRepo(dataData, logger)
	httpServer := server.NewHTTPServer(confData, logger, userService,
		commentService, blogService, tagService, friendService, photoService,
		filterRepo)
	app := newApp(logger, httpServer, r)
	return app, func() {
	}, nil
}

// Cron 定时任务
func Cron(cli blog.BlogClient) {
	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() {
		go cli.UpdateBlogVisit(context.Background(), &blog.UpdateBlogVisitRq{})
	})
	c.Start()
}
