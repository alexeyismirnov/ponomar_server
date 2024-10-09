package api

import "time"

func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func LeapStart(year int) time.Time {
	return time.Date(year, time.Month(2), 29, 0, 0, 0, 0, time.UTC)
}

func LeapEnd(year int) time.Time {
	return time.Date(year, time.Month(3), 13, 0, 0, 0, 0, time.UTC)
}

func DateTime(day int, month int, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
