/**
 * @author ljn
 * @date 2024-03-21 17:41:16
 * @dev 资源服务器上传文件使用
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	InitRouter(r)
	r.StaticFS("/img", http.Dir("tool/assets"))
	UpdatePhoto()
	err := r.Run(":8099")
	if err != nil {
		return
	}
}

func init() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() {
		go UpdatePhoto()
	})
	if err != nil {
		return
	}
	c.Start()
}

func UpdatePhoto() {
	go func() {
		Current.mu.Lock()
		defer Current.mu.Unlock()
		Current.lastUpdateTime = time.Now()
		c := make(chan string, 1)
		err := BingPhoto(c)
		if err != nil {
			log.Println(err)
			return
		}
		Current.Url = <-c
		log.Println(Current.Url)
	}()
}
