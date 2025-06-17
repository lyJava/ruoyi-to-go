package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysNoticeService 系统公告服务接口
type SysNoticeService interface {
	Insert(e echo.Context, dept *domain.SysNotice) (*domain.SysNotice, error)
	Update(e echo.Context, dept *domain.SysNotice) (int64, error)
	BatchSave(e echo.Context, list []*domain.SysNotice) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysNotice, error)
	SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysNotice, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectAll(e echo.Context) ([]*domain.SysNotice, error)
}
