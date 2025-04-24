package router

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"net/http"
	"ry-go/business/controller"
	"ry-go/business/routers"
	"ry-go/configuation"
)

var ProviderSet = wire.NewSet(NewRouter)

//goland:noinspection GoUnhandledErrorResult
func NewRouter(
	sysUser *controller.UserController,
	sysDept *controller.DeptController,
	sysNotice *controller.SysNoticeController,
	sysMenu *controller.SysMenuController,
	sysRole *controller.SysRoleController,
	dictData *controller.SysDictDataController,
	fileController *controller.FileController,
	barcode *controller.BarcodeController,
	qrcode *controller.QrcodeController,
	sysUserRole *controller.SysUserRoleController,
	sysPost *controller.SysPostController,
	dictType *controller.SysDictTypeController,
	login *controller.LoginController,
	loginRecord *controller.SysLoginRecordController,
	sysConfig *controller.SysConfigController,
	operation *controller.SysOperationLogController,
	monitor *controller.SysMonitorController,
	sysJob *controller.SysJobController,
	sysJobLog *controller.SysJobLogController,
) (*echo.Echo, error) {

	e := echo.New()

	// 加入zap日志处理
	e.Use(EnsureRequestID(), EchoLogger(configuation.ZapLog), EchoRecovery(configuation.ZapLog, true))

	serverConfigItem := configuation.ServerConfig.Server

	// 设置echo服务请求前缀
	group := e.Group(serverConfigItem.Path)
	group.Use(AuthMiddleware(serverConfigItem.Version))
	{
		routers.SysUserRouterInit(group, sysUser)
		routers.SysDeptRouterInit(group, sysDept)
		routers.SysNoticeRouterInit(group, sysNotice)
		routers.SysMenuRouterInit(group, sysMenu)
		routers.SysRoleRouterInit(group, sysRole)
		routers.SysDictDataRouterInit(group, dictData)
		routers.FileRouterInit(group, fileController)
		routers.BarcodeRouterInit(group, barcode)
		routers.QrcodeRouterInit(group, qrcode)
		routers.SysUserRoleRouterInit(group, sysUserRole)
		routers.SysPostRouterInit(group, sysPost)
		routers.SysDictTypeRouterInit(group, dictType)
		routers.LoginRouterInit(group, login)
		routers.SysLoginRecordRouterInit(group, loginRecord)
		routers.SysConfigRouterInit(group, sysConfig)
		routers.SysOperationLogRouterInit(group, operation)
		routers.SysMonitorRouterInit(group, monitor)
		routers.SysJobRouterInit(group, sysJob)
		routers.SysJobLogRouterInit(group, sysJobLog)
	}

	// 设置静态资源目录
	e.Static("/static", "./resources/static")

	e.RouteNotFound("/*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"msg": "404 - 页面未找到",
		})
	})

	return e, nil
}
