package utils

import (
	"encoding/base64"
	"errors"
	"mime"
	"net/http"
	"ry-go/common/response"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 生成验证码
func GenerateCaptcha(e echo.Context) error {
	// 分别是验证码图片高度，验证码图片高度，验证码长度，偏斜系数，背景噪声数量
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 20)
	captcha := base64Captcha.NewCaptcha(driver, store)

	id, base64Str, _, err := captcha.Generate()
	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "生成二维码失败")
		return err
	}

	decodeStr, err := base64.StdEncoding.DecodeString(strings.Split(base64Str, ",")[1])
	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "生成二维码失败")
		return err
	}

	e.Response().Header().Set("captcha-id", id)
	if err = e.Blob(http.StatusOK, mime.TypeByExtension(".png"), decodeStr); err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "获取二维码失败")
		return err
	}
	return nil
}

// VerifyCaptcha 校验验证码
func VerifyCaptcha(e echo.Context) error {
	id := e.FormValue("id")
	code := e.FormValue("code")

	if id == "" || code == "" {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "未找到验证码")
		return errors.New("未找到验证码")
	}

	// 验证码是否过期
	if store.Get(id, false) == "" {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "验证码已过期")
		return errors.New("验证码已过期")
	}

	if store.Verify(id, code, true) {
		response.NewRespCodeMsg(e, http.StatusOK, "验证码正确")
		return nil
	} else {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "验证码错误")
		return errors.New("验证码错误")
	}
}
