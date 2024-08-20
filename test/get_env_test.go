/*
@Time : 2024/8/19 下午6:34
@Author : ljn
@File : get_env_test
@Software: GoLand
*/

package test

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	godotenv.Load("../.env")
	blogConfPath := os.Getenv("BLOG_CONF_PATH")
	commentConfPath := os.Getenv("COMMENT_CONF_PATH")

	// 使用环境变量
	log.Println("Blog config path:", blogConfPath)
	log.Println("Comment config path:", commentConfPath)
}
