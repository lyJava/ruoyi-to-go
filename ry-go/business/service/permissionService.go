package service

import (
	"github.com/labstack/echo/v4"
)

type PermissionService interface {
	SelectRolePermission(e echo.Context, userId int64) ([]string, error)
	SelectMenuPermission(e echo.Context, userId int64) ([]string, error)
}
