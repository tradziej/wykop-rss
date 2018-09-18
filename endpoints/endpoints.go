package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tradziej/wykop-rss/api"
	"github.com/tradziej/wykop-rss/config"
	"github.com/tradziej/wykop-rss/rss"
	"github.com/tradziej/wykop-rss/utils"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "wykop-rss"})
}

func Promoted(c *gin.Context) {
	params := rss.Params{
		ChannelTitle: "Wykop.pl - Strona Główna",
	}
	handler(c, &params)
}

func Upcoming(c *gin.Context) {
	params := rss.Params{
		ChannelTitle: "Wykop.pl - Wykopalisko",
	}
	handler(c, &params)
}

func Popular(c *gin.Context) {
	params := rss.Params{
		ChannelTitle: "Wykop.pl - Ostatnio Popularne",
	}
	handler(c, &params)
}

func handler(c *gin.Context, p *rss.Params) {
	requestURL := c.Request.URL.String()
	p.AtomLink = config.Get().AppURL + requestURL

	links, err := api.GetLinks(requestURL[1:])

	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	if len(links.Data) > 0 {
		link := links.Data[0]

		c.Writer.Header().Set("Last-Modified", utils.StringToDate("http", link.GetCreatedAt()))
	}

	rss := rss.Generate(links, p)

	c.XML(http.StatusOK, rss)
}
