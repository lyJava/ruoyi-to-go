package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysUserRouterInit 系统用户路由初始化
func SysUserRouterInit(group *echo.Group, userController *controller.UserController) {
	sysUserRouterGroup := group.Group("/system/user")
	sysUserRouterGroup.POST("", userController.SaveHandler)
	sysUserRouterGroup.PUT("", userController.UpdateHandler)
	sysUserRouterGroup.GET("/:id", userController.DetailHandler)
	sysUserRouterGroup.DELETE("/:id", userController.DeleteHandler)
	sysUserRouterGroup.GET("/list", userController.SelectPageHandler)
	sysUserRouterGroup.POST("/batchSave", userController.BatchInsertHandler)
	sysUserRouterGroup.POST("/batchDelete", userController.BatchDeleteHandler)
	sysUserRouterGroup.DELETE("/batchDelete", userController.BatchDeleteHandler)
	sysUserRouterGroup.GET("/export/excel", userController.ExportExcelHandler)
	sysUserRouterGroup.GET("/captcha", userController.CaptchaHandler)
	sysUserRouterGroup.GET("/captcha/verify", userController.VerifyCaptchaHandler)
	sysUserRouterGroup.GET("/captcha/redis", userController.CaptchaRedisHandler)
	sysUserRouterGroup.GET("/captcha/redis/verify", userController.VerifyCaptchaRedisHandler)
	sysUserRouterGroup.GET("/export", userController.DownloadExcelBufferHandler)
}
