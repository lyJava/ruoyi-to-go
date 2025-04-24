package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"net/http"
	"ry-go/business/domain"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/request"
	"ry-go/common/response"
)

// SysOperationLogController 操作日志控制器
type SysOperationLogController struct {
	service service.SysOperationLogService
}

// SysOperationLogControllerInit 操作日志控制器初始化
func SysOperationLogControllerInit(s *serviceImpl.SysOperationLogServiceImpl) *SysOperationLogController {
	return &SysOperationLogController{
		service: s,
	}
}

// SaveHandler 保存
func (controller *SysOperationLogController) SaveHandler(c echo.Context) error {
	var sysOperationLog *domain.SysOperationLog
	if err := c.Bind(&sysOperationLog); err != nil {
		response.NewResponse(c, 500, "操作日志新增参数解析错误", nil)
		return err
	}

	result, err := controller.service.Insert(c, sysOperationLog)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	response.NewResponse(c, 200, "操作日志信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysOperationLogController) UpdateHandler(c echo.Context) error {
	var sysOperationLog *domain.SysOperationLog
	if err := c.Bind(&sysOperationLog); err != nil {
		response.NewResponse(c, 500, "操作日志更新参数解析错误", nil)
		return err
	}

	if sysOperationLog.Id == 0 {
		response.NewResponse(c, 500, "操作日志ID不能为空", nil)
		return errors.New("操作日志ID不能为空")
	}

	result, err := controller.service.Update(c, sysOperationLog)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	response.NewResponse(c, 200, "操作日志信息更新成功", result)
	return nil
}

func (controller *SysOperationLogController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysOperationLog
	if err := c.Bind(&list); err != nil {
		response.NewResponse(c, 500, "操作日志批量新增参数解析错误", err)
		return err
	}

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	response.NewResponse(c, 200, "操作日志批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysOperationLogController) SelectPageHandler(c echo.Context) error {
	param := new(request.SysNoticePageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "操作日志分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, 200, "操作日志分页查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysOperationLogController) SelectOneHandler(c echo.Context) error {
	id := c.Param("id")

	data, err := controller.service.SelectById(c, cast.ToInt64(id))
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysOperationLogController) BatchDeleteHandler(c echo.Context) error {
	ids, err := GlobalDeleteHandler(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	count, err := controller.service.BatchDelete(c, ids)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	response.NewResponse(c, 200, "批量删除成功", count)
	return nil
}
