package utils

import (
	"net/http"
	"time"
)

func StringToDate(fmt string, input time.Time) string {
	switch fmt {
	case "rss":
		return input.Format(time.RFC1123Z)
	case "http":
		return input.Format(http.TimeFormat)
	default:
		return input.Format(time.RFC1123Z)
	}
}
