package domain

import "gorm.io/gorm"

// SysConfig 系统配置结构
//
//goland:noinspection GoUnusedExportedType
type SysConfig struct {
	ConfigId    int64         `gorm:"column:config_id;primaryKey;autoIncrement" json:"id,string"` // 参数主键
	ConfigName  string        `gorm:"column:config_name" json:"configName"`                       // 参数名称
	ConfigKey   string        `gorm:"column:config_key" json:"configKey"`                         // 参数键名
	ConfigValue string        `gorm:"column:config_value" json:"configValue"`                     // 参数键值
	ConfigType  string        `gorm:"column:config_type" json:"configType"`                       // 系统内置（Y是 N否）
	CreateBy    string        `gorm:"column:create_by" json:"createBy"`                           // 创建者
	CreateTime  LocalDateTime `gorm:"column:create_time" json:"createTime"`                       // 创建时间
	UpdateBy    string        `gorm:"column:update_by" json:"updateBy"`                           // 更新者
	UpdateTime  LocalDateTime `gorm:"column:update_time" json:"updateTime"`                       // 更新时间
	Remark      string        `gorm:"column:remark" json:"remark"`                                // 备注
}

// TableName gorm指定表名
func (SysConfig) TableName() string {
	return "sys_config"
}

func (u *SysConfig) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysConfig) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}

// SysConfigExcel 系统配置导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysConfigExcel struct {
	ConfigId    string `gorm:"column:config_id" json:"id"`                       // 参数主键
	ConfigName  string `gorm:"column:config_name" json:"configName,omitempty"`   // 参数名称
	ConfigKey   string `gorm:"column:config_key" json:"configKey,omitempty"`     // 参数键名
	ConfigValue string `gorm:"column:config_value" json:"configValue,omitempty"` // 参数键值
	ConfigType  string `gorm:"column:config_type" json:"configType,omitempty"`   // 系统内置（Y是 N否）
	CreateBy    string `gorm:"column:create_by" json:"createBy,omitempty"`       // 创建者
	CreateTime  string `gorm:"column:create_time" json:"createTime,omitempty"`   // 创建时间
	UpdateBy    string `gorm:"column:update_by" json:"updateBy,omitempty"`       // 更新者
	UpdateTime  string `gorm:"column:update_time" json:"updateTime,omitempty"`   // 更新时间
	Remark      string `gorm:"column:remark" json:"remark,omitempty"`            // 备注
}
