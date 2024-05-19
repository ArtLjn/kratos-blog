/**
 * @author ljn
 * @date 2024-03-21 17:41:16
 * @dev 资源服务器上传文件使用
 */

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var ProxyPath string

func main() {
	r := gin.Default()
	InitRouter(r)
	r.StaticFS("/img", http.Dir(ProxyPath))
	UpdatePhoto()
	err := r.Run(":8099")
	if err != nil {
		return
	}
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
