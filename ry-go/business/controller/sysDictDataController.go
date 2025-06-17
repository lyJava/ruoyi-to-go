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
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// SysDictDataController 字典数据控制器
type SysDictDataController struct {
	service     service.SysDictDataService
	redisClient *redis.Client
}

// SysDictDataControllerInit 字典数据控制器初始化
func SysDictDataControllerInit(s *serviceImpl.SysDictDataServiceImpl, client *redis.Client) *SysDictDataController {
	return &SysDictDataController{
		service:     s,
		redisClient: client,
	}
}

// SaveHandler 保存
func (controller *SysDictDataController) SaveHandler(c echo.Context) error {
	var sysDictData *domain.SysDictData
	if err := c.Bind(&sysDictData); err != nil {
		response.NewRespCodeMsg(c, 500, "字典数据新增参数解析错误")
		return err
	}

	result, err := controller.service.Insert(c, sysDictData)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	response.NewResponse(c, 200, "字典数据信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysDictDataController) UpdateHandler(c echo.Context) error {
	var sysDictData *domain.SysDictData
	if err := c.Bind(&sysDictData); err != nil {
		zap.L().Sugar().Errorf("字典数据更新参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "字典数据更新参数解析错误")
		return err
	}

	if sysDictData.Id == 0 {
		response.NewRespCodeMsg(c, 500, "字典数据ID不能为空")
		return errors.New("字典数据ID不能为空")
	}

	result, err := controller.service.Update(c, sysDictData)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	response.NewResponse(c, 200, "字典数据信息更新成功", result)
	return nil
}

func (controller *SysDictDataController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysDictData
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("字典数据批量新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "字典数据批量新增参数解析错误")
		return err
	}

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	response.NewResponse(c, 200, "字典数据批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysDictDataController) SelectPageHandler(c echo.Context) error {
	param := new(request.DictDataPageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, param)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, 200, "查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysDictDataController) SelectOneHandler(c echo.Context) error {

	data, err := controller.service.SelectById(c, cast.ToInt64(c.Param("id")))
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysDictDataController) BatchDeleteHandler(c echo.Context) error {
	param := c.Param("id")
	if param == "" {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "参数验证失败")
		return errors.New("参数验证失败")
	}

	ids := strings.Split(param, ",")
	if len(ids) == 0 {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "参数验证失败")
		return errors.New("参数验证失败")
	}

	count, err := controller.service.BatchDelete(c, utils.StringSliceToAnySlice(ids))
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	response.NewResponse(c, 200, "批量删除成功", count)
	return nil
}

func (controller *SysDictDataController) SelectDictDataByTypeHandler(c echo.Context) error {
	dictType := c.Param("dictType")

	var dataList []*domain.SysDictData

	cacheList, _ := utils.GetTyped(c.Request().Context(), controller.redisClient, "sys_dict:"+dictType, dataList)

	if len(cacheList) <= 0 {
		dataList, err := controller.service.SelectDictDataByType(c, dictType)
		if err != nil {
			response.NewRespCodeErr(c, 500, err)
			return err
		}

		if len(dataList) > 0 {
			_, err = utils.Set(c.Request().Context(), controller.redisClient, "sys_dict:"+dictType, dataList, time.Hour)
			if err != nil {
				response.NewRespCodeErr(c, 500, err)
				return err
			}
			response.NewResponse(c, 200, "查询成功", dataList)
			return nil
		}
	} else {
		response.NewResponse(c, 200, "查询成功", cacheList)
		return nil
	}

	return nil
}


// DownloadExcelBufferHandler 导出Excel，使用缓冲流
func (controller *SysDictDataController) DownloadExcelBufferHandler(c echo.Context) error {
	data, err := controller.service.SelectAll(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	headers := []string{"字典ID", "排序", "字典标签", "字典键值", "字典类型", "样式属性", "表格回显样式", "是否默认", "字典状态（0正常 1停用）", "创建人", "创建时间", "修改人", "修改时间", "备注信息"}
	return DownloadExcelBuffer(c, "字典数据_"+time.Now().Format("20060102150405")+".xlsx", "字典数据", headers, data, true)
}
