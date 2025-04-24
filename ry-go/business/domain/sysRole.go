package domain

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// SysRole 系统角色结构体
type SysRole struct {
	Id                int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id,string"`           // 角色ID
	RoleName          string `gorm:"column:role_name" json:"roleName"`                              // 角色名称
	RoleKey           string `gorm:"column:role_key" json:"roleKey"`                                // 角色权限字符串
	RoleSort          int    `gorm:"column:role_sort" json:"roleSort"`                              // 显示顺序
	DataScope         string `gorm:"column:data_scope" json:"dataScope,omitempty"`                  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly string `gorm:"column:menu_check_strictly" json:"menuCheckStrictly,omitempty"` // 菜单树选择项是否关联显示
	DeptCheckStrictly string `gorm:"column:dept_check_strictly" json:"deptCheckStrictly,omitempty"` // 部门树选择项是否关联显示
	RoleStatus        string `gorm:"column:role_status" json:"roleStatus"`                          // 角色状态（0正常 1停用）
	DelFlag           string `gorm:"column:del_flag" json:"delFlag,omitempty"`                      // 删除标志（0代表存在 2代表删除）
	BaseDomain
	MenuIds []int64 `gorm:"-" json:"menuIds,omitempty"` // 角色菜单
	DeptIds []int64 `gorm:"-" json:"deptIds,omitempty"` // 部门组-数据权限

}

// TableName gorm指定表名
func (SysRole) TableName() string {
	return "sys_role"
}

func (u *SysRole) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysRole) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}

// SysRoleExcel 角色信息导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysRoleExcel struct {
	RoleId            string `gorm:"column:role_id" json:"roleId,string"`                           // 角色ID
	RoleName          string `gorm:"column:role_name" json:"roleName"`                              // 角色名称
	RoleKey           string `gorm:"column:role_key" json:"roleKey"`                                // 角色权限字符串
	RoleSort          string `gorm:"column:role_sort" json:"roleSort"`                              // 显示顺序
	DataScope         string `gorm:"column:data_scope" json:"dataScope,omitempty"`                  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly string `gorm:"column:menu_check_strictly" json:"menuCheckStrictly,omitempty"` // 菜单树选择项是否关联显示
	DeptCheckStrictly string `gorm:"column:dept_check_strictly" json:"deptCheckStrictly,omitempty"` // 部门树选择项是否关联显示
	Status            string `gorm:"column:status" json:"status"`                                   // 角色状态（0正常 1停用）
	DelFlag           string `gorm:"column:del_flag" json:"delFlag,omitempty"`                      // 删除标志（0代表存在 2代表删除）
	CreateBy          string `gorm:"column:create_by" json:"createBy,omitempty"`                    // 创建者
	CreateTime        string `gorm:"column:create_time" json:"createTime,omitempty"`                // 创建时间
	UpdateBy          string `gorm:"column:update_by" json:"updateBy,omitempty"`                    // 更新者
	UpdateTime        string `gorm:"column:update_time" json:"updateTime,omitempty"`                // 更新时间
	Remark            string `gorm:"column:remark" json:"remark,omitempty"`                         // 备注
}

// IsAdminRole 角色是否是管理员
func IsAdminRole(roleId int64) bool {
	return roleId == 1
}

type SysRoleInsert struct {
	Id                string  `json:"id,omitempty"`                // 角色名称
	RoleName          string  `json:"roleName"`                    // 角色名称
	RoleKey           string  `json:"roleKey"`                     // 角色权限字符串
	RoleSort          int     `json:"roleSort"`                    // 显示顺序
	DataScope         string  `json:"dataScope,omitempty"`         // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly bool    `json:"menuCheckStrictly,omitempty"` // 菜单树选择项是否关联显示
	DeptCheckStrictly bool    `json:"deptCheckStrictly,omitempty"` // 部门树选择项是否关联显示
	RoleStatus        string  `json:"roleStatus"`                  // 角色状态（0正常 1停用）
	Remarks           string  `json:"remarks"`                     // 备注
	MenuIds           []int64 `json:"menuIds"`                     // 角色菜单
	DeptIds           []int64 `json:"deptIds"`                     // 部门组-数据权限
}

func SysRoleInsertConvertToEntity(obj *SysRoleInsert) *SysRole {
	role := &SysRole{
		Id:         cast.ToInt64(obj.Id),
		RoleName:   obj.RoleName,
		RoleKey:    obj.RoleKey,
		RoleStatus: obj.RoleStatus,
		DataScope:  obj.DataScope,
		RoleSort:   obj.RoleSort,
		MenuIds:    obj.MenuIds,
		DeptIds:    obj.DeptIds,
	}
	role.Remarks = obj.Remarks

	if obj.MenuCheckStrictly {
		role.MenuCheckStrictly = "1"
	} else {
		role.MenuCheckStrictly = "0"
	}

	if obj.DeptCheckStrictly {
		role.DeptCheckStrictly = "1"
	} else {
		role.DeptCheckStrictly = "0"
	}

	return role
}

type SysRoleDataScope struct {
	Id                string  `json:"id,omitempty"`        // 角色名称
	DataScope         string  `json:"dataScope,omitempty"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	DeptIds           []int64 `json:"deptIds"`
	MenuCheckStrictly string  `json:"menuCheckStrictly,omitempty"` // 菜单树选择项是否关联显示
	DeptCheckStrictly string  `json:"deptCheckStrictly,omitempty"` // 部门树选择项是否关联显示
}

func SysRoleDataScopeConvertToEntity(obj *SysRoleDataScope) *SysRole {
	sysRole := &SysRole{
		Id:                cast.ToInt64(obj.Id),
		DataScope:         obj.DataScope,
		DeptIds:           obj.DeptIds,
		MenuCheckStrictly: obj.MenuCheckStrictly,
		DeptCheckStrictly: obj.DeptCheckStrictly,
	}

	sysRole.UpdateTime = LocalDateTimeNow("")
	return sysRole
}
