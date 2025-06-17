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
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// SysDictTypeController 字典类型控制器
type SysDictTypeController struct {
	service service.SysDictTypeService
}

// SysDictTypeControllerInit 字典类型控制器初始化
func SysDictTypeControllerInit(s *serviceImpl.SysDictTypeServiceImpl) *SysDictTypeController {
	return &SysDictTypeController{
		service: s,
	}
}

// SaveHandler 保存
func (controller *SysDictTypeController) SaveHandler(c echo.Context) error {
	var sysDictType *domain.SysDictType
	if err := c.Bind(&sysDictType); err != nil {
		zap.L().Sugar().Errorf("字典数据新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "字典数据新增参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("字典数据新增入参===\n%s", utils.ToJsonFormat(sysDictType))

	result, err := controller.service.Insert(c, sysDictType)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Sugar().Infof("\n字典数据新增请求ID:%s\nresult:\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "字典数据信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysDictTypeController) UpdateHandler(c echo.Context) error {
	var sysDictType *domain.SysDictType
	if err := c.Bind(&sysDictType); err != nil {
		zap.L().Sugar().Errorf("字典数据更新参数解析错误===%+v", err)
		response.NewResponse(c, 500, "字典数据更新参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("字典数据更新入参===\n%s", utils.ToJsonFormat(sysDictType))

	if sysDictType.DictId == 0 {
		response.NewResponse(c, 500, "字典数据ID不能为空", nil)
		return errors.New("字典数据ID不能为空")
	}

	result, err := controller.service.Update(c, sysDictType)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	zap.L().Sugar().Infof("字典数据更新返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "字典数据信息更新成功", result)
	return nil
}

func (controller *SysDictTypeController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysDictType
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("字典数据批量新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "字典数据批量新增参数解析错误", err)
		return err
	}

	zap.L().Sugar().Infof("字典数据批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("字典数据批量新增请求ID===%s\n返回===\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "字典数据批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysDictTypeController) SelectPageHandler(c echo.Context) error {
	param := new(request.DictTypePageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "分页参数解析失败")
		return err
	}

	zap.L().Sugar().Infof("字典数据分页查询入参===\n%s", utils.ToJsonFormat(param))

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
func (controller *SysDictTypeController) SelectOneHandler(c echo.Context) error {
	data, err := controller.service.SelectById(c, cast.ToInt64(c.Param("id")))
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysDictTypeController) BatchDeleteHandler(c echo.Context) error {
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

func (controller *SysDictTypeController) ClearCacheHandler(c echo.Context) error {
	err := controller.service.ClearCache(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewRespCodeMsg(c, 200, "缓存清理成功")
	return nil
}

func (controller *SysDictTypeController) RefreshCacheHandler(c echo.Context) error {
	err := controller.service.RefreshCache(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewRespCodeMsg(c, 200, "缓存刷新成功")
	return nil
}

// DownloadExcelBufferHandler 导出Excel，使用缓冲流
func (controller *SysDictTypeController) DownloadExcelBufferHandler(c echo.Context) error {
	data, err := controller.service.SelectAll(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	headers := []string{"字典ID", "字典名称", "字典类型", "字典状态（0正常 1停用）", "创建人", "创建时间", "修改人", "修改时间", "备注信息"}
	return DownloadExcelBuffer(c, "字典信息_"+time.Now().Format("20060102150405")+".xlsx", "字典信息", headers, data, true)
}
