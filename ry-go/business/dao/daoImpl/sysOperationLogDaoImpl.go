package daoImpl

import (
	"context"
	"errors"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"

	"gorm.io/gorm"
)

// SysOperationLogDaoImpl 操作日志Dao实现结构体
type SysOperationLogDaoImpl struct {
	Gorm *gorm.DB
}

// NewSysOperationLogDaoImpl 操作日志Dao实现创建
func NewSysOperationLogDaoImpl(db *gorm.DB) *SysOperationLogDaoImpl {
	return &SysOperationLogDaoImpl{
		Gorm: db,
	}
}

func (impl *SysOperationLogDaoImpl) Insert(ctx context.Context, sysOperationLog *domain.SysOperationLog) (*domain.SysOperationLog, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysOperationLog).Error; err != nil {
		return nil, errors.New("操作日志新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysOperationLog
	sysOperationLog, err := impl.SelectById(ctx, sysOperationLog.Id)
	if err != nil {
		return nil, errors.New("操作日志新增失败")
	}
	return sysOperationLog, nil
}

func (impl *SysOperationLogDaoImpl) Update(ctx context.Context, sysOperationLog *domain.SysOperationLog) (int64, error) {
	result := impl.Gorm.WithContext(ctx).
		Where("id = ?", sysOperationLog.Id).
		Omit("id").
		Save(&sysOperationLog)
	if result.Error != nil {
		return 0, errors.New("操作日志修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysOperationLogDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysOperationLog) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("操作日志批量新增参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		return nil, errors.New("操作日志批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.Id)
	}

	return ids, nil
}

func (impl *SysOperationLogDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysOperationLog, error) {
	var sysOperationLog *domain.SysOperationLog
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysOperationLog{}).First(&sysOperationLog, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("操作日志未找到")
		}
		return nil, errors.New("操作日志查询失败")
	}
	return sysOperationLog, nil
}

func (impl *SysOperationLogDaoImpl) SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysOperationLog, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysOperationLog{})
	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, 0, errors.New("获取操作日志总记录数失败")
	}

	// 查询用户列表
	var list []*domain.SysOperationLog
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		return nil, 0, 0, errors.New("获取操作日志列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysOperationLogDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("操作日志批量删除参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).Model(&domain.SysOperationLog{}).Where("id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		return 0, errors.New("批量删除失败")
	}

	return result.RowsAffected, nil
}
