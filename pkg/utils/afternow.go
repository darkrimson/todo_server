package utils

import "time"

func AfterNow(now, date time.Time) bool {
	return date.After(now)
}
