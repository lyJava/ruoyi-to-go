package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysDictTypeServiceImpl 字典类型服务实现
type SysDictTypeServiceImpl struct {
	SysDictTypeDao dao.SysDictTypeDao
}

// NewSysDictTypeServiceImpl 字典类型服务实现创建
func NewSysDictTypeServiceImpl(s *daoImpl.SysDictTypeDaoImpl) *SysDictTypeServiceImpl {
	return &SysDictTypeServiceImpl{
		SysDictTypeDao: s,
	}
}

func (s *SysDictTypeServiceImpl) Insert(e echo.Context, dept *domain.SysDictType) (*domain.SysDictType, error) {
	return s.SysDictTypeDao.Insert(e.Request().Context(), dept)
}

func (s *SysDictTypeServiceImpl) Update(e echo.Context, dept *domain.SysDictType) (int64, error) {
	return s.SysDictTypeDao.Update(e.Request().Context(), dept)
}

func (s *SysDictTypeServiceImpl) BatchInsert(e echo.Context, list []*domain.SysDictType) ([]int64, error) {
	return s.SysDictTypeDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysDictTypeServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysDictType, error) {
	if id == 0 {
		return nil, errors.New("字典数据详情参数验证失败")
	}
	return s.SysDictTypeDao.SelectById(e.Request().Context(), id)
}

func (s *SysDictTypeServiceImpl) SelectPage(e echo.Context, param *request.DictTypePageParam) ([]*domain.SysDictType, int64, int64, error) {
	return s.SysDictTypeDao.SelectPage(e.Request().Context(), param)
}

func (s *SysDictTypeServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("字典数据批量新增参数验证失败")
	}
	return s.SysDictTypeDao.BatchDelete(e.Request().Context(), ids)
}

func (s *SysDictTypeServiceImpl) SelectAll(e echo.Context) ([]*domain.SysDictType, error) {
	return s.SysDictTypeDao.SelectAll(e.Request().Context())
}

func (s *SysDictTypeServiceImpl) ClearCache(e echo.Context) error {
	return s.SysDictTypeDao.ClearCache(e.Request().Context())
}

func (s *SysDictTypeServiceImpl) RefreshCache(e echo.Context) error {
	return s.SysDictTypeDao.RefreshCache(e.Request().Context())
}
