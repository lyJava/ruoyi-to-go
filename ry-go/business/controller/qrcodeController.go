package controller

import (
	"errors"
	"net/http"
	"ry-go/common/response"
	"ry-go/utils"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// QrcodeController 二维码控制器
type QrcodeController struct {
}

// QrcodeControllerInit 二维码控制器初始化
func QrcodeControllerInit() *QrcodeController {
	return &QrcodeController{}
}

// CodeHandler 生成二维码处理器
func (controller *QrcodeController) CodeHandler(c echo.Context) error {
	content := c.QueryParam("content")
	size := c.QueryParam("size")
	if content == "" || size == "" {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "二维码参数验证失败")
		return errors.New("二维码参数验证失败")
	}

	if err := utils.GenerateQrCodeToResp(c, content, cast.ToInt(size)); err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	return nil
}
