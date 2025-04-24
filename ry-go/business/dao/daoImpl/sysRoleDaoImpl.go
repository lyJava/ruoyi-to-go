package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"ry-go/business/dao"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"
	"strings"
	"time"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// SysRoleDaoImpl 系统角色Dao实现结构体
type SysRoleDaoImpl struct {
	Gorm        *gorm.DB
	RoleMenuDao dao.SysRoleMenuDao
	RoleDeptDao dao.SysRoleDeptDao
}

// NewSysRoleDaoImpl 系统角色Dao实现创建
func NewSysRoleDaoImpl(db *gorm.DB, rm *SysRoleMenuDaoImpl, rd *SysRoleDeptDaoImpl) *SysRoleDaoImpl {
	return &SysRoleDaoImpl{
		Gorm:        db,
		RoleMenuDao: rm,
		RoleDeptDao: rd,
	}
}

func (impl *SysRoleDaoImpl) Insert(ctx context.Context, sysRole *domain.SysRole) (*domain.SysRole, error) {
	sysRole.Id = 0
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysRole{}).Create(sysRole).Error; err != nil {
		return nil, errors.New("系统角色新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysRole
	sysRoleInert, err := impl.SelectById(ctx, sysRole.Id)
	if err != nil {
		return nil, errors.New("系统角色新增失败")
	}

	if sysRoleInert != nil && len(sysRole.MenuIds) > 0 {
		sysRoleId := sysRoleInert.Id

		_, menuFelErr := impl.RoleMenuDao.DeleteByRoleId(ctx, sysRoleId)
		if menuFelErr != nil {
			return nil, menuFelErr
		}

		var roleMenuList []*domain.SysRoleMenu
		for _, menuId := range sysRole.MenuIds {
			roleMenuList = append(roleMenuList, &domain.SysRoleMenu{
				RoleId: sysRoleId,
				MenuId: cast.ToInt64(menuId),
			})
		}
		_, menuInsertErr := impl.RoleMenuDao.BatchInsert(ctx, roleMenuList)
		if menuInsertErr != nil {
			return nil, fmt.Errorf("角色:%d 菜单批量新增错误, %w", sysRoleId, menuInsertErr)
		}

	}

	return sysRoleInert, nil
}

func (impl *SysRoleDaoImpl) Update(ctx context.Context, sysRole *domain.SysRole) (int64, error) {
	sysRoleId := sysRole.Id
	result := impl.Gorm.WithContext(ctx).
		Where("id = ?", sysRoleId).
		Omit("id").
		Save(&sysRole)
	if result.Error != nil {
		return 0, errors.New("系统角色修改失败")
	}
	count := result.RowsAffected

	if count > 0 {
		_, err := impl.RoleMenuDao.DeleteByRoleId(ctx, sysRoleId)
		if err != nil {
			return 0, err
		}
		if len(sysRole.MenuIds) > 0 {
			var roleMenuList []*domain.SysRoleMenu
			for _, menuId := range sysRole.MenuIds {
				roleMenuList = append(roleMenuList, &domain.SysRoleMenu{
					RoleId: sysRoleId,
					MenuId: cast.ToInt64(menuId),
				})
			}
			_, menuInsertErr := impl.RoleMenuDao.BatchInsert(ctx, roleMenuList)
			if menuInsertErr != nil {
				return 0, fmt.Errorf("用户:%d 菜单批量新增错误, %w", sysRoleId, menuInsertErr)
			}
		}
	}

	return count, nil
}

func (impl *SysRoleDaoImpl) UpdateWithOutMenuAndDataScope(ctx context.Context, sysRole *domain.SysRole) (int64, error) {
	roleId := sysRole.Id
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysRole{}).
		Where("id = ?", roleId).Updates(&sysRole)
	if result.Error != nil {
		return 0, errors.New("系统角色修改失败")
	}
	_, err := impl.RoleDeptDao.DeleteByRoleId(ctx, roleId)
	if err != nil {
		return 0, fmt.Errorf("角色:%d 关联部门删除错误，%w", roleId, err)
	}

	if len(sysRole.DeptIds) > 0 {
		var roleDeptList []*domain.SysRoleDept
		for _, deptId := range sysRole.DeptIds {
			roleDeptList = append(roleDeptList, &domain.SysRoleDept{
				RoleId: roleId,
				DeptId: cast.ToInt64(deptId),
			})
		}
		_, err = impl.RoleDeptDao.BatchInsert(ctx, roleDeptList)
		if err != nil {
			return 0, fmt.Errorf("角色:%d 关联部门新增错误，%w", roleId, err)
		}
	}

	return result.RowsAffected, nil
}

func (impl *SysRoleDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysRole) ([]int64, error) {
	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		return nil, errors.New("系统角色批量新增失败")
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.Id)
	}

	return ids, nil
}

func (impl *SysRoleDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysRole, error) {
	var sysRole *domain.SysRole
	// 主键为id的时候First(&sysRole, id)默认使用主键ID查询
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysRole{}).First(&sysRole, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("系统角色未找到")
		}
		return nil, errors.New("角色查询失败")
	}
	return sysRole, nil
}

func (impl *SysRoleDaoImpl) SelectPage(ctx context.Context, param *request.RolePageParam) ([]*domain.SysRole, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysRole{})
	// 动态构建查询条件
	if param.RoleName != "" {
		db = db.Where("role_name ILIKE ?", "%"+param.RoleName+"%")
	}
	if param.RoleKey != "" {
		db = db.Where("role_key ILIKE ?", "%"+param.RoleKey+"%")
	}
	if param.RoleStatus != "" {
		db = db.Where("role_status = ?", param.RoleStatus)
	}
	if param.BeginTime != "" && param.EndTime != "" {
		// 时间格式校验（示例）
		if _, err := time.Parse(time.DateOnly, param.BeginTime); err != nil {
			return nil, 0, 0, errors.New("时间格式错误（应为 YYYY-MM-DD）")
		}
		if _, err := time.Parse(time.DateOnly, param.EndTime); err != nil {
			return nil, 0, 0, errors.New("时间格式错误（应为 YYYY-MM-DD）")
		}
		db = db.Where("create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}

	// 安全处理排序
	if param.Column != "" && param.Order != "" {
		// 校验列名是否合法（示例：只允许特定字段排序）
		validColumns := map[string]bool{"role_name": true, "create_time": true}
		if !validColumns[param.Column] {
			return nil, 0, 0, errors.New("非法的排序字段")
		}
		// 校验排序方向
		if param.Order != "ASC" && param.Order != "DESC" {
			param.Order = "ASC" // 默认升序
		}
		db = db.Order(utils.CamelToSnakeCase(param.Column) + " " + param.Order)
	}

	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, 0, errors.New("角色分页总记录失败")
	}

	// 查询角色列表
	var list []*domain.SysRole
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		return nil, 0, 0, errors.New("角色分页列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysRoleDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysRole{}).Where("id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		return 0, errors.New("批量删除失败")
	}
	return result.RowsAffected, nil
}

func (impl *SysRoleDaoImpl) SelectList(ctx context.Context, ids []int64) ([]*domain.SysRole, error) {
	var list []*domain.SysRole
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysRole{}).Find(&list, ids).Error; err != nil {
		return nil, errors.New("查询角色失败")
	}

	return list, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysRoleDaoImpl) SelectIdsByUserId(ctx context.Context, userId any) ([]int64, error) {
	var list []int64
	if err := impl.Gorm.WithContext(ctx).
		Raw(`
				SELECT
					r.id
				FROM
					sys_role r
					LEFT JOIN sys_user_role ur ON ur.role_id = r.id
					LEFT JOIN sys_user u ON u.id = ur.user_id
				WHERE u.id = ?`, userId).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询角色集合失败")
	}
	return list, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysRoleDaoImpl) SelectRoleKeyListByUserId(ctx context.Context, userId any) ([]string, error) {
	var list []string
	if err := impl.Gorm.WithContext(ctx).
		Raw(`
				SELECT
					r.role_key
				FROM
					sys_role r
					LEFT JOIN sys_user_role ur ON ur.role_id = r.id
					LEFT JOIN sys_user u ON u.id = ur.user_id
				WHERE u.id = ?`, userId).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询角色集合失败")
	}
	return list, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysRoleDaoImpl) SelectAll(ctx context.Context) ([]*domain.SysRole, error) {
	var list []*domain.SysRole
	if err := impl.Gorm.WithContext(ctx).
		Raw(`SELECT DISTINCT
					r.id,
					r.role_name,
					r.role_key,
					r.role_sort,
					r.data_scope,
					r.menu_check_strictly,
					r.dept_check_strictly,
					r.role_status,
					r.del_flag,
					r.create_time,
					r.remarks
				FROM
					sys_role r
					LEFT JOIN sys_user_role ur ON ur.role_id = r.id
					LEFT JOIN sys_user u ON u.id = ur.user_id
					LEFT JOIN sys_dept d ON u.dept_id = d.id
				WHERE
					r.del_flag = ?`, "0").
		Find(&list).Error; err != nil {
		return nil, errors.New("查询所有角色失败")
	}
	return list, nil
}

func (impl *SysRoleDaoImpl) SelectAllFilterAdmin(list []*domain.SysRole, isAdmin bool) ([]*domain.SysRole, error) {
	var noAdminList []*domain.SysRole
	if isAdmin {
		for _, role := range list {
			if !domain.IsAdminRole(role.Id) {
				noAdminList = append(noAdminList, role)
			}
		}
	} else {
		noAdminList = list
	}
	return noAdminList, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysRoleDaoImpl) SelectRoleNameListByUserId(ctx context.Context, userId any) ([]string, error) {
	var list []string
	if err := impl.Gorm.WithContext(ctx).
		Raw(`
				SELECT
					r.role_name
				FROM
					sys_role r
					LEFT JOIN sys_user_role ur ON ur.role_id = r.id
					LEFT JOIN sys_user u ON u.id = ur.user_id
				WHERE u.id = ?`, userId).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询角色名称失败")
	}
	return list, nil
}

func (impl *SysRoleDaoImpl) SelectRoleByRoleName(ctx context.Context, sysRole *domain.SysRole) (*domain.SysRole, error) {
	var role *domain.SysRole
	if err := impl.Gorm.WithContext(ctx).Model(domain.SysRole{}).Where("role_name = ?", sysRole.RoleName).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("查询角色名称失败")
	}
	return role, nil
}

func (impl *SysRoleDaoImpl) SelectCountByRoleKey(ctx context.Context, sysRole *domain.SysRole) (*domain.SysRole, error) {
	var role *domain.SysRole
	if err := impl.Gorm.WithContext(ctx).Model(domain.SysRole{}).Where("role_key = ?", sysRole.RoleKey).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("查询权限字符失败")
	}
	return role, nil
}

//goland:noinspection SqlResolve,SqlCaseVsIf,SqlNoDataSourceInspection
func (impl *SysRoleDaoImpl) SelectRolePermsByUserId(ctx context.Context, userId int64) ([]string, error) {
	var roleList []*domain.SysRole
	querySql :=
		`SELECT DISTINCT
			r.id,
			r.role_name,
			r.role_key,
			r.role_sort,
			r.data_scope,
			r.menu_check_strictly,
			r.dept_check_strictly,
			r.role_status,
			r.del_flag,
			r.create_time,
			r.remark
		FROM
			sys_role r
			LEFT JOIN sys_user_role ur ON ur.role_id = r.id
			LEFT JOIN sys_user u ON u.id = ur.user_id
			LEFT JOIN sys_dept d ON u.dept_id = d.id
		WHERE
			r.del_flag = '0'
			AND ur.user_id = ?`
	if err := impl.Gorm.WithContext(ctx).Raw(querySql, userId).Find(&roleList).Error; err != nil {
		return nil, errors.New("查询角色权限失败")
	}

	var permissions []string
	if len(roleList) > 0 {

		permSet := make(map[string]struct{})
		for _, role := range roleList {
			if role == nil {
				continue
			}

			trimmedRoleKey := strings.TrimSpace(role.RoleKey)
			if trimmedRoleKey == "" {
				continue
			}
			// 使用 FieldsFunc 分割并自动处理空格
			keys := strings.FieldsFunc(trimmedRoleKey, func(r rune) bool {
				return r == ','
			})
			for _, key := range keys {
				if key != "" {
					permSet[key] = struct{}{}
				}
			}
		}
		for k := range permSet {
			permissions = append(permissions, k)
		}
	}

	return permissions, nil
}
