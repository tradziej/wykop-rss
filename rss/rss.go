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
	// RSS contains the fields associated with the rss tag in the rss document.
	RSS struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		NS      string   `xml:"xmlns:atom,attr"`
		Channel channel  `xml:"channel"`
	}

	channel struct {
		XMLName       xml.Name `xml:"channel"`
		Generator     string   `xml:"generator"`
		Docs          string   `xml:"docs"`
		Title         string   `xml:"title"`
		Description   string   `xml:"description"`
		Link          string   `xml:"link"`
		AtomLink      atomLink `xml:"atom:link"`
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

	atomLink struct {
		Reference    string `xml:"href,attr"`
		Relationship string `xml:"rel,attr,omitempty"`
		Type         string `xml:"type,attr,omitempty"`
	}

	// Params contains RSS feed parameters.
	Params struct {
		AtomLink     string
		ChannelTitle string
	}
)

// Generate generate RSS feed for given links and feed params.
func Generate(links *api.Links, p *Params) *RSS {
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

	rss := RSS{
		Version: "2.0",
		NS:      "http://www.w3.org/2005/Atom",
		Channel: channel{
			Generator:     "wykop-rss",
			Docs:          generatorURL,
			Title:         p.ChannelTitle,
			Description:   "Wykop - serwis tworzony przez użykowników",
			Link:          wykopURL,
			AtomLink:      atomLink{p.AtomLink, "self", "application/rss+xml"},
			LastBuildDate: utils.StringToDate("rss", time.Now().UTC()),
			Items:         items,
		},
	}

	return &rss
}
