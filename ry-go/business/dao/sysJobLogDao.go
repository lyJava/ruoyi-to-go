package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysJobLogDao 系统任务日志Dao接口
type SysJobLogDao interface {
	Insert(ctx context.Context, sysJobLog *domain.SysJobLog) (*domain.SysJobLog, error)
	Update(ctx context.Context, sysJobLog *domain.SysJobLog) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysJobLog) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysJobLog, error)
	SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysJobLog, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
}
