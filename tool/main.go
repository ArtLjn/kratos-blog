/**
 * @author ljn
 * @date 2024-03-21 17:41:16
 * @dev 资源服务器上传文件使用
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kratos-blog/pkg/m_logs"
	"log"
	"net/http"
	"time"
)

var (
	ProxyPath string
	Host      string
	Port      string
)

func main() {
	m_logs.InitGinLog("tool")
	r := gin.Default()
	r.StaticFS("/img", http.Dir(ProxyPath))
	UpdatePhoto()

	options := NewFilterOptions(
		WithOriginFilter(),
		WithIpWhiteFilter(),
		WithApiKeyFilter(),
	)
	if len(options.filters) != 0 {
		for _, filter := range options.filters {
			r.Use(filter.Apply())
		}
	}

	InitRouter(r)
	err := r.Run(fmt.Sprintf("%s:%s", Host, Port))
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
