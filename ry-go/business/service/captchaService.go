package service

import (
	"github.com/labstack/echo/v4"
)

type CaptchaService interface {
	// GenerateCaptcha 生成验证码
	//
	// 参数
	//
	// 	-e: echo上下文对象
	//
	// 返回
	//
	//  -error 错误信息
	GenerateCaptcha(e echo.Context) error

	// VerifyCaptcha 校验验证码
	//
	// 参数
	//
	//  -e: echo上下文对象
	//
	// 返回
	//
	//  -error 错误信息
	VerifyCaptcha(e echo.Context) error

	// GenerateLoginCaptcha 生成登陆验证码
	//
	// 参数
	//
	//  -e: echo上下文对象
	//
	// 返回
	//
	//  -error 错误信息
	//  -map[string]string 验证码信息
	GenerateLoginCaptcha(e echo.Context) error

	// VerifyLoginCaptcha 校验登陆验证码
	//
	// 参数
	//
	//  -e: echo上下文对象
	//  -code: 验证码
	//  -uuid: 验证码ID
	//
	// 返回
	//
	//  -error 错误信息
	VerifyLoginCaptcha(e echo.Context, code, uuid string) error
}
