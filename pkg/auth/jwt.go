package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"time"
)

type LoginClaims struct {
	Username string
	jwt.RegisteredClaims
}

var (
	signKey = "blog_key"
	letters = []rune("@(@&LJQLFQNLUEKssldjfoweasfdss0")
)

func randStr(strLen int) string {
	randBytes := make([]rune, strLen)
	for i := range randBytes {
		randBytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(randBytes)
}

/**
 * 生成签名，30分钟过期
 * @param username 用户名
 * @param loginTime 登录时间
 * @return 生成的token
 */

func Sign(username string) (string, error) {
	claim := LoginClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                          // 签发者
			Subject:   "Auth",                                                 // 签发对象
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),        //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                         //签发时间
			ID:        randStr(10),                                            // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(signKey))
	return token, err
}

/**
 * 检验token是否正确
 * @param token 需要校验的token
 * @return 校验是否成功
 */

func Verify(tokenString string) (*LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*LoginClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

/**
 * 获取token中包含的登录用户名
 * @param token 需要校验的token
 * @return 登录用户名
 */

func GetLoginName(token string) string {
	claims, err := Verify(token)
	if err != nil {
		return ""
	}
	return claims.Username
}
