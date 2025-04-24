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

// SysJobLogController 系统任务日志控制器
type SysJobLogController struct {
	service service.SysJobLogService
}

// SysJobLogControllerInit 系统任务日志控制器初始化
func SysJobLogControllerInit(s *serviceImpl.SysJobLogServiceImpl) *SysJobLogController {
	return &SysJobLogController{
		service: s,
	}
}

// SaveHandler 保存
func (controller *SysJobLogController) SaveHandler(c echo.Context) error {
	var sysJobLog *domain.SysJobLog
	if err := c.Bind(&sysJobLog); err != nil {
		response.NewResponse(c, 500, "系统任务日志新增参数解析错误", nil)
		return err
	}

	result, err := controller.service.Insert(c, sysJobLog)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	response.NewResponse(c, 200, "系统任务日志信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysJobLogController) UpdateHandler(c echo.Context) error {
	var sysJobLog *domain.SysJobLog
	if err := c.Bind(&sysJobLog); err != nil {
		response.NewResponse(c, 500, "系统任务日志更新参数解析错误", nil)
		return err
	}

	if sysJobLog.JobLogId == 0 {
		response.NewResponse(c, 500, "系统任务日志ID不能为空", nil)
		return errors.New("系统任务日志ID不能为空")
	}

	result, err := controller.service.Update(c, sysJobLog)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	response.NewResponse(c, 200, "系统任务日志信息更新成功", result)
	return nil
}

func (controller *SysJobLogController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysJobLog
	if err := c.Bind(&list); err != nil {
		response.NewResponse(c, 500, "系统任务日志批量新增参数解析错误", err)
		return err
	}

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	response.NewResponse(c, 200, "系统任务日志批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysJobLogController) SelectPageHandler(c echo.Context) error {
	param := new(request.SysNoticePageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "系统任务日志分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, 200, "系统任务日志分页查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysJobLogController) SelectOneHandler(c echo.Context) error {
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
func (controller *SysJobLogController) BatchDeleteHandler(c echo.Context) error {
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
