package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysDeptRouterInit 系统部门路由初始化
func SysDeptRouterInit(group *echo.Group, dept *controller.DeptController) {
	sysDeptRouterGroup := group.Group("/system/dept")
	sysDeptRouterGroup.GET("/check/exist", dept.CheckDept)
	sysDeptRouterGroup.POST("", dept.SaveHandler)
	sysDeptRouterGroup.PUT("", dept.UpdateDept)
	sysDeptRouterGroup.POST("/batchSave", dept.BatchSaveDept)
	sysDeptRouterGroup.POST("/batchSaveList", dept.BatchSaveDeptList)
	sysDeptRouterGroup.POST("/import", dept.ImportData)
	sysDeptRouterGroup.GET("/list", dept.SelectListHandler)
	sysDeptRouterGroup.GET("/:deptId", dept.SelectByIdHandler)
	sysDeptRouterGroup.GET("/list/exclude/:deptId", dept.ExcludeDetailHandler)
	sysDeptRouterGroup.GET("/page", dept.SelectPageHandler)
	sysDeptRouterGroup.DELETE("/:id", dept.BatchDelete)
	sysDeptRouterGroup.DELETE("/batchDel/:id", dept.BatchDelete)
	sysDeptRouterGroup.GET("/treeSelect", dept.TreeSelectHandler)
	sysDeptRouterGroup.GET("/roleDeptTreeSelect/:roleId", dept.RoleDeptTreeSelectHandler)
	sysDeptRouterGroup.GET("/export", dept.DownloadExcelBufferHandler)
}
