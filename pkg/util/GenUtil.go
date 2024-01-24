package util

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func RandomInt() int {
	rand.Seed(time.Now().UnixNano())

	//生成一个随机6位数
	randomNumber := rand.Intn(900000) + 100000
	return randomNumber
}

func GenerateToken() string {
	token := uuid.New().String()
	return token
}
