package service

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysLoginRecordService 登录记录服务接口
type SysLoginRecordService interface {
	Insert(ctx context.Context, sysLoginRecord *domain.SysLoginRecord) (*domain.SysLoginRecord, error)
	Update(e echo.Context, sysLoginRecord *domain.SysLoginRecord) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysLoginRecord) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysLoginRecord, error)
	SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysLoginRecord, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectLastRecode(e echo.Context, username string) (*domain.SysLoginRecord, error)
}
