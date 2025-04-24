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

// SysMenuController 系统菜单控制器
type SysMenuController struct {
	service      service.SysMenuService
	tokenService service.TokenService
}

// SysMenuControllerInit 系统菜单控制器初始化
func SysMenuControllerInit(s *serviceImpl.SysMenuServiceImpl, t *serviceImpl.TokenServiceImpl) *SysMenuController {
	return &SysMenuController{
		service:      s,
		tokenService: t,
	}
}

// SaveHandler 保存
func (controller *SysMenuController) SaveHandler(e echo.Context) error {
	var sysMenu *domain.SysMenu
	if err := e.Bind(&sysMenu); err != nil {
		zap.L().Sugar().Errorf("系统菜单新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(e, 500, "系统菜单新增参数解析错误")
		return err
	}

	result, err := controller.service.Insert(e, sysMenu)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	response.NewResponse(e, 200, "新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysMenuController) UpdateHandler(e echo.Context) error {
	var sysMenu *domain.SysMenu
	if err := e.Bind(&sysMenu); err != nil {
		zap.L().Sugar().Errorf("系统菜单更新参数解析错误===%+v", err)
		response.NewRespCodeMsg(e, 500, "系统菜单更新参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统菜单更新入参===\n%s", utils.ToJsonFormat(sysMenu))

	if sysMenu.Id == 0 {
		response.NewRespCodeMsg(e, 500, "系统菜单ID不能为空")
		return errors.New("系统菜单ID不能为空")
	}

	result, err := controller.service.Update(e, sysMenu)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	response.NewResponse(e, 200, "系统菜单信息更新成功", result)
	return nil
}

func (controller *SysMenuController) BatchSaveHandler(e echo.Context) error {
	var list []*domain.SysMenu
	if err := e.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("系统菜单批量新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(e, 500, "系统菜单批量新增参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统菜单批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchInsert(e, list)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	zap.L().Sugar().Infof("系统菜单批量新增请求ID===%s\n返回===\n%s",
		utils.GetValForContext(e.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(e, 200, "系统菜单批量新增成功", result)
	return nil
}

// SelectTreeTableListHandler 不分页树型表格数据
func (controller *SysMenuController) SelectTreeTableListHandler(e echo.Context) error {
	req := new(request.MenuPageParam)
	if err := e.Bind(req); err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "菜单表格树参数解析失败")
		return err
	}

	loginUser, err := controller.tokenService.GetLoginUser(e)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	if loginUser != nil {
		req.UserId = loginUser.UserId
	}

	list, err := controller.service.SelectList(e, req)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	response.NewResponse(e, 200, "查询成功", list)
	return nil
}

// SelectPageHandler 分页处理
func (controller *SysMenuController) SelectPageHandler(e echo.Context) error {
	req := new(request.MenuPageParam)
	if err := e.Bind(req); err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "菜单分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(e, req)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(e, 200, "系统公告分页查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysMenuController) SelectOneHandler(e echo.Context) error {
	id := e.Param("id")
	zap.L().Sugar().Infof("系统菜单通过ID查询入参===%s", id)

	data, err := controller.service.SelectById(e, cast.ToInt64(id))
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	response.NewResponse(e, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysMenuController) BatchDeleteHandler(e echo.Context) error {
	ids, err := GlobalDeleteHandler(e)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	count, err := controller.service.BatchDelete(e, ids)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	response.NewResponse(e, 200, "批量删除成功", count)
	return nil
}

// TreeSelectHandler 获取菜单下拉树列表
func (controller *SysMenuController) TreeSelectHandler(e echo.Context) error {
	loginUser, err := controller.tokenService.GetLoginUser(e)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	menuList, err := controller.service.SelectListByMenuAndUserId(e, &domain.SysMenu{}, loginUser.UserId)

	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}

	/*menuTreeList, err := controller.service.BuildMenuTree(menuList)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}*/

	response.NewResponse(e, http.StatusOK, "查询成功", controller.service.BuildMenuTreeSelect(menuList))
	return nil
}

// RoleMenuTreeRoleIdHandler 对应角色菜单列表树
func (controller *SysMenuController) RoleMenuTreeRoleIdHandler(e echo.Context) error {
	roleId := e.Param("roleId")
	if roleId == "" {
		response.NewRespCodeMsg(e, 500, "角色ID不能为空")
		return errors.New("角色ID不能为空")
	}

	loginUser, err := controller.tokenService.GetLoginUser(e)
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	roleMenuTreeSelect, err := controller.service.BuildRoleMenuTreeSelect(e, loginUser.UserId, cast.ToInt64(roleId))
	if err != nil {
		response.NewRespCodeErr(e, 500, err)
		return err
	}
	response.NewResponse(e, http.StatusOK, "查询成功", roleMenuTreeSelect)
	return nil
}
