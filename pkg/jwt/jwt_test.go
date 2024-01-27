package jwt

import (
	"context"
	"fmt"
	"testing"
)

var key = "sssssss"

func TestGenerateJWT(t *testing.T) {
	fmt.Println(EncryptUserInfoToJwtToken([]byte(key), "user001"))
}

func TestParseJWT(t *testing.T) {
	token, _ := EncryptUserInfoToJwtToken([]byte(key), "user001")
	ctx := context.WithValue(context.Background(), "username", token)
	result := ExtractUserInfoFromClaims(ctx)
	fmt.Println(result)
}
