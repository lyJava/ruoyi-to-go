package serviceImpl

import (
	"ry-go/utils"

	"github.com/labstack/echo/v4"
)

type FileServiceImpl struct {
}

func NewFileServiceImpl() *FileServiceImpl {
	return &FileServiceImpl{}
}

func (s *FileServiceImpl) Upload(e echo.Context, name, prefix, downloadUrl string) (map[string]string, error) {
	return utils.HandlerEchoUpload(e, name, prefix, downloadUrl)
}

func (s *FileServiceImpl) UploadFiles(e echo.Context, name, prefix, downloadUrl string) (map[string][]string, error) {
	return utils.HandlerEchoUploadFlies(e, name, prefix, downloadUrl)
}

func (s *FileServiceImpl) Download(e echo.Context, name string) error {
	return utils.HandlerEchoDownload(e, name)
}

func (s *FileServiceImpl) DownloadFiles(e echo.Context, list []string) error {
	return utils.DownloadFiles(e, list)
}

func (s *FileServiceImpl) DownloadFilesByStream(e echo.Context, list []string, rename bool) error {
	return utils.DownloadFilesByStream(e, list, rename)
}
