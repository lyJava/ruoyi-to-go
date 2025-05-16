package controller

import (
	"errors"
	"github.com/spf13/cast"
	"net/http"
	"ry-go/business/domain"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/request"
	"ry-go/common/response"
	"ry-go/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// SysConfigController 系统配置控制器
type SysConfigController struct {
	service service.SysConfigService
}

// SysConfigControllerInit 系统配置控制器初始化
func SysConfigControllerInit(s *serviceImpl.SysConfigServiceImpl) *SysConfigController {
	return &SysConfigController{
		service: s,
	}
}

// SaveHandler 保存
func (controller *SysConfigController) SaveHandler(c echo.Context) error {
	var sysConfig *domain.SysConfig
	if err := c.Bind(&sysConfig); err != nil {
		zap.L().Sugar().Errorf("系统配置新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "系统配置新增参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("系统配置新增入参===\n%s", utils.ToJsonFormat(sysConfig))

	result, err := controller.service.Insert(c, sysConfig)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("\n系统配置新增请求ID:%s\nresult:\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "系统配置信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysConfigController) UpdateHandler(c echo.Context) error {
	var sysConfig *domain.SysConfig
	if err := c.Bind(&sysConfig); err != nil {
		zap.L().Sugar().Errorf("系统配置更新参数解析错误===%+v", err)
		response.NewResponse(c, 500, "系统配置更新参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("系统配置更新入参===\n%s", utils.ToJsonFormat(sysConfig))

	if sysConfig.ConfigId == 0 {
		response.NewResponse(c, 500, "系统配置ID不能为空", nil)
		return errors.New("系统配置ID不能为空")
	}

	result, err := controller.service.Update(c, sysConfig)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	zap.L().Sugar().Infof("系统配置更新返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "系统配置信息更新成功", result)
	return nil
}

func (controller *SysConfigController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysConfig
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("系统配置批量新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "系统配置批量新增参数解析错误", err)
		return err
	}

	zap.L().Sugar().Infof("系统配置批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("系统配置批量新增请求ID===%s\n返回===\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "系统配置批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysConfigController) SelectPageHandler(c echo.Context) error {
	var param = new(request.ConfigPageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "参数配置分页查询参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	zap.L().Sugar().Infof("系统配置分页查询返回===\n%s", utils.ToJsonFormat(pageData))
	response.NewResponse(c, 200, "系统公告分页查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysConfigController) SelectOneHandler(c echo.Context) error {
	id := c.Param("id")
	zap.L().Sugar().Infof("系统配置通过ID查询入参===%s", id)

	data, err := controller.service.SelectById(c, cast.ToInt64(id))
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Info("系统配置通过ID查询返回", zap.String("id", id), zap.Any("data", data))
	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysConfigController) BatchDeleteHandler(c echo.Context) error {
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

func (controller *SysConfigController) SelectByKeyHandler(c echo.Context) error {
	key := c.Param("key")
	configData, err := controller.service.SelectConfigByKey(c, key)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewResponse(c, 200, "查询成功", configData)
	return nil
}

func (controller *SysConfigController) RefreshCacheHandler(c echo.Context) error {
	configList, err := controller.service.SelectConfigList(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	if err = controller.service.RefreshConfigCache(c, configList); err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewRespCodeMsg(c, 200, "刷新成功")
	return nil
}

func (controller *SysConfigController) ClearCacheHandler(c echo.Context) error {
	configList, err := controller.service.SelectConfigList(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	if err = controller.service.ClearConfigCache(c, configList); err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewRespCodeMsg(c, 200, "删除成功")
	return nil
}
