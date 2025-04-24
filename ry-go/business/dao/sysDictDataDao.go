package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysDictDataDao 字典数据Dao接口
type SysDictDataDao interface {
	Insert(ctx context.Context, sysDictData *domain.SysDictData) (*domain.SysDictData, error)
	Update(ctx context.Context, sysDictData *domain.SysDictData) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysDictData) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysDictData, error)
	SelectPage(ctx context.Context, param *request.DictDataPageParam) ([]*domain.SysDictData, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectDictDataByType(ctx context.Context, dictType string) ([]*domain.SysDictData, error)
}
