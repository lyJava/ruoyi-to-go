package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysOperationLogRouterInit 操作日志路由初始化
func SysOperationLogRouterInit(group *echo.Group, sysOperationLog *controller.SysOperationLogController) {
	sysOperationLogRouterGroup := group.Group("/monitor/operlog")
	sysOperationLogRouterGroup.POST("", sysOperationLog.SaveHandler)
	sysOperationLogRouterGroup.PUT("", sysOperationLog.UpdateHandler)
	sysOperationLogRouterGroup.POST("/batchSave", sysOperationLog.BatchSaveHandler)
	sysOperationLogRouterGroup.GET("/list", sysOperationLog.SelectPageHandler)
	sysOperationLogRouterGroup.GET("/:id", sysOperationLog.SelectOneHandler)
	sysOperationLogRouterGroup.DELETE("/:id", sysOperationLog.BatchDeleteHandler)
}
