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

// SysJobController 系统任务控制器
type SysJobController struct {
	service     service.SysJobService
	cornService service.CronSchedulerService
}

// SysJobControllerInit 系统任务控制器初始化
func SysJobControllerInit(s *serviceImpl.SysJobServiceImpl, c *serviceImpl.CronSchedulerServiceImpl) *SysJobController {
	return &SysJobController{
		service:     s,
		cornService: c,
	}
}

// SaveHandler 保存
func (controller *SysJobController) SaveHandler(c echo.Context) error {
	var sysJob *domain.SysJob
	if err := c.Bind(&sysJob); err != nil {
		response.NewResponse(c, 500, "系统任务新增参数解析错误", nil)
		return err
	}

	result, err := controller.service.Insert(c, sysJob)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	if err = controller.cornService.AddTask(c.Request().Context(), result); err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	return nil
}

// UpdateHandler 更新
func (controller *SysJobController) UpdateHandler(c echo.Context) error {
	var sysJob *domain.SysJob
	if err := c.Bind(&sysJob); err != nil {
		response.NewResponse(c, 500, "系统任务更新参数解析错误", nil)
		return err
	}

	if sysJob.JobId == 0 {
		response.NewResponse(c, 500, "系统任务ID不能为空", nil)
		return errors.New("系统任务ID不能为空")
	}

	result, err := controller.service.Update(c, sysJob)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	response.NewResponse(c, 200, "系统任务信息更新成功", result)
	return nil
}

func (controller *SysJobController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysJob
	if err := c.Bind(&list); err != nil {
		response.NewResponse(c, 500, "系统任务批量新增参数解析错误", err)
		return err
	}

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	response.NewResponse(c, 200, "系统任务批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysJobController) SelectPageHandler(c echo.Context) error {
	param := new(request.SysNoticePageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "系统任务分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	if len(list) > 0 {
		for _, job := range list {
			if job.Status == "0" {
				controller.cornService.AddTask(c.Request().Context(), job)
			}
			continue
		}
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, 200, "系统任务分页查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysJobController) SelectOneHandler(c echo.Context) error {
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
func (controller *SysJobController) BatchDeleteHandler(c echo.Context) error {
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

func (controller *SysJobController) ChangeStatusHandler(c echo.Context) error {
	var param *request.JobChangeParam
	if err := c.Bind(&param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "系统任务状态改变参数解析失败")
		return err
	}

	if param.Status == "1" {
		if err := controller.cornService.TaskStop(c.Request().Context(), param.JobId); err != nil {
			response.NewRespCodeMsg(c, http.StatusInternalServerError, "任务停止失败")
			return err
		}

		_, err := controller.service.UpdateJobStatus(c, cast.ToInt64(param.JobId), param.Status)
		if err != nil {
			response.NewResponse(c, 500, err.Error(), nil)
			return err
		}
		response.NewRespCodeMsg(c, http.StatusOK, "任务停止成功")
	}

	if param.Status == "0" {
		if err := controller.cornService.TaskStart(c.Request().Context(), param.JobId); err != nil {
			response.NewRespCodeMsg(c, http.StatusInternalServerError, "任务启动失败")
			return err
		}
		_, err := controller.service.UpdateJobStatus(c, cast.ToInt64(param.JobId), param.Status)
		if err != nil {
			response.NewResponse(c, 500, err.Error(), nil)
			return err
		}
		response.NewRespCodeMsg(c, http.StatusOK, "任务启动成功")
	}

	return nil
}

func (controller *SysJobController) ChangeStatusBatchHandler(c echo.Context) error {
	var list []*request.JobChangeParam
	if err := c.Bind(&list); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "系统任务状态改变参数解析失败")
		return err
	}

	if len(list) > 0 {
		var errList []error
		for _, param := range list {
			if param.Status == "1" {
				if err := controller.cornService.TaskStop(c.Request().Context(), param.JobId); err != nil {
					errList = append(errList, err)
				}
				_, err := controller.service.UpdateJobStatus(c, cast.ToInt64(param.JobId), param.Status)
				if err != nil {
					errList = append(errList, err)
				}
			}

			if param.Status == "0" {
				if err := controller.cornService.TaskStart(c.Request().Context(), param.JobId); err != nil {
					errList = append(errList, err)
				}
				_, err := controller.service.UpdateJobStatus(c, cast.ToInt64(param.JobId), param.Status)
				if err != nil {
					errList = append(errList, err)
				}
			}
		}
		if len(errList) > 0 {
			return errors.New("任务状态改动失败")
		}
	}
	response.NewRespCodeMsg(c, http.StatusOK, "任务状态改动成功")
	return nil
}
