package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/tradziej/wykop-rss/config"
)

const (
	wykopAPI    = "https://a2.wykop.pl/"
	wykoLinkURL = "https://www.wykop.pl/link/"
)

var (
	promotedEndpoint = wykopAPI + "Links/Promoted/" + "appkey/" + config.Get().WykopAppKey
	upcomingEndpoint = wykopAPI + "Links/Upcoming/" + "appkey/" + config.Get().WykopAppKey
	popularEndpoint  = wykopAPI + "Hits/Popular/" + "appkey/" + config.Get().WykopAppKey
)

var wykopAPIClient = http.Client{
	Timeout: 10 * time.Second,
}

type Links struct {
	Data []Link
}

type Link struct {
	ID            int
	Title         string
	Description   string
	Tags          string
	URL           string `json:"source_url"`
	VoteCount     int    `json:"vote_count"`
	BuryCount     int    `json:"bury_count"`
	RelatedCount  int    `json:"related_count"`
	CommentsCount int    `json:"comments_count"`
	Date          string `json:"date"`
	Author        string `json:"author:login"`
}

func (link Link) GetCreatedAt() time.Time {
	if d, err := time.Parse("2006-01-02 15:04:05", link.Date); err == nil {
		return d
	} else {
		return time.Now()
	}
}

func (link Link) GetGUID() string {
	return wykoLinkURL + strconv.Itoa(link.ID)
}

func GetLinks(endpoint string) (*Links, error) {
	var e string
	switch endpoint {
	case "promoted":
		e = promotedEndpoint
	case "upcoming":
		e = upcomingEndpoint
	case "popular":
		e = popularEndpoint
	default:
		e = promotedEndpoint
	}

	resp, err := wykopAPIClient.Get(e)

	if err != nil {
		return nil, errors.New("Error requesting Wykop.pl API")
	}
	defer resp.Body.Close()

	var parsed Links

	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, errors.New("Invalid JSON")
	}

	return &parsed, nil
}
