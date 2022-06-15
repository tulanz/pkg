package util

import (
	"fmt"
	"time"
)

const (
	YYYYMMDD        = "20060102"
	YYYYMMDD2       = "2006-01-02"
	YYYYMMDDHHMMSS  = "20060102150405"
	YYYYMMDDHHMMSS2 = "2006-01-02 15:04:05"
	YYYYMMDDHH      = "2006010215"
	YYYYMMDDHH2     = "2006-01-02 15"
	YYYYMMDDHHMM    = "200601021504"
	YYYYMMDDHHMM3   = "2006-01-02 15:04"
	YYYYMM          = "200601"
	YYYYMM2         = "2006-01"
	YYYYMMDDHHMM2   = "2006年01月02日15点04分"
)

func CurrentTimeString() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(YYYYMMDDHHMMSS2)
}

func ParseTimestamp(timestamp int64, layout string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(layout)
}

func CurrentTimeWithFormat(format string) string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func TimeFormat2(t time.Time, format string) string {
	timestamp := t.Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func TimeFormatInLocation(t time.Time, format string, zone string) string {
	timestamp := t.Unix()
	tm := time.Unix(timestamp, 0)
	location, _ := time.LoadLocation(zone)
	tm = tm.In(location)
	return tm.Format(format)
}

func StrToTime(timeStr string, layout string) time.Time {
	loc, _ := time.LoadLocation("Local")
	val, err := time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		return time.Now()
	}
	return val
}

//精确到毫秒
func StartTimeOfDay(parseTime *time.Time) int64 {
	timeStr := fmt.Sprintf("%s %s", time.Now().Format(YYYYMMDD2), "00:00:00")
	if parseTime != nil {
		timeStr = fmt.Sprintf("%s %s", parseTime.Format(YYYYMMDD2), "00:00:00")
	}
	t, _ := time.ParseInLocation(YYYYMMDDHHMMSS2, timeStr, time.Local)
	return t.UnixNano() / int64(time.Millisecond)
}

//精确到毫秒
func EndTimeOfDay() int64 {
	timeStr := fmt.Sprintf("%s %s", time.Now().Format(YYYYMMDD2), "23:59:59")
	t, _ := time.ParseInLocation(YYYYMMDDHHMMSS2, timeStr, time.Local)
	return t.UnixNano() / int64(time.Millisecond)
}

func TimeFormat(parseTime string, layout string, format string) string {
	val, err := time.Parse(layout, parseTime)
	if err != nil {
		return ""
	}
	return val.Format(format)
}

func RangeTime(startTime time.Time, endTime time.Time, format string, field time.Duration) []string {
	var times []string
	for {
		startTime = startTime.Add(field)
		if startTime.UnixNano() >= endTime.UnixNano() {
			return times
		}
		times = append(times, TimeFormat2(startTime, format))
	}
}

func GetTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func GetTimestampSecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Second)
}

func CurrentTimeMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func CurrentTimeSecond() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

func ParseTime(param time.Time, layout string) string {
	timestamp := param.Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(layout)
}

func BeforeTime(duration time.Duration) time.Time {
	return time.Now().Add(-duration)
}

func BeforeDate(years, months, days int) time.Time {
	return time.Now().AddDate(-years, -months, -days)
}

func AfterDate(years, months, days int) time.Time {
	return time.Now().AddDate(years, months, days)
}

func AfterTime(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}

func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func LastMonth(months int) time.Time {
	return time.Now().AddDate(0, months, 0)
}

func FirstDayOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	//获取某一天的0点时间
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func LastDayOfMonth(t time.Time) time.Time {
	return FirstDayOfMonth(t).AddDate(0, 1, -1)
}

//layout:2006-01-02 15:04:05
//parseTime:2006-01-02 15:04:05
func TimestrToUnix(parseTime string, layout string, location string) int64 {
	loc, _ := time.LoadLocation(location)
	theTime, _ := time.ParseInLocation(layout, parseTime, loc)
	return theTime.UnixNano() / int64(time.Millisecond)
}
