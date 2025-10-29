package utils

import "time"

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func GetDaysInMonth(month time.Month, year int) int {
	daysInMonth := map[time.Month]int{
		time.January:   31,
		time.February:  28,
		time.March:     31,
		time.April:     30,
		time.May:       31,
		time.June:      30,
		time.July:      31,
		time.August:    31,
		time.September: 30,
		time.October:   31,
		time.November:  30,
		time.December:  31,
	}
	days := daysInMonth[month]
	if month == time.February && IsLeapYear(year) {
		days = 29
	}
	return days
}

// GetWeekOfMonth gets the week in a month (0 to 4)
func GetWeekOfMonth(date time.Time) int {
	dayOfMonth := date.Day()
	if dayOfMonth >= 1 && dayOfMonth <= 7 {
		return 0
	} else if dayOfMonth >= 8 && dayOfMonth <= 14 {
		return 1
	} else if dayOfMonth >= 15 && dayOfMonth <= 21 {
		return 2
	} else if dayOfMonth >= 22 && dayOfMonth <= 28 {
		return 3
	} else if dayOfMonth >= 29 && dayOfMonth <= 31 {
		return 4
	} else {
		return -1 // should never happen
	}
}

func CombineDateAndTime(date time.Time, timeOfDay time.Time) time.Time {
	return time.Date(
		date.Year(), date.Month(), date.Day(),
		timeOfDay.Hour(), timeOfDay.Minute(), timeOfDay.Second(), timeOfDay.Nanosecond(),
		date.Location(),
	)
}
