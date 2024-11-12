package auth

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"kratos-blog/pkg/server"
)

// Token 管理模块 ** 存储|过期｜续期

const TokenMangerKey = server.TokenMangerKey

type TokenManager interface {
	SaveToken(string) string
	LogOutToken(string)
	VerifyToken(string) error
}

type Token struct {
	rdb *redis.Client
	l   *log.Helper
}

func NewToken(rdb *redis.Client, l *log.Helper) TokenManager {
	return &Token{
		rdb: rdb,
		l:   l,
	}
}

func (t *Token) SaveToken(username string) string {
	token, err := Sign(username)
	if err != nil {
		t.l.Error("❌ 生成token失败")
		return ""
	}
	_, err = t.rdb.HSet(context.Background(), TokenMangerKey, username, token).Result()
	if err != nil {
		t.l.Error("❌ 存储token失败")
		return ""
	}
	return token
}

func (t *Token) LogOutToken(username string) {
	_, err := t.rdb.HDel(context.Background(), TokenMangerKey, username).Result()
	if err != nil {
		t.l.Error("❌ 删除token失败")
	}
}

func (t *Token) VerifyToken(token string) error {
	username := GetLoginName(token)
	if username == "" {
		t.l.Error("❌ token无效")
		return fmt.Errorf("❌ token无效")
	}
	cacheToken, err := t.rdb.HGet(context.Background(), TokenMangerKey, username).Result()
	if err != nil {
		t.l.Error("❌ token不存在")
		return fmt.Errorf("❌ token不存在")
	}
	if token != cacheToken {
		t.l.Error("❌ token无效")
		return fmt.Errorf("❌ token无效")
	}
	return nil
}
