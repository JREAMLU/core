package com

import "time"

const (
	Time_Format_default  = "2006-01-02 15:04:05"
	Time_Format_yymmdd   = "060102"
	Time_Format_yyyymmdd = "20060102"
	Time_Format_yyyymm   = "200601"
)

//Today 当天，时分秒为0
func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func CurrHourUnix() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()).Unix()
}

func Formatyymmdd(date time.Time) string {
	return date.Format(Time_Format_yymmdd)
}

func Formatyyyymmdd(date time.Time) string {
	return date.Format(Time_Format_yyyymmdd)
}

func Formatyyyymm(date time.Time) string {
	return date.Format(Time_Format_yyyymm)
}

func FormatDefault(date time.Time) string {
	return date.Format(Time_Format_default)
}

//TicksToTime c#中的时间Ticks转成time.Time
func TicksToTime(ticks int64) time.Time {
	ticks = ticks / 10
	n := int64(1000000)
	return time.Unix(ticks/n, ticks-(ticks/n)*n).AddDate(-1969, 0, 0).Add(-8 * time.Hour)
}

func TimeToTicks(t time.Time) int64 {
	return t.AddDate(1969, 0, 0).Unix() * 10000000

}

//TicksToUnixNano c#中的时间Ticks转成UnixNano
func TicksToUnixNano(ticks int64) int64 {
	return TicksToTime(ticks).UnixNano()
}
