package controller

import (
	"context"
	"net/http"
	"ry-go/business/domain"
	"ry-go/business/service"
	"ry-go/business/service/serviceImpl"
	"ry-go/common/response"
	"ry-go/configuation"
	"ry-go/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// LoginController 登陆控制器
type LoginController struct {
	userService       service.SysUserService
	captchaService    service.CaptchaService
	deptService       service.SysDeptService
	roleService       service.SysRoleService
	menuService       service.SysMenuService
	tokenService      service.TokenService
	recordService     service.SysLoginRecordService
	permissionService service.PermissionService
}

// LoginControllerInit 登陆控制器初始化
func LoginControllerInit(
	s *serviceImpl.SysUserServiceImpl,
	c *serviceImpl.CaptchaServiceImpl,
	d *serviceImpl.SysDeptServiceImpl,
	r *serviceImpl.SysRoleServiceImpl,
	m *serviceImpl.SysMenuServiceImpl,
	t *serviceImpl.TokenServiceImpl,
	lr *serviceImpl.SysLoginRecordServiceImpl,
	p *serviceImpl.PermissionServiceImpl,
) *LoginController {
	return &LoginController{
		userService:       s,
		captchaService:    c,
		deptService:       d,
		roleService:       r,
		menuService:       m,
		tokenService:      t,
		recordService:     lr,
		permissionService: p,
	}
}

// CaptchaImageHandler 登陆验证码处理器
func (controller *LoginController) CaptchaImageHandler(e echo.Context) error {
	// 设置响应头禁用缓存
	e.Request().Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	e.Request().Header.Set("Pragma", "no-cache")
	e.Request().Header.Set("Expires", "0")
	return controller.captchaService.GenerateLoginCaptcha(e)
}

// LoginHandler 用户登录处理器
func (controller *LoginController) LoginHandler(c echo.Context) error {
	var loginParam *domain.UserLogin
	if err := c.Bind(&loginParam); err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "用户登陆参数解析错误")
		return err
	}

	validate, err := LoginValidator(loginParam, validatorMap)
	if err != nil {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "用户登陆参数校验失败")
		return err
	}

	if !validate {
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "用户登陆参数校验失败")
		return err
	}

	if err = controller.captchaService.VerifyLoginCaptcha(c, loginParam.Code, loginParam.Uuid); err != nil {
		return err
	}

	user, err := controller.userService.Login(c, loginParam)
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}

	dept, err := controller.deptService.SelectById(c, user.DeptId)
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		return err
	}

	deptName := ""
	if dept != nil {
		deptName = dept.DeptName
	}

	var loginUser = &domain.LoginUser{
		Uuid:       utils.ShortUUID(),
		UserId:     user.Id,
		Username:   user.Username,
		DeptId:     user.DeptId,
		DeptName:   deptName,
		LoginTime:  time.Now().Unix(),
		ExpireTime: time.Now().Add(time.Hour).Unix(),
	}

	userAgent := utils.GetUserAgent(c)
	recode := &domain.SysLoginRecord{
		UserName:  loginUser.Username,
		Browser:   cast.ToString(userAgent["browser"]),
		Ipaddr:    cast.ToString(userAgent["ip"]),
		Os:        cast.ToString(userAgent["os"]),
		LoginTime: domain.LocalDateTimeNow(""),
	}

	loginUser.IP = recode.Ipaddr
	loginUser.Browser = recode.Browser
	loginUser.Os = recode.Os

	// 创建token
	token, err := controller.tokenService.CreateToken(c, loginUser)
	if err != nil {
		response.NewRespCodeErr(c, http.StatusInternalServerError, err)
		recode.Status = "1"
		recode.Msg = "返回凭证错误"
		controller.asyncLoginRecord(recode)
		return err
	}

	utils.SetValueToResp(c, configuation.JwtConfig.Item.RespHeaderKey, configuation.JwtConfig.Item.AuthPrefix+" "+token)

	// 凭证存入redis
	if err = controller.tokenService.CacheToken(c, loginUser.Uuid, token); err != nil {
		recode.Status = "1"
		recode.Msg = "凭证处理错误"
		controller.asyncLoginRecord(recode)
		response.NewRespCodeMsg(c, http.StatusInternalServerError, "登录失败，请稍后重试")
		return err
	}

	recode.Token = token
	recode.Status = "0"
	recode.Msg = "登录成功"

	controller.asyncLoginRecord(recode)
	// 设置有效期为7天
	utils.SetJwtTokenToCookie(c, configuation.JwtConfig.Item.RespHeaderKey, token, 24*60*60*7)

	response.NewResponse(c, http.StatusOK, "登陆成功", map[string]string{
		"token": token,
	})
	return nil
}

// asyncLoginRecord 登陆记录使用异步方式
func (controller *LoginController) asyncLoginRecord(record *domain.SysLoginRecord) {
	// 深拷贝记录对象
	copiedRecord := deepCopyRecord(record)

	go func() {
		// 使用独立上下文
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if _, err := controller.recordService.Insert(ctx, copiedRecord); err != nil {
			return
		}

	}()
}

func deepCopyRecord(src *domain.SysLoginRecord) *domain.SysLoginRecord {
	return &domain.SysLoginRecord{
		UserName:  src.UserName,
		Browser:   src.Browser,
		Ipaddr:    src.Ipaddr,
		Os:        src.Os,
		LoginTime: src.LoginTime,
		Status:    src.Status,
		Msg:       src.Msg,
		Token:     src.Token,
	}
}

// LoginInfoHandler 获取用户信息处理器
func (controller *LoginController) LoginInfoHandler(e echo.Context) error {
	loginUser, err := controller.tokenService.GetLoginUser(e)
	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "获取用户信息失败，请稍后重试")
		return err
	}

	user, err := controller.userService.SelectByUsername(e, loginUser.Username)
	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "获取用户信息失败，请稍后重试")
		return err
	}

	user.IsAdmin = domain.IsAdmin(user.Id)

	dept, err := controller.deptService.SelectById(e, user.DeptId)
	if err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return err
	}
	user.Dept = dept

	roleIds, err := controller.roleService.SelectIdsByUserId(e, user.Id)
	if err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return err
	}

	roleList, err := controller.roleService.SelectList(e, roleIds)
	if err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return err
	}
	user.Roles = roleList

	permissions, err := controller.permissionService.SelectMenuPermission(e, user.Id)
	if err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return err
	}

	response.NewResponse(e, http.StatusOK, "获取用户信息成功",
		&response.UserInfo{
			User:        user,
			Roles:       cast.ToStringSlice(roleIds),
			Permissions: permissions,
		},
	)
	return nil
}

// RoutersHandler 用户路由处理器
func (controller *LoginController) RoutersHandler(e echo.Context) error {
	loginUser, err := controller.tokenService.GetLoginUser(e)
	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "获取路由信息失败，请稍后重试")
		return err
	}

	menus, err := controller.menuService.SelectTreeByUserId(e, loginUser.UserId)
	if err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return err
	}
	routerVoList, err := controller.menuService.BuildMenus(menus)
	if err != nil {
		response.NewRespCodeErr(e, http.StatusInternalServerError, err)
		return err
	}

	response.NewRespNoMsg(e, http.StatusOK, routerVoList)
	return nil
}

// LogoutHandler 用户登出/注销管理器
func (controller *LoginController) LogoutHandler(e echo.Context) error {
	err := controller.tokenService.DeleteLoginUser(e)
	if err != nil {
		response.NewRespCodeMsg(e, http.StatusInternalServerError, "用户退出/注销失败，请稍后重试")
		return err
	}
	response.NewRespCodeMsg(e, http.StatusOK, "用户退出/注销成功")
	return nil
}
