package jwt

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// ExtractUserInfoFromClaims 从jwt载荷中提取出用户信息
func ExtractUserInfoFromClaims(ctx context.Context) string {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return ""
	}

	str := claims.(jwtv4.MapClaims)["username"]
	if str == nil {
		return ""
	}
	switch str.(type) {
	case string:
		break
	default:
		return ""
	}
	return str.(string)
}

// EncryptUserInfoToJwtToken 将用户名加密到jwt载荷中
func EncryptUserInfoToJwtToken(key []byte, username string) (string, error) {
	claims := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256,
		jwtv4.MapClaims{
			"username": username,
		})
	return claims.SignedString(key)
}