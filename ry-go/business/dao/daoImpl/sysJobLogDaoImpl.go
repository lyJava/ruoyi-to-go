package daoImpl

import (
	"context"
	"errors"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"

	"gorm.io/gorm"
)

// SysJobLogDaoImpl 系统任务日志Dao实现结构体
type SysJobLogDaoImpl struct {
	Gorm *gorm.DB
}

// NewSysJobLogDaoImpl 系统任务日志Dao实现创建
func NewSysJobLogDaoImpl(db *gorm.DB) *SysJobLogDaoImpl {
	return &SysJobLogDaoImpl{
		Gorm: db,
	}
}

func (impl *SysJobLogDaoImpl) Insert(ctx context.Context, sysJobLog *domain.SysJobLog) (*domain.SysJobLog, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysJobLog).Error; err != nil {
		return nil, errors.New("系统任务日志新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysJobLog
	sysJobLog, err := impl.SelectById(ctx, sysJobLog.JobLogId)
	if err != nil {
		return nil, errors.New("系统任务日志新增失败")
	}
	return sysJobLog, nil
}

func (impl *SysJobLogDaoImpl) Update(ctx context.Context, sysJobLog *domain.SysJobLog) (int64, error) {
	result := impl.Gorm.WithContext(ctx).
		Where("job_log_id = ?", sysJobLog.JobLogId).
		Omit("job_log_id").
		Save(&sysJobLog)
	if result.Error != nil {
		return 0, errors.New("系统任务日志修改失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysJobLogDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysJobLog) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("系统任务日志批量新增参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		return nil, errors.New("系统任务日志批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.JobLogId)
	}

	return ids, nil
}

func (impl *SysJobLogDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysJobLog, error) {
	var sysJobLog *domain.SysJobLog
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysJobLog{}).First(&sysJobLog, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("系统任务日志未找到")
		}
		return nil, errors.New("系统任务日志查询失败")
	}
	return sysJobLog, nil
}

func (impl *SysJobLogDaoImpl) SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysJobLog, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysJobLog{})
	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, 0, errors.New("获取系统任务日志总记录数失败")
	}

	// 查询用户列表
	var list []*domain.SysJobLog
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).
		Find(&list).Error; err != nil {
		return nil, 0, 0, errors.New("获取系统任务日志列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysJobLogDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	if len(ids) <= 0 {
		return 0, errors.New("系统任务日志批量删除参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).Model(&domain.SysJobLog{}).Where("job_log_id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		return 0, errors.New("批量删除失败")
	}

	return result.RowsAffected, nil
}
