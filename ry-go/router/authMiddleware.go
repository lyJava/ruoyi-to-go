package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"ry-go/business/domain"
	"ry-go/common/response"
	"ry-go/configuation"
	"ry-go/utils"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware 请求凭证中间件
func AuthMiddleware(version string) echo.MiddlewareFunc {
	whiteURLs := configuation.ViperConfig.GetStringSlice("request.white-url")
	zap.L().Sugar().Infof("请求白名单:%s", utils.ToJsonFormat(whiteURLs))
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 在响应头中添加版本信息
			c.Response().Header().Set("X-Api-Version", version)

			currentUrl := c.Request().URL.Path
			requestId := utils.GetValForContext(c.Request().Context(), "requestId")
			zap.L().Info("当前请求", zap.String("url", currentUrl), zap.String("requestId", requestId))

			// 白名单检查
			if isWhiteList(currentUrl, whiteURLs) {
				return next(c)
			}

			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				zap.L().Sugar().Errorf("请求===%s头缺失Authorization,requestId===%s", currentUrl, requestId)
				response.NewRespCodeMsg(c, http.StatusUnauthorized, "请求头凭证不能为空")
				return errors.New("请求头凭证不能为空")
			}

			parts := strings.SplitN(authorization, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				zap.L().Sugar().Errorf("请求头===%s凭证===%s格式错误,requestId===%s", currentUrl, authorization, requestId)
				response.NewRespCodeMsg(c, http.StatusUnauthorized, "请求头凭证格式错误")
				return errors.New("请求头凭证格式错误")
			}

			token := parts[1]

			if err := verifyToken(c, token); err != nil {
				return err
			}
			return next(c)
		}
	}
}

// verifyToken 验证请求凭证是否生效过期等等
//
// 参数
//
//	-c: echo上下文
//	-tokenStr: jwt的请求凭证字符串
//
// 返回
//
//	-error: 错误信息
func verifyToken(c echo.Context, tokenStr string) error {
	myClaims, err := utils.ParseToken(tokenStr, configuation.JwtConfig.Item.Secret)
	if err != nil {
		response.NewRespCodeErr(c, http.StatusUnauthorized, err)
		return err
	}

	redisKey := configuation.CacheConfig.Cache.Login.Prefix + myClaims.Uuid
	// 从缓存中获取token
	tokenData, err := utils.Get(c.Request().Context(), configuation.GlobalRedisClient, redisKey)
	switch {
	case errors.Is(err, redis.Nil):
		response.NewRespCodeMsg(c, http.StatusUnauthorized, "请求凭证已过期，请重新登陆")
		return err
	case err != nil:
		response.NewRespCodeMsg(c, http.StatusUnauthorized, "当前会话已过期，请重新登陆")
		return err
	}
	if tokenData == "" {
		response.NewRespCodeMsg(c, http.StatusUnauthorized, "请求凭证已过期")
		return errors.New("请求凭证已过期，请重新登陆")
	}

	// 用户退出/注销请求不刷新token
	if !strings.Contains(c.Request().URL.Path, "/logout") {
		var loginUser *domain.LoginUser
		if err = json.Unmarshal([]byte(tokenData), &loginUser); err != nil {
			zap.L().Error("用户缓存反序列化用户错误===", zap.Error(err), zap.String("key", redisKey))
		}

		// 用户令牌有效期小于20分钟，刷新用户信息缓存
		if loginUser != nil && loginUser.ExpireTime-time.Now().Unix() < 20*time.Minute.Milliseconds() {
			loginUser.LoginTime = time.Now().Unix()
			loginUser.ExpireTime = loginUser.LoginTime + 30*time.Minute.Milliseconds()
			result, err := utils.Set(c.Request().Context(), configuation.GlobalRedisClient, redisKey, loginUser, 30*time.Minute)
			if err != nil {
				zap.L().Error(fmt.Sprintf("用户===%s令牌刷新错误===", loginUser.Username), zap.Error(err), zap.String("key", redisKey))
			}
			zap.L().Error(fmt.Sprintf("用户===%s令牌刷新错误===", loginUser.Username), zap.String("result", result))
		}
	}

	return nil
}

// isWhiteList 使用正则表达式精确匹配路径
func isWhiteList(path string, patterns []string) bool {
	for _, p := range patterns {
		matched, _ := regexp.MatchString("^"+regexp.QuoteMeta(p)+"(/.*)?$", path)
		if matched {
			return true
		}
	}
	return false
}
