package request

import (
	"github.com/spf13/cast"
)

// BasePageParam 分页公有参数
type BasePageParam struct {
	Page   int64  `json:"page"`   // 当前页码
	Size   int64  `json:"size"`   // 每页条数
	Column string `json:"column"` // 排序的列
	Order  string `json:"order"`  // 降序或者升序
}

func NewPageParam(pageNum, pageSize, column, order string) *BasePageParam {
	return &BasePageParam{
		Page:   cast.ToInt64(pageNum),
		Size:   cast.ToInt64(pageSize),
		Column: column,
		Order:  order,
	}
}

type MenuPageParam struct {
	CommonPageParam
	MenuName   string `query:"menuName"` // 菜单名称
	UpdateBy   string `query:"updateBy"`
	MenuStatus string `query:"menuStatus"`
	UserId     int64  `query:"userId"`
}

type DeptPageParam struct {
	CommonPageParam
	DeptName string `query:"deptName"` // 部门名称
	Remarks  string `query:"remarks"`  // 备注
	UpdateBy string `query:"updateBy"`
	Status   string `query:"status"`
}

// CommonPageParam 公用分页与排序时间范围参数
type CommonPageParam struct {
	Page      int64  `query:"pageNum"`
	Size      int64  `query:"pageSize"`
	Column    string `query:"column"`
	Order     string `query:"order"`
	BeginTime string `query:"params[beginTime]"`
	EndTime   string `query:"params[endTime]"`
}

// SysNoticePageParam 通知公告分页查询参数
type SysNoticePageParam struct {
	CommonPageParam
	Title    string `query:"noticeTitle"`  // 公共标题
	Type     string `query:"noticeType"`   // 类型
	CreateBy string `query:"createBy"`     // 内容
	Status   string `query:"noticeStatus"` // 状态
}

// UserPageParam 用户分页查询参数
type UserPageParam struct {
	CommonPageParam
	Username   string `query:"username"` // 用户名称
	Nickname   string `query:"nickname"` // 用户名称
	PhoneNo    string `query:"phoneNo"`  // 手机
	Sex        string `query:"sex"`
	UserStatus string `query:"userStatus"` // 状态
	DeptId     string `query:"deptId"`     // 部门ID
}

// RolePageParam 角色分页查询参数
type RolePageParam struct {
	CommonPageParam
	RoleName   string `query:"roleName"`
	RoleStatus string `query:"status"`
	RoleKey    string `query:"roleKey"`
}

type PostPageParam struct {
	CommonPageParam
	PostCode   string `query:"postCode"`   // 岗位编码
	PostName   string `query:"postName"`   // 岗位名称
	PostStatus string `query:"postStatus"` // 状态
}

type DictTypePageParam struct {
	CommonPageParam
	DictName   string `query:"dictName"` // 字典名称
	DictType   string `query:"dictType"` // 字典类型
	DictStatus string `query:"status"`   // 状态
}

// DictDataPageParam 字典数据分页参数
type DictDataPageParam struct {
	CommonPageParam
	DictType   string `query:"dictType"`   // 字典类型
	DictLabel  string `query:"dictLabel"`  // 字典标签
	DictStatus string `query:"dictStatus"` // 状态
}

// ConfigPageParam 参数配置分页查询参数
type ConfigPageParam struct {
	CommonPageParam
	ConfigName string `query:"configName"` // 参数名称
	ConfigKey  string `query:"configKey"`  // 参数键
	ConfigType string `query:"configType"` // 参数类型
}

type OnlinePageParam struct {
	CommonPageParam
	Username string `query:"userName"` // 用户名称
	Ip       string `query:"ip"`       // 用户名称
}

type JobChangeParam struct {
	JobId  string `json:"jobId"`
	Status string `json:"status"`
}
