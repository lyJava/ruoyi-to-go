package domain

import "gorm.io/gorm"

// SysPost 岗位信息结构
//
//goland:noinspection GoUnusedExportedType
type SysPost struct {
	Id         int64         `gorm:"column:id;primaryKey;autoIncrement" json:"id,string"` // 岗位ID
	PostCode   string        `gorm:"column:post_code" json:"postCode"`                    // 岗位编码
	PostName   string        `gorm:"column:post_name" json:"postName"`                    // 岗位名称
	PostSort   int64         `gorm:"column:post_sort" json:"postSort"`                    // 显示顺序
	PostStatus string        `gorm:"column:post_status" json:"postStatus"`                // 状态（0正常 1停用）
	CreateBy   string        `gorm:"column:create_by" json:"createBy,omitempty"`          // 创建者
	CreateTime LocalDateTime `gorm:"column:create_time" json:"createTime,omitempty"`      // 创建时间
	UpdateBy   string        `gorm:"column:update_by" json:"updateBy,omitempty"`          // 更新者
	UpdateTime LocalDateTime `gorm:"column:update_time" json:"updateTime,omitempty"`      // 更新时间
	Remarks    string        `gorm:"column:remarks" json:"remarks,omitempty"`             // 备注
}

// TableName gorm指定表名
func (SysPost) TableName() string {
	return "sys_post"
}

func (u *SysPost) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysPost) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}

// SysPostExcel 岗位信息导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysPostExcel struct {
	Id         string `gorm:"column:id" json:"id,string"`                     // 岗位ID
	PostCode   string `gorm:"column:post_code" json:"postCode"`               // 岗位编码
	PostName   string `gorm:"column:post_name" json:"postName"`               // 岗位名称
	PostSort   string `gorm:"column:post_sort" json:"postSort"`               // 显示顺序
	PostStatus string `gorm:"column:post_status" json:"postStatus"`           // 状态（0正常 1停用）
	CreateBy   string `gorm:"column:create_by" json:"createBy,omitempty"`     // 创建者
	CreateTime string `gorm:"column:create_time" json:"createTime,omitempty"` // 创建时间
	UpdateBy   string `gorm:"column:update_by" json:"updateBy,omitempty"`     // 更新者
	UpdateTime string `gorm:"column:update_time" json:"updateTime,omitempty"` // 更新时间
	Remarks    string `gorm:"column:remarks" json:"remarks,omitempty"`        // 备注
}
