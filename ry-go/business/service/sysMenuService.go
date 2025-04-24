package service

import (
	"ry-go/business/domain"
	"ry-go/common/model"
	"ry-go/common/request"

	"github.com/labstack/echo/v4"
)

// SysMenuService 系统菜单服务接口
type SysMenuService interface {
	Insert(e echo.Context, sysMenu *domain.SysMenu) (*domain.SysMenu, error)
	Update(e echo.Context, sysMenu *domain.SysMenu) (int64, error)
	BatchInsert(e echo.Context, list []*domain.SysMenu) ([]int64, error)
	SelectById(e echo.Context, id int64) (*domain.SysMenu, error)
	SelectList(e echo.Context, param *request.MenuPageParam) ([]*domain.SysMenu, error)
	SelectPage(e echo.Context, param *request.MenuPageParam) ([]*domain.SysMenu, int64, int64, error)
	BatchDelete(e echo.Context, ids []any) (int64, error)
	SelectListByUserId(e echo.Context, userId int64, menuName, visible, beginTime, endTime string) ([]*domain.SysMenu, error)
	SelectMenuPermsByUserId(e echo.Context, userId int64) ([]string, error)
	SelectTreeAll(e echo.Context) ([]*domain.SysMenu, error)
	SelectTreeByUserId(e echo.Context, userId int64) ([]*domain.SysMenu, error)
	BuildMenus(menuList []*domain.SysMenu) ([]*model.RouterVo, error)
	BuildMenuTree(menuList []*domain.SysMenu) ([]*domain.SysMenu, error)
	SelectListByMenuAndUserId(e echo.Context, menu *domain.SysMenu, userId int64) ([]*domain.SysMenu, error)
	BuildMenuTreeSelect(menuList []*domain.SysMenu) []*domain.TreeSelect
	BuildRoleMenuTreeSelect(e echo.Context, userId, roleId int64) (*domain.RoleMenuTreeSelect, error)
}
