package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"ry-go/business/domain"
	"ry-go/common/response"
	"ry-go/utils"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

// RegisterCustomValidator 注册自定义验证器
func RegisterCustomValidator(key string, value govalidator.CustomTypeValidator) {
	if key == "" || value == nil {
		return
	}
	govalidator.CustomTypeTagMap.Set(key, value)
}

// RegisterCustomValidators 注册自定义验证器
func RegisterCustomValidators(validatorMap map[string]govalidator.CustomTypeValidator) {
	if len(validatorMap) == 0 {
		return
	}
	for key, value := range validatorMap {
		RegisterCustomValidator(key, value)
	}
}

// ValidatorStruct 验证结构体
func ValidatorStruct(t any) error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		if validatorErrors, ok := err.(govalidator.Errors); ok {
			var errList []string
			for _, e := range validatorErrors {
				errList = append(errList, e.Error())
			}
			return fmt.Errorf(strings.Join(errList, ";"))
		}
	}
	return nil
}

// UseCustomValidator 单个验证
//
// 参数
//
//	-t: 泛型结构体
//	-validatorMap: 名称
//	-e: echo.Context echo上下文
//
// 返回
//
//	-error: 错误
func UseCustomValidator[T any](t T, validatorMap map[string]govalidator.CustomTypeValidator, e echo.Context) error {
	RegisterCustomValidators(validatorMap)
	if err := ValidatorStruct(t); err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, err.Error())
		return errors.New("校验失败")
	}
	return nil
}

func LoginValidator[T any](t T, validatorMap map[string]govalidator.CustomTypeValidator) (bool, error) {
	RegisterCustomValidators(validatorMap)
	if err := ValidatorStruct(t); err != nil {
		return false, errors.New("校验失败")
	}
	return true, nil
}

// UseCustomValidators 切片验证
//
// 参数
//
//	-list: 泛型切片
//	-validatorMap: 名称
//	-e: echo.Context echo上下文
//	-duplicate: 是否去除重复验证信息
//
// 返回
//
//	-error: 错误
func UseCustomValidators[T any](list []T, validatorMap map[string]govalidator.CustomTypeValidator, e echo.Context, duplicate bool) error {
	if len(list) <= 0 {
		return errors.New("验证数据不能为空")
	}

	RegisterCustomValidators(validatorMap)
	// 错误切片
	var errList []string
	for _, t := range list {
		_, err := govalidator.ValidateStruct(t)
		if err != nil {
			s := err.Error()
			split := strings.Split(s, ";")
			if len(split) <= 1 {
				errList = append(errList, s)
			} else {
				for _, single := range split {
					errList = append(errList, single)
				}
			}
		}
	}

	log.Printf("原始验证错误信息===%+v", errList)

	if len(errList) > 0 {
		if duplicate {
			errList = utils.DeduplicateStrings(errList)
		}
		response.NewRespCodeMsg(e, http.StatusInternalServerError, strings.Join(errList, ";"))
		return errors.New("校验失败")
	}
	return nil
}

func VerifyRoleAllowed(role *domain.SysRole) bool {
	if role == nil {
		return true
	}

	if role.Id != 0 && role.Id == int64(1) {
		return false
	}
	return true
}

// GetSinglePathParam 获取单个restful请求path参数
func GetSinglePathParam(c echo.Context, paramName string) (string, error) {
	value := c.Param(paramName)
	if value == "" {
		return "", errors.New("参数验证失败")
	}

	return value, nil
}
