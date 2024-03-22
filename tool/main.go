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
	r.Run(":8099")
}

func init() {
	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		go UpdatePhoto()
	})
	c.Start()
}

func UpdatePhoto() {
	go func() {
		Current.mu.Lock()
		defer Current.mu.Unlock()
		Current.lastUpdateTime = time.Now()
		BingPhoto(&Current.Url)
		log.Println(Current)
	}()
}
