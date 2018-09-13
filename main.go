package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	wykopAPI = "https://a2.wykop.pl/"
)

func Newest(c *gin.Context) {
	c.JSON(http.StatusOK, wykopAPI)
}

func main() {
	router := gin.Default()

	router.GET("/newest", Newest)
	log.Fatal(router.Run("localhost:9001"))
}
