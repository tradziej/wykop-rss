package utils

import (
	"net/http"
	"time"
)

func StringToDate(fmt string, t time.Time) string {
	loc, err := time.LoadLocation("Europe/Warsaw")
	if err == nil {
		t = t.In(loc)
	}

	switch fmt {
	case "rss":
		return t.Format(time.RFC1123Z)
	case "http":
		return t.Format(http.TimeFormat)
	default:
		return t.Format(time.RFC1123Z)
	}
}
