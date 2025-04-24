package serviceImpl

import (
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/common/model"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysMenuServiceImpl 系统菜单服务实现
type SysMenuServiceImpl struct {
	SysMenuDao dao.SysMenuDao
}

// NewSysMenuServiceImpl 系统菜单服务实现注入
func NewSysMenuServiceImpl(s *daoImpl.SysMenuDaoImpl) *SysMenuServiceImpl {
	return &SysMenuServiceImpl{
		SysMenuDao: s,
	}
}

func (s *SysMenuServiceImpl) Insert(e echo.Context, dept *domain.SysMenu) (*domain.SysMenu, error) {
	return s.SysMenuDao.Insert(e.Request().Context(), dept)
}

func (s *SysMenuServiceImpl) Update(e echo.Context, dept *domain.SysMenu) (int64, error) {
	return s.SysMenuDao.Update(e.Request().Context(), dept)
}

func (s *SysMenuServiceImpl) BatchInsert(e echo.Context, list []*domain.SysMenu) ([]int64, error) {
	return s.SysMenuDao.BatchInsert(e.Request().Context(), list)
}

func (s *SysMenuServiceImpl) SelectById(e echo.Context, id int64) (*domain.SysMenu, error) {
	return s.SysMenuDao.SelectById(e.Request().Context(), id)
}

func (s *SysMenuServiceImpl) SelectList(e echo.Context, param *request.MenuPageParam) ([]*domain.SysMenu, error) {
	return s.SysMenuDao.SelectList(e.Request().Context(), param)
}

func (s *SysMenuServiceImpl) SelectPage(e echo.Context, param *request.MenuPageParam) ([]*domain.SysMenu, int64, int64, error) {
	return s.SysMenuDao.SelectPage(e.Request().Context(), param)
}

func (s *SysMenuServiceImpl) BatchDelete(e echo.Context, ids []any) (int64, error) {
	return s.SysMenuDao.BatchDelete(e.Request().Context(), ids)
}

func (s *SysMenuServiceImpl) SelectListByUserId(e echo.Context, userId int64, menuName, visible, beginTime, endTime string) ([]*domain.SysMenu, error) {
	return s.SysMenuDao.SelectListByUserId(e.Request().Context(), userId, menuName, visible, beginTime, endTime)
}

func (s *SysMenuServiceImpl) SelectMenuPermsByUserId(e echo.Context, userId int64) ([]string, error) {
	return s.SysMenuDao.SelectMenuPermsByUserId(e.Request().Context(), userId)
}

func (s *SysMenuServiceImpl) SelectTreeAll(e echo.Context) ([]*domain.SysMenu, error) {
	return s.SysMenuDao.SelectTreeAll(e.Request().Context())
}

func (s *SysMenuServiceImpl) SelectTreeByUserId(e echo.Context, userId int64) ([]*domain.SysMenu, error) {
	return s.SysMenuDao.SelectTreeByUserId(e.Request().Context(), userId)
}

func (s *SysMenuServiceImpl) BuildMenus(menuList []*domain.SysMenu) ([]*model.RouterVo, error) {
	return s.SysMenuDao.BuildMenus(menuList)
}

func (s *SysMenuServiceImpl) BuildMenuTree(menuList []*domain.SysMenu) ([]*domain.SysMenu, error) {
	return s.SysMenuDao.BuildMenuTree(menuList)
}

func (s *SysMenuServiceImpl) SelectListByMenuAndUserId(e echo.Context, menu *domain.SysMenu, userId int64) ([]*domain.SysMenu, error) {
	return s.SysMenuDao.SelectListByMenuAndUserId(e.Request().Context(), menu, userId)
}

func (s *SysMenuServiceImpl) BuildMenuTreeSelect(menuList []*domain.SysMenu) []*domain.TreeSelect {
	return s.SysMenuDao.BuildMenuTreeSelect(menuList)
}

func (s *SysMenuServiceImpl) BuildRoleMenuTreeSelect(e echo.Context, userId, roleId int64) (*domain.RoleMenuTreeSelect, error) {
	return s.SysMenuDao.BuildRoleMenuTreeSelect(e.Request().Context(), userId, roleId)
}
