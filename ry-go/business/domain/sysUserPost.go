package domain

import "gorm.io/gorm"

// SysUserPost 用户岗位
type SysUserPost struct {
	UserId     int64         `gorm:"column:user_id;primaryKey" json:"userId,string"` // 用户ID
	PostId     int64         `gorm:"column:post_id;primaryKey" json:"postId,string"` // 岗位ID
	CreateTime LocalDateTime `gorm:"column:create_time" json:"createTime,omitempty"` // 创建时间
	UpdateTime LocalDateTime `gorm:"column:update_time" json:"updateTime,omitempty"` // 修改时间
}

// TableName gorm指定表名
func (SysUserPost) TableName() string {
	return "sys_user_post"
}

func (u *SysUserPost) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysUserPost) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}
