package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysConfigDao 系统配置Dao接口
type SysConfigDao interface {
	Insert(ctx context.Context, sysConfig *domain.SysConfig) (*domain.SysConfig, error)
	Update(ctx context.Context, sysConfig *domain.SysConfig) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysConfig) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysConfig, error)
	SelectPage(ctx context.Context, param *request.ConfigPageParam) ([]*domain.SysConfig, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectConfigByKey(ctx context.Context, key string) (*domain.SysConfig, error)
}
