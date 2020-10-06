package utils

import "time"

// 时间字串转时间对象，固定格式：2006-01-02 15:04:05
func DateTimeStrToTime(datetime string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
	return t
}

// 获取当天凌晨时间对象
func GetMidNightObj(t time.Time) time.Time {
	zero_tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return zero_tm
}
