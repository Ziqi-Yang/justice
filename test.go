package main

import (
	// "fmt"
	"time"
)

func isInCompetitionTime(begin string, end string) bool {
	// though users can change time to pass the program

	// timeLocal, _ := time.LoadLocation("Asia/Shanghai")
	timeLocal := time.FixedZone("CST", 8*3600)
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", begin, timeLocal)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", end, timeLocal)
	nowTime := time.Now()
	if nowTime.After(startTime) && nowTime.Before(endTime) {
		return true
	}
	return false
}

func main() {
	print(isInCompetitionTime("2021-11-04 00:00:00", "2021-11-04 10:22:00"))
}
