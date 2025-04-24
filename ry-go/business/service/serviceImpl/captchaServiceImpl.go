package serviceImpl

import (
	"encoding/base64"
	"errors"
	"mime"
	"net/http"
	"ry-go/common/response"
	"ry-go/utils"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/mojocn/base64Captcha"
)

type CaptchaServiceImpl struct {
	redisClient *redis.Client
	redisStore  *utils.RedisStore
}

func NewCaptchaServiceImpl(client *redis.Client, store *utils.RedisStore) *CaptchaServiceImpl {
	return &CaptchaServiceImpl{
		redisClient: client,
		redisStore:  store,
	}
}

func (s *CaptchaServiceImpl) GenerateCaptcha(e echo.Context) error {
	// 分别是验证码图片高度，验证码图片高度，验证码长度，偏斜系数，背景噪声数量
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 20)
	captcha := base64Captcha.NewCaptcha(driver, s.redisStore)

	id, base64Str, _, err := captcha.Generate()

	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "生成验证码失败")
		return err
	}

	decodeStr, err := base64.StdEncoding.DecodeString(strings.Split(base64Str, ",")[1])
	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "生成验证码失败")
		return err
	}

	e.Response().Header().Set("captcha-id", id)
	if err = e.Blob(http.StatusOK, mime.TypeByExtension(".png"), decodeStr); err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "获取验证码失败")
		return err
	}
	return nil
}

func (s *CaptchaServiceImpl) VerifyCaptcha(e echo.Context) error {
	id := e.FormValue("id")
	code := e.FormValue("code")

	if id == "" || code == "" {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "未找到验证码")
		return errors.New("未找到验证码")
	}

	return validatorCaptcha(e, code, id, s)
}

func (s *CaptchaServiceImpl) GenerateLoginCaptcha(e echo.Context) error {
	// 分别是验证码图片高度，验证码图片高度，验证码长度，偏斜系数，背景噪声数量
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 20)
	captcha := base64Captcha.NewCaptcha(driver, s.redisStore)

	id, base64Str, _, err := captcha.Generate()

	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "生成验证码失败")
		return err
	}

	response.NewResponse(e, http.StatusOK, "生成成功", map[string]string{
		"img":  base64Str,
		"uuid": id,
	})
	return nil
}

func (s *CaptchaServiceImpl) VerifyLoginCaptcha(e echo.Context, code, uuid string) error {
	return validatorCaptcha(e, code, uuid, s)
}

func validatorCaptcha(e echo.Context, code string, uuid string, s *CaptchaServiceImpl) error {
	result, _ := utils.Get(e.Request().Context(), s.redisClient, s.redisStore.Prefix+uuid)
	if result == "" {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "验证码过期了")
		return errors.New("验证码过期了")
	}

	verify := s.redisStore.Verify(uuid, code, true)
	if verify == false {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "验证码错误")
		return errors.New("验证码错误")
	}
	return nil
}
