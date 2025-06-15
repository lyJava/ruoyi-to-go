package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysLoginRecordRouterInit 登录记录路由初始化
func SysLoginRecordRouterInit(group *echo.Group, sysLoginRecord *controller.SysLoginRecordController) {
	sysLoginRecordRouterGroup := group.Group("/monitor/logininfor")
	sysLoginRecordRouterGroup.POST("/save", sysLoginRecord.SaveHandler)
	sysLoginRecordRouterGroup.PUT("/update", sysLoginRecord.UpdateHandler)
	sysLoginRecordRouterGroup.POST("/batchSave", sysLoginRecord.BatchSaveHandler)
	sysLoginRecordRouterGroup.GET("/list", sysLoginRecord.SelectPageHandler)
	sysLoginRecordRouterGroup.GET("/:id", sysLoginRecord.SelectOneHandler)
	sysLoginRecordRouterGroup.POST("/batchDelete", sysLoginRecord.BatchDeleteHandler)
	sysLoginRecordRouterGroup.DELETE("/:id", sysLoginRecord.BatchDeleteHandler)
	sysLoginRecordRouterGroup.GET("/lastLogin", sysLoginRecord.SelectLastLoginRecodeHandler)
	sysLoginRecordRouterGroup.GET("/export", sysLoginRecord.DownloadExcelBufferHandler)
}
