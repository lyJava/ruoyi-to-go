package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysDictDataRouterInit 字典数据路由初始化
func SysDictDataRouterInit(group *echo.Group, sysDictData *controller.SysDictDataController) {
	sysDictDataRouterGroup := group.Group("/system/dict/data")
	sysDictDataRouterGroup.POST("", sysDictData.SaveHandler)
	sysDictDataRouterGroup.PUT("", sysDictData.UpdateHandler)
	sysDictDataRouterGroup.POST("/batchSave", sysDictData.BatchSaveHandler)
	sysDictDataRouterGroup.GET("/list", sysDictData.SelectPageHandler)
	sysDictDataRouterGroup.GET("/:id", sysDictData.SelectOneHandler)
	sysDictDataRouterGroup.DELETE("/:id", sysDictData.BatchDeleteHandler)
	sysDictDataRouterGroup.GET("/index/:dictType", sysDictData.SelectDictDataByTypeHandler)
	sysDictDataRouterGroup.GET("/type/:dictType", sysDictData.SelectDictDataByTypeHandler)
}
