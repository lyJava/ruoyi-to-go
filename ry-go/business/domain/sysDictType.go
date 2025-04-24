package domain

import (
	"gorm.io/gorm"
)

// SysDictType 字典类型结构
//
//goland:noinspection GoUnusedExportedType
type SysDictType struct {
	DictId     int64         `gorm:"column:dict_id;primaryKey;autoIncrement" json:"dictId,string"` // 主键
	DictName   string        `gorm:"column:dict_name" json:"dictName,omitempty"`                   // 字典名称
	DictType   string        `gorm:"column:dict_type" json:"dictType,omitempty"`                   // 字典类型
	Status     string        `gorm:"column:status" json:"status,omitempty"`                        // 状态（0正常 1停用）
	CreateBy   string        `gorm:"column:create_by" json:"createBy,omitempty"`                   // 创建者
	CreateTime LocalDateTime `gorm:"column:create_time" json:"createTime,omitempty"`               // 创建时间
	UpdateBy   string        `gorm:"column:update_by" json:"updateBy,omitempty"`                   // 更新者
	UpdateTime LocalDateTime `gorm:"column:update_time" json:"updateTime,omitempty"`               // 更新时间
	Remark     string        `gorm:"column:remark" json:"remark,omitempty"`                        // 备注
}

// TableName gorm指定表名
func (SysDictType) TableName() string {
	return "sys_dict_type"
}

func (u *SysDictType) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysDictType) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}

// SysDictTypeExcel 字典类型结构导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysDictTypeExcel struct {
	DictId     string `gorm:"column:dict_id" json:"dictId,string"`            // 主键
	DictName   string `gorm:"column:dict_name" json:"dictName,omitempty"`     // 字典名称
	DictType   string `gorm:"column:dict_type" json:"dictType,omitempty"`     // 字典类型
	Status     string `gorm:"column:status" json:"status,omitempty"`          // 状态（0正常 1停用）
	CreateBy   string `gorm:"column:create_by" json:"createBy,omitempty"`     // 创建者
	CreateTime string `gorm:"column:create_time" json:"createTime,omitempty"` // 创建时间
	UpdateBy   string `gorm:"column:update_by" json:"updateBy,omitempty"`     // 更新者
	UpdateTime string `gorm:"column:update_time" json:"updateTime,omitempty"` // 更新时间
	Remark     string `gorm:"column:remark" json:"remark,omitempty"`          // 备注
}
