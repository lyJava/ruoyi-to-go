package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysConfigService 系统配置服务接口
type SysConfigService interface {
	Insert(e echo.Context, sysConfig *domain.SysConfig) (*domain.SysConfig, error)
	Update(e echo.Context, sysConfig *domain.SysConfig) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysConfig) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysConfig, error)
	SelectPage(e echo.Context, param *request.ConfigPageParam) ([]*domain.SysConfig, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectConfigByKey(e echo.Context, key string) (*domain.SysConfig, error)
	SelectConfigList(e echo.Context) ([]*domain.SysConfig, error)
	RefreshConfigCache(e echo.Context, list []*domain.SysConfig) error
	ClearConfigCache(e echo.Context, list []*domain.SysConfig) error
}
