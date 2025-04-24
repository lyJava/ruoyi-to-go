package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysUserRoleRouterInit 用户角色信息路由初始化
func SysUserRoleRouterInit(group *echo.Group, sysUserRole *controller.SysUserRoleController) {
	sysUserRoleRouterGroup := group.Group("/system/userRole")
	sysUserRoleRouterGroup.POST("/save", sysUserRole.SaveHandler)
	sysUserRoleRouterGroup.PUT("/update", sysUserRole.UpdateHandler)
	sysUserRoleRouterGroup.POST("/batchSave", sysUserRole.BatchSaveHandler)
	sysUserRoleRouterGroup.POST("/page", sysUserRole.SelectPageHandler)
	sysUserRoleRouterGroup.GET("/:id", sysUserRole.SelectOneHandler)
	sysUserRoleRouterGroup.POST("/batchDelete", sysUserRole.BatchDeleteHandler)
	sysUserRoleRouterGroup.DELETE("/batchDelete", sysUserRole.BatchDeleteHandler)
}
