package utils

import (
	"errors"
	"fmt"
	"net/http"
	"ry-go/common/response"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// CreateJwtToken 创建jwt凭证
//
// 参数
//
//	-userId (int64): 用户主键ID
//	-username (string): 用户姓名
//	-expiration (time.Duration): 有效时间
//	-secret ([]byte): 加密字符串
//
// 返回
//
//	-string: 凭证字符串
//	-error: 凭证错误
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func CreateJwtToken(userId int64, username, uuid string, expiration time.Duration, secret []byte) (string, error) {
	userIdStr := fmt.Sprintf("%d", userId)
	claims := response.MyClaims{
		UserId:   userId,
		Username: username,
		Uuid:     uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ry-go-server",
			Subject:   userIdStr,
			Audience:  jwt.ClaimStrings{"test", "api-service"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                 // 可选：生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                 // 签发时间
			ID:        ShortUUID(),
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := tokenClaim.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func SetValueToResp(e echo.Context, key, value string) {
	e.Response().Header().Set(key, value)
}

// GetJwtTokenFromEcho 从echo请求中获取jwt的凭证
//
// 参数
//
//	-e: echo上下文
//	-tokenKey: 请求凭证key
//	-authPrefix: 请求凭证开头
//
// 返回
//
//	-string: 凭证信息
//	-error: 错误
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func GetJwtTokenFromEcho(c echo.Context, tokenKey, authPrefix string) (string, error) {
	//authorization := c.Request.Header.Get("Authorization")
	authorization := c.Request().Header.Get(tokenKey)
	if authorization == "" {
		return "", errors.New("当前请求头中无凭证")
	} else {
		// 按空格分割
		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == authPrefix) {
			return authorization, fmt.Errorf("请求头中的t凭证:%s格式错误", authorization)
		} else {
			authorization = parts[1]
			return authorization, nil
		}
	}
}

// SetJwtTokenToCookie 设置请求凭证到cookie中
//
// 参数
//
//	-ctx: echo.Context echo上下文
//	-name: 名称
//	-token: 请求凭证
//	-maxAge: 过期时间(秒)
//
// 返回
//
//	-string: 小写开头驼峰字符串
//	-error: 凭证错误
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func SetJwtTokenToCookie(ctx echo.Context, name, token string, maxAge int) {
	if token != "" {
		ctx.SetCookie(&http.Cookie{
			Name:     name,        // cookie名称
			Value:    token,       // cookie值
			MaxAge:   maxAge,      // 过期时间，单位为秒
			Path:     "/",         // 可访问cookie的路径
			Domain:   "localhost", // 可访问cookie的域名
			Secure:   false,       // 是否为Secure，true表示cookie只能通过HTTPS发送
			HttpOnly: false,       // 是否为HttpOnly，true表示cookie只能通过HTTP(S)发送，JavaScript无法访问
		})
	}
}

// ValidateJwtToken 验证jwt凭证有效性
//
// 参数
//
//	-tokenStr: 凭证字符串
//	-secret: 密钥字符串
//
// 返回
//
//	-*jwt.Token: 小写开头驼峰字符串
//	-error: 凭证错误
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func ValidateJwtToken(tokenStr, secret string) (*jwt.Token, error) {
	tokenClaim, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("验证token失败 %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("请求头中的凭证格式有效性验证失败")
	}

	if !tokenClaim.Valid {
		return nil, errors.New("请求头中的凭证验证失败")
	}

	return tokenClaim, nil
}

// ParseToken 解析凭证
//
// 参数
//
//	-tokenString: 请求凭证字符串
//	-secret: jwt密钥字符串
//
// 返回
//
//	-*response.MyClaims: 自定义jwt返回结构
//	-error: 错误信息
func ParseToken(tokenString string, secret string) (*response.MyClaims, error) {
	var secretKey = []byte(secret)
	token, err := jwt.ParseWithClaims(tokenString, &response.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("请求凭证已过期，请重新登陆")
		}
		return nil, fmt.Errorf("解析凭证错误==%w", err)
	}

	if claims, ok := token.Claims.(*response.MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("凭证解析失败")
}
