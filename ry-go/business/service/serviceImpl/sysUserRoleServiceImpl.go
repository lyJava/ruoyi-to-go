package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysUserRoleServiceImpl 用户角色信息服务实现
type SysUserRoleServiceImpl struct {
	SysUserRoleDao dao.SysUserRoleDao
}

// NewSysUserRoleServiceImpl 用户角色信息服务实现创建
func NewSysUserRoleServiceImpl(s *daoImpl.SysUserRoleDaoImpl) *SysUserRoleServiceImpl {
	return &SysUserRoleServiceImpl{
		SysUserRoleDao: s,
	}
}

func (s *SysUserRoleServiceImpl) Insert(e echo.Context, dept *domain.SysUserRole) (*domain.SysUserRole, error) {
	return s.SysUserRoleDao.Insert(e.Request().Context(), dept)
}

func (s *SysUserRoleServiceImpl) Update(e echo.Context, dept *domain.SysUserRole) (int64, error) {
	return s.SysUserRoleDao.Update(e.Request().Context(), dept)
}

func (s *SysUserRoleServiceImpl) BatchInsert(e echo.Context, list []*domain.SysUserRole) ([]int64, error) {
	return s.SysUserRoleDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysUserRoleServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysUserRole, error) {
	if id == 0 {
		return nil, errors.New("用户角色信息详情参数验证失败")
	}
	return s.SysUserRoleDao.SelectById(e.Request().Context(), id)
}

func (s *SysUserRoleServiceImpl) SelectPage(e echo.Context, param *request.SysNoticePageParam) ([]*domain.SysUserRole, int64, int64, error) {
	return s.SysUserRoleDao.SelectPage(e.Request().Context(), param)
}

func (s *SysUserRoleServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("用户角色信息批量新增参数验证失败")
	}
	return s.SysUserRoleDao.BatchDelete(e.Request().Context(), ids)
}
