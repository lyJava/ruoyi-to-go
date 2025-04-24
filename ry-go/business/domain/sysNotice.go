package domain

import (
	"gorm.io/gorm"
)

// SysNotice 系统公告
type SysNotice struct {
	Id            int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id,string"` // 主键ID(设置返回json类型为字符串，防止丢失精度)
	NoticeTitle   string `gorm:"column:notice_title" json:"noticeTitle"`              // 公告标题(设置返回json类型为字符串)
	NoticeType    string `gorm:"column:notice_type" json:"noticeType"`                // 公告类型（1通知 2公告）
	NoticeContent string `gorm:"column:notice_content" json:"noticeContent"`          // 公告内容
	NoticeStatus  string `gorm:"column:notice_status" json:"noticeStatus"`            // 公告状态（0正常 1关闭）
	BaseDomain
}

// TableName gorm指定表名
func (SysNotice) TableName() string {
	return "sys_notice"
}

// BeforeCreate 在创建记录之前调用
func (n *SysNotice) BeforeCreate(tx *gorm.DB) error {
	n.CreateTime = LocalDateTimeNow("")
	return nil
}

// AfterCreate 只有创建后才触发
func (n *SysNotice) AfterCreate(tx *gorm.DB) error {
	return nil
}

func (n *SysNotice) BeforeUpdate(tx *gorm.DB) (err error) {
	n.UpdateTime = LocalDateTimeNow("")
	return
}
