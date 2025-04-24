package utils

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"ry-go/common/response"
	"strconv"
	"strings"
)

// OpenFile 打开文件
//
// 参数
//
//	-fileName (string): 文件名
func OpenFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		zap.L().Sugar().Errorf("Open File error===%+v", err)
		return &os.File{}, err
	}
	return file, nil
}

// GetFileInfo 获取文件信息
//
// 参数
//
//	-file(*os.File):文件对象
//
// 返回
//
//	-os.FileInfo: 文件信息
//	-error: 错误
func GetFileInfo(file *os.File) (os.FileInfo, error) {
	// 获取临时文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		zap.L().Sugar().Errorf("Get File Info error===%+v", err)
		return nil, err
	}

	contentLength := strconv.FormatInt(fileInfo.Size(), 10)
	zap.L().Sugar().Infof("文件长度===%s", contentLength)
	fileName := filepath.Base(fileInfo.Name())
	zap.L().Sugar().Infof("文件名称===%s", fileName)
	return fileInfo, nil
}

// HandlerEchoUpload 处理echo上传
// 参数
//
//	-ctx: echo上下文
//	-name: 表单文件参数名称
//	-prefix: 文件名前缀
//	-downloadUrl: 文件下载链接
//
// 返回
//
//	-string: 上传文件路径
//	-error: 上传失败错误
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func HandlerEchoUpload(ctx echo.Context, name, prefix, downloadUrl string) (map[string]string, error) {
	file, err := ctx.FormFile(name)
	if err != nil {
		zap.L().Sugar().Errorf("未获取到上传文件:%+v", err)
		return nil, errors.New("未获取到上传信息")
	}

	// 缓存文件名
	filename := file.Filename
	fileSize := file.Size

	zap.L().Sugar().Infof("上传文件名称===%s", filename)
	zap.L().Sugar().Infof("上传文件大小===%s", FormatFileSize(fileSize))

	// 定义上传目录和文件路径
	uploadDir := "./upload"
	filePath := filepath.Join(uploadDir, prefix+"_"+filename)

	ext := filepath.Ext(filename)
	zap.L().Sugar().Infof("获取上传文件后缀:%s", ext)

	currentPath, err := os.Getwd()
	if err != nil {
		zap.L().Sugar().Errorf("获取当前目录错误===%+v", err)
		return nil, err
	}
	zap.L().Sugar().Infof("获取当前工作目录路径===%s", currentPath)

	zap.L().Sugar().Infof("获取当前文件filePath===%s", filePath)

	// 检查并创建上传目录
	if err = os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		zap.L().Sugar().Errorf("创建上传目录失败:%+v", err)
		return nil, errors.New("文件上传失败")
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		zap.L().Sugar().Errorf("打开上传文件失败:%+v", err)
		return nil, errors.New("文件上传失败")
	}
	defer src.Close()

	if err = SaveFile(src, filePath); err != nil {
		return nil, errors.New("文件上传失败")
	}

	_, err = os.Stat(filePath)
	if err != nil {
		zap.L().Sugar().Errorf("获取文件信息出错:%+v", err)
		return nil, errors.New("文件上传失败")
	}

	// 获取文件的绝对路径
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		zap.L().Sugar().Errorf("获取文件绝对路径出错:%+v", err)
		return nil, errors.New("文件上传失败")
	}

	// 记录上传成功日志
	/*zap.L().Info("文件上传成功",
		zap.String("文件名", fileInfo.Name()),
		zap.String("相对路径", filePath),
		zap.String("绝对路径", absoluteFilePath),
	)*/

	fileName := filepath.Base(filePath)
	var dataMap map[string]string
	if downloadUrl == "" {
		dataMap = map[string]string{
			"filePath": fileName,
			"absPath":  absPath,
		}
	} else {
		dataMap = map[string]string{
			"name":        fileName,
			"downloadUrl": fmt.Sprintf(downloadUrl, url.QueryEscape(fileName)),
		}
	}
	return dataMap, nil
}

func HandlerEchoUploadFlies(ctx echo.Context, name, prefix, downloadUrl string) (map[string][]string, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, errors.New("多文件上传失败")
	}

	// 获取所有文件
	files := form.File[name] // name前端上传文件的字段名
	fileCount := len(files)

	var dataMap map[string][]string

	if fileCount > 0 {

		fileNameList := make([]string, 0, fileCount)
		absPathList := make([]string, 0, fileCount)
		downloadUrlList := make([]string, 0, fileCount)

		for _, file := range files {

			// 打开上传的文件
			src, err := file.Open()
			if err != nil {
				zap.L().Sugar().Errorf("打开上传文件失败:%+v", err)
				return nil, errors.New("文件上传失败")
			}
			func(src multipart.File) {
				err := src.Close()
				if err != nil {
					zap.L().Sugar().Errorf("关闭上传文件失败:%+v", err)
				}
			}(src)

			// 缓存文件名
			filename := file.Filename
			fileSize := file.Size

			zap.L().Sugar().Infof("上传文件名称===%s", filename)
			zap.L().Sugar().Infof("上传文件大小===%s", FormatFileSize(fileSize))

			// 定义上传目录和文件路径
			uploadDir := "./upload"
			filePath := filepath.Join(uploadDir, prefix+"_"+filename)

			ext := filepath.Ext(filename)
			zap.L().Sugar().Infof("获取上传文件后缀:%s", ext)

			currentPath, err := os.Getwd()
			if err != nil {
				zap.L().Sugar().Errorf("获取当前目录错误===%+v", err)
				return nil, err
			}
			zap.L().Sugar().Infof("获取当前工作目录路径===%s", currentPath)

			zap.L().Sugar().Infof("获取当前文件filePath===%s", filePath)

			// 检查并创建上传目录
			if err = os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				zap.L().Sugar().Errorf("创建上传目录失败:%+v", err)
				return nil, errors.New("文件上传失败")
			}

			if err = SaveFile(src, filePath); err != nil {
				return nil, errors.New("文件上传失败")
			}

			fileInfo, err := os.Stat(filePath)
			if err != nil {
				zap.L().Sugar().Errorf("获取文件信息出错:%+v", err)
				return nil, errors.New("文件上传失败")
			}

			// 获取文件的绝对路径
			absoluteFilePath, err := filepath.Abs(filePath)
			if err != nil {
				zap.L().Sugar().Errorf("获取文件绝对路径出错:%+v", err)
				return nil, errors.New("文件上传失败")
			}
			fileName := filepath.Base(filePath)
			log.Printf("文件名===%s", fileName)
			log.Printf("下载路径===%s", fmt.Sprintf(downloadUrl, url.QueryEscape(fileName)))

			fileNameList = append(fileNameList, fileName)
			absPathList = append(absPathList, absoluteFilePath)
			downloadUrlList = append(downloadUrlList, fmt.Sprintf(downloadUrl, url.QueryEscape(fileName)))
			// 记录上传成功日志
			zap.L().Info("文件上传成功",
				zap.String("文件名", fileInfo.Name()),
				zap.String("相对路径", filePath),
				zap.String("绝对路径", absoluteFilePath),
			)
		}
		if downloadUrl == "" {
			dataMap = map[string][]string{
				"filePath": fileNameList,
				"absPath":  absPathList,
			}
		} else {
			dataMap = map[string][]string{
				"name":        fileNameList,
				"downloadUrl": downloadUrlList,
			}
		}
	}

	log.Printf("dataMap===%v", dataMap)

	return dataMap, nil
}

// HandlerEchoDownload 处理echo文件下载
// 参数
//
//	-ctx: echo上下文
//	-name: 文件名
//
// 返回
//
//	-error: 下载错误
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func HandlerEchoDownload(ctx echo.Context, name string) error {
	fileName := ctx.QueryParam(name)
	if fileName == "" {
		return errors.New("下载文件名不能为空")
	}

	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}

	filePath := filepath.Join(currentPath+"/upload", fileName)

	// 检查文件是否存在
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		zap.L().Sugar().Errorf("文件不存在: %s，error===%+v", filePath, err)
		return errors.New("文件不存在")
	}

	if err = ctx.Attachment(filePath, NormalUUID()+filepath.Ext(fileName)); err != nil {
		zap.L().Sugar().Errorf("echo文件下载错误===%+v", err)
		return errors.New("文件下载失败")
	}

	return nil
}

func GetCurrentPath() (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("获取当前目录错误===%w", err)
	}
	return currentPath, nil
}

// checkFileExists 检测文件是否存在
func checkFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil // 文件存在
	}
	if os.IsNotExist(err) {
		return false, nil // 文件不存在
	}
	return false, err // 其他错误
}

// DownloadFilesByStream echo多文件下载(流式方式)
//
// 参数
//
// -ctx: echo上下文
// -nameList: 文件名切片
// -rename: zip文件是否重命名
//
// 返回
// -error: 错误
//
//goland:noinspection GoUnhandledErrorResult
func DownloadFilesByStream(ctx echo.Context, nameList []string, rename bool) error {
	if len(nameList) == 0 {
		return errors.New("下载文件名不能为空")
	}

	// 获取当前工作目录
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	// 创建管道
	pr, pw := io.Pipe()
	// 创建 ZIP 写入器
	zipWriter := zip.NewWriter(pw)

	// 启动 Goroutine 写入 ZIP 文件
	go func() {
		defer pw.Close()
		defer zipWriter.Close()

		// 遍历文件名列表，将每个文件添加到 ZIP 中
		for _, fileName := range nameList {
			filePath := filepath.Join(currentPath+"/upload", fileName)

			// 检查文件是否存在
			if _, err = os.Stat(filePath); os.IsNotExist(err) {
				pw.CloseWithError(fmt.Errorf("文件 %s 不存在", fileName))
				return
			}

			// 打开文件
			file, err := os.Open(filePath)
			if err != nil {
				pw.CloseWithError(fmt.Errorf("打开文件 %s 失败===%+v", filePath, err))
				return
			}
			defer file.Close()

			// 获取文件信息
			fileInfo, err := file.Stat()
			if err != nil {
				pw.CloseWithError(fmt.Errorf("获取文件 %s 信息失败===%+v", filePath, err))
				return
			}

			// 创建 ZIP 文件头
			header, err := zip.FileInfoHeader(fileInfo)
			if err != nil {
				pw.CloseWithError(fmt.Errorf("创建 ZIP 文件头失败: %s===%+v", filePath, err))
				return
			}

			if rename {
				// 生成 UUID 作为文件名
				uuidFileName := NormalUUID()
				// 保留文件扩展名
				ext := filepath.Ext(fileName) // 获取文件扩展名
				if ext != "" {
					uuidFileName += ext // 将扩展名附加到 UUID 后
				}
				// 设置 ZIP 文件头的文件名
				header.Name = uuidFileName
			} else {
				// 设置 ZIP 文件头的文件名
				header.Name = fileName
			}

			// 创建 ZIP 文件写入器
			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				pw.CloseWithError(fmt.Errorf("创建 ZIP 文件写入器失败: %s===%+v", header.Name, err))
				return
			}

			// 将文件内容写入 ZIP
			if _, err = io.Copy(writer, file); err != nil {
				pw.CloseWithError(fmt.Errorf("写入 ZIP 文件失败: %s===%+v", header.Name, err))
				return
			}
		}
	}()

	zipName := ShortUUID() + ".zip"
	// 设置响应头
	ctx.Response().Header().Set(echo.HeaderContentType, "application/zip")
	ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+zipName)

	// 将 ZIP 文件内容写入响应
	if _, err = io.Copy(ctx.Response(), pr); err != nil {
		return errors.New("写入响应失败")
	}

	return nil
}

// DownloadFiles echo多文件下载(直接写入磁盘方式)
//
// 参数
//
// -ctx: echo上下文
// -nameList: 文件名切片
//
// 返回
// -error: 错误
//
//goland:noinspection GoUnhandledErrorResult
func DownloadFiles(ctx echo.Context, nameList []string) error {
	if len(nameList) == 0 {
		return errors.New("下载文件名不能为空")
	}

	// 获取当前工作目录
	currentPath, err := os.Getwd()
	if err != nil {
		zap.L().Error("获取当前目录错误", zap.Error(err))
		return err
	}

	zipName := ShortUUID() + ".zip"
	// 创建临时 ZIP 文件
	zipFilePath := filepath.Join(currentPath, zipName)
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		zap.L().Error("创建 ZIP 文件失败", zap.Error(err))
		return errors.New("创建 ZIP 文件失败")
	}
	defer func(zipFile *os.File) {
		if err := zipFile.Close(); err != nil {
			zap.L().Error("关闭ZI文件", zap.Error(err))
		}
	}(zipFile)

	// 创建 ZIP 写入器
	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		if err = zipWriter.Close(); err != nil {
			zap.L().Error("关闭ZIP写入器%+v失败===%+v", zap.Error(err))
		}
	}(zipWriter)

	// 遍历文件名列表，将每个文件添加到 ZIP 中
	for _, fileName := range nameList {
		filePath := filepath.Join(currentPath+"/upload", fileName)

		// 检查文件是否存在
		if _, err = os.Stat(filePath); os.IsNotExist(err) {
			zap.L().Error("文件不存在", zap.String("filePath", filePath), zap.Error(err))
			return fmt.Errorf("文件:%s 不存在", fileName)
		}

		// 打开文件
		file, err := os.Open(filePath)
		if err != nil {
			zap.L().Error("打开文件失败", zap.String("filePath", filePath), zap.Error(err))
			return fmt.Errorf("打开文件 %s 失败", fileName)
		}

		defer file.Close()

		// 获取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			zap.L().Error("获取文件信息失败", zap.String("filePath", filePath), zap.Error(err))
			return fmt.Errorf("获取文件 %s 信息失败", fileName)
		}

		// 创建 ZIP 文件头
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			zap.L().Error("创建 ZIP 文件头失败", zap.String("filePath", filePath), zap.Error(err))
			return fmt.Errorf("创建 ZIP 文件头失败: %s", fileName)
		}

		// 设置 ZIP 文件头的文件名
		header.Name = fileName

		// 创建 ZIP 文件写入器
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			zap.L().Error("创建 ZIP 文件写入器失败", zap.String("filePath", filePath), zap.Error(err))
			return fmt.Errorf("创建 ZIP 文件写入器失败: %s", fileName)
		}

		// 将文件内容写入 ZIP
		if _, err = io.Copy(writer, file); err != nil {
			zap.L().Error("写入 ZIP 文件失败", zap.String("filePath", filePath), zap.Error(err))
			return fmt.Errorf("写入 ZIP 文件失败: %s", fileName)
		}
	}

	// 打开 ZIP 文件
	zipFileToSend, err := os.Open(zipFilePath)
	if err != nil {
		return errors.New("打开ZIP文件失败")
	}
	defer func(zipFileToSend *os.File) {
		if err = zipFileToSend.Close(); err != nil {
			zap.L().Error("写入 ZIP 文件失败", zap.String("filePath", zipFilePath), zap.Error(err))
		}
	}(zipFileToSend)

	/*// 设置响应头
	//ctx.Response().Header().Set(echo.HeaderContentType, "application/zip")
	//ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+zipName)

	// 将 ZIP 文件内容写入响应
	if _, err = io.Copy(ctx.Response(), zipFileToSend); err != nil {
		zap.L().Error("写入响应失败", zap.Error(err))
		return errors.New("写入响应失败")
	}*/

	// 将 ZIP 文件内容写入响应
	//http.ServeFile(ctx.Response(), ctx.Request(), zipFilePath)
	if err = ctx.Attachment(zipFilePath, zipName); err != nil {
		return err
	}

	// 删除临时 ZIP 文件
	if err = os.Remove(zipFilePath); err != nil {
		zap.L().Sugar().Errorf("删除临时 ZIP 文件失败===%+v", err)
	}

	return nil
}

// SaveUploadedFile 保存上传的文件到指定路径
//
//goland:noinspection GoUnusedFunction,GoUnhandledErrorResult,GoUnusedExportedFunction
func SaveUploadedFile(fileHeader *multipart.FileHeader, dst string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, src); err != nil {
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	return nil
}

// SaveFile 保存上传的文件到指定路径
//
//goland:noinspection GoUnusedFunction,GoUnhandledErrorResult,GoUnusedExportedFunction
func SaveFile(src io.Reader, targetPath string) error {
	dstFile, err := os.Create(filepath.Clean(targetPath))
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, src); err != nil {
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	return nil
}

// SaveUploadFile 保存上传文件
//
//goland:noinspection GoUnusedFunction,GoUnhandledErrorResult,GoUnusedExportedFunction
func SaveUploadFile(fileHeader *multipart.FileHeader, dst string, maxFileSize int64) error {
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dstFile.Close()

	if _, err := io.CopyN(dstFile, src, maxFileSize); err != nil {
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	return nil
}

// SaveUploadBigFile 保存上传的大文件
//
//goland:noinspection GoUnusedFunction,GoUnhandledErrorResult,GoUnusedExportedFunction
func SaveUploadBigFile(fileHeader *multipart.FileHeader, dst string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1*1024*1024) // 1MB 缓冲区
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("读取文件失败: %w", err)
		}
		if n == 0 {
			break
		}

		if _, err := dstFile.Write(buf[:n]); err != nil {
			return fmt.Errorf("写入文件失败: %w", err)
		}
	}

	return nil
}

// saveUploadedFileAsync 保存上传大文件，使用携程
//
//goland:noinspection GoUnusedFunction,GoUnhandledErrorResult
func saveUploadedFileAsync(fileHeader *multipart.FileHeader, dst string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dstFile.Close()

	// 使用 Goroutine 异步复制文件
	done := make(chan error)
	go func() {
		_, err := io.Copy(dstFile, src)
		done <- err
	}()

	// 等待复制完成
	if err := <-done; err != nil {
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	return nil
}

// DownloadFile 下载文件
//
// 参数
//
//	-fileNamePath (string): 文件名称路径
//	-e (echo.Context): echo请求上下文
//
// 返回
//
//	-error: 错误
//
//goland:noinspection GoUnhandledErrorResult
func DownloadFile(fileNamePath string, e echo.Context) error {
	file, err := OpenFile(fileNamePath)
	if err != nil {
		zap.L().Sugar().Errorf("Download File error===%+v", err)
		response.NewResponse(e, http.StatusInternalServerError, "文件下载异常", nil)
		return err
	}

	defer file.Close()

	fileInfo, err := GetFileInfo(file)
	if err != nil {
		zap.L().Sugar().Errorf("Download File error===%+v", err)
		response.NewResponse(e, http.StatusInternalServerError, "文件下载异常", nil)
		return err
	}

	var fileMineType string

	// 获取文件扩展名
	suffix := strings.ToLower(filepath.Ext(fileInfo.Name()))
	// 检查文件扩展名是否为.xlsx 或.xls
	if suffix == ".xlsx" {
		// excel 2007格式
		fileMineType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	} else if suffix == ".xls" {
		fileMineType = "application/vnd.ms-excel"
	} else {
		// 读取文件的前512字节
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			zap.L().Sugar().Errorf("Read File error===%+v", err)
			response.NewResponse(e, http.StatusInternalServerError, "读取文件失败", nil)
			return err
		}
		fileMineType = http.DetectContentType(buffer)
	}
	zap.L().Sugar().Infof("当前文件MIME==%s", fileMineType)
	respHeader := e.Response().Header()
	respHeader.Set("Content-Type", fileMineType)
	// 对文件名进行编码处理，避免在firefox或postman中请求下载变成response.bin，filepath.Base从路径中获取文件名
	respHeader.Set("Content-Disposition", "attachment; filename*=utf-8''"+url.QueryEscape(filepath.Base(fileNamePath)))
	respHeader.Set("Content-Transfer-Encoding", "binary")
	/*if err = e.File(fileNamePath); err != nil {
		zap.L().Sugar().Errorf("文件下载异常==%+v", err)
		return err
	}*/

	// 使用流式传输文件
	http.ServeContent(e.Response(), e.Request(), fileInfo.Name(), fileInfo.ModTime(), file)
	// 导入成功后删除excel文件
	defer func(name string) {
		if err = os.Remove(name); err != nil {
			zap.L().Sugar().Errorf("删除===%s出错:%+v", fileNamePath, err)
		}
		zap.L().Sugar().Infof("导出后成功删除===%s", fileNamePath)
	}(fileNamePath)
	return nil
}
