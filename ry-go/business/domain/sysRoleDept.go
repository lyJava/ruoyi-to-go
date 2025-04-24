package domain

import "gorm.io/gorm"

// SysRoleDept 角色部门
type SysRoleDept struct {
	RoleId     int64         `gorm:"column:role_id;primaryKey" json:"roleId,string"` // 角色ID
	DeptId     int64         `gorm:"column:dept_id;primaryKey" json:"deptId,string"` // 部门ID
	CreateTime LocalDateTime `gorm:"column:create_time" json:"createTime,omitempty"` // 创建时间
	UpdateTime LocalDateTime `gorm:"column:update_time" json:"updateTime,omitempty"` // 修改时间
}

// TableName gorm指定表名
func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}

func (u *SysRoleDept) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysRoleDept) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}
