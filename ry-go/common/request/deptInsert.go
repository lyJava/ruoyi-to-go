package request

import (
	"time"
)

// DeptInsert 系统部门新增结构体
type DeptInsert struct {
	Id         int64     `json:"id"`                                             // 主键ID
	ParentId   int64     `json:"parentId"`                                       // 父部门ID
	Ancestors  string    `json:"ancestors"`                                      // 组级列表
	DeptName   string    `json:"deptName,omitempty"`                             // 部门主键ID
	OrderNum   int64     `json:"sortNum,omitempty"`                              // 部门排序
	DeptStatus string    `json:"deptStatus,omitempty"`                           // 部门状态（0正常 1停用）
	CreateBy   string    `json:"createBy,omitempty"`                             // 创建人
	CreateTime time.Time `json:"createTime,omitempty" gorm:"column:create_time"` // 创建时间
	UpdateBy   string    `json:"updateBy,omitempty"`                             // 修改人
	Remarks    string    `json:"remarks,omitempty"`                              // 备注
	DelFlag    string    `json:"delFlag,omitempty"`                              // 是否删除(0:正常；1:删除)
}
