package helper

import "time"

// Format string date
const (
	DDMMYYYYhhmmss = "20060102150405"
	StartHour      = "000000"
	EndHour        = "595959"
)

func GetCurrentDate() string {
	now := time.Now()
	var getCurrentDate = string(now.Format(DDMMYYYYhhmmss))
	return getCurrentDate

}

func GetBeforeDate() string {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	var getCurrentDate = string(yesterday.Format(DDMMYYYYhhmmss))
	return getCurrentDate

}

var (
	BeginCurrentDate = GetCurrentDate()[0:8] + StartHour
	EndCurrentDate   = GetCurrentDate()[0:8] + EndHour
)
