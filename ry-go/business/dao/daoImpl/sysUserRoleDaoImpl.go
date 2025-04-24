package daoImpl

import (
	"context"
	"errors"
	"fmt"
	domain2 "ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SysUserRoleDaoImpl 用户角色信息Dao实现结构体
type SysUserRoleDaoImpl struct {
	Gorm *gorm.DB
}

// NewSysUserRoleDaoImpl 用户角色信息Dao实现创建
func NewSysUserRoleDaoImpl(db *gorm.DB) *SysUserRoleDaoImpl {
	return &SysUserRoleDaoImpl{
		Gorm: db,
	}
}

func (impl *SysUserRoleDaoImpl) Insert(ctx context.Context, sysUserRole *domain2.SysUserRole) (*domain2.SysUserRole, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysUserRole).Error; err != nil {
		zap.L().Sugar().Error("用户角色信息新增异常===", zap.Error(err))
		return nil, errors.New("用户角色信息新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysUserRole
	sysUserRole, err := impl.SelectById(ctx, sysUserRole.RoleId)
	if err != nil {
		return nil, errors.New("用户角色信息新增失败")
	}
	return sysUserRole, nil
}

func (impl *SysUserRoleDaoImpl) Update(ctx context.Context, sysUserRole *domain2.SysUserRole) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Model(&domain2.SysUserRole{}).
		Where("id = ?", sysUserRole.RoleId).
		Omit("id").
		Updates(&sysUserRole)
	if result.Error != nil {
		zap.L().Sugar().Error("用户角色信息修改事务执行错误===", zap.Error(result.Error))
		return 0, errors.New("用户角色信息修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysUserRoleDaoImpl) BatchInsert(ctx context.Context, list []*domain2.SysUserRole) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("用户角色信息批量新增参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		zap.L().Sugar().Error("用户角色信息批量新增BatchInsert错误===", zap.Error(result.Error))
		return nil, errors.New("用户角色信息批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.RoleId)
	}

	zap.L().Sugar().Infof("用户角色信息批量新增BatchInsert成功,成功条数===%d,返回ids:===%v", result.RowsAffected, ids)
	return ids, nil
}

func (impl *SysUserRoleDaoImpl) SelectById(ctx context.Context, id int64) (*domain2.SysUserRole, error) {
	var sysUserRole *domain2.SysUserRole
	// 主键为id的时候First(&sysUserRole, id)默认使用主键ID查询
	if err := impl.Gorm.WithContext(ctx).Model(&domain2.SysUserRole{}).First(&sysUserRole, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户角色信息未找到")
		}
		zap.L().Error("查询SysUserRole失败", zap.Error(err))
		return nil, errors.New("用户角色信息查询失败")
	}
	return sysUserRole, nil
}

func (impl *SysUserRoleDaoImpl) SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain2.SysUserRole, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain2.SysUserRole{})
	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		zap.L().Sugar().Error("获取用户角色信息总记录数异常:", zap.Error(err))
		return nil, 0, 0, errors.New("获取用户角色信息总记录数失败")
	}

	// 查询用户列表
	var list []*domain2.SysUserRole
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		zap.L().Error("获取用户角色信息列表异常",
			zap.Error(err),
			zap.String("function", "SelectPage"),
			zap.Any("param", param),
		)
		return nil, 0, 0, errors.New("获取用户角色信息列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysUserRoleDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("用户角色信息批量删除参数验证失败")
	}
	zap.L().Sugar().Infof("用户角色信息批量删除参数:ids===%v", ids)
	result := impl.Gorm.WithContext(ctx).Model(&domain2.SysUserRole{}).Where("id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		zap.L().Sugar().Error("用户角色信息批量删除异常", zap.Error(result.Error))
		return 0, errors.New("批量删除失败")
	}
	zap.L().Sugar().Infof("用户角色信息批量删除成功执行条数===%d", result.RowsAffected)
	return result.RowsAffected, nil
}

func (impl *SysUserRoleDaoImpl) DeleteByUserId(ctx context.Context, userId int64) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Model(&domain2.SysUserRole{}).Where("user_id = ?", userId).Delete(nil)
	if result.Error != nil {
		return 0, fmt.Errorf("通过用户ID===%d删除绑定角色失败", userId)
	}
	return result.RowsAffected, nil
}

func (impl *SysUserRoleDaoImpl) DeleteByUserIdList(ctx context.Context, userIds []int64) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Model(&domain2.SysUserRole{}).Where("user_id IN ?", userIds).Delete(nil)
	if result.Error != nil {
		return 0, fmt.Errorf("通过用户ID===%v删除绑定角色失败,%w", userIds, result.Error)
	}
	return result.RowsAffected, nil
}
