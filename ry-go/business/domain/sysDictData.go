package domain

import "gorm.io/gorm"

// SysDictData 数据字典结构体
type SysDictData struct {
	Id         int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id,string"` // 主键ID
	DictSort   int    `gorm:"column:dict_sort" json:"dictSort"`                    // 字典排序
	DictLabel  string `gorm:"column:dict_label" json:"dictLabel"`                  // 字典标签
	DictValue  string `gorm:"column:dict_value" json:"dictValue"`                  // 字典键值
	DictType   string `gorm:"column:dict_type" json:"dictType"`                    // 字典类型
	CssClass   string `gorm:"column:css_class" json:"cssClass"`                    // 样式属性（其他样式扩展）
	ListClass  string `gorm:"column:list_class" json:"listClass"`                  // 表格回显样式
	IsDefault  string `gorm:"column:is_default" json:"isDefault"`                  // 是否默认（Y是 N否）
	DictStatus string `gorm:"column:dict_status" json:"dictStatus"`                // 状态（0正常 1停用）
	BaseDomain
}

// TableName gorm指定表名
func (SysDictData) TableName() string {
	return "sys_dict_data"
}

func (u *SysDictData) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysDictData) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}
