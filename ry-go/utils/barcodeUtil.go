package utils

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/labstack/echo/v4"
)

// GenerateBarcode 生成条形码并直接输出到浏览器
func GenerateBarcode(ctx echo.Context, content string, width, height int) error {
	// 生成条形码
	bc128, err := code128.Encode(content)
	if err != nil {
		return errors.New("生成条形码失败")
	}

	// 将 code128.Barcode 转换为 barcode.Barcode
	barcodeImg := barcode.Barcode(bc128)

	// 调整条形码大小
	scaledBarcode, err := barcode.Scale(barcodeImg, width, height)
	if err != nil {
		return errors.New("调整条形码大小失败：" + err.Error())
	}

	// 将条形码图片编码为 PNG 字节数据
	var buf bytes.Buffer
	if err = png.Encode(&buf, scaledBarcode); err != nil {
		return errors.New("编码条形码失败")
	}

	// 设置响应头
	ctx.Response().Header().Set(echo.HeaderContentType, "image/png")
	ctx.Response().Header().Set(echo.HeaderContentLength, strconv.Itoa(buf.Len()))

	// 将条形码图片写入响应
	/*if err = png.Encode(ctx.Response(), scaledBarcode); err != nil {
		return err
	}*/

	// 将字节数据写入响应
	if _, err = ctx.Response().Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

// cropWhiteBorder 裁剪二维码图片的白边
func cropWhiteBorder(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// 找到二维码的实际边界
	minX, minY, maxX, maxY := width, height, 0, 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r != 0xffff || g != 0xffff || b != 0xffff { // 如果不是白色
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// 裁剪图片
	croppedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(minX, minY, maxX+1, maxY+1))

	return croppedImg
}
