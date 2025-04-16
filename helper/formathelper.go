package helper

import "time"


var TIMESTAMP_LAYOUT = "2006-01-02 15:04:05"

func DateToString(date time.Time) string {
	return date.Format(TIMESTAMP_LAYOUT)
}

func StringToDate(date string) (time.Time, error) {
	return time.Parse(TIMESTAMP_LAYOUT, date)
}