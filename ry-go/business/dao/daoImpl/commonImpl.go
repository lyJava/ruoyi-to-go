package daoImpl

import (
	"errors"
	"time"
)

// parseTimeRange 判断时间并解析设置范围
func parseTimeRange(beginStr, endStr string) (time.Time, time.Time, error) {
	begin, err1 := time.Parse(time.DateOnly, beginStr)
	end, err2 := time.Parse(time.DateOnly, endStr)

	if err1 != nil || err2 != nil {
		return time.Time{}, time.Time{}, errors.New("时间格式错误")
	}
	return begin, end.AddDate(0, 0, 1).Add(-time.Nanosecond), nil
}
