package util

import (
	"strconv"
	"time"
)

func FormatAmount(amount float64) string {
	return strconv.FormatFloat(amount, 'f', 2, 64)
}

func FormatAmountInt(amount float64) int64 {
	data := strconv.FormatFloat(amount*100, 'f', 0, 64)
	dataint, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0
	}
	return dataint
}

func FormatTime(time time.Time) string {
	rTime := time.Format("2006-01-02 15:04:05")
	if rTime == "0001-01-01 00:00:00" {
		return ""
	}
	return rTime
}
