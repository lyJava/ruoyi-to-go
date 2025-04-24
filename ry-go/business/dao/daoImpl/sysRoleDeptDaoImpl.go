package daoImpl

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"ry-go/business/domain"
)

type SysRoleDeptDaoImpl struct {
	Gorm *gorm.DB
}

func NewSysRoleDeptDaoImpl(db *gorm.DB) *SysRoleDeptDaoImpl {
	return &SysRoleDeptDaoImpl{
		Gorm: db,
	}
}

func (impl *SysRoleDeptDaoImpl) DeleteByRoleId(ctx context.Context, roleId int64) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Where("role_id = ?", roleId).
		Model(&domain.SysRoleDept{}).Delete(nil)
	if result.Error != nil {
		return 0, fmt.Errorf("删除角色与部门关联失败，(角色ID=%d): %w", roleId, result.Error)
	}

	return result.RowsAffected, nil
}

func (impl *SysRoleDeptDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysRoleDept) (int64, error) {
	if len(list) == 0 {
		return 0, errors.New("角色部门数据不能为空")
	}
	for _, item := range list {
		if item == nil {
			return 0, errors.New("角色部门数据不能存在空值")
		}
	}

	result := impl.Gorm.WithContext(ctx).Model(&domain.SysRoleDept{}).Create(list)
	if result.Error != nil {
		return 0, fmt.Errorf("角色部门关联批量插入失败: %w", result.Error)
	}
	return result.RowsAffected, nil
}
