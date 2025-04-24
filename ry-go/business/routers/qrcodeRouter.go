package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// QrcodeRouterInit 二维码控制器路由初始化
func QrcodeRouterInit(group *echo.Group, qrcodeController *controller.QrcodeController) {
	group.Group("/qrcode").POST("/generate", qrcodeController.CodeHandler)
}
