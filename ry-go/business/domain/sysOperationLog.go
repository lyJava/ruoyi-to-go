package domain

import (
	"gorm.io/gorm"
)

// SysOperationLog 操作日志结构
//
//goland:noinspection GoUnusedExportedType
type SysOperationLog struct {
	Id              int64         `gorm:"column:id;primaryKey;autoIncrement" json:"id,string"`                               // 日志主键
	ModuleName      string        `gorm:"column:module_name;primaryKey;autoIncrement" json:"moduleName,omitempty"`           // 模块标题
	BusinessType    int           `gorm:"column:business_type;primaryKey;autoIncrement" json:"businessType,omitempty"`       // 业务类型（0其它 1新增 2修改 3删除）
	MethodName      string        `gorm:"column:method_name;primaryKey;autoIncrement" json:"methodName,omitempty"`           // 方法名称
	RequestMethod   string        `gorm:"column:request_method;primaryKey;autoIncrement" json:"requestMethod,omitempty"`     // 请求方式
	OperationType   int           `gorm:"column:operation_type;primaryKey;autoIncrement" json:"operationType,omitempty"`     // 操作类别（0其它 1后台用户 2手机端用户）
	OperationUser   string        `gorm:"column:operation_user;primaryKey;autoIncrement" json:"operationUser,omitempty"`     // 操作人员
	DeptName        string        `gorm:"column:dept_name;primaryKey;autoIncrement" json:"deptName,omitempty"`               // 部门名称
	RequestUrl      string        `gorm:"column:request_url;primaryKey;autoIncrement" json:"requestUrl,omitempty"`           // 请求URL
	RequestIp       string        `gorm:"column:request_ip;primaryKey;autoIncrement" json:"requestIp,omitempty"`             // 主机地址
	RequestLocation string        `gorm:"column:request_location;primaryKey;autoIncrement" json:"requestLocation,omitempty"` // 操作地点
	RequestParam    string        `gorm:"column:request_param;primaryKey;autoIncrement" json:"requestParam,omitempty"`       // 请求参数
	ResponseResult  string        `gorm:"column:response_result;primaryKey;autoIncrement" json:"responseResult,omitempty"`   // 返回参数
	OperationStatus int           `gorm:"column:operation_status;primaryKey;autoIncrement" json:"operationStatus,omitempty"` // 操作状态（0正常 1异常）
	ErrorMsg        string        `gorm:"column:error_msg;primaryKey;autoIncrement" json:"errorMsg,omitempty"`               // 错误消息
	OperationTime   LocalDateTime `gorm:"column:operation_time;primaryKey;autoIncrement" json:"operationTime,omitempty"`     // 操作时间
}

// TableName gorm指定表名
func (SysOperationLog) TableName() string {
	return "sys_operation_log"
}

func (u *SysOperationLog) BeforeCreate(tx *gorm.DB) (err error) {
	u.OperationTime = LocalDateTimeNow("")
	return
}

// SysOperationLogExcel 操作日志导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysOperationLogExcel struct {
	Id              string `gorm:"column:id;primaryKey;autoIncrement" json:"id,string"`                               // 日志主键
	ModuleName      string `gorm:"column:module_name;primaryKey;autoIncrement" json:"moduleName,omitempty"`           // 模块标题
	BusinessType    string `gorm:"column:business_type;primaryKey;autoIncrement" json:"businessType,omitempty"`       // 业务类型（0其它 1新增 2修改 3删除）
	MethodName      string `gorm:"column:method_name;primaryKey;autoIncrement" json:"methodName,omitempty"`           // 方法名称
	RequestMethod   string `gorm:"column:request_method;primaryKey;autoIncrement" json:"requestMethod,omitempty"`     // 请求方式
	OperationType   string `gorm:"column:operation_type;primaryKey;autoIncrement" json:"operationType,omitempty"`     // 操作类别（0其它 1后台用户 2手机端用户）
	OperationUser   string `gorm:"column:operation_user;primaryKey;autoIncrement" json:"operationUser,omitempty"`     // 操作人员
	DeptName        string `gorm:"column:dept_name;primaryKey;autoIncrement" json:"deptName,omitempty"`               // 部门名称
	RequestUrl      string `gorm:"column:request_url;primaryKey;autoIncrement" json:"requestUrl,omitempty"`           // 请求URL
	RequestIp       string `gorm:"column:request_ip;primaryKey;autoIncrement" json:"requestIp,omitempty"`             // 主机地址
	RequestLocation string `gorm:"column:request_location;primaryKey;autoIncrement" json:"requestLocation,omitempty"` // 操作地点
	RequestParam    string `gorm:"column:request_param;primaryKey;autoIncrement" json:"requestParam,omitempty"`       // 请求参数
	ResponseResult  string `gorm:"column:response_result;primaryKey;autoIncrement" json:"responseResult,omitempty"`   // 返回参数
	OperationStatus string `gorm:"column:operation_status;primaryKey;autoIncrement" json:"operationStatus,omitempty"` // 操作状态（0正常 1异常）
	ErrorMsg        string `gorm:"column:error_msg;primaryKey;autoIncrement" json:"errorMsg,omitempty"`               // 错误消息
	OperationTime   string `gorm:"column:operation_time;primaryKey;autoIncrement" json:"operationTime,omitempty"`     // 操作时间
}
