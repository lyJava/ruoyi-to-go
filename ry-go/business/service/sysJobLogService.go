package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysJobLogService 系统任务日志服务接口
type SysJobLogService interface {
	Insert(e echo.Context, sysJobLog *domain.SysJobLog) (*domain.SysJobLog, error)
	Update(e echo.Context, sysJobLog *domain.SysJobLog) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysJobLog) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysJobLog, error)
	SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysJobLog, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
}
