package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysPostRouterInit 岗位信息路由初始化
func SysPostRouterInit(group *echo.Group, sysPost *controller.SysPostController) {
	sysPostRouterGroup := group.Group("/system/post")
	sysPostRouterGroup.POST("", sysPost.SaveHandler)
	sysPostRouterGroup.PUT("", sysPost.UpdateHandler)
	sysPostRouterGroup.POST("/batchSave", sysPost.BatchSaveHandler)
	sysPostRouterGroup.GET("/list", sysPost.SelectPageHandler)
	sysPostRouterGroup.GET("/:id", sysPost.SelectOneHandler)
	sysPostRouterGroup.POST("/batchDelete", sysPost.BatchDeleteHandler)
	sysPostRouterGroup.DELETE("/:id", sysPost.BatchDeleteHandler)
	sysPostRouterGroup.GET("/selectAll", sysPost.SelectAllHandler)
	sysPostRouterGroup.GET("/export", sysPost.DownloadExcelBufferHandler)
}
