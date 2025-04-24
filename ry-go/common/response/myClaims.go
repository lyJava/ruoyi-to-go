package response

import "github.com/golang-jwt/jwt/v4"

// MyClaims 自定义jwt的token返回字段
type MyClaims struct {
	UserId               int64  `json:"userId"`
	Username             string `json:"username"`
	Uuid                 string `json:"uuid"`
	jwt.RegisteredClaims        //jwt的Claims
}
