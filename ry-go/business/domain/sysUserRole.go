package domain

// SysUserRole 用户角色信息结构
//
//goland:noinspection GoUnusedExportedType
type SysUserRole struct {
	UserId     int64         `gorm:"column:user_id;primaryKey" json:"userId,string"` // 用户ID
	RoleId     int64         `gorm:"column:role_id;primaryKey" json:"roleId,string"` // 角色ID
	CreateTime LocalDateTime `gorm:"column:create_time" json:"createTime,omitempty"` // 创建时间
	UpdateTime LocalDateTime `gorm:"column:update_time" json:"updateTime,omitempty"` // 修改时间
}

// TableName gorm指定表名
func (SysUserRole) TableName() string {
	return "sys_user_role"
}

// SysUserRoleExcel 用户角色信息导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysUserRoleExcel struct {
	UserId     string `gorm:"column:user_id" json:"userId,string"`            // 用户ID
	RoleId     string `gorm:"column:role_id" json:"roleId,string"`            // 角色ID
	CreateTime string `gorm:"column:create_time" json:"createTime,omitempty"` // 创建时间
	UpdateTime string `gorm:"column:update_time" json:"updateTime,omitempty"` // 修改时间
}
