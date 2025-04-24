package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysJobLogRouterInit 系统任务日志路由初始化
func SysJobLogRouterInit(group *echo.Group, sysJobLog *controller.SysJobLogController) {
	sysJobLogRouterGroup := group.Group("/monitor/jobLog")
	sysJobLogRouterGroup.POST("", sysJobLog.SaveHandler)
	sysJobLogRouterGroup.PUT("", sysJobLog.UpdateHandler)
	sysJobLogRouterGroup.POST("/batchSave", sysJobLog.BatchSaveHandler)
	sysJobLogRouterGroup.GET("/list", sysJobLog.SelectPageHandler)
	sysJobLogRouterGroup.GET("/:id", sysJobLog.SelectOneHandler)
	sysJobLogRouterGroup.DELETE("/:id", sysJobLog.BatchDeleteHandler)
}
