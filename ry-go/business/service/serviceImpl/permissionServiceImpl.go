package serviceImpl

import (
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/utils"

	"github.com/labstack/echo/v4"
)

type PermissionServiceImpl struct {
	RoleDao dao.SysRoleDao
	MenuDao dao.SysMenuDao
}

func NewPermissionServiceImpl(s *daoImpl.SysRoleDaoImpl, m *daoImpl.SysMenuDaoImpl) *PermissionServiceImpl {
	return &PermissionServiceImpl{
		RoleDao: s,
		MenuDao: m,
	}
}

func (impl *PermissionServiceImpl) SelectRolePermission(e echo.Context, userId int64) ([]string, error) {
	var permissions []string
	isAdmin := domain.IsAdmin(userId)
	if isAdmin {
		permissions = append(permissions, "admin")
	} else {
		permissionList, err := impl.RoleDao.SelectRolePermsByUserId(e.Request().Context(), userId)
		if err != nil {
			return nil, err
		}
		permissions = permissionList
	}
	return permissions, nil
}

func (impl *PermissionServiceImpl) SelectMenuPermission(e echo.Context, userId int64) ([]string, error) {
	var permissions []string
	isAdmin := domain.IsAdmin(userId)
	if isAdmin {
		permissions = append(permissions, "*:*:*")
	} else {
		permissionList, err := impl.MenuDao.SelectMenuPermsByUserId(e.Request().Context(), userId)
		if err != nil {
			return nil, err
		}
		permissions = utils.DeduplicateStrings(permissionList)
	}
	return permissions, nil
}
