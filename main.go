package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tradziej/wykop-rss/endpoints"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("html/index.html")

	router.GET("/promoted", endpoints.Promoted)
	router.GET("/upcoming", endpoints.Upcoming)
	router.GET("/popular", endpoints.Popular)

	router.GET("/", endpoints.Index)

	log.Fatal(router.Run("0.0.0.0:9001"))
}
