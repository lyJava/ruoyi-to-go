package domain

// SysLoginRecord 登录记录结构
//
//goland:noinspection GoUnusedExportedType
type SysLoginRecord struct {
	InfoId        int64         `gorm:"column:info_id;primaryKey;autoIncrement" json:"infoId,string"` // 访问ID
	UserName      string        `gorm:"column:user_name" json:"userName,omitempty"`                   // 用户账号
	Ipaddr        string        `gorm:"column:ipaddr" json:"ipaddr,omitempty"`                        // 登录IP地址
	LoginLocation string        `gorm:"column:login_location" json:"loginLocation,omitempty"`         // 登录地点
	Browser       string        `gorm:"column:browser" json:"browser,omitempty"`                      // 浏览器类型
	Os            string        `gorm:"column:os" json:"os,omitempty"`                                // 操作系统
	Status        string        `gorm:"column:status" json:"status,omitempty"`                        // 登录状态（0成功 1失败）
	Msg           string        `gorm:"column:msg" json:"msg,omitempty"`                              // 提示消息
	LoginTime     LocalDateTime `gorm:"column:login_time" json:"loginTime,omitempty"`                 // 访问时间
	Token         string        `gorm:"column:token" json:"token,omitempty"`                          // 登录token
}

// TableName gorm指定表名
func (SysLoginRecord) TableName() string {
	return "sys_login_record"
}

// SysLoginRecordExcel 登录记录导出excel结构
//
//goland:noinspection GoUnusedExportedType
type SysLoginRecordExcel struct {
	InfoId        string `gorm:"column:info_id" json:"infoId,string"`                  // 访问ID
	UserName      string `gorm:"column:user_name" json:"userName,omitempty"`           // 用户账号
	Ipaddr        string `gorm:"column:ipaddr" json:"ipaddr,omitempty"`                // 登录IP地址
	LoginLocation string `gorm:"column:login_location" json:"loginLocation,omitempty"` // 登录地点
	Browser       string `gorm:"column:browser" json:"browser,omitempty"`              // 浏览器类型
	Os            string `gorm:"column:os" json:"os,omitempty"`                        // 操作系统
	Status        string `gorm:"column:status" json:"status,omitempty"`                // 登录状态（0成功 1失败）
	Msg           string `gorm:"column:msg" json:"msg,omitempty"`                      // 提示消息
	LoginTime     string `gorm:"column:login_time" json:"loginTime,omitempty"`         // 访问时间
	Token         string `gorm:"column:token" json:"token,omitempty"`                  // 登录token
}
