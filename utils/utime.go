package utils

import "time"

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentDate() string {
	return time.Now().Format("20060102")
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
