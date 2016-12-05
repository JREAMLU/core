package com

import "time"

const (
	// TimeFormatDefault default
	TimeFormatDefault = "2006-01-02 15:04:05"
	// TimeFormatyymmdd yymmdd
	TimeFormatyymmdd = "060102"
	// TimeFormatyyyymmdd yyyymmdd
	TimeFormatyyyymmdd = "20060102"
	// TimeFormatyyyymm yyyymm
	TimeFormatyyyymm = "200601"
)

//Today 当天，时分秒为0
func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// CurrHourUnix current hour unix
func CurrHourUnix() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()).Unix()
}

// Formatyymmdd yymmdd
func Formatyymmdd(date time.Time) string {
	return date.Format(TimeFormatyymmdd)
}

// Formatyyyymmdd yyyymmdd
func Formatyyyymmdd(date time.Time) string {
	return date.Format(TimeFormatyyyymmdd)
}

// Formatyyyymm yyyymm
func Formatyyyymm(date time.Time) string {
	return date.Format(TimeFormatyyyymm)
}

// FormatDefault default
func FormatDefault(date time.Time) string {
	return date.Format(TimeFormatDefault)
}

//TicksToTime c#中的时间Ticks转成time.Time
func TicksToTime(ticks int64) time.Time {
	ticks = ticks / 10
	n := int64(1000000)
	return time.Unix(ticks/n, ticks-(ticks/n)*n).AddDate(-1969, 0, 0).Add(-8 * time.Hour)
}

// TimeToTicks time to ticks
func TimeToTicks(t time.Time) int64 {
	return t.AddDate(1969, 0, 0).Unix() * 10000000
}

//TicksToUnixNano c#中的时间Ticks转成UnixNano
func TicksToUnixNano(ticks int64) int64 {
	return TicksToTime(ticks).UnixNano()
}
