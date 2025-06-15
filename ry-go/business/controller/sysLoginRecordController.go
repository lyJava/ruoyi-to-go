package controller

import (
	"context"
	"errors"
	"net/http"
	"ry-go/business/domain"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/request"
	"ry-go/common/response"
	"ry-go/utils"
	"time"

	"github.com/spf13/cast"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// SysLoginRecordController 登录记录控制器
type SysLoginRecordController struct {
	recordService service.SysLoginRecordService
	tokenService  service.TokenService
}

// SysLoginRecordControllerInit 登录记录控制器初始化
func SysLoginRecordControllerInit(
	s *serviceImpl.SysLoginRecordServiceImpl,
	t *serviceImpl.TokenServiceImpl,
) *SysLoginRecordController {
	return &SysLoginRecordController{
		recordService: s,
		tokenService:  t,
	}
}

// SaveHandler 保存
func (controller *SysLoginRecordController) SaveHandler(c echo.Context) error {
	var sysLoginRecord *domain.SysLoginRecord
	if err := c.Bind(&sysLoginRecord); err != nil {
		zap.L().Sugar().Errorf("登录记录新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "登录记录新增参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("登录记录新增入参===\n%s", utils.ToJsonFormat(sysLoginRecord))

	// 使用独立上下文
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := controller.recordService.Insert(ctx, sysLoginRecord)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("\n登录记录新增请求ID:%s\nresult:\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "登录记录信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysLoginRecordController) UpdateHandler(c echo.Context) error {
	var sysLoginRecord *domain.SysLoginRecord
	if err := c.Bind(&sysLoginRecord); err != nil {
		zap.L().Sugar().Errorf("登录记录更新参数解析错误===%+v", err)
		response.NewResponse(c, 500, "登录记录更新参数解析错误", nil)
		return err
	}

	zap.L().Sugar().Infof("登录记录更新入参===\n%s", utils.ToJsonFormat(sysLoginRecord))

	if sysLoginRecord.InfoId == 0 {
		response.NewResponse(c, 500, "登录记录ID不能为空", nil)
		return errors.New("登录记录ID不能为空")
	}

	result, err := controller.recordService.Update(c, sysLoginRecord)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	zap.L().Sugar().Infof("登录记录更新返回===\n%s", utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "登录记录信息更新成功", result)
	return nil
}

func (controller *SysLoginRecordController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysLoginRecord
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("登录记录批量新增参数解析错误===%+v", err)
		response.NewResponse(c, 500, "登录记录批量新增参数解析错误", err)
		return err
	}

	zap.L().Sugar().Infof("登录记录批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.recordService.BatchInsert(c, list)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Sugar().Infof("登录记录批量新增请求ID===%s\n返回===\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "登录记录批量新增成功", result)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysLoginRecordController) SelectPageHandler(c echo.Context) error {
	param := new(request.SysNoticePageParam)
	if err := c.Bind(param); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "登录日志分页参数解析失败")
		return err
	}

	zap.L().Sugar().Infof("登录记录分页查询入参===\n%s", utils.ToJsonFormat(param))

	list, total, page, err := controller.recordService.SelectPage(c, param)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, 200, "系统公告分页查询成功", pageData)
	// 使用自定义上下文
	//cc := c.(*CustomContext)

	return nil
}

// SelectOneHandler 查询单个
func (controller *SysLoginRecordController) SelectOneHandler(c echo.Context) error {
	id := c.Param("id")
	zap.L().Sugar().Infof("登录记录通过ID查询入参===%s", id)

	data, err := controller.recordService.SelectById(c, cast.ToInt64(id))
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}

	zap.L().Info("登录记录通过ID查询返回", zap.String("id", id), zap.Any("data", data))
	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysLoginRecordController) BatchDeleteHandler(c echo.Context) error {
	ids, err := GlobalDeleteHandler(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	count, err := controller.recordService.BatchDelete(c, ids)
	if err != nil {
		response.NewResponse(c, 500, err.Error(), nil)
		return err
	}
	response.NewResponse(c, 200, "批量删除成功", count)
	return nil
}

func (controller *SysLoginRecordController) SelectLastLoginRecodeHandler(c echo.Context) error {
	loginUser, err := controller.tokenService.GetLoginUser(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if loginUser == nil {
		response.NewRespCodeMsg(c, 500, "用户登录信息已失效")
		return errors.New("用户登录信息已失效")
	}

	record, recordErr := controller.recordService.SelectLastRecode(c, loginUser.Username)
	if recordErr != nil {
		response.NewRespCodeErr(c, 500, recordErr)
		return recordErr
	}

	response.NewResponse(c, http.StatusOK, "查询成功", utils.TimeForHuman(record.LoginTime.Local()))
	return nil
}


// DownloadExcelBufferHandler 导出Excel，使用缓冲流
func (controller *SysLoginRecordController) DownloadExcelBufferHandler(c echo.Context) error {
	data, err := controller.recordService.SelectAll(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	headers := []string{"ID", "账号", "登陆IP地址", "登陆地点", "浏览器类型", "操作系统", "登录状态", "提示消息", "访问时间", "登录凭证"}
	return DownloadExcelBuffer(c, "登陆日志_" + time.Now().Format("20060102150405") + ".xlsx", headers, data)
}
