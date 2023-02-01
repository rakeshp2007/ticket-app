package commonfunctions

import (
	"time"
)

func ConvertToUtcDateTime(dateTime string, timeZone string, format string) string {
	loc, _ := time.LoadLocation(timeZone)
	t, _ := time.ParseInLocation(format, dateTime, loc)
	return (t.UTC().Format(format))
}

func ConvertUtcDateTime(dateTime string, timeZone string, format string) string {
	t, _ := time.Parse(format, dateTime)
	loc, _ := time.LoadLocation(timeZone)
	t = t.In(loc)
	return (t.Format(format))
}
