package models

import (
	"github.com/beego/beego/v2/core/logs"
	"time"
)

// 时间戳转时间
func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

// 时间转时间戳
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		logs.Info(err)
		return 0
	}

	return t.Unix()
}
