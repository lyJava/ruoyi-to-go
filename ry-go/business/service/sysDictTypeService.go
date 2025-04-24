package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysDictTypeService 字典类型服务接口
type SysDictTypeService interface {
	Insert(e echo.Context, sysDictType *domain.SysDictType) (*domain.SysDictType, error)
	Update(e echo.Context, sysDictType *domain.SysDictType) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysDictType) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysDictType, error)
	SelectPage(e echo.Context, param *request.DictTypePageParam) ([]*domain.SysDictType, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectAll(e echo.Context) ([]*domain.SysDictType, error)
	ClearCache(e echo.Context) error
	RefreshCache(e echo.Context) error
}
