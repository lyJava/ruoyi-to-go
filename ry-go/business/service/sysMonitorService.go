package service

import (
	"github.com/labstack/echo/v4"
	"ry-go/business/domain"
	"ry-go/common/model"
)

type SysMonitorService interface {
	OnlineList(e echo.Context, keyPrefix string) ([]*domain.LoginUser, error)
	OnlineForceLogout(e echo.Context, cacheKey string) (bool, error)
	CacheList(e echo.Context) (map[string]any, error)
	CacheListByKey(e echo.Context, cacheKey string) ([]string, error)
	CacheDetail(e echo.Context, prefix, suffix string) (*model.Cache, error)
	CacheClear(e echo.Context, prefix string) error
	CacheClearAll(e echo.Context) error
}
