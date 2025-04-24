package domain

import (
	"gorm.io/gorm"
)

// SysJob 系统任务结构
//
//goland:noinspection GoUnusedExportedType
type SysJob struct {
	JobId          int64         `gorm:"column:job_id;primaryKey;autoIncrement" json:"jobId,string"` // 任务ID
	JobName        string        `gorm:"column:job_name;idx:job_name" json:"jobName"`                // 任务名称
	JobGroup       string        `gorm:"column:job_group;idx:job_group" json:"jobGroup"`             // 任务组名
	InvokeTarget   string        `gorm:"column:invoke_target" json:"invokeTarget"`                   // 调用目标字符串
	CronExpression string        `gorm:"column:cron_expression" json:"cronExpression"`               // cron执行表达式
	MisfirePolicy  string        `gorm:"column:misfire_policy" json:"misfirePolicy"`                 // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     string        `gorm:"column:concurrent" json:"concurrent"`                        // 是否并发执行（0允许 1禁止）
	Status         string        `gorm:"column:status" json:"status"`                                // 状态（0正常 1暂停）
	CreateBy       string        `gorm:"column:create_by" json:"createBy"`                           // 创建者
	CreateTime     LocalDateTime `gorm:"column:create_time" json:"createTime"`                       // 创建时间
	UpdateBy       string        `gorm:"column:update_by" json:"updateBy"`                           // 更新者
	UpdateTime     LocalDateTime `gorm:"column:update_time" json:"updateTime"`                       // 更新时间
	Remark         string        `gorm:"column:remark" json:"remark"`                                // 备注信息
}

// TableName gorm指定表名
func (SysJob) TableName() string {
	return "sys_job"
}

func (u *SysJob) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysJob) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}

// SysJobExcel 系统任务导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysJobExcel struct {
	JobId          string `gorm:"column:job_id;primaryKey;autoIncrement" json:"jobId,string"` // 任务ID
	JobName        string `gorm:"column:job_name;primaryKey" json:"jobName"`                  // 任务名称
	JobGroup       string `gorm:"column:job_group;primaryKey" json:"jobGroup"`                // 任务组名
	InvokeTarget   string `gorm:"column:invoke_target" json:"invokeTarget"`                   // 调用目标字符串
	CronExpression string `gorm:"column:cron_expression" json:"cronExpression"`               // cron执行表达式
	MisfirePolicy  string `gorm:"column:misfire_policy" json:"misfirePolicy"`                 // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     string `gorm:"column:concurrent" json:"concurrent"`                        // 是否并发执行（0允许 1禁止）
	Status         string `gorm:"column:status" json:"status"`                                // 状态（0正常 1暂停）
	CreateBy       string `gorm:"column:create_by" json:"createBy"`                           // 创建者
	CreateTime     string `gorm:"column:create_time" json:"createTime"`                       // 创建时间
	UpdateBy       string `gorm:"column:update_by" json:"updateBy"`                           // 更新者
	UpdateTime     string `gorm:"column:update_time" json:"updateTime"`                       // 更新时间
	Remark         string `gorm:"column:remark" json:"remark"`                                // 备注信息
}
