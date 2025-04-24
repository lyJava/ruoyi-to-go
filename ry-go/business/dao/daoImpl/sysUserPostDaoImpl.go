package daoImpl

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"ry-go/business/domain"
)

type SysUserPostDaoImpl struct {
	Gorm *gorm.DB
}

func NewSysUserPostDaoImpl(db *gorm.DB) *SysUserPostDaoImpl {
	return &SysUserPostDaoImpl{
		Gorm: db,
	}
}

func (impl *SysUserPostDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysUserPost) (int64, error) {
	if len(list) == 0 {
		return 0, errors.New("用户岗位切片不能为空")
	}
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysUserPost{}).Create(list)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (impl *SysUserPostDaoImpl) DeleteByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	if err := impl.Gorm.WithContext(ctx).Where("user_id = ?", userId).Delete(&domain.SysUserPost{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
