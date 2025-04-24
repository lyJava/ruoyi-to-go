package service

import (
	"context"
	"ry-go/business/domain"
)

type CronSchedulerService interface {
	// AddTask 添加任务后的处理，根据规则处理任务
	AddTask(ctx context.Context, job *domain.SysJob) error
	// TaskStart 启动任务
	TaskStart(ctx context.Context, jobId string) error
	// TaskStop 停止任务
	TaskStop(ctx context.Context, jobId string) error
	// TaskPaused 暂停任务 （更新数据库状态）
	TaskPaused(ctx context.Context, jobId string) error
	// TaskResume 恢复任务 （更新数据库状态并重新加载）
	TaskResume(ctx context.Context, jobId string) error
}
