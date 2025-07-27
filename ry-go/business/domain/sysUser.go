package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"ry-go/utils"
)

type User struct {
	Id            int64         `gorm:"column:id;primaryKey;autoIncrement" json:"id,string,omitempty"`                                                                              // 主键ID
	DeptId        int64         `gorm:"column:dept_id" json:"deptId"`                                                                                                               // 部门ID
	Username      string        `gorm:"column:username" json:"username" valid:"required~用户名不能为空,length(6|20)~用户名长度必须是6到20位"`                                                        // 用户名
	Nickname      string        `gorm:"column:nickname" json:"nickname" valid:"required~昵称不能为空,length(6|20)~昵称长度必须是6到20位"`                                                          // 昵称
	UserType      string        `gorm:"column:user_type" json:"userType"`                                                                                                           // 用户类型
	Email         string        `gorm:"column:email" json:"email" valid:"required~邮箱不能为空,email~邮箱格式不正确"`                                                                            // 邮箱
	PhoneNo       string        `gorm:"column:phone_no" json:"phoneNo" valid:"required~手机号不能为空,matches(^1(3\\d|4[5-9]|5[0-35-9]|6[0-9]|7[0-35-8]|8\\d|9[0-35-9])\\d{8}$)~手机号格式不正确"` // 手机号
	Sex           string        `gorm:"column:sex" json:"sex" valid:"required~性别不能为空,customerSexValidator~性别参数只能为0，1，2"`                                                            // 用户性别（0男 1女 2未知）
	Avatar        string        `gorm:"column:avatar" json:"avatar"`                                                                                                                // 头像
	UserPwd       string        `gorm:"column:user_pwd" json:"password" valid:"required~密码不能为空,length(6|20)~密码长度必须是6到20位"`                                                          // 密码
	UserStatus    string        `gorm:"column:user_status" json:"userStatus" valid:"required~帐号状态不能为空,customerStatusValidator~帐号状态参数只能为0或1"`                                        // 帐号状态（0正常 1停用）
	CreateBy      string        `gorm:"column:create_by" json:"createBy"`                                                                                                           // 创建人
	CreateTime    LocalDateTime `gorm:"column:create_time" json:"createTime"`                                                                                                       // 创建时间
	UpdateBy      string        `gorm:"column:update_by" json:"updateBy"`                                                                                                           // 修改人
	UpdateTime    LocalDateTime `gorm:"column:update_time" json:"updateTime"`                                                                                                       // 更新时间
	Remarks       string        `gorm:"column:remarks" json:"remarks"`
	DelFlag       string        `gorm:"column:del_flag" json:"delFlag"`
	IsAdmin       bool          `gorm:"-" json:"isAdmin"`        // 是否是管理员
	Dept          *Dept         `gorm:"-" json:"dept,omitempty"` // 部门信息
	Roles         []*SysRole    `gorm:"-" json:"roles"`          // 所有角色
	Posts         []*SysPost    `gorm:"-" json:"posts"`          // 所有岗位
	RoleIds       []int64       `gorm:"-" json:"roleIds"`        // 角色组
	PostIds       []int64       `gorm:"-" json:"postIds"`        // 岗位组
	DeptName      string        `gorm:"-" json:"deptName"`       // 部门名称
	PostNameArray []string      `gorm:"-" json:"postNameArray"`  // 岗位名称切片
	RoleNameArray []string      `gorm:"-" json:"roleNameArray"`  // 角色名称切片
}

// TableName gorm指定表名
func (User) TableName() string {
	return "sys_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = LocalDateTimeNow("")
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateTime = LocalDateTimeNow("")
	return
}

type UserForExcel struct {
	Id         int64  `json:"id,string" excel:"yes"`                                 // 主键ID
	DeptId     int64  `json:"deptId,string,omitempty" excel:"yes"`                   // 部门ID
	Username   string `gorm:"column:username" json:"username,omitempty" excel:"yes"` // 用户名
	Nickname   string `gorm:"column:nickname" json:"nickname,omitempty" excel:"yes"` // 昵称
	UserType   string `json:"userType,omitempty" excel:"yes"`                        // 用户类型
	Email      string `json:"email,omitempty" excel:"yes"`                           // 邮箱
	PhoneNo    string `json:"phoneNo,omitempty" excel:"yes"`                         // 手机号
	Sex        string `json:"sex,omitempty" excel:"yes"`                             // 用户性别（0男 1女 2未知）
	Avatar     string `json:"avatar,omitempty" excel:"yes"`                          // 头像
	UserPwd    string `json:"userPwd,omitempty" excel:"yes"`                         // 密码
	UserStatus string `json:"userStatus,omitempty" excel:"yes"`                      // 帐号状态（0正常 1停用）
	CreateBy   string `json:"createBy,omitempty" excel:"yes"`                        // 创建人
	CreateTime string `json:"createTime,omitempty" excel:"yes"`                      // 创建时间
	UpdateBy   string `json:"updateBy,omitempty" excel:"yes"`                        // 修改人
	UpdateTime string `json:"updateTime,omitempty" excel:"yes"`                      // 更新时间
	Remarks    string `json:"remarks,omitempty" excel:"yes"`                         // 备注
	DelFlag    string `json:"delFlag,omitempty" excel:"yes"`                         // 是否删除(0:正常；1:删除)
	DeptName   string `gorm:"-" json:"DeptName,string,omitempty" excel:"yes"`        // 部门名称
}

// UserLogin 用户登陆
type UserLogin struct {
	Username string `json:"username,omitempty" valid:"required~用户名不能为空"`    // 用户名
	UserPwd  string `json:"password,omitempty" valid:"required~登陆密码密码不能为空"` // 密码
	Uuid     string `json:"uuid,omitempty" valid:"required~唯一ID不能为空"`       // 用户唯一ID
	Code     string `json:"code,omitempty" valid:"required~验证码不能为空"`
}

// LoginUser 用户登录返回结构
type LoginUser struct {
	UserId        int64      `json:"id,string,omitempty"`     // 用户ID
	Uuid          string     `json:"uuid,omitempty"`          // 用户token的uuid
	Username      string     `json:"username,omitempty"`      // 用户名称
	DeptId        int64      `json:"deptId,string,omitempty"` // 部门ID
	Token         string     `json:"token,omitempty"`         // 凭证信息
	LoginTime     int64      `json:"loginTime,omitempty"`     // 登陆时间
	ExpireTime    int64      `json:"expireTime,omitempty"`    // 过期时间
	Permissions   []string   `json:"permissions,omitempty"`   // 用户权限
	IsAdmin       bool       `json:"isAdmin,omitempty"`       // 是否是管理员
	Dept          *Dept      `json:"dept,omitempty"`          // 用户部门信息
	Roles         []*SysRole `json:"role,omitempty"`          // 用户角色信息
	User          *User      `json:"user,omitempty"`          // 用户信息
	DeptName      string     `json:"deptName"`                // 部门信息
	IP            string     `json:"ip,omitempty"`            // 登陆IP
	Browser       string     `json:"browser,omitempty"`       // 浏览器信息
	Os            string     `json:"os,omitempty"`            // 系统信息
	LoginLocation string     `json:"loginLocation,omitempty"` // 登陆地点
}

type UserUpdate struct {
	Id         int64   `json:"id,string"` // 主键ID
	DeptId     int64   `json:"deptId"`    // 部门ID
	Username   string  `json:"username" valid:"required~用户名不能为空"`
	UserType   string  `json:"userType"`                                                                                                            // 用户类型
	Avatar     string  `json:"avatar"`                                                                                                              // 头像 // 用户名
	Nickname   string  `json:"nickname" valid:"required~昵称不能为空"`                                                                                    // 昵称
	Email      string  `json:"email" valid:"required~邮箱不能为空,email~邮箱格式不正确"`                                                                         // 邮箱
	PhoneNo    string  `json:"phoneNo" valid:"required~手机号不能为空,matches(^1(3\\d|4[5-9]|5[0-35-9]|6[0-9]|7[0-35-8]|8\\d|9[0-35-9])\\d{8}$)~手机号格式不正确"` // 手机号
	Sex        string  `json:"sex" valid:"required~性别不能为空,customerSexValidator~性别参数只能为0，1，2"`                                                       // 用户性别（0男 1女 2未知）
	UserStatus string  `json:"userStatus" valid:"required~帐号状态不能为空,customerStatusValidator~帐号状态参数只能为0或1"`                                           // 帐号状态（0正常 1停用）
	Remarks    string  `json:"remarks"`
	RoleIds    []int64 `json:"roleIds"` // 角色组
	PostIds    []int64 `json:"postIds"` // 岗位组
}

// MyClaims 自定义jwt的token返回字段
type MyClaims struct {
	Uuid                 string `json:"uuid"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        //jwt的Claims
}

var sexMap = map[string]string{
	"0": "女",
	"1": "男",
	"2": "未知",
}

var statusMap = map[string]string{
	"0": "正常",
	"1": "停用",
}

func CustomerSexValidator(i any, ctx any) bool {
	sex, ok := i.(string)
	if !ok {
		return false
	}
	return utils.MapContainsKey(sexMap, sex)
}

func CustomerStatusValidator(i any, ctx any) bool {
	status, ok := i.(string)
	if !ok {
		return false
	}
	return utils.MapContainsKey(statusMap, status)
}

type LoginUserResp struct {
	User *User    `json:"user,omitempty"`      // 用户信息
	Role []string `json:"rolesInfo,omitempty"` // 用户信息
}

func IsAdmin(id int64) bool {
	return id == 1
}
