package egb

import (
	"fmt"
	"time"
	"strconv"
)

//TimeYear return now year string.
//eg:2016
func TimeYear() string {
	now := time.Now()
	year, _, _ := now.Date()
	return fmt.Sprintf("%d", year)
}

//TimeMonth return now month string.
//eg:8
func TimeMonth() string {
	now := time.Now()
	_, month, _ := now.Date()
	return fmt.Sprintf("%d", month)
}

//TimeDay return now day string.
//eg:5
func TimeDay() string {
	now := time.Now()
	_, _, day := now.Date()
	return fmt.Sprintf("%d", day)
}

//TimeWeekDay return now week day string.
//eg:Friday
func TimeWeekDay() string {
	now := time.Now()
	return fmt.Sprintf("%s", now.Weekday().String())
}

//TimeFromUnix return normal format time from unix time string.
func TimeFromUnix(unix string) string {
	i, _ := strconv.ParseInt(unix, 10, 64)
	str_time := time.Unix(i, 0).Format("2006-01-02 15:04:05")
	return str_time
}

//TimeFromUnixNano return normal format time from unix nano time string.
func TimeFromUnixNano(unix string) string {
	i, _ := strconv.ParseInt(unix, 10, 64)
	str_time := time.Unix(0, i).Format("2006-01-02 15:04:05")
	return str_time
}

//TimeNowDate return now date time string.
//eg:2016-08-05 15:04:05
func TimeNowDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//TimeNowDateDay return now date day time string.
//eg:2016-08-05
func TimeNowDateDay() string {
	return time.Now().Format("2006-01-02")
}

//TimeNowUnix return now unix time string.
func TimeNowUnix() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

//TimeNowUnixMs return now unix ms time string.
func TimeNowUnixMs() string {
	return StringSubStr(TimeNowUnixNano(), 0, 13)
}

//TimeNowUnixNano return now unix nano time string.
//eg:1471226178 882 997 341
func TimeNowUnixNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

//TimeDayToUnix return unix time string by given format-time.
//input:2016-08-05 output:1470355200
func TimeDayToUnix(daytime string) string {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, daytime, loc) //使用模板在对应时区转化为time.time类型
	unix := theTime.Unix()
	return strconv.FormatInt(unix, 10)
}

//TimeSecondToUnix return unix time string by given format-time.
//input:2016/8/5 15:04:02 output:1470380642
func TimeSecondToUnix(stime string) string {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, stime, loc) //使用模板在对应时区转化为time.time类型
	unix := theTime.Unix()
	return strconv.FormatInt(unix, 10)
}

