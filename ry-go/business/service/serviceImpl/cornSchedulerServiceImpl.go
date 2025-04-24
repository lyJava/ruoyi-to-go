package serviceImpl

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"ry-go/business/dao"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
)

type CronSchedulerServiceImpl struct {
	Gorm    *gorm.DB
	Cron    *cron.Cron
	Entries sync.Map
	JobDao  dao.SysJobDao
}

func NewCornSchedulerServiceImpl(pg *gorm.DB, entryMap sync.Map, jd *daoImpl.SysJobDaoImpl) *CronSchedulerServiceImpl {
	impl := &CronSchedulerServiceImpl{
		Gorm: pg,
		Cron: cron.New(
			cron.WithSeconds(),
			cron.WithLocation(time.Local),
			cron.WithChain(
				cron.SkipIfStillRunning(cron.DefaultLogger),
				cron.Recover(cron.DefaultLogger),
			),
		),
		Entries: entryMap,
		JobDao:  jd,
	}

	return impl
}

func (impl *CronSchedulerServiceImpl) AddTask(ctx context.Context, job *domain.SysJob) error {
	jobId := cast.ToString(job.JobId)

	if entryID, ok := impl.Entries.Load(jobId); ok {
		impl.Cron.Remove(entryID.(cron.EntryID))
	}

	entryID, err := impl.Cron.AddFunc(job.CronExpression, func() {
		// 记录任务开始时间
		startTime := time.Now()
		zap.L().Sugar().Infof("执行任务ID: %s, 时间: %s", jobId, startTime.Format(time.DateTime))
	})
	if err != nil {
		zap.L().Sugar().Error("任务ID: %s 添加失败", jobId)
		return err
	}

	impl.Entries.Store(jobId, entryID)

	zap.L().Sugar().Infof("任务Id: %s 执行成功", jobId)

	// 处理调度策略
	switch job.MisfirePolicy {
	case "1", "2":
		impl.Cron.Start() // 立即开始调度
	case "3":
		impl.Cron.Stop() // 停止调度
	default:
		impl.Cron.Start()
	}

	zap.L().Sugar().Infof("Success AddFunc entryID===%+v", entryID)

	return nil
}

func (impl *CronSchedulerServiceImpl) TaskStart(ctx context.Context, jobId string) error {
	job, err := impl.JobDao.SelectById(ctx, cast.ToInt64(jobId))
	if err != nil || job == nil {
		return fmt.Errorf("任务不存在: %w", err)
	}

	// 如果任务已存在，先移除旧条目
	if entryID, ok := impl.Entries.Load(jobId); ok {
		impl.Cron.Remove(entryID.(cron.EntryID))
	}

	// 添加新任务
	entryID, err := impl.Cron.AddFunc(job.CronExpression, impl.wrapJobLogic(ctx, job))
	if err != nil {
		return fmt.Errorf("添加任务失败: %w", err)
	}

	impl.Entries.Store(jobId, entryID)
	// 处理调度策略
	switch job.MisfirePolicy {
	case "1", "2":
		impl.Cron.Start() // 立即开始调度
	case "3":
		impl.Cron.Stop() // 停止调度
	default:
		impl.Cron.Start()
	}
	return nil
}

func (impl *CronSchedulerServiceImpl) TaskStop(ctx context.Context, jobId string) error {

	if entryID, ok := impl.Entries.Load(jobId); ok {
		impl.Cron.Remove(entryID.(cron.EntryID))
		impl.Entries.Delete(jobId)
		zap.L().Sugar().Infof("任务ID: %s 成功停止", jobId)
		return nil
	}

	return fmt.Errorf("任务未在运行")
}

func (impl *CronSchedulerServiceImpl) TaskPaused(ctx context.Context, jobId string) error {
	// 复用停止逻辑
	return impl.TaskStop(ctx, jobId)
}

func (impl *CronSchedulerServiceImpl) TaskResume(ctx context.Context, jobId string) error {
	// 复用启动逻辑
	return impl.TaskStart(ctx, jobId)
}

// 封装任务逻辑
func (impl *CronSchedulerServiceImpl) wrapJobLogic(ctx context.Context, job *domain.SysJob) func() {
	return func() {
		//start := time.Now()

	}
}
