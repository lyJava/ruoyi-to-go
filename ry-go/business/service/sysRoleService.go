package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysRoleService 系统角色服务接口
type SysRoleService interface {
	Insert(e echo.Context, sysRole *domain.SysRole) (*domain.SysRole, error)
	Update(e echo.Context, sysRole *domain.SysRole) (int64, error)
	UpdateWithOutMenuAndDataScope(e echo.Context, sysRole *domain.SysRole) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysRole) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysRole, error)
	SelectPage(e echo.Context, param *request.RolePageParam) ([]*domain.SysRole, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectList(e echo.Context, ids []int64) ([]*domain.SysRole, error)
	SelectIdsByUserId(e echo.Context, userId any) ([]int64, error)
	SelectRoleKeyListByUserId(e echo.Context, userId any) ([]string, error)
	SelectRoleNameListByUserId(e echo.Context, userId any) ([]string, error)
	SelectAll(e echo.Context) ([]*domain.SysRole, error)
	VerifyRoleNameExist(e echo.Context, sysRole *domain.SysRole) (bool, error)
	VerifyRoleKeyExist(e echo.Context, sysRole *domain.SysRole) (bool, error)
	VerifyRoleIsAllowed(role *domain.SysRole) bool
}
