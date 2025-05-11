package utils

import (
	"time"
)

func AddHour(h int) string {
	// h = 24

	now_time := time.Now()

	duration_hour := time.Duration(h) * time.Hour

	res := now_time.Add(duration_hour).Format(time.DateTime)

	return res
}
