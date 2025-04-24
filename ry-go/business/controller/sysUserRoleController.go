package controller

import (
	"errors"
	"ry-go/business/domain"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/request"
	"ry-go/common/response"
	"ry-go/utils"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// SysUserRoleController 用户角色信息控制器
type SysUserRoleController struct {
	service service.SysUserRoleService
}

// SysUserRoleControllerInit 用户角色信息控制器初始化
func SysUserRoleControllerInit(s *serviceImpl.SysUserRoleServiceImpl) *SysUserRoleController {
	return &SysUserRoleController{
		service: s,
	}
}

// SaveHandler 保存
func (controller *SysUserRoleController) SaveHandler(c echo.Context) error {
	var sysUserRole *domain.SysUserRole
	if err := c.Bind(&sysUserRole); err != nil {
		zap.L().Sugar().Errorf("用户角色信息新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "用户角色信息新增参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("用户角色信息新增入参===\n%s", utils.ToJsonFormat(sysUserRole))

	result, err := controller.service.Insert(c, sysUserRole)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("\n用户角色信息新增请求ID:%s\nresult:\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "用户角色信息信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysUserRoleController) UpdateHandler(c echo.Context) error {
	var sysUserRole *domain.SysUserRole
	if err := c.Bind(&sysUserRole); err != nil {
		zap.L().Sugar().Errorf("用户角色信息更新参数解析错误===%+v", err)
		response.NewResponse(c, 500, "用户角色信息更新参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("用户角色信息更新入参===\n%s", utils.ToJsonFormat(sysUserRole))

	if sysUserRole.UserId == 0 {
		response.NewResponse(c, 500, "用户角色信息ID不能为空", nil)
		return errors.New("用户角色信息ID不能为空")
	}

	result, err := controller.service.Update(c, sysUserRole)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	zap.L().Sugar().Infof("用户角色信息更新返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "用户角色信息信息更新成功", result)
	return nil
}

func (controller *SysUserRoleController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysUserRole
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("用户角色信息批量新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "用户角色信息批量新增参数解析错误", err)
		return err
	}

	zap.L().Sugar().Infof("用户角色信息批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("用户角色信息批量新增请求ID===%s\n返回===\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "用户角色信息批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysUserRoleController) SelectPageHandler(c echo.Context) error {
	var param *request.SysNoticePageParam
	if err := c.Bind(&param); err != nil {
		zap.L().Sugar().Error("用户角色信息分页查询参数解析错误", zap.Error(err))
		response.NewResponse(c, 500, "用户角色信息分页查询参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("用户角色信息分页查询入参===\n%s", utils.ToJsonFormat(param))

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	zap.L().Sugar().Infof("用户角色信息分页查询返回===\n%s", utils.ToJsonFormat(pageData))
	response.NewResponse(c, 200, "系统公告分页查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysUserRoleController) SelectOneHandler(c echo.Context) error {
	id := c.Param("id")
	zap.L().Sugar().Infof("用户角色信息通过ID查询入参===%s", id)

	data, err := controller.service.SelectById(c, cast.ToInt64(id))
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Info("用户角色信息通过ID查询返回", zap.String("id", id), zap.Any("data", data))
	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysUserRoleController) BatchDeleteHandler(c echo.Context) error {
	var ids []any
	if err := c.Bind(&ids); err != nil {
		zap.L().Sugar().Errorf("用户角色信息批量删除参数解析错误===%+v", err)
		response.NewResponse(c, 500, "用户角色信息系统公告批量删除参数解析错误", nil)
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
