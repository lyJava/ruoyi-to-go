package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysDictTypeRouterInit 字典类型路由初始化
func SysDictTypeRouterInit(group *echo.Group, sysDictType *controller.SysDictTypeController) {
	sysDictTypeRouterGroup := group.Group("/system/dict/type")
	sysDictTypeRouterGroup.POST("", sysDictType.SaveHandler)
	sysDictTypeRouterGroup.PUT("", sysDictType.UpdateHandler)
	sysDictTypeRouterGroup.POST("/batchSave", sysDictType.BatchSaveHandler)
	sysDictTypeRouterGroup.GET("/list", sysDictType.SelectPageHandler)
	sysDictTypeRouterGroup.GET("/:id", sysDictType.SelectOneHandler)
	sysDictTypeRouterGroup.DELETE("/:id", sysDictType.BatchDeleteHandler)
	sysDictTypeRouterGroup.DELETE("/clearCache", sysDictType.ClearCacheHandler)
	sysDictTypeRouterGroup.GET("/refreshCache", sysDictType.RefreshCacheHandler)
	sysDictTypeRouterGroup.GET("/export", sysDictType.DownloadExcelBufferHandler)
}
