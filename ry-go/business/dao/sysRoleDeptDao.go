package dao

import (
	"context"
	"ry-go/business/domain"
)

type SysRoleDeptDao interface {
	DeleteByRoleId(ctx context.Context, roleId int64) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysRoleDept) (int64, error)
}
