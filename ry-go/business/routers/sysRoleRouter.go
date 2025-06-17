package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysRoleRouterInit 系统角色路由初始化
func SysRoleRouterInit(group *echo.Group, sysRole *controller.SysRoleController) {
	sysRoleRouterGroup := group.Group("/system/role")
	sysRoleRouterGroup.POST("", sysRole.SaveHandler)
	sysRoleRouterGroup.PUT("", sysRole.UpdateHandler)
	sysRoleRouterGroup.POST("/batchSave", sysRole.BatchSaveHandler)
	sysRoleRouterGroup.GET("/list", sysRole.SelectPageHandler)
	sysRoleRouterGroup.GET("/:id", sysRole.SelectOneHandler)
	sysRoleRouterGroup.POST("/batchDelete", sysRole.BatchDeleteHandler)
	sysRoleRouterGroup.DELETE("/:id", sysRole.BatchDeleteHandler)
	sysRoleRouterGroup.POST("/list", sysRole.SelectListHandler)
	sysRoleRouterGroup.GET("/list/user", sysRole.SelectIdListHandler)
	sysRoleRouterGroup.GET("/selectAll", sysRole.SelectAllHandler)
	sysRoleRouterGroup.PUT("/dataScope", sysRole.DataScopeHandler)
	sysRoleRouterGroup.GET("/export", sysRole.DownloadExcelBufferHandler)
}
