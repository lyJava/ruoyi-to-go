package controller

import (
	"net/http"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/response"
	"ry-go/utils"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// FileController 文件控制器
type FileController struct {
	Service service.FileService
}

// FileControllerInit 文件控制器初始化
func FileControllerInit(s *serviceImpl.FileServiceImpl) *FileController {
	return &FileController{
		Service: s,
	}
}

// UploadHandler 文件上传
//
//goland:noinspection GoUnhandledErrorResult
func (c *FileController) UploadHandler(e echo.Context) error {
	data, err := c.Service.Upload(e, "file", time.Now().Format("20060102150405"),
		"http://localhost:8086/system/user/file/download?name=%s")
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	response.NewResponse(e, 200, "上传成功", data)
	return nil
}

// UploadFilesHandler 多文件上传
//
//goland:noinspection GoUnhandledErrorResult
func (c *FileController) UploadFilesHandler(e echo.Context) error {
	data, err := c.Service.UploadFiles(e, "files", time.Now().Format("20060102150405"),
		"http://localhost:8086/system/user/file/download?name=%s")
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	response.NewResponse(e, 200, "上传成功", data)
	return nil
}

// DownloadHandler 文件下载
//
//goland:noinspection GoUnhandledErrorResult
func (c *FileController) DownloadHandler(e echo.Context) error {
	if err := utils.HandlerEchoDownload(e, "name"); err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	return nil
}

// DownloadFilesHandler 多文件下载
//
//goland:noinspection GoUnhandledErrorResult
func (c *FileController) DownloadFilesHandler(e echo.Context) error {
	var list []string
	if err := e.Bind(&list); err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "多文件下载参数转换失败")
		return err
	}

	if err := c.Service.DownloadFiles(e, list); err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	return nil
}

// DownloadFilesStreamHandler 多文件下载(流处理方式)
//
//goland:noinspection GoUnhandledErrorResult
func (c *FileController) DownloadFilesStreamHandler(e echo.Context) error {
	param := e.QueryParam("file")
	if param == "" {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "文件名参数验证失败")
		return nil
	}

	if err := c.Service.DownloadFilesByStream(e, strings.Split(param, ","), cast.ToBool(e.QueryParam("rename"))); err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	return nil
}
