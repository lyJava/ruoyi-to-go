package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysLoginRecordDao 登录记录Dao接口
type SysLoginRecordDao interface {
	Insert(ctx context.Context, sysLoginRecord *domain.SysLoginRecord) (*domain.SysLoginRecord, error)
	Update(ctx context.Context, sysLoginRecord *domain.SysLoginRecord) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysLoginRecord) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysLoginRecord, error)
	SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysLoginRecord, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectLastRecode(ctx context.Context, username string) (*domain.SysLoginRecord, error)
}
