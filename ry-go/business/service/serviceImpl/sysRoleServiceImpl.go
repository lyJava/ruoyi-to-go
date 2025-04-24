package serviceImpl

import (
	"errors"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysRoleServiceImpl 系统角色服务实现
type SysRoleServiceImpl struct {
	SysRoleDao dao.SysRoleDao
}

// NewSysRoleServiceImpl 系统角色服务实现注入
func NewSysRoleServiceImpl(s *daoImpl.SysRoleDaoImpl) *SysRoleServiceImpl {
	return &SysRoleServiceImpl{
		SysRoleDao: s,
	}
}

func (s *SysRoleServiceImpl) Insert(e echo.Context, dept *domain.SysRole) (*domain.SysRole, error) {
	return s.SysRoleDao.Insert(e.Request().Context(), dept)
}

func (s *SysRoleServiceImpl) Update(e echo.Context, dept *domain.SysRole) (int64, error) {
	return s.SysRoleDao.Update(e.Request().Context(), dept)
}

func (s *SysRoleServiceImpl) UpdateWithOutMenuAndDataScope(e echo.Context, sysRole *domain.SysRole) (int64, error) {
	return s.SysRoleDao.UpdateWithOutMenuAndDataScope(e.Request().Context(), sysRole)
}

func (s *SysRoleServiceImpl) BatchInsert(e echo.Context, list []*domain.SysRole) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("系统角色批量新增参数验证失败")
	}
	return s.SysRoleDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysRoleServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysRole, error) {
	if id == 0 {
		return nil, errors.New("查询参数验证失败")
	}
	return s.SysRoleDao.SelectById(e.Request().Context(), id)
}

func (s *SysRoleServiceImpl) SelectPage(e echo.Context, param *request.RolePageParam) ([]*domain.SysRole, int64, int64, error) {
	return s.SysRoleDao.SelectPage(e.Request().Context(), param)
}

func (s *SysRoleServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统角色批量删除参数验证失败")
	}
	return s.SysRoleDao.BatchDelete(e.Request().Context(), ids)
}

func (s *SysRoleServiceImpl) SelectList(e echo.Context, ids []int64) ([]*domain.SysRole, error) {
	if len(ids) <= 0 {
		return nil, errors.New("系统角色查询数验证失败")
	}
	return s.SysRoleDao.SelectList(e.Request().Context(), ids)
}

func (s *SysRoleServiceImpl) SelectIdsByUserId(e echo.Context, userId any) ([]int64, error) {
	if userId == nil {
		return nil, errors.New("系统角色查询角色ID不能为空")
	}
	return s.SysRoleDao.SelectIdsByUserId(e.Request().Context(), userId)
}

func (s *SysRoleServiceImpl) SelectRoleKeyListByUserId(e echo.Context, userId any) ([]string, error) {
	if userId == nil {
		return nil, errors.New("系统角色查询角色ID不能为空")
	}
	return s.SysRoleDao.SelectRoleKeyListByUserId(e.Request().Context(), userId)
}

func (s *SysRoleServiceImpl) SelectRoleNameListByUserId(e echo.Context, userId any) ([]string, error) {
	if userId == nil {
		return nil, errors.New("系统角色查询角色ID不能为空")
	}
	return s.SysRoleDao.SelectRoleNameListByUserId(e.Request().Context(), userId)
}

func (s *SysRoleServiceImpl) SelectAll(e echo.Context) ([]*domain.SysRole, error) {
	return s.SysRoleDao.SelectAll(e.Request().Context())
}

func (s *SysRoleServiceImpl) VerifyRoleNameExist(e echo.Context, role *domain.SysRole) (bool, error) {
	sysRole, err := s.SysRoleDao.SelectRoleByRoleName(e.Request().Context(), role)
	if err != nil {
		return false, err
	}
	if sysRole == nil {
		return false, nil
	}
	if sysRole.Id == sysRole.Id {
		return false, err
	}
	return true, nil
}

func (s *SysRoleServiceImpl) VerifyRoleKeyExist(e echo.Context, role *domain.SysRole) (bool, error) {
	sysRole, err := s.SysRoleDao.SelectCountByRoleKey(e.Request().Context(), role)
	if err != nil {
		return false, err
	}
	if sysRole == nil {
		return false, nil
	}
	if sysRole.Id == sysRole.Id {
		return false, err
	}
	return true, nil
}

func (s *SysRoleServiceImpl) VerifyRoleIsAllowed(role *domain.SysRole) bool {
	if role == nil {
		return false
	}
	if domain.IsAdminRole(role.Id) {
		return true
	} else {
		return false
	}
}
