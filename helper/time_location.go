package helper

import "time"

func IDNLocalTime() (loc *time.Location) {
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")
	return timeLoc
}
