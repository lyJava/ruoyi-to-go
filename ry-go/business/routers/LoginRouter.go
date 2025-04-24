package routers

import (
	"ry-go/business/controller"

	"github.com/labstack/echo/v4"
)

// LoginRouterInit 登陆路由初始化
func LoginRouterInit(group *echo.Group, login *controller.LoginController) {
	loginRouterGroup := group.Group("")
	loginRouterGroup.GET("/captchaImage", login.CaptchaImageHandler)
	loginRouterGroup.POST("/login", login.LoginHandler)
	loginRouterGroup.GET("/getInfo", login.LoginInfoHandler)
	loginRouterGroup.GET("/getRouters", login.RoutersHandler)
	loginRouterGroup.POST("/logout", login.LogoutHandler)

}
