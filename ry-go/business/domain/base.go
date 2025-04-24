package domain

import (
	"database/sql/driver"
	"fmt"
	"log"
	"time"
)

// BaseDomain 公共字段
type BaseDomain struct {
	CreateBy   string        `gorm:"column:create_by" json:"createBy,omitempty"`     // 创建人
	CreateTime LocalDateTime `gorm:"column:create_time" json:"createTime,omitempty"` // 创建时间
	UpdateBy   string        `gorm:"column:update_by" json:"updateBy,omitempty"`     // 修改人
	UpdateTime LocalDateTime `gorm:"column:update_time" json:"updateTime,omitempty"` // 更新时间
	Remarks    string        `gorm:"column:remarks" json:"remarks,omitempty"`        // 备注
}

// LocalDate 自定义日期 下边四个函数保证序列化与反序列化的返回格式，这里会覆盖默认的格式
type LocalDate struct {
	time.Time
}

// LocalDateTime 自定义日期时间 下边四个函数保证序列化与反序列化的返回格式，这里会覆盖默认的格式
type LocalDateTime struct {
	time.Time
}

func LoadDefaultLocation() *time.Location {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return loc
}

func (t *LocalDate) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalDate{Time: time.Time{}}
		return
	}
	now, err := time.ParseInLocation(`"`+time.DateOnly+`"`, string(data), LoadDefaultLocation())
	*t = LocalDate{Time: now}
	return
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Format(time.DateOnly))), nil
}

func (t LocalDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *LocalDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to date", v)
}

func (t *LocalDateTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalDateTime{Time: time.Time{}}
		return
	}
	now, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), LoadDefaultLocation())
	*t = LocalDateTime{Time: now}
	return
}

func (t LocalDateTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		//return []byte("null"), nil
		// 返回空字符串
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Format(time.DateTime))), nil
}

func (t LocalDateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *LocalDateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalDateTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %+v to datetime", v)
}

// LocalDateTimeNow 设置当前日期时间
//
// name: 时区名称，默认Asia/Shanghai
//
// LocalDateTime： 当前日期时间
func LocalDateTimeNow(name string) LocalDateTime {
	if name == "" {
		name = "Asia/Shanghai"
	}
	location, err := time.LoadLocation(name)
	if err != nil {
		log.Printf("LocalDateTime LoadLocation error %+v", err)
	}
	return LocalDateTime{Time: time.Now().In(location)}
}

// LocalDateNow 设置当前日期
//
// name: 时区名称，默认Asia/Shanghai
//
// LocalDate： 当前日期
func LocalDateNow(name string) LocalDate {
	if name == "" {
		name = "Asia/Shanghai"
	}
	location, err := time.LoadLocation(name)
	if err != nil {
		log.Printf("LocalDateTime LoadLocation error %+v", err)
	}
	return LocalDate{Time: time.Now().In(location)}
}

// String 返回默认的yyyy-MM-dd字符串日期
func (t LocalDate) String() string {
	if t.IsZero() {
		return ""
	}
	// 使用时区转换确保输出正确
	return t.Time.In(LoadDefaultLocation()).Format(time.DateOnly)
}

// String 返回默认的yyyy-MM-dd HH:mm:ss字符串日期
func (t LocalDateTime) String() string {
	if t.IsZero() {
		return ""
	}
	// 使用时区转换确保输出正确
	return t.Time.In(LoadDefaultLocation()).Format(time.DateTime)
}
