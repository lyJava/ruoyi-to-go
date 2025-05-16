package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysConfigRouterInit 系统配置路由初始化
func SysConfigRouterInit(group *echo.Group, sysConfig *controller.SysConfigController) {
	sysConfigRouterGroup := group.Group("/system/config")
	sysConfigRouterGroup.POST("", sysConfig.SaveHandler)
	sysConfigRouterGroup.PUT("", sysConfig.UpdateHandler)
	sysConfigRouterGroup.POST("/batchSave", sysConfig.BatchSaveHandler)
	sysConfigRouterGroup.GET("/list", sysConfig.SelectPageHandler)
	sysConfigRouterGroup.GET("/:id", sysConfig.SelectOneHandler)
	sysConfigRouterGroup.DELETE("/:id", sysConfig.BatchDeleteHandler)
	sysConfigRouterGroup.GET("/configKey/:key", sysConfig.SelectByKeyHandler)
	sysConfigRouterGroup.GET("/refreshCache", sysConfig.RefreshCacheHandler)
	sysConfigRouterGroup.DELETE("/clearCache", sysConfig.ClearCacheHandler)
}
