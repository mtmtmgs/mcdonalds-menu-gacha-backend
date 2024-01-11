package utils

import (
	"time"
)

func NowTimeJst() time.Time {
	now := time.Now()
	jstLocation, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}
	}
	nowJst := now.In(jstLocation)
	return nowJst
}

func ConvertTimeUtcToJst(utcTime time.Time) time.Time {
	jstLocation, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}
	}
	jstTime := utcTime.In(jstLocation)
	return jstTime
}

func TimeFormat(time time.Time, format string) string {
	switch format {
	case "yyyyMMddHHmmss":
		return time.Format("2006-01-02 15:04:05")
	case "yyyyMMdd":
		return time.Format("2006-01-02")
	case "HHmmss":
		return time.Format("15:04:05")
	default:
		return time.Format("2006-01-02 15:04:05")
	}
}
