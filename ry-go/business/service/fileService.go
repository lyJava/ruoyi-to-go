package service

import "github.com/labstack/echo/v4"

type FileService interface {
	Upload(e echo.Context, name, prefix, downloadUrl string) (map[string]string, error)
	UploadFiles(e echo.Context, name, prefix, downloadUrl string) (map[string][]string, error)
	Download(e echo.Context, name string) error
	DownloadFiles(e echo.Context, List []string) error
	DownloadFilesByStream(e echo.Context, List []string, rename bool) error
}
