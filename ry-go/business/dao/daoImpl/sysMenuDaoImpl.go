package daoImpl

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"ry-go/business/dao"
	"ry-go/business/domain"
	"ry-go/common/model"
	"ry-go/common/request"
	"ry-go/utils"
	"strings"

	"gorm.io/gorm"
)

// SysMenuDaoImpl 系统菜单Dao实现结构体
type SysMenuDaoImpl struct {
	Gorm    *gorm.DB
	RoleDao dao.SysRoleDao
}

// NewSysMenuDaoImpl 系统菜单Dao实现创建
func NewSysMenuDaoImpl(db *gorm.DB, r *SysRoleDaoImpl) *SysMenuDaoImpl {
	return &SysMenuDaoImpl{
		Gorm:    db,
		RoleDao: r,
	}
}

func (impl *SysMenuDaoImpl) Insert(ctx context.Context, sysMenu *domain.SysMenu) (*domain.SysMenu, error) {
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysMenu{}).Create(&sysMenu).Error; err != nil {
		return nil, errors.New("系统菜单新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysMenu
	sysMenu, err := impl.SelectById(ctx, sysMenu.Id)
	if err != nil {
		return nil, errors.New("系统菜单新增失败")
	}
	return sysMenu, nil
}

func (impl *SysMenuDaoImpl) Update(ctx context.Context, sysMenu *domain.SysMenu) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Save(&sysMenu)
	if result.Error != nil {
		return 0, errors.New("系统菜单修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysMenuDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysMenu) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("系统菜单批量新增参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		return nil, errors.New("系统菜单批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.Id)
	}

	return ids, nil
}

func (impl *SysMenuDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysMenu, error) {
	var sysMenu *domain.SysMenu
	// 主键为id的时候First(&sysMenu, id)默认使用主键ID查询，且会增加limit 1
	if err := impl.Gorm.WithContext(ctx).
		Model(&domain.SysMenu{}).
		Select("*, coalesce(perms, '') AS perms").
		Find(&sysMenu, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("系统菜单未找到")
		}
		return nil, errors.New("系统菜单查询失败")
	}
	return sysMenu, nil
}

func (impl *SysMenuDaoImpl) SelectList(ctx context.Context, param *request.MenuPageParam) ([]*domain.SysMenu, error) {
	var list []*domain.SysMenu
	if domain.IsAdmin(param.UserId) {
		menuList, err := impl.SelectMenuList(ctx, &domain.SysMenu{
			MenuName:   param.MenuName,
			MenuStatus: param.MenuStatus,
			BeginTime:  param.BeginTime,
			EndTime:    param.EndTime,
		})

		if err != nil {
			return nil, errors.New("系统菜单查询失败")
		}

		list = menuList
	} else {
		menuList, err := impl.SelectMenuListMenuAndUserId(ctx, &domain.SysMenu{
			MenuName:   param.MenuName,
			MenuStatus: param.MenuStatus,
			BeginTime:  param.BeginTime,
			EndTime:    param.EndTime,
		}, param.UserId)
		if err != nil {
			return nil, errors.New("系统菜单查询失败")
		}

		list = menuList
	}
	return list, nil
}

func (impl *SysMenuDaoImpl) SelectPage(ctx context.Context, param *request.MenuPageParam) ([]*domain.SysMenu, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysMenu{})
	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		zap.L().Sugar().Error("获取系统菜单总记录数异常:", zap.Error(err))
		return nil, 0, 0, errors.New("获取系统菜单总记录数失败")
	}

	// 查询用户列表
	var list []*domain.SysMenu
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		zap.L().Error("获取系统菜单列表异常",
			zap.Error(err),
			zap.String("function", "SelectPage"),
			zap.Any("param", param),
		)
		return nil, 0, 0, errors.New("获取系统菜单列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysMenuDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统菜单批量删除参数验证失败")
	}
	zap.L().Sugar().Infof("系统菜单批量删除参数:ids===%v", ids)

	result := impl.Gorm.WithContext(ctx).Model(&domain.SysMenu{}).Where("id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		zap.L().Sugar().Error("系统菜单批量删除异常", zap.Error(result.Error))
		return 0, errors.New("批量删除失败")
	}
	zap.L().Sugar().Infof("系统菜单批量删除成功执行条数===%d", result.RowsAffected)
	return result.RowsAffected, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysMenuDaoImpl) SelectListByUserId(ctx context.Context, userId int64, menuName, visible, beginTime, endTime string) ([]*domain.SysMenu, error) {
	const baseQuery = `
        SELECT 
            m.id,
            m.parent_id,
            m.menu_name,
            m.router_path,
            m.component,
            m.visible,
            m.menu_status,
            COALESCE(m.perms, '') AS perms,
            m.is_frame,
            m.is_cache,
            m.menu_type,
            m.icon,
            m.order_num,
            m.create_time
        FROM sys_menu m
        WHERE EXISTS (
            SELECT 1
            FROM sys_role_menu rm
            JOIN sys_user_role ur ON rm.role_id = ur.role_id
            WHERE ur.user_id = @userId
                AND rm.menu_id = m.id
        )`

	// 使用 strings.Builder 安全构建查询
	var queryBuilder strings.Builder
	queryBuilder.WriteString(baseQuery)

	params := map[string]interface{}{
		"userId": userId,
	}

	// 动态添加过滤条件
	if menuName != "" {
		queryBuilder.WriteString(" AND m.menu_name LIKE CONCAT('%%' || @menuName || '%%')")
		params["menuName"] = menuName
	}

	if visible != "" {
		queryBuilder.WriteString(" AND m.visible = @visible")
		params["visible"] = visible
	}

	if beginTime != "" && endTime != "" {
		queryBuilder.WriteString(" AND m.create_time BETWEEN @beginTime AND @endTime")
		params["beginTime"] = beginTime
		params["endTime"] = endTime
	}

	queryBuilder.WriteString(" ORDER BY m.parent_id, m.order_num")

	var list []*domain.SysMenu
	if err := impl.Gorm.WithContext(ctx).Raw(queryBuilder.String(), params).Find(&list).Error; err != nil {
		return list, errors.New("查询用户权限集合失败")
	}
	return list, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysMenuDaoImpl) SelectMenuPermsByUserId(ctx context.Context, userId int64) ([]string, error) {
	var list []string
	if err := impl.Gorm.WithContext(ctx).
		Raw(`
				SELECT DISTINCT
					m.perms
				FROM
					sys_menu m
					LEFT JOIN sys_role_menu rm ON m.id = rm.menu_id
					LEFT JOIN sys_user_role ur ON rm.role_id = ur.role_id
					LEFT JOIN sys_role r ON r.id = ur.role_id
				WHERE
					m.menu_status = '0'
					AND r.role_status = '0'
				    AND m.perms IS NOT NULL
    				AND m.perms <> ''
					AND ur.user_id = ?`, userId).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询角色集合失败")
	}

	return list, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysMenuDaoImpl) SelectTreeAll(ctx context.Context) ([]*domain.SysMenu, error) {
	var list []*domain.SysMenu
	if err := impl.Gorm.WithContext(ctx).
		Raw(
			`SELECT DISTINCT
					m.id,
					m.parent_id,
					m.menu_name,
					m.router_path,
					m.component,
					m.visible,
					m.menu_status,
					COALESCE(m.perms, '') AS perms,
					m.is_frame,
					m.is_cache,
					m.menu_type,
					m.icon,
					m.order_num,
					m.create_time
				FROM
					sys_menu m
				WHERE
					m.menu_type IN ('M', 'C')
					AND m.menu_status = '0'
				ORDER BY
					m.parent_id,
					m.order_num`).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysMenuDaoImpl) SelectTreeByUserId(ctx context.Context, userId int64) ([]*domain.SysMenu, error) {
	var list []*domain.SysMenu
	query := impl.Gorm.WithContext(ctx).Table("sys_menu m").
		Select(`DISTINCT m.id, m.parent_id, m.menu_name, m.router_path, m.component, m.visible, m.menu_status, 
		coalesce(m.perms, '') AS perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time`).
		Joins("LEFT JOIN sys_role_menu rm ON m.id = rm.menu_id").
		Joins("LEFT JOIN sys_user_role ur ON rm.role_id = ur.role_id").
		Joins("LEFT JOIN sys_role ro ON ur.role_id = ro.id").
		Joins("LEFT JOIN sys_user u ON ur.user_id = u.id").
		Where("m.menu_type IN ('M', 'C') AND m.menu_status = '0'").
		Order("m.parent_id, m.order_num")

	if !domain.IsAdmin(userId) {
		query.Where("u.id = ? AND ro.role_status = '0'", userId)
	}

	if err := query.Find(&list).Error; err != nil {
		return nil, err
	}

	return getChildPerms(list, 0), nil
}

// BuildMenus 构建前端路由所需要的菜单
func (impl *SysMenuDaoImpl) BuildMenus(menuList []*domain.SysMenu) ([]*model.RouterVo, error) {
	if len(menuList) <= 0 {
		return nil, errors.New("菜单数据不能为空")
	}
	var routerList []*model.RouterVo

	for _, menu := range menuList {
		routerVo := &model.RouterVo{
			Hidden:    menu.Visible == "1",
			Name:      getRouterName(menu),
			Path:      getRouterPath(menu),
			Component: getComponent(menu),
			Meta:      model.NewMetaVo(menu.MenuName, menu.Icon, menu.IsCache == 1),
		}

		childMenuList := menu.Children
		if len(childMenuList) > 0 && menu.MenuType == "M" {
			routerVo.AlwaysShow = true
			routerVo.Redirect = "noRedirect"
			// 修复递归调用：取第一个返回值忽略错误
			children, _ := impl.BuildMenus(childMenuList)
			routerVo.Children = children
		} else if isMenuFrame(menu) {
			routerChildList := make([]*model.RouterVo, 0)
			// 必须初始化子节点
			children := &model.RouterVo{
				Path:      menu.RouterPath,
				Component: menu.Component,
				Name:      capitalize(menu.MenuName),
				Meta:      model.NewMetaVo(menu.MenuName, menu.Icon, menu.IsCache == 1),
			}
			routerChildList = append(routerList, children)
			routerVo.Children = routerChildList
		}
		routerList = append(routerList, routerVo)
	}
	return routerList, nil
}

func (impl *SysMenuDaoImpl) BuildMenuTree(menuList []*domain.SysMenu) ([]*domain.SysMenu, error) {
	var returnList []*domain.SysMenu
	var tempList []any
	if len(menuList) <= 0 {
		return nil, errors.New("菜单数据不能为空")
	}

	for _, menu := range menuList {
		tempList = append(tempList, menu.Id)
	}

	for _, menu := range menuList {
		if !utils.Contains(tempList, menu.ParentId) {
			recursionFn(menuList, menu)
			returnList = append(returnList, menu)
		}
	}

	if len(returnList) == 0 {
		returnList = menuList
	}

	return returnList, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysMenuDaoImpl) SelectMenuList(ctx context.Context, menu *domain.SysMenu) ([]*domain.SysMenu, error) {
	var menuList []*domain.SysMenu
	querySql :=
		`SELECT
			m.id,
			m.menu_name,
			m.parent_id,
			m.order_num,
			m.router_path,
			m.component,
			m.is_frame,
			m.is_cache,
			m.menu_type,
			m.visible,
			m.menu_status,
			coalesce(m.perms, '') AS perms,
			m.icon,
			m.create_time,
			m.update_by,
			m.update_time
		FROM
			sys_menu m`
	db := impl.Gorm.WithContext(ctx)
	// 动态条件
	if menu.MenuName != "" {
		db.Where("m.menu_name ILIKE ?", "%"+menu.MenuName+"%")
	}

	if menu.Visible != "" {
		db.Where("m.visible = ?", menu.Visible)
	}

	if menu.MenuStatus != "" {
		db.Where("m.menu_status = ?", menu.MenuStatus)
	}

	// 时间范围处理
	if menu.BeginTime != "" && menu.EndTime != "" {
		if begin, end, err := parseTimeRange(menu.BeginTime, menu.EndTime); err == nil {
			db.Where("create_time BETWEEN ? AND ?", begin, end)
		}
	}

	if err := db.Raw(querySql).Find(&menuList).Order("m.parent_id, m.order_num").Error; err != nil {
		return nil, err
	}
	return menuList, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysMenuDaoImpl) SelectMenuListMenuAndUserId(ctx context.Context, menu *domain.SysMenu, userId int64) ([]*domain.SysMenu, error) {
	var menuList []*domain.SysMenu

	querySql :=
		`SELECT DISTINCT
			m.id,
			m.parent_id,
			m.menu_name,
			m.router_path,
			m.component,
			m.visible,
			m.menu_status,
			coalesce(m.perms, '') AS perms,
			m.is_frame,
			m.is_cache,
			m.menu_type,
			m.icon,
			m.order_num,
			m.create_time
		FROM
			sys_menu m
			LEFT JOIN sys_role_menu rm ON m.id = rm.menu_id
			LEFT JOIN sys_user_role ur ON rm.role_id = ur.role_id
			LEFT JOIN sys_role ro ON ur.role_id = ro.id`

	db := impl.Gorm.WithContext(ctx)

	// 动态条件
	if userId != 0 {
		db.Where("ur.user_id = ?", userId)
	}

	if menu.MenuName != "" {
		db.Where("m.menu_name ILIKE ?", "%"+menu.MenuName+"%")
	}

	if menu.Visible != "" {
		db.Where("m.visible = ?", menu.Visible)
	}

	if menu.MenuStatus != "" {
		db.Where("m.menu_status = ?", menu.MenuStatus)
	}

	// 时间范围处理
	if menu.BeginTime != "" && menu.EndTime != "" {
		if begin, end, err := parseTimeRange(menu.BeginTime, menu.EndTime); err == nil {
			db.Where("create_time BETWEEN ? AND ?", begin, end)
		}
	}
	if err := db.Raw(querySql).Find(&menuList).Order("m.parent_id, m.order_num").Error; err != nil {
		return nil, err
	}
	return menuList, nil
}

func (impl *SysMenuDaoImpl) SelectListByMenuAndUserId(ctx context.Context, menu *domain.SysMenu, userId int64) ([]*domain.SysMenu, error) {
	var menuList []*domain.SysMenu
	if domain.IsAdmin(userId) {
		allMenuList, err := impl.SelectMenuList(ctx, menu)
		if err != nil {
			return nil, err
		}
		menuList = allMenuList
	} else {
		notAdminMenuList, err := impl.SelectMenuListMenuAndUserId(ctx, menu, userId)
		if err != nil {
			return nil, err
		}
		menuList = notAdminMenuList
	}
	return menuList, nil
}

func (impl *SysMenuDaoImpl) BuildMenuTreeSelect(menuList []*domain.SysMenu) []*domain.TreeSelect {
	menuTreeList, err := impl.BuildMenuTree(menuList)
	if err != nil {
		return nil
	}
	var treeSelectList []*domain.TreeSelect
	for _, menu := range menuTreeList {
		treeSelect := domain.NewTreeSelectByMenu(menu)
		treeSelectList = append(treeSelectList, treeSelect)
	}
	return treeSelectList
}

func (impl *SysMenuDaoImpl) BuildRoleMenuTreeSelect(ctx context.Context, userId, roleId int64) (*domain.RoleMenuTreeSelect, error) {
	listMenuAndUserId, _ := impl.SelectListByMenuAndUserId(ctx, &domain.SysMenu{}, userId)
	checkedMenuIds, _ := impl.SelectMenuListByRoleId(ctx, roleId)

	return &domain.RoleMenuTreeSelect{
		CheckedKeys: checkedMenuIds,
		Menus:       impl.BuildMenuTreeSelect(listMenuAndUserId),
	}, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysMenuDaoImpl) SelectMenuListByRoleId(ctx context.Context, roleId int64) ([]int64, error) {
	roleInfo, err := impl.RoleDao.SelectById(ctx, roleId)
	if err != nil {
		return nil, err
	}

	var menuIds []int64

	if roleInfo == nil {
		return menuIds, nil
	}

	querySql :=
		`SELECT
			m.id
		FROM
			sys_menu m
			LEFT JOIN sys_role_menu rm ON m.id = rm.menu_id
		WHERE
			rm.role_id = ?`

	db := impl.Gorm.WithContext(ctx)

	if roleInfo.MenuCheckStrictly == "1" {
		db.Where(`AND m.id NOT IN (
						SELECT
							m.parent_id
						FROM
							sys_menu m
							INNER JOIN sys_role_menu rm ON m.id = rm.menu_id
							AND rm.role_id = ?
						)`, roleInfo.Id)
	}
	if err = db.Raw(querySql, roleInfo.Id).Find(&menuIds).Error; err != nil {
		return nil, err
	}

	return menuIds, nil
}

// getChildPerms 根据父节点的ID获取所有子节点
func getChildPerms(list []*domain.SysMenu, parentId int64) []*domain.SysMenu {
	menuList := make([]*domain.SysMenu, 0)
	for _, menu := range list {
		if menu.ParentId == parentId {
			recursionFn(list, menu)
			menuList = append(menuList, menu)
		}
	}
	return menuList
}

// recursionFn 递归列表
//
// 参数
//
//	-list: 菜单切片
//	-menu: 菜单新信息
func recursionFn(list []*domain.SysMenu, menu *domain.SysMenu) {
	// 得到子节点列表
	childList := getChildList(list, menu)
	menu.Children = childList
	for _, child := range childList {
		if hasChild(list, child) {
			recursionFn(list, child)
		}
	}
}

// getChildList 得到子节点列表
func getChildList(menuList []*domain.SysMenu, t *domain.SysMenu) []*domain.SysMenu {
	var tList []*domain.SysMenu
	for _, menu := range menuList {
		if menu.ParentId == t.Id {
			tList = append(tList, menu)
		}
	}
	return tList
}

/**
 * 判断是否有子节点
 */
func hasChild(list []*domain.SysMenu, t *domain.SysMenu) bool {
	return len(getChildList(list, t)) > 0
}

// getRouterName 获取路由名称
func getRouterName(menu *domain.SysMenu) string {
	routerName := capitalize(menu.MenuName)
	// 非外链并且是一级目录（类型为目录）
	if isMenuFrame(menu) {
		return ""
	}
	return routerName
}

// getRouterPath 获取路由地址
func getRouterPath(menu *domain.SysMenu) string {
	routerPath := menu.RouterPath
	// 非外链并且是一级目录（类型为目录）
	if menu.ParentId == 0 && menu.MenuType == "M" && menu.IsFrame == 1 {
		routerPath = "/" + menu.RouterPath
	} else if isMenuFrame(menu) {
		routerPath = "/"
	}
	return routerPath
}

// getComponent 获取组件名称
func getComponent(menu *domain.SysMenu) string {
	component := "Layout"
	menuComponent := menu.Component
	if menuComponent != "" && !isMenuFrame(menu) {
		component = menuComponent
	} else if menuComponent == "" && isParentView(menu) {
		component = "ParentView"
	}
	return component
}

// 首字母大写处理（支持空字符串）
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	// 转换为 rune 切片以支持 Unicode
	runes := []rune(s)
	first := strings.ToUpper(string(runes[0]))
	if len(runes) > 1 {
		return first + string(runes[1:])
	}
	return first
}

// isMenuFrame 是否为菜单内部跳转
func isMenuFrame(menu *domain.SysMenu) bool {
	return menu.ParentId == 0 && menu.MenuType == "C" && menu.IsFrame == 1
}

// isParentView 是否为parent_view组件
func isParentView(menu *domain.SysMenu) bool {
	return menu.ParentId != 0 && menu.MenuType == "M"
}
