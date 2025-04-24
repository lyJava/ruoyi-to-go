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

// SysNoticeController 系统公告控制器
type SysNoticeController struct {
	service service.SysNoticeService
}

// SysNoticeControllerInit 系统系统公告控制器初始化
func SysNoticeControllerInit(s *serviceImpl.SysNoticeServiceImpl) *SysNoticeController {
	return &SysNoticeController{
		service: s,
	}
}

// SaveHandler 保存
func (controller *SysNoticeController) SaveHandler(c echo.Context) error {
	var notice *domain.SysNotice
	if err := c.Bind(&notice); err != nil {
		zap.L().Sugar().Errorf("系统公告新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "系统公告新增参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统公告新增入参===\n%s", utils.ToJsonFormat(notice))

	result, err := controller.service.Insert(c, notice)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Sugar().Infof("\n系统公告新增请求ID:%s\nresult:\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "系统公告信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysNoticeController) UpdateHandler(c echo.Context) error {
	var notice *domain.SysNotice
	if err := c.Bind(&notice); err != nil {
		zap.L().Sugar().Errorf("系统公告更新参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "系统公告更新参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统公告更新入参===\n%s", utils.ToJsonFormat(notice))

	if notice.Id == 0 {
		response.NewRespCodeMsg(c, 500, "系统公告ID不能为空")
		return errors.New("系统公告ID不能为空")
	}

	result, err := controller.service.Update(c, notice)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	zap.L().Sugar().Infof("系统公告更新返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "系统公告信息更新成功", result)
	return nil
}

func (controller *SysNoticeController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysNotice
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("系统公告批量新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "系统公告批量新增参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统公告批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchSave(c, list)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Sugar().Infof("系统公告批量新增请求ID===%s\n返回===\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "系统公告批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysNoticeController) SelectPageHandler(c echo.Context) error {
	var param *request.SysNoticePageParam
	if err := c.Bind(&param); err != nil {
		zap.L().Sugar().Error("系统公告分页查询参数解析错误", zap.Error(err))
		response.NewRespCodeMsg(c, 500, "系统公告分页查询参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统公告分页查询入参===\n%s", utils.ToJsonFormat(param))

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	zap.L().Sugar().Infof("系统公告分页查询返回===\n%s", utils.ToJsonFormat(pageData))
	response.NewResponse(c, 200, "系统公告分页查询成功", pageData)
	return nil
}

// SelectPageGetHandler 分页处理
func (controller *SysNoticeController) SelectPageGetHandler(c echo.Context) error {
	param := new(request.SysNoticePageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "通知公告分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	pageData := response.NewPageData(list, total, cast.ToInt64(page))

	zap.L().Sugar().Infof("系统公告分页查询返回===\n%s", utils.ToJsonFormat(pageData))
	response.NewResponse(c, 200, "系统公告分页查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysNoticeController) SelectOneHandler(c echo.Context) error {
	id := c.Param("id")

	zap.L().Sugar().Infof("系统公告通过ID查询入参===%s", id)

	data, err := controller.service.SelectById(c, cast.ToInt64(id))
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Info("系统公告通过ID查询返回", zap.String("id", id), zap.Any("data", data))
	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysNoticeController) BatchDeleteHandler(c echo.Context) error {
	ids, err := GlobalDeleteHandler(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	count, err := controller.service.BatchDelete(c, ids)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewResponse(c, 200, "批量删除成功", count)
	return nil
}
