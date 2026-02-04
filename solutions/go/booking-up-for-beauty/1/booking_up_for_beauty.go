package booking

import (
	"time"
)

func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	return parsed
}

func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	return time.Now().After(parsed)
}

func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	return parsed.Hour() >= 12 && parsed.Hour() < 18
}

func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	formatted := parsed.Format("Monday, January 2, 2006, at 15:04")
	return "You have an appointment on " + formatted + "."
}

func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
