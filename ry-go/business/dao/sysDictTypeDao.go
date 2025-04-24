package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysDictTypeDao 字典类型Dao接口
type SysDictTypeDao interface {
	Insert(ctx context.Context, sysDictType *domain.SysDictType) (*domain.SysDictType, error)
	Update(ctx context.Context, sysDictType *domain.SysDictType) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysDictType) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysDictType, error)
	SelectPage(ctx context.Context, param *request.DictTypePageParam) ([]*domain.SysDictType, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectAll(ctx context.Context) ([]*domain.SysDictType, error)
	ClearCache(ctx context.Context) error
	RefreshCache(ctx context.Context) error
}
