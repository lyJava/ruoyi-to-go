package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysRoleDao 系统角色Dao接口
type SysRoleDao interface {
	Insert(ctx context.Context, sysRole *domain.SysRole) (*domain.SysRole, error)
	Update(ctx context.Context, sysRole *domain.SysRole) (int64, error)
	UpdateWithOutMenuAndDataScope(ctx context.Context, sysRole *domain.SysRole) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysRole) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysRole, error)
	SelectPage(ctx context.Context, param *request.RolePageParam) ([]*domain.SysRole, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectList(ctx context.Context, ids []int64) ([]*domain.SysRole, error)
	SelectIdsByUserId(ctx context.Context, userId any) ([]int64, error)
	SelectRoleKeyListByUserId(ctx context.Context, userId any) ([]string, error)
	SelectAll(ctx context.Context) ([]*domain.SysRole, error)
	SelectAllFilterAdmin(list []*domain.SysRole, isAdmin bool) ([]*domain.SysRole, error)
	SelectRoleNameListByUserId(ctx context.Context, userId any) ([]string, error)
	SelectRoleByRoleName(ctx context.Context, sysRole *domain.SysRole) (*domain.SysRole, error)
	SelectCountByRoleKey(ctx context.Context, sysRole *domain.SysRole) (*domain.SysRole, error)
	// SelectRolePermsByUserId 根据ID查询用户角色切片
	SelectRolePermsByUserId(ctx context.Context, userId int64) ([]string, error)
}
