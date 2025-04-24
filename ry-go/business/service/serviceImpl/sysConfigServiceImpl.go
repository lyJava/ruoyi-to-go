package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysConfigServiceImpl 系统配置服务实现
type SysConfigServiceImpl struct {
	SysConfigDao dao.SysConfigDao
}

// NewSysConfigServiceImpl 系统配置服务实现创建
func NewSysConfigServiceImpl(s *daoImpl.SysConfigDaoImpl) *SysConfigServiceImpl {
	return &SysConfigServiceImpl{
		SysConfigDao: s,
	}
}

func (s *SysConfigServiceImpl) Insert(e echo.Context, dept *domain.SysConfig) (*domain.SysConfig, error) {
	return s.SysConfigDao.Insert(e.Request().Context(), dept)
}

func (s *SysConfigServiceImpl) Update(e echo.Context, dept *domain.SysConfig) (int64, error) {
	return s.SysConfigDao.Update(e.Request().Context(), dept)
}

func (s *SysConfigServiceImpl) BatchInsert(e echo.Context, list []*domain.SysConfig) ([]int64, error) {
	return s.SysConfigDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysConfigServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysConfig, error) {
	if id == 0 {
		return nil, errors.New("系统配置详情参数验证失败")
	}
	return s.SysConfigDao.SelectById(e.Request().Context(), id)
}

func (s *SysConfigServiceImpl) SelectPage(e echo.Context, param *request.ConfigPageParam) ([]*domain.SysConfig, int64, int64, error) {
	return s.SysConfigDao.SelectPage(e.Request().Context(), param)
}

func (s *SysConfigServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统配置批量新增参数验证失败")
	}
	return s.SysConfigDao.BatchDelete(e.Request().Context(), ids)
}

func (s *SysConfigServiceImpl) SelectConfigByKey(e echo.Context, key string) (*domain.SysConfig, error) {
	return s.SysConfigDao.SelectConfigByKey(e.Request().Context(), key)
}
