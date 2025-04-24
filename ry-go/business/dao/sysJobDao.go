package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"
)

// SysJobDao 系统任务Dao接口
type SysJobDao interface {
	Insert(ctx context.Context, sysJob *domain.SysJob) (*domain.SysJob, error)
	Update(ctx context.Context, sysJob *domain.SysJob) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysJob) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysJob, error)
	SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysJob, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	UpdateJobStatus(ctx context.Context, id int64, status string) (int64, error)
	UpdateJobStatusBatch(ctx context.Context, list []*request.JobChangeParam) (int64, error)
}
