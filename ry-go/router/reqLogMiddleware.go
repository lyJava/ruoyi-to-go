package router

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"ry-go/common/response"
	"ry-go/utils"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// EchoLogger 接收echo框架默认的日志
func EchoLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now()
			path := c.Request().URL.Path
			// 读取请求体
			var requestBody string
			if c.Request().Body != nil {
				bodyBytes, err := io.ReadAll(c.Request().Body)
				if err != nil {
					logger.Error("读取请求体失败", zap.Error(err))
					return err
				}
				requestBody = string(bodyBytes)
				// 将body重新设置回请求体，以便后续处理
				c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}

			// 调用下一个中间件或路由处理器
			err := next(c)

			// 记录请求日志
			duration := time.Since(startTime)
			ginRequestLog := map[string]interface{}{
				"path":     path,
				"method":   c.Request().Method,
				"duration": duration.String(),
				"query":    c.Request().URL.RawQuery,
				"ip":       c.RealIP(),
				"headers":  c.Request().Header,
				"status":   c.Response().Status,
			}

			if strings.HasPrefix(requestBody, "[") && strings.HasSuffix(requestBody, "]") {
				ginRequestLog["body"] = utils.ConvertJsonStrToSlice(requestBody)
			} else {
				params := c.QueryParams()
				if len(params) > 0 {
					ginRequestLog["pathParam"] = params
				} else {
					ginRequestLog["body"] = cast.ToStringMap(requestBody)
				}
			}

			zap.L().Sugar().Infof("请求日志%+v", ginRequestLog)
			return err
		}
	}
}

// EchoRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func EchoRecovery(logger *zap.Logger, stack bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 获取 X-Request-Id
			// requestId := c.Request().Context().Value("requestId").(string)
			requestId := utils.GetValForContext(c.Request().Context(), "requestId")

			defer func() {
				if err := recover(); err != nil {
					// 检查是否是 broken pipe 错误
					var brokenPipe bool
					if ne, ok := err.(*net.OpError); ok {
						if se, ok := ne.Err.(*os.SyscallError); ok {
							if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
								brokenPipe = true
							}
						}
					}

					httpRequest, _ := httputil.DumpRequest(c.Request(), false)
					if brokenPipe {
						logger.Error("broken pipe 错误",
							zap.String("request_id", requestId),
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
						)
						c.Error(err.(error))
						return
					}

					// 记录 panic 日志
					logger.Error("Recovery from panic",
						zap.String("request_id", requestId),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)

					// 返回 500 错误
					response.NewRespCodeMsg(c, http.StatusInternalServerError, "服务器内部错误")
				}
			}()

			// 调用下一个中间件或路由处理器
			return next(c)
		}
	}
}
