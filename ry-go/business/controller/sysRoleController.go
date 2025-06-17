package controller

import (
	"errors"
	"fmt"
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

// SysRoleController 系统角色控制器
type SysRoleController struct {
	service           service.SysRoleService
	userService       service.SysUserService
	tokenService      service.TokenService
	permissionService service.PermissionService
}

// SysRoleControllerInit 系统角色控制器初始化
func SysRoleControllerInit(
	s *serviceImpl.SysRoleServiceImpl,
	u *serviceImpl.SysUserServiceImpl,
	t *serviceImpl.TokenServiceImpl,
	p *serviceImpl.PermissionServiceImpl,
) *SysRoleController {
	return &SysRoleController{
		service:           s,
		tokenService:      t,
		userService:       u,
		permissionService: p,
	}
}

// SaveHandler 角色新增
func (controller *SysRoleController) SaveHandler(c echo.Context) error {
	var insertObj *domain.SysRoleInsert
	if err := c.Bind(&insertObj); err != nil {
		zap.L().Sugar().Errorf("系统角色新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "系统角色新增参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统角色新增入参===\n%s", utils.ToJsonFormat(insertObj))

	sysRole := domain.SysRoleInsertConvertToEntity(insertObj)

	roleNameExist, err := controller.service.VerifyRoleNameExist(c, sysRole)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if roleNameExist {
		response.NewRespCodeMsg(c, 500, "名称已存在")
		return err
	}

	roleKeyExist, err := controller.service.VerifyRoleKeyExist(c, sysRole)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if roleKeyExist {
		response.NewRespCodeMsg(c, 500, "权限字符已存在")
		return err
	}

	user, err := controller.tokenService.GetLoginUser(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if user != nil {
		sysRole.CreateBy = user.Username
	}
	sysRole.CreateTime = domain.LocalDateTimeNow("")

	result, err := controller.service.Insert(c, sysRole)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Sugar().Infof("\n系统角色新增请求ID:%s\nresult:\n%s",
		utils.GetValForContext(c.Request().Context(), "requestId"),
		utils.ToJsonFormat(result))
	response.NewResponse(c, 200, "系统角色信息新增成功", result)
	return nil
}

// UpdateHandler 更新
func (controller *SysRoleController) UpdateHandler(c echo.Context) error {
	var sysRole *domain.SysRole
	if err := c.Bind(&sysRole); err != nil {
		zap.L().Sugar().Errorf("系统角色更新参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "系统角色更新参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统角色更新入参===\n%s", utils.ToJsonFormat(sysRole))

	if sysRole.Id == 0 {
		response.NewRespCodeMsg(c, 500, "系统角色ID不能为空")
		return errors.New("系统角色ID不能为空")
	}

	roleIsAllowed := controller.service.VerifyRoleIsAllowed(sysRole)
	if roleIsAllowed {
		response.NewRespCodeMsg(c, 500, "管理员不允许操作")
		return errors.New("管理员不允许操作")
	}

	roleNameExist, err := controller.service.VerifyRoleNameExist(c, sysRole)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if roleNameExist {
		response.NewRespCodeMsg(c, 500, "名称已存在")
		return errors.New("名称已存在")
	}

	roleKeyExist, err := controller.service.VerifyRoleKeyExist(c, sysRole)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	if roleKeyExist {
		response.NewRespCodeMsg(c, 500, "权限字符已存在")
		return errors.New("权限字符已存在")
	}

	result, err := controller.service.Update(c, sysRole)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	// 更新用户缓存
	loginUser, err := controller.tokenService.GetLoginUser(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	if loginUser != nil && domain.IsAdmin(loginUser.UserId) {
		permissions, menuPermissionErr := controller.permissionService.SelectMenuPermission(c, loginUser.UserId)
		if menuPermissionErr != nil {
			response.NewRespCodeErr(c, 500, menuPermissionErr)
			return errors.New("获取用户菜单失败")
		}
		loginUser.Permissions = permissions
		user, userErr := controller.userService.SelectByUsername(c, loginUser.Username)
		if userErr != nil {
			response.NewRespCodeErr(c, 500, err)
			return errors.New("获取用户信息失败")
		}
		user.IsAdmin = domain.IsAdmin(user.Id)
		loginUser.User = user
		if serUserErr := controller.tokenService.SetLoginUser(c, loginUser); serUserErr != nil {
			response.NewRespCodeErr(c, 500, serUserErr)
			return errors.New("设置用户信息失败")
		}
	}

	response.NewResponse(c, 200, "系统角色信息更新成功", result)
	return nil
}

func (controller *SysRoleController) BatchSaveHandler(c echo.Context) error {
	var list []*domain.SysRole
	if err := c.Bind(&list); err != nil {
		zap.L().Sugar().Errorf("系统角色批量新增参数解析错误===%+v", err)
		response.NewRespCodeMsg(c, 500, "系统角色批量新增参数解析错误")
		return err
	}

	zap.L().Sugar().Infof("系统角色批量新增入参===\n%s", utils.ToJsonFormat(cast.ToSlice(list)))

	result, err := controller.service.BatchInsert(c, list)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	response.NewResponse(c, 200, "系统角色批量新增成功", result)
	return nil
}

// SelectPageHandler 分页查询处理器
func (controller *SysRoleController) SelectPageHandler(c echo.Context) error {
	req := new(request.RolePageParam)
	if err := c.Bind(req); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "分页参数解析失败")
		return err
	}

	list, total, page, err := controller.service.SelectPage(c, req)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	pageData := response.NewPageData(list, total, page)

	response.NewResponse(c, 200, "系统公告分页查询成功", pageData)
	return nil
}

// SelectOneHandler 查询单个
func (controller *SysRoleController) SelectOneHandler(c echo.Context) error {
	id := c.Param("id")
	zap.L().Sugar().Infof("系统角色通过ID查询入参===%s", id)

	data, err := controller.service.SelectById(c, cast.ToInt64(id))
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}

	zap.L().Info("系统角色通过ID查询返回", zap.String("id", id), zap.Any("data", data))
	response.NewResponse(c, 200, "查询成功", data)
	return nil
}

// BatchDeleteHandler 批量删除
func (controller *SysRoleController) BatchDeleteHandler(c echo.Context) error {
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

// SelectListHandler 查询角色分片处理器
func (controller *SysRoleController) SelectListHandler(c echo.Context) error {
	var ids []int64
	if err := c.Bind(&ids); err != nil {
		response.NewRespCodeMsg(c, 500, "系统角色查询参数解析错误")
		return err
	}

	list, err := controller.service.SelectList(c, ids)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewResponse(c, 200, "查询成功", list)
	return nil
}

// SelectIdListHandler 查询角色ID切片处理器
func (controller *SysRoleController) SelectIdListHandler(c echo.Context) error {
	list, err := controller.service.SelectIdsByUserId(c, c.QueryParam("userId"))
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewResponse(c, 200, "查询成功", list)
	return nil
}

func (controller *SysRoleController) SelectAllHandler(c echo.Context) error {
	allRole, err := controller.service.SelectAll(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	response.NewResponse(c, 200, "查询成功", allRole)
	return nil
}

func (controller *SysRoleController) DataScopeHandler(c echo.Context) error {
	var dataScope *domain.SysRoleDataScope
	if err := c.Bind(&dataScope); err != nil {
		response.NewRespCodeMsg(c, 500, "分配数据权限参数错误")
		return err
	}

	roleDataScope := domain.SysRoleDataScopeConvertToEntity(dataScope)

	allowed := VerifyRoleAllowed(roleDataScope)
	if !allowed {
		response.NewRespCodeMsg(c, 500, "不允许修改管理员数据权限")
		return errors.New("不允许修改管理员数据权限")
	}

	result, err := controller.service.UpdateWithOutMenuAndDataScope(c, roleDataScope)
	if err != nil {
		response.NewRespCodeMsg(c, 500, "分配数据权限操作错误")
		return err
	}
	response.NewRespCodeMsg(c, 200, fmt.Sprintf("操作成功，更新了：%d条数据", result))
	return nil
}

// DownloadExcelBufferHandler 导出Excel，使用缓冲流
func (controller *SysRoleController) DownloadExcelBufferHandler(c echo.Context) error {
	data, err := controller.service.SelectAll(c)
	if err != nil {
		response.NewRespCodeErr(c, 500, err)
		return err
	}
	headers := []string{"ID", "角色名称", "权限字符", "排序", "数据范围", "菜单树选择项是否关联显示", "部门树选择项是否关联显示", "角色状态（0正常 1停用）", "删除标志（0代表存在 2代表删除）", "创建人", "创建时间", "修改人", "修改时间", "备注信息"}
	return DownloadExcelBuffer(c, "角色信息_"+time.Now().Format("20060102150405")+".xlsx", "角色信息", headers, data, true)
}
