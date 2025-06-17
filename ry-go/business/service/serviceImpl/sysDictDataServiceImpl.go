package serviceImpl

import (
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysDictDataServiceImpl 字典数据服务实现
type SysDictDataServiceImpl struct {
	SysDictDataDao dao.SysDictDataDao
}

// NewSysDictDataServiceImpl 字典数据服务实现创建
func NewSysDictDataServiceImpl(s *daoImpl.SysDictDataDaoImpl) *SysDictDataServiceImpl {
	return &SysDictDataServiceImpl{
		SysDictDataDao: s,
	}
}

func (s *SysDictDataServiceImpl) Insert(e echo.Context, dept *domain.SysDictData) (*domain.SysDictData, error) {
	return s.SysDictDataDao.Insert(e.Request().Context(), dept)
}

func (s *SysDictDataServiceImpl) Update(e echo.Context, dept *domain.SysDictData) (int64, error) {
	return s.SysDictDataDao.Update(e.Request().Context(), dept)
}

func (s *SysDictDataServiceImpl) BatchInsert(e echo.Context, list []*domain.SysDictData) ([]int64, error) {
	return s.SysDictDataDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysDictDataServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysDictData, error) {
	return s.SysDictDataDao.SelectById(e.Request().Context(), id)
}

func (s *SysDictDataServiceImpl) SelectPage(e echo.Context, param *request.DictDataPageParam) ([]*domain.SysDictData, int64, int64, error) {
	return s.SysDictDataDao.SelectPage(e.Request().Context(), param)
}

func (s *SysDictDataServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	return s.SysDictDataDao.BatchDelete(e.Request().Context(), ids)
}

func (s *SysDictDataServiceImpl) SelectDictDataByType(e echo.Context, dictType string) ([]*domain.SysDictData, error) {
	return s.SysDictDataDao.SelectDictDataByType(e.Request().Context(), dictType)
}

func (s *SysDictDataServiceImpl) SelectAll(e echo.Context) ([]*domain.SysDictData, error) {
	return s.SysDictDataDao.SelectAll(e.Request().Context())
}
