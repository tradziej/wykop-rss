package main

import (
	"net/http"

	"github.com/tradziej/wykop-rss/endpoints"
)

func main() {
	http.HandleFunc("/", endpoints.Index)

	http.HandleFunc("/promoted", endpoints.Promoted)
	http.HandleFunc("/upcoming", endpoints.Upcoming)
	http.HandleFunc("/popular", endpoints.Popular)

	http.ListenAndServe(":9001", nil)
}
