package service

import (
	"ry-go/business/domain"

	"github.com/labstack/echo/v4"
)

type SysMonitorService interface {
	OnlineList(e echo.Context, keyPrefix string) ([]*domain.LoginUser, error)
	OnlineForceLogout(e echo.Context, cacheKey string) (bool, error)
}
