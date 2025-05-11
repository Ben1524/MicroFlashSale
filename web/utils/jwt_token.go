package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// UserToken 定义用户令牌结构体
type UserToken struct {
	jwt.StandardClaims
	// 自定义的用户信息
	UserName string `json:"user_name"`
}

// 前端用户的过期时间和加密秘钥
var FrontUserExpireDuration = time.Hour
var FrontUserSecretKey = []byte("User")

// 后端用户的过期时间和加密秘钥
var AdminUserExpireDuration = time.Hour * 2
var AdminUserSecretKey = []byte("Admin")

// GenToken 生成 JWT 令牌
func GenToken(UserName string, expireDuration time.Duration, secretKey []byte) (string, error) {
	// 生成一个新的 token
	user := UserToken{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(), // 过期时间
			Issuer:    "ZJQ",                                 // 签发人
		},
		UserName, // 使用了结构体隐式字段，这里要按照顺序来
	}
	// 创建一个带有声明的 JWT 令牌，即JWT的
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user) //jwt.NewWithClaims(jwt.SigningMethodHS256, user) 会自动生成包含算法（HS256）和类型（JWT）的头部。
	// 对令牌进行签名，jwt签名和加密
	return token.SignedString(secretKey)
}

func AuthToken(tokenString string, secretKey []byte) (*UserToken, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserToken); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
