package domain

import "gorm.io/gorm"

// SysRoleMenu 角色菜单
type SysRoleMenu struct {
	RoleId     int64         `gorm:"column:role_id;primaryKey" json:"roleId,string"` // 角色ID
	MenuId     int64         `gorm:"column:menu_id;primaryKey" json:"menuId,string"` // 菜单ID
	CreateTime LocalDateTime `gorm:"column:create_time" json:"createTime,omitempty"` // 创建时间
	UpdateTime LocalDateTime `gorm:"column:update_time" json:"updateTime,omitempty"` // 修改时间
}

// TableName gorm指定表名
func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

func (u *SysRoleMenu) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysRoleMenu) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}
