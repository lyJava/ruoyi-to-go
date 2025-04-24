package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysPostServiceImpl 岗位信息服务实现
type SysPostServiceImpl struct {
	SysPostDao dao.SysPostDao
}

// NewSysPostServiceImpl 岗位信息服务实现创建
func NewSysPostServiceImpl(s *daoImpl.SysPostDaoImpl) *SysPostServiceImpl {
	return &SysPostServiceImpl{
		SysPostDao: s,
	}
}

func (s *SysPostServiceImpl) Insert(e echo.Context, dept *domain.SysPost) (*domain.SysPost, error) {
	return s.SysPostDao.Insert(e.Request().Context(), dept)
}

func (s *SysPostServiceImpl) Update(e echo.Context, dept *domain.SysPost) (int64, error) {
	return s.SysPostDao.Update(e.Request().Context(), dept)
}

func (s *SysPostServiceImpl) BatchInsert(e echo.Context, list []*domain.SysPost) ([]int64, error) {
	return s.SysPostDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysPostServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysPost, error) {
	if id == 0 {
		return nil, errors.New("岗位信息详情参数验证失败")
	}
	return s.SysPostDao.SelectById(e.Request().Context(), id)
}

func (s *SysPostServiceImpl) SelectPage(e echo.Context, param *request.PostPageParam) ([]*domain.SysPost, int64, int64, error) {
	return s.SysPostDao.SelectPage(e.Request().Context(), param)
}

func (s *SysPostServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("岗位信息批量新增参数验证失败")
	}
	return s.SysPostDao.BatchDelete(e.Request().Context(), ids)
}

func (s *SysPostServiceImpl) SelectAll(e echo.Context) ([]*domain.SysPost, error) {
	return s.SysPostDao.SelectAll(e.Request().Context())
}
