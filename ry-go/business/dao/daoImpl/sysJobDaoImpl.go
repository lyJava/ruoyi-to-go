package daoImpl

import (
	"context"
	"errors"
	"ry-go/business/domain"
	"ry-go/common/request"
	"ry-go/utils"
	"time"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// SysJobDaoImpl 系统任务Dao实现结构体
type SysJobDaoImpl struct {
	Gorm *gorm.DB
}

// NewSysJobDaoImpl 系统任务Dao实现创建
func NewSysJobDaoImpl(db *gorm.DB) *SysJobDaoImpl {
	return &SysJobDaoImpl{
		Gorm: db,
	}
}

func (impl *SysJobDaoImpl) Insert(ctx context.Context, sysJob *domain.SysJob) (*domain.SysJob, error) {
	if err := impl.Gorm.WithContext(ctx).Create(&sysJob).Error; err != nil {
		return nil, errors.New("系统任务新增失败")
	}
	// 如果这里边有使用钩子函数，那么这里需要重新查询而不是直接返回sysJob
	newSysJob, err := impl.SelectById(ctx, sysJob.JobId)
	if err != nil {
		return nil, errors.New("系统任务新增失败")
	}
	return newSysJob, nil
}

func (impl *SysJobDaoImpl) Update(ctx context.Context, sysJob *domain.SysJob) (int64, error) {
	result := impl.Gorm.WithContext(ctx).
		Omit("job_id").
		Save(&sysJob)
	if result.Error != nil {
		return 0, errors.New("系统任务修改失败")
	}
	return result.RowsAffected, nil
}

func (impl *SysJobDaoImpl) BatchInsert(ctx context.Context, list []*domain.SysJob) ([]int64, error) {
	if len(list) <= 0 {
		return nil, errors.New("系统任务批量新增参数验证失败")
	}

	result := impl.Gorm.WithContext(ctx).CreateInBatches(list, 50)
	if result.Error != nil {
		return nil, errors.New("系统任务批量新增失败")
	}

	var ids []int64
	for _, item := range list {
		ids = append(ids, item.JobId)
	}

	return ids, nil
}

func (impl *SysJobDaoImpl) SelectById(ctx context.Context, id int64) (*domain.SysJob, error) {
	var sysJob *domain.SysJob
	if err := impl.Gorm.WithContext(ctx).Model(&domain.SysJob{}).First(&sysJob, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("系统任务未找到")
		}
		return nil, errors.New("系统任务查询失败")
	}
	return sysJob, nil
}

func (impl *SysJobDaoImpl) SelectPage(ctx context.Context, param *request.SysNoticePageParam) ([]*domain.SysJob, int64, int64, error) {
	page, size := utils.SetPageDefault(param.Page, param.Size)
	param.Page = page
	param.Size = size

	// 计算分页参数
	offset := (param.Page - 1) * param.Size
	// 初始化 GORM 查询
	db := impl.Gorm.WithContext(ctx).Model(&domain.SysJob{})
	// 查询总记录数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, 0, errors.New("获取系统任务总记录数失败")
	}

	// 查询用户列表
	var list []*domain.SysJob
	if err := db.
		Offset(int(offset)).Limit(int(param.Size)).Order("job_id").
		Find(&list).Error; err != nil {
		return nil, 0, 0, errors.New("获取系统任务列表失败")
	}

	return list, total, utils.CalcTotalPage(total, param.Size), nil
}

func (impl *SysJobDaoImpl) BatchDelete(ctx context.Context, ids []any) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysJob{}).Where("job_id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		return 0, errors.New("批量删除失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysJobDaoImpl) UpdateJobStatus(ctx context.Context, id int64, status string) (int64, error) {
	result := impl.Gorm.WithContext(ctx).Model(&domain.SysJob{}).
		Where("job_id = ?", id).
		Updates(map[string]any{
			"status":      status,
			"update_time": time.Now(),
		})
	if result.Error != nil {
		return 0, errors.New("系统任务修改状态失败")
	}

	return result.RowsAffected, nil
}

func (impl *SysJobDaoImpl) UpdateJobStatusBatch(ctx context.Context, list []*request.JobChangeParam) (int64, error) {
	count := int64(0)

	for _, job := range list {
		result, err := impl.UpdateJobStatus(ctx, cast.ToInt64(job.JobId), job.Status)
		if err != nil {
			return 0, err
		}
		count += result
	}

	return count, nil
}
