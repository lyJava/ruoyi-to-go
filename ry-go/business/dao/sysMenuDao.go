package dao

import (
	"context"
	"ry-go/business/domain"
	"ry-go/common/model"
	"ry-go/common/request"
)

// SysMenuDao 系统菜单Dao接口
type SysMenuDao interface {
	Insert(ctx context.Context, sysMenu *domain.SysMenu) (*domain.SysMenu, error)
	Update(ctx context.Context, sysMenu *domain.SysMenu) (int64, error)
	BatchInsert(ctx context.Context, list []*domain.SysMenu) ([]int64, error)
	SelectById(ctx context.Context, id int64) (*domain.SysMenu, error)
	SelectList(ctx context.Context, param *request.MenuPageParam) ([]*domain.SysMenu, error)
	SelectPage(ctx context.Context, param *request.MenuPageParam) ([]*domain.SysMenu, int64, int64, error)
	BatchDelete(ctx context.Context, ids []any) (int64, error)
	SelectListByUserId(ctx context.Context, userId int64, menuName, visible, beginTime, endTime string) ([]*domain.SysMenu, error)
	SelectMenuPermsByUserId(ctx context.Context, userId int64) ([]string, error)
	SelectTreeAll(ctx context.Context) ([]*domain.SysMenu, error)
	SelectTreeByUserId(ctx context.Context, userId int64) ([]*domain.SysMenu, error)
	// BuildMenus 构建前端路由所需要的菜单 构建前端路由所需要的菜单
	//
	// 参数
	//
	// 	-menuList: 菜单切片
	//
	// 返回
	//
	// 	-[]*model.RouterVo: 路由列表
	// 	-error: 错误信息
	BuildMenus(menuList []*domain.SysMenu) ([]*model.RouterVo, error)
	// BuildMenuTree 构建前端所需要树结构
	//
	// 参数
	//
	// 	-menuList: 菜单切片
	//
	// 返回
	//
	// 	-[]*domain.SysMenu: 菜单切片
	BuildMenuTree(menuList []*domain.SysMenu) ([]*domain.SysMenu, error)
	SelectMenuList(ctx context.Context, menu *domain.SysMenu) ([]*domain.SysMenu, error)
	SelectMenuListMenuAndUserId(ctx context.Context, menu *domain.SysMenu, userId int64) ([]*domain.SysMenu, error)
	SelectListByMenuAndUserId(ctx context.Context, menu *domain.SysMenu, userId int64) ([]*domain.SysMenu, error)
	BuildMenuTreeSelect(menuList []*domain.SysMenu) []*domain.TreeSelect
	BuildRoleMenuTreeSelect(ctx context.Context, userId, roleId int64) (*domain.RoleMenuTreeSelect, error)
	SelectMenuListByRoleId(ctx context.Context, roleId int64) ([]int64, error)
}
