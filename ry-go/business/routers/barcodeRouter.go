package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// BarcodeRouterInit 条形码控制器路由初始化
func BarcodeRouterInit(group *echo.Group, barcodeController *controller.BarcodeController) {
	group.Group("/barcode").GET("/generate", barcodeController.CodeHandler)
}
