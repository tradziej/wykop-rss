package rss

import (
	"encoding/xml"
	"time"

	"github.com/tradziej/wykop-rss/api"
	"github.com/tradziej/wykop-rss/utils"
)

const (
	generatorURL = "https://github.com/tradziej/wykop-rss"
	wykopURL     = "https://www.wykop.pl"
)

type (
	rss struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Channel channel  `xml:"channel"`
	}

	channel struct {
		XMLName       xml.Name `xml:"channel"`
		Generator     string   `xml:"generator"`
		Docs          string   `xml:"docs"`
		Title         string   `xml:"title"`
		Description   string   `xml:"description"`
		Link          string   `xml:"link"`
		LastBuildDate string   `xml:"lastBuildDate"`
		Items         []item   `xml:"item"`
	}

	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        guid     `xml:"guid"`
		Comments    string   `xml:"comments"`
	}

	guid struct {
		Value       string `xml:",chardata"`
		IsPermaLink string `xml:"isPermaLink,attr"`
	}
)

func Generate(links *api.Links) *rss {
	items := []item{}

	for _, link := range links.Data {
		item := item{
			PubDate:     utils.StringToDate("rss", link.GetCreatedAt()),
			Title:       link.Title,
			Description: link.Description,
			Link:        link.URL,
			GUID:        guid{link.GetGUID(), "false"},
			Comments:    link.GetGUID(),
		}
		items = append(items, item)
	}

	rss := rss{
		Version: "2.0",
		Channel: channel{
			Generator:     "wykop-rss",
			Docs:          generatorURL,
			Title:         "Wykop.pl - Strona Główna",
			Description:   "Wykop - serwis tworzony przez użykowników",
			Link:          wykopURL,
			LastBuildDate: utils.StringToDate("rss", time.Now().UTC()),
			Items:         items,
		},
	}

	return &rss
}
