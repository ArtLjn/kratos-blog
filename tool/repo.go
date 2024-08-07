/*
@Time : 2024/7/30 上午10:20
@Author : ljn
@File : repo
@Software: GoLand
*/

package main

import (
	"sync"
	"time"
)

type UploadRepo struct {
	UploadPath string //上传路径
	Domain     string // 域名
	MaxSize    string // 最大文件大小（字节）
	Url        string // 外部uri
}

func NewUploadRepo(data ...string) *UploadRepo {
	return &UploadRepo{UploadPath: data[0], Domain: data[1], MaxSize: data[2], Url: data[3]}
}

// CurrentBing :dev 今日Bing图片
type CurrentBing struct {
	Url            string
	lastUpdateTime time.Time
	mu             sync.Mutex
}
