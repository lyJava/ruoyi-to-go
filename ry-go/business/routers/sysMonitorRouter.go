package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// SysMonitorRouterInit 在线用户请求路由
func SysMonitorRouterInit(group *echo.Group, monitor *controller.SysMonitorController) {
	sysMonitorRouterGroup := group.Group("/monitor")
	sysMonitorRouterGroup.GET("/online/list", monitor.OnlineListHandler)
	sysMonitorRouterGroup.DELETE("/online/:uuid", monitor.OnlineForceLogoutHandler)
	sysMonitorRouterGroup.GET("/server", monitor.ServerInfoHandler)

}
