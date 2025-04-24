package service

import (
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysDictDataService 字典数据服务接口
type SysDictDataService interface {
	Insert(e echo.Context, sysDictData *domain.SysDictData) (*domain.SysDictData, error)
	Update(e echo.Context, sysDictData *domain.SysDictData) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysDictData) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysDictData, error)
	SelectPage(e echo.Context, param *request.DictDataPageParam) ([]*domain.SysDictData, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectDictDataByType(e echo.Context, dictType string) ([]*domain.SysDictData, error)
}
