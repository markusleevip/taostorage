package utils

import "time"

const (
	date        = "200601"
)

func GetDateYYYYMM(t time.Time) string {
	return t.Format(date)
}
