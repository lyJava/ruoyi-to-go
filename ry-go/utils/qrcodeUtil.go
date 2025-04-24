package utils

import (
	"errors"
	"image/png"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
)

func GenerateQrCode(ctx echo.Context, content string, size int) error {
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}

	qrcodeName := ShortUUID() + ".png"
	qrcodeRoot := filepath.Join(currentPath, "/upload")

	// 检查并创建上传目录
	if err = os.MkdirAll(qrcodeRoot, os.ModePerm); err != nil {
		return errors.New("创建上传目录失败")
	}
	qrcodePath := qrcodeRoot + "/" + qrcodeName

	if err = qrcode.WriteFile(content, qrcode.High, size, qrcodePath); err != nil {
		return errors.New("生成二维码失败")
	}

	// 设置响应头
	ctx.Response().Header().Set(echo.HeaderContentType, "image/png")

	if err = ctx.File(qrcodePath); err != nil {
		return err
	}

	// 删除临时文件
	if err = os.Remove(qrcodePath); err != nil {
		return err
	}
	return nil
}

// GenerateQrCodeToResp 直接生成二维码输出到响应流，无需保留二维码图片
func GenerateQrCodeToResp(ctx echo.Context, content string, size int) error {
	// 生成二维码
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return errors.New("生成二维码失败")
	}
	// 将二维码编码为 PNG 图片
	img := qr.Image(size)

	// 裁剪二维码图片（去掉白边）
	//croppedImg := cropWhiteBorder(img)

	// 设置响应头
	ctx.Response().Header().Set(echo.HeaderContentType, "image/png")

	// 将二维码图片写入响应
	if err = png.Encode(ctx.Response(), img); err != nil {
		return err
	}

	return nil
}
