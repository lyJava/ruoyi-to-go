package dao

import (
	"context"
	"ry-go/business/domain"
)

type SysRoleMenuDao interface {
	DeleteByRoleId(ctx context.Context, roleId int64) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysRoleMenu) (int64, error)
}
