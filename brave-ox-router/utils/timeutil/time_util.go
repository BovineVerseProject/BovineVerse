package timeutil

import (
	"time"

	"github.com/araddon/dateparse"
)

const (
	MillisPerSecond = 1000

	MillisPerMinute = 60 * MillisPerSecond

	MillisPerHour = 60 * MillisPerMinute

	MillisPerDay = 24 * MillisPerHour

	MillisPerWeek = 7 * MillisPerDay

	DATE_LAYOUT = "2006-01-02 15:04:05"
)

var (
	defaultTimeZone = time.Local
	weekStartDay    = time.Monday
)

func SetTimeZone(timeZone string) {
	defaultTimeZone, _ = time.LoadLocation(timeZone)
}

func NowSec() int64 {
	return time.Now().Unix()
}

func NowMs() int64 {
	return TimeToMs(time.Now())
}

func TimeUnix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec).In(defaultTimeZone)
}

func TimeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func DateToTime(date string) (time.Time, error) {
	return dateparse.ParseIn(date, defaultTimeZone)
}

func NowFormat() string {
	return time.Now().Format(DATE_LAYOUT)
}

func DateToMs(date string, args ...string) int64 {
	t, err := DateToTime(date)
	if err != nil {
		return 0
	}
	return t.UnixNano() / int64(time.Millisecond)
}

func TodaySec() int64 {
	now := time.Now()
	return now.Unix() - With(now).StartOfDay().Unix()
}

func StartOfDay() time.Time {
	return With(time.Now()).StartOfDay()
}

func EndOfDay() time.Time {
	return With(time.Now()).EndOfDay()
}

func StartOfWeek() time.Time {
	return With(time.Now()).StartOfWeek()
}

func EndOfWeek() time.Time {
	return With(time.Now()).EndOfWeek()
}

func StartOfMonth() time.Time {
	return With(time.Now()).StartOfMonth()
}

func EndOfMonth() time.Time {
	return With(time.Now()).EndOfMonth()
}

func Between(start, end time.Time) bool {
	return With(time.Now()).Between(start, end)
}

func GetDaysBetween(start, end time.Time) int64 {
	d := end.Sub(start)
	return int64(d / (time.Hour * 24))
}

func Time(t interface{}) time.Time {
	switch v := t.(type) {
	case int64:
		return time.Unix(v, 0)
	case string:
		time, _ := DateToTime(v)
		return time
	case time.Time:
		return v
	}
	return time.Time{}
}

func GetTomorrowOnHour(hour int) time.Time {
	return With(time.Now()).GetTomorrowOnHour(hour)
}

func GetNextWeekDay(week int) time.Time {
	return With(time.Now()).GetNextWeekDay(week)
}

type TimeUtil struct {
	time.Time
	ext time.Time
}

func With(t time.Time, ext ...time.Time) *TimeUtil {
	if len(ext) > 0 {
		return &TimeUtil{Time: t, ext: ext[0]}
	}
	return &TimeUtil{Time: t}
}

func (timeUtil *TimeUtil) StartOfDay() time.Time {
	y, m, d := timeUtil.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, timeUtil.Location())
}

func StartOfDay2(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func (timeUtil *TimeUtil) EndOfDay() time.Time {
	y, m, d := timeUtil.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), defaultTimeZone)
}

func (timeUtil *TimeUtil) StartOfWeek() time.Time {
	t := timeUtil.StartOfDay()
	weekday := int(timeUtil.Weekday())
	startday := int(weekStartDay)
	if startday != int(time.Sunday) {
		if weekday < startday {
			weekday = weekday + 7 - startday
		} else {
			weekday -= startday
		}
	}
	return t.AddDate(0, 0, -weekday)
}

func (timeUtil *TimeUtil) EndOfWeek() time.Time {
	return timeUtil.StartOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (timeUtil *TimeUtil) StartOfMonth() time.Time {
	y, m, _ := timeUtil.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, timeUtil.Location())
}

func (timeUtil *TimeUtil) EndOfMonth() time.Time {
	return timeUtil.StartOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (timeUtil *TimeUtil) IsSameDay() bool {
	return timeUtil.YearDay() == timeUtil.ext.YearDay()
}

func (timeUtil *TimeUtil) IsSameHour() bool {
	return timeUtil.IsSameDay() && timeUtil.Hour() == timeUtil.ext.Hour()
}

func (timeUtil *TimeUtil) IsSameWeek() bool {
	year1, week1 := timeUtil.ISOWeek()
	year2, week2 := timeUtil.ext.ISOWeek()
	return year1 == year2 && week1 == week2
}

func (timeUtil *TimeUtil) IsSameYear() bool {
	return timeUtil.Year() == timeUtil.ext.Year()
}

func (timeUtil *TimeUtil) Between(start, end time.Time) bool {
	return timeUtil.After(start) && timeUtil.Before(end)
}

func (timeUtil *TimeUtil) GetTomorrowOnHour(hour int) time.Time {
	tomorrow := timeUtil.AddDate(0, 0, 1)
	y, m, d := tomorrow.Date()
	return time.Date(y, m, d, hour, 0, 0, 0, timeUtil.Location())
}

func (timeUtil *TimeUtil) GetNextWeekDay(week int) time.Time {
	t := timeUtil.StartOfDay()
	weekday := int(timeUtil.Weekday())
	if week == int(time.Sunday) {
		week = 7
	}
	if week < weekday {
		week += 7
	}
	weekday = week - weekday
	return t.AddDate(0, 0, weekday)
}
