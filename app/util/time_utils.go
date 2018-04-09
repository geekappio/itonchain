package util

import "time"

const (
	DEFAULT_TIME_FORMAT = "2006-01-02 03:04:05"
)

func TimeFormat(t time.Time) string {
	return t.Format(string(DEFAULT_TIME_FORMAT))
}
