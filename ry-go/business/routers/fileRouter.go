package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

func FileRouterInit(group *echo.Group, fileController *controller.FileController) {
	sysUserRouterGroup := group.Group("/system/file")
	sysUserRouterGroup.POST("/upload", fileController.UploadHandler)
	sysUserRouterGroup.POST("/batchUpload", fileController.UploadFilesHandler)
	sysUserRouterGroup.GET("/download", fileController.DownloadHandler)
	sysUserRouterGroup.POST("/batchDownload", fileController.DownloadFilesHandler)
	sysUserRouterGroup.GET("/batchDownloadStream", fileController.DownloadFilesStreamHandler)
	sysUserRouterGroup.POST("/batchDownloadStream", fileController.DownloadFilesStreamHandler)
}
