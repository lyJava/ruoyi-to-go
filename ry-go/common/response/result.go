package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Result 响应结果
type Result struct {
	Code    int    `json:"code"`              // 响应码
	Message string `json:"message,omitempty"` // 响应信息
	Data    any    `json:"data,omitempty"`    // 响应数据
}

// NewResult 创建响应结果
func NewResult(code int, message string, data any) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewMsgResult 创建响应结果
//
//goland:noinspection GoUnusedExportedFunction
func NewMsgResult(message string, data any) *Result {
	return &Result{
		Message: message,
		Data:    data,
	}
}

// NewCodeMsg 创建响应结果
//
//goland:noinspection GoUnusedExportedFunction
func NewCodeMsg(code int, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
	}
}

// NewMsgResultNoData 创建响应结果
func NewMsgResultNoData(message string, data any) *Result {
	return &Result{
		Message: message,
		Data:    data,
	}
}

// NewMsgResultNoMsg 创建响应结果(无message，只有code与data)
func NewMsgResultNoMsg(code int, data any) *Result {
	return &Result{
		Code: code,
		Data: data,
	}
}

// NewRespNoMsg 创建echo响应结果
//
//goland:noinspection GoUnhandledErrorResult
func NewRespNoMsg(e echo.Context, code int, data any) {
	SetCharset(e, "")
	e.JSON(http.StatusOK, NewMsgResultNoMsg(code, data))
}

// NewResponse 创建echo响应结果
//
//goland:noinspection GoUnhandledErrorResult
func NewResponse(e echo.Context, code int, msg string, data any) {
	SetCharset(e, "")
	e.JSON(http.StatusOK, NewResult(code, msg, data))
}

// NewRespCodeMsg 创建echo响应结果
//
// 参数
//
//	-e: echo上下文
//	-code: 响应码
//	-message: 响应信息
//
//goland:noinspection GoUnhandledErrorResult
func NewRespCodeMsg(e echo.Context, code int, msg string) {
	SetCharset(e, "")
	e.JSON(http.StatusOK, NewCodeMsg(code, msg))
}

// NewRespCodeErr 创建echo响应结果
//
// 参数
//
//	-e: echo上下文
//	-code: 响应码
//	-err: 错误接口
//
//goland:noinspection GoUnhandledErrorResult
func NewRespCodeErr(e echo.Context, code int, err error) {
	SetCharset(e, "")
	e.JSON(http.StatusOK, NewCodeMsg(code, err.Error()))
}

// PageData 分页数据结构体
type PageData[T any] struct {
	Content []T   `json:"content"`        // 数据集合
	Records int64 `json:"records,string"` // 总条数(返回字符串)
	Pages   int64 `json:"pages,string"`   // 总页数(返回字符串)
}

// NewPageData 创建分页数据结构体
func NewPageData[T any](data []T, records, pages int64) PageData[T] {
	return PageData[T]{
		Content: data,
		Records: records,
		Pages:   pages,
	}
}

// SetCharset 设置返回json字符集，默认UTF-8
func SetCharset(e echo.Context, charset string) {
	if charset == "" {
		charset = "application/json;charset=UTF-8"
	}
	e.Response().Header().Set("Content-Type", charset)
}

//goland:noinspection GoUnhandledErrorResult
type LoginResult struct {
	Code    int    `json:"code"`            // 响应码
	Message string `json:"msg,omitempty"`   // 响应信息
	Token   string `json:"token,omitempty"` // 响应凭证
}

func NewLoginResult(e echo.Context, code int, msg string, token string) {
	SetCharset(e, "")
	e.JSON(http.StatusOK, &LoginResult{
		Code:    code,
		Message: msg,
		Token:   token,
	})
}

type UserInfo struct {
	User        any      `json:"user,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

func NewLoginInfoResult[T any](e echo.Context, data T) {
	SetCharset(e, "")
	e.JSON(http.StatusOK, data)
}
