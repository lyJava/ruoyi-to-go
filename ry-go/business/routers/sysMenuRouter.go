package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysMenuRouterInit 系统菜单路由初始化
func SysMenuRouterInit(group *echo.Group, sysMenu *controller.SysMenuController) {
	sysMenuRouterGroup := group.Group("/system/menu")
	sysMenuRouterGroup.POST("", sysMenu.SaveHandler)
	sysMenuRouterGroup.PUT("", sysMenu.UpdateHandler)
	sysMenuRouterGroup.POST("/batchSave", sysMenu.BatchSaveHandler)
	sysMenuRouterGroup.GET("/list", sysMenu.SelectTreeTableListHandler)
	sysMenuRouterGroup.GET("/page", sysMenu.SelectPageHandler)
	sysMenuRouterGroup.GET("/:id", sysMenu.SelectOneHandler)
	sysMenuRouterGroup.DELETE("/:id", sysMenu.BatchDeleteHandler)
	sysMenuRouterGroup.GET("/treeSelect", sysMenu.TreeSelectHandler)
	sysMenuRouterGroup.GET("/roleMenuTreeSelect/:roleId", sysMenu.RoleMenuTreeRoleIdHandler)
}
