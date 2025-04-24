package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysJobRouterInit 系统任务路由初始化
func SysJobRouterInit(group *echo.Group, sysJob *controller.SysJobController) {
	sysJobRouterGroup := group.Group("/monitor/job")
	sysJobRouterGroup.POST("", sysJob.SaveHandler)
	sysJobRouterGroup.PUT("", sysJob.UpdateHandler)
	sysJobRouterGroup.POST("/batchSave", sysJob.BatchSaveHandler)
	sysJobRouterGroup.GET("/list", sysJob.SelectPageHandler)
	sysJobRouterGroup.GET("/:id", sysJob.SelectOneHandler)
	sysJobRouterGroup.DELETE("/:id", sysJob.BatchDeleteHandler)
	sysJobRouterGroup.PUT("/changeStatus", sysJob.ChangeStatusHandler)
	sysJobRouterGroup.PUT("/changeStatus/batch", sysJob.ChangeStatusBatchHandler)
}
