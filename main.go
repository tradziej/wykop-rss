package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tradziej/wykop-rss/api"
	"github.com/tradziej/wykop-rss/rss"
	"github.com/tradziej/wykop-rss/utils"
)

func newest(c *gin.Context) {
	links, err := api.GetLinks()

	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	if len(links.Data) > 0 {
		link := links.Data[0]

		c.Writer.Header().Set("Last-Modified", utils.StringToDate("http", link.GetCreatedAt()))
	}

	rss := rss.Generate(links)

	c.XML(http.StatusOK, rss)
}

func main() {
	router := gin.Default()

	router.GET("/newest", newest)
	log.Fatal(router.Run("localhost:9001"))
}
