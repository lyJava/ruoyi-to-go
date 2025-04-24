package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"ry-go/business/domain"
)

type SysRoleMenuDaoImpl struct {
	Gorm *gorm.DB
}

func NewSysRoleMenuDaoImpl(db *gorm.DB) *SysRoleMenuDaoImpl {
	return &SysRoleMenuDaoImpl{
		Gorm: db,
	}
}

func (impl *SysRoleMenuDaoImpl) DeleteByRoleId(ctx context.Context, roleId int64) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Where("role_id = ?", roleId).
		Model(&domain.SysRoleMenu{}).Delete(nil)
	if result.Error != nil {
		return 0, fmt.Errorf("删除角色菜单关联失败(角色ID=%d): %w", roleId, result.Error)
	}

	return result.RowsAffected, nil
}

func (impl *SysRoleMenuDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysRoleMenu) (int64, error) {
	if len(list) == 0 {
		return 0, errors.New("角色菜单数据不能为空")
	}
	for _, item := range list {
		if item == nil {
			return 0, errors.New("角色菜单数据不能存在空值")
		}
	}

	result := impl.Gorm.WithContext(ctx).Model(&domain.SysRoleMenu{}).Omit(clause.Associations).Create(list)
	if result.Error != nil {
		return 0, fmt.Errorf("角色菜单关联批量插入失败: %w", result.Error)
	}
	return result.RowsAffected, nil
}
