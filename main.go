package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tradziej/wykop-rss/endpoints"
)

func main() {
	router := gin.Default()

	router.GET("/promoted", endpoints.Promoted)

	log.Fatal(router.Run("0.0.0.0:9001"))
}
