package domain

import "gorm.io/gorm"

// SysMenu 系统菜单结构体
type SysMenu struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id,string"` // 菜单ID
	MenuName   string     `gorm:"column:menu_name" json:"menuName"`                    // 菜单名称
	ParentId   int64      `gorm:"column:parent_id" json:"parentId,string"`             // 父菜单ID(这里的返回json的类型必须与前端一致，不然会影响表格树形的折叠展开)
	OrderNum   int64      `gorm:"column:order_num" json:"orderNum"`                    // 显示顺序
	RouterPath string     `gorm:"column:router_path" json:"routerPath"`                // 路由地址
	Component  string     `gorm:"column:component" json:"component"`                   // 组件路径
	IsFrame    int64      `gorm:"column:is_frame" json:"isFrame"`                      // 是否为外链（0是 1否）
	IsCache    int64      `gorm:"column:is_cache" json:"isCache"`                      // 是否缓存（0缓存 1不缓存）
	MenuType   string     `gorm:"column:menu_type" json:"menuType"`                    // 菜单类型（M目录 C菜单 F按钮）
	Visible    string     `gorm:"column:visible" json:"visible"`                       // 菜单状态（0显示 1隐藏）
	MenuStatus string     `gorm:"column:menu_status" json:"menuStatus"`                // 菜单状态（0正常 1停用）
	Perms      string     `gorm:"column:perms" json:"perms"`                           // 权限标识
	Icon       string     `gorm:"column:icon" json:"icon"`                             // 菜单图标
	Children   []*SysMenu `gorm:"-" json:"children"`
	BaseDomain
	BeginTime string `gorm:"-" json:"beginTime"`
	EndTime   string `gorm:"-" json:"endTime"`
}

// TableName gorm指定表名
func (SysMenu) TableName() string {
	return "sys_menu"
}

func (u *SysMenu) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *SysMenu) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}

// SysMenuExcel 系统菜单结构体
type SysMenuExcel struct {
	Id         int64  `gorm:"column:id" json:"id,string" excel:"yes"`                        // 菜单ID
	MenuName   string `gorm:"column:menu_name" json:"menuName" excel:"yes"`                  // 菜单名称
	ParentId   int64  `gorm:"column:parent_id" json:"parentId,string,omitempty" excel:"yes"` // 父菜单ID
	OrderNum   int    `gorm:"column:order_num" json:"orderNum,omitempty" excel:"yes"`        // 显示顺序
	RouterPath string `gorm:"column:router_path" json:"routerPath,omitempty" excel:"yes"`    // 路由地址
	Component  string `gorm:"column:component" json:"component,omitempty" excel:"yes"`       // 组件路径
	IsFrame    int    `gorm:"column:is_frame" json:"isFrame,omitempty" excel:"yes"`          // 是否为外链（0是 1否）
	IsCache    int    `gorm:"column:is_cache" json:"isCache,omitempty" excel:"yes"`          // 是否缓存（0缓存 1不缓存）
	MenuType   string `gorm:"column:menu_type" json:"menuType,omitempty" excel:"yes"`        // 菜单类型（M目录 C菜单 F按钮）
	Visible    string `gorm:"column:visible" json:"visible,omitempty" excel:"yes"`           // 菜单状态（0显示 1隐藏）
	MenuStatus string `gorm:"column:menu_status" json:"menuStatus,omitempty" excel:"yes"`    // 菜单状态（0正常 1停用）
	Perms      string `gorm:"column:perms" json:"perms,omitempty" excel:"yes"`               // 权限标识
	Icon       string `gorm:"column:icon" json:"icon,omitempty" excel:"yes"`                 // 菜单图标
	BaseDomain
}
