package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplitEmail(t *testing.T) {
	email := "example@qq.com"
	emailPrefix := strings.Split(email, "@")[1]
	fmt.Println(emailPrefix)
}
