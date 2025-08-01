package controller

import (
	"errors"
	"net/http"
	"ry-go/common/response"
	"ry-go/utils"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// BarcodeController 条形码控制器
type BarcodeController struct {
}

// BarcodeControllerInit 条形码控制器初始化
func BarcodeControllerInit() *BarcodeController {
	return &BarcodeController{}
}

// CodeHandler 生成二维码处理器
func (controller *BarcodeController) CodeHandler(c echo.Context) error {
	content := c.QueryParam("content")
	width := c.QueryParam("width")
	height := c.QueryParam("height")
	if content == "" {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "条形码参数验证失败")
		return errors.New("条形码参数验证失败")
	}

	if err := utils.GenerateBarcode(c, content, cast.ToInt(width), cast.ToInt(height)); err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	return nil
}
