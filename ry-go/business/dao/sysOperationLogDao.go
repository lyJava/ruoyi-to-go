package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysOperationLogDao 操作日志Dao接口
type SysOperationLogDao interface {
	Insert(ctx context.Context, sysOperationLog *domain.SysOperationLog) (*domain.SysOperationLog, error)
	Update(ctx context.Context, sysOperationLog *domain.SysOperationLog) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysOperationLog) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysOperationLog, error)
	SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysOperationLog, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
}
