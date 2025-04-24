package dao

import (
	"context"
	"ry-go/business/domain"
)

type SysUserPostDao interface {
	BatchInsert(ctx context.Context, list []*domain.SysUserPost) (int64, error)
	DeleteByUserId(ctx context.Context, userId int64) (int64, error)
}
