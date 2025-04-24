package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysJobService 系统任务服务接口
type SysJobService interface {
	Insert(e echo.Context, sysJob *domain.SysJob) (*domain.SysJob, error)
	Update(e echo.Context, sysJob *domain.SysJob) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysJob) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysJob, error)
	SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysJob, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	UpdateJobStatus(e echo.Context, id int64, status string) (int64, error)
	UpdateJobStatusBatch(e echo.Context, list []*request.JobChangeParam) (int64, error)
}
