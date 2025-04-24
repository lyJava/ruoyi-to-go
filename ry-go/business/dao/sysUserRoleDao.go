package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysUserRoleDao 用户角色信息Dao接口
type SysUserRoleDao interface {
	Insert(ctx context.Context, sysUserRole *domain.SysUserRole) (*domain.SysUserRole, error)
	Update(ctx context.Context, sysUserRole *domain.SysUserRole) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysUserRole) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysUserRole, error)
	SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysUserRole, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	DeleteByUserId(ctx context.Context, userId int64) (int64, error)
	DeleteByUserIdList(ctx context.Context, userIds []int64) (int64, error)
}
