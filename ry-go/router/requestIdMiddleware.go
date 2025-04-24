package router

import (
	"ry-go/utils"

	"github.com/labstack/echo/v4"
)

// EnsureRequestID 请求ID中间件
func EnsureRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 检查请求头中是否已经存在 X-Request-Id
			requestID := c.Request().Header.Get("X-Request-Id")
			if requestID == "" {
				// 如果不存在，生成一个新的 Request ID
				requestID = utils.ShortUUID()
				// 设置到请求头中
				c.Request().Header.Set("X-Request-Id", requestID)
			}
			// 将requestID设置到响应头中，方便客户端追踪
			c.Response().Header().Set("X-Request-Id", requestID)
			c.SetRequest(c.Request().WithContext(utils.SetValForContext(c.Request().Context(), "requestId", requestID)))

			return next(c)
		}
	}
}
