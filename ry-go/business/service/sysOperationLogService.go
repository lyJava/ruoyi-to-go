package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysOperationLogService 操作日志服务接口
type SysOperationLogService interface {
	Insert(e echo.Context, sysOperationLog *domain.SysOperationLog) (*domain.SysOperationLog, error)
	Update(e echo.Context, sysOperationLog *domain.SysOperationLog) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysOperationLog) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysOperationLog, error)
	SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysOperationLog, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
}
