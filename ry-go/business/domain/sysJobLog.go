package domain

import (
	"gorm.io/gorm"
)

// SysJobLog 系统任务日志结构
//
//goland:noinspection GoUnusedExportedType
type SysJobLog struct {
	JobLogId      int64         `gorm:"column:job_log_id;primaryKey;autoIncrement" json:"jobLogId,string"` // 任务日志ID
	JobName       string        `gorm:"column:job_name" json:"jobName"`                                    // 任务名称
	JobGroup      string        `gorm:"column:job_group" json:"jobGroup"`                                  // 任务组名
	InvokeTarget  string        `gorm:"column:invoke_target" json:"invokeTarget"`                          // 调用目标字符串
	JobMessage    string        `gorm:"column:job_message" json:"jobMessage"`                              // 日志信息
	Status        string        `gorm:"column:status" json:"status"`                                       // 执行状态（0正常 1失败）
	ExceptionInfo string        `gorm:"column:exception_info" json:"exceptionInfo"`                        // 异常信息
	CreateTime    LocalDateTime `gorm:"column:create_time" json:"createTime"`                              // 创建时间
}

// TableName gorm指定表名
func (SysJobLog) TableName() string {
	return "sys_job_log"
}

func (u *SysJobLog) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

// SysJobLogExcel 系统任务日志导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysJobLogExcel struct {
	JobLogId      string `gorm:"column:job_log_id;primaryKey;autoIncrement" json:"jobLogId,string"` // 任务日志ID
	JobName       string `gorm:"column:job_name" json:"jobName"`                                    // 任务名称
	JobGroup      string `gorm:"column:job_group" json:"jobGroup"`                                  // 任务组名
	InvokeTarget  string `gorm:"column:invoke_target" json:"invokeTarget"`                          // 调用目标字符串
	JobMessage    string `gorm:"column:job_message" json:"jobMessage"`                              // 日志信息
	Status        string `gorm:"column:status" json:"status"`                                       // 执行状态（0正常 1失败）
	ExceptionInfo string `gorm:"column:exception_info" json:"exceptionInfo"`                        // 异常信息
	CreateTime    string `gorm:"column:create_time" json:"createTime"`                              // 创建时间
}
