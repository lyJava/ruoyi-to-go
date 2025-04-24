package domain

// Dept 部门信息
type Dept struct {
	Id         int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id,string"` // 部门主键ID(设置返回json类型为字符串，防止丢失精度)
	ParentId   int64  `gorm:"column:parent_id" json:"parentId,string"`             // 父部门ID(设置返回json类型为字符串)
	Ancestors  string `gorm:"column:ancestors" json:"ancestors"`                   // 组级列表
	DeptName   string `gorm:"column:dept_name" json:"deptName"`                    // 部门名称
	OrderNum   int64  `gorm:"column:order_num" json:"orderNum"`                    // 部门排序
	DeptStatus string `gorm:"column:dept_status" json:"deptStatus"`                // 部门状态（0正常 1停用）
	CreateBy   string `gorm:"column:create_by" json:"createBy"`                    // 创建人
	BaseDomain
	Leader     string  `gorm:"column:leader" json:"leader"`    // 负责人
	Phone      string  `gorm:"column:phone" json:"phone"`      // 手机
	Email      string  `gorm:"column:email" json:"email"`      // 邮箱
	DelFlag    string  `gorm:"column:del_flag" json:"delFlag"` // 是否删除(0:正常；1:删除)
	Children   []*Dept `gorm:"-" json:"children"`
	ParentName string  `gorm:"-" json:"parentName"` // 上级部门名称
}

// TableName gorm指定表名
func (Dept) TableName() string {
	return "sys_dept"
}

type TreeSelect struct {
	Id       int64         `json:"id"`
	Label    string        `json:"label"`
	Children []*TreeSelect `json:"children"`
}

func NewTreeSelectByDept(dept *Dept) *TreeSelect {
	deptTreeData := &TreeSelect{
		Id:    dept.Id,
		Label: dept.DeptName,
	}
	// 子节点
	children := dept.Children
	if len(children) > 0 {
		var treeList []*TreeSelect
		for _, child := range children {
			// 递归生成子节点的TreeSelect
			childTree := NewTreeSelectByDept(child)
			treeList = append(treeList, childTree)
		}
		deptTreeData.Children = treeList
	}
	return deptTreeData
}

func NewTreeSelectByMenu(menu *SysMenu) *TreeSelect {
	menuTreeData := &TreeSelect{
		Id:    menu.Id,
		Label: menu.MenuName,
	}
	// 子节点
	children := menu.Children
	if len(children) > 0 {
		var treeList []*TreeSelect
		for _, child := range children {
			// 递归生成子节点的TreeSelect
			childTree := NewTreeSelectByMenu(child)
			treeList = append(treeList, childTree)
		}
		menuTreeData.Children = treeList
	}

	return menuTreeData
}

// RoleMenuTreeSelect 角色菜单treeSelect
type RoleMenuTreeSelect struct {
	CheckedKeys []int64       `json:"checkedKeys"`
	Menus       []*TreeSelect `json:"menus"`
}
