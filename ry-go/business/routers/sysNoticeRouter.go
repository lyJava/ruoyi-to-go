package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysNoticeRouterInit 系统公告路由初始化
func SysNoticeRouterInit(group *echo.Group, notice *controller.SysNoticeController) {
	sysNoticeRouterGroup := group.Group("/system/notice")
	sysNoticeRouterGroup.POST("", notice.SaveHandler)
	sysNoticeRouterGroup.PUT("", notice.UpdateHandler)
	sysNoticeRouterGroup.POST("/batchSave", notice.BatchSaveHandler)
	sysNoticeRouterGroup.GET("/list", notice.SelectPageGetHandler)
	sysNoticeRouterGroup.GET("/:id", notice.SelectOneHandler)
	sysNoticeRouterGroup.DELETE("/:id", notice.BatchDeleteHandler)
}
