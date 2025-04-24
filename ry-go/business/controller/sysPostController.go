package controller

import (
	"errors"
	"net/http"
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

// SysPostController 岗位信息控制器
type SysPostController struct {
	service service.SysPostService
}

// SysPostControllerInit 岗位信息控制器初始化
func SysPostControllerInit(s *serviceImpl.SysPostServiceImpl) *SysPostController {
	return &SysPostController{
		service: s,
	}
}

// SaveHandler 保存
func (controller *SysPostController) SaveHandler(c echo.Context) error {
	var sysPost *domain.SysPost
	if err := c.Bind(&sysPost); err != nil {
		zap.L().Sugar().Errorf("岗位信息新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "岗位信息新增参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("岗位信息新增入参===\n%s", utils.ToJsonFormat(sysPost))

	result, err := controller.service.Insert(c, sysPost)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("\n岗位信息新增请求ID:%s\nresult:\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "岗位信息信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysPostController) UpdateHandler(c echo.Context) error {
	var sysPost *domain.SysPost
	if err := c.Bind(&sysPost); err != nil {
		zap.L().Sugar().Errorf("岗位信息更新参数解析错误===%+v", err)
		response.NewResponse(c, 500, "岗位信息更新参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("岗位信息更新入参===\n%s", utils.ToJsonFormat(sysPost))

	if sysPost.Id == 0 {
		response.NewResponse(c, 500, "岗位信息ID不能为空", nil)
		return errors.New("岗位信息ID不能为空")
	}

	result, err := controller.service.Update(c, sysPost)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	zap.L().Sugar().Infof("岗位信息更新返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "岗位信息信息更新成功", result)
	return nil
}

func (controller *SysPostController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysPost
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("岗位信息批量新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "岗位信息批量新增参数解析错误", err)
		return err
	}

	zap.L().Sugar().Infof("岗位信息批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("岗位信息批量新增请求ID===%s\n返回===\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "岗位信息批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysPostController) SelectPageHandler(c echo.Context) error {
	param := new(request.PostPageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "岗位分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, 200, "查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysPostController) SelectOneHandler(c echo.Context) error {
	id := c.Param("id")
	zap.L().Sugar().Infof("岗位信息通过ID查询入参===%s", id)

	data, err := controller.service.SelectById(c, cast.ToInt64(id))
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Info("岗位信息通过ID查询返回", zap.String("id", id), zap.Any("data", data))
	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysPostController) BatchDeleteHandler(c echo.Context) error {
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

func (controller *SysPostController) SelectAllHandler(c echo.Context) error {
	allPost, err := controller.service.SelectAll(c)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	response.NewResponse(c, 200, "查询成功", allPost)
	return nil
}
