package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysUserRoleService 用户角色信息服务接口
type SysUserRoleService interface {
	Insert(e echo.Context, sysUserRole *domain.SysUserRole) (*domain.SysUserRole, error)
	Update(e echo.Context, sysUserRole *domain.SysUserRole) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysUserRole) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysUserRole, error)
	SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysUserRole, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
}
