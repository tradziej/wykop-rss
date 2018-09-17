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

var promotedEndpoint = wykopAPI + "Links/Promoted/" + "appkey/" + config.Get().WykopAppKey

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
	loc, _ := time.LoadLocation("Europe/Warsaw")

	if d, err := time.ParseInLocation("2006-01-02 15:04:05", link.Date, loc); err != nil {
		return d
	} else {
		return time.Now().UTC()
	}
}

func (link Link) GetGUID() string {
	return wykoLinkURL + strconv.Itoa(link.ID)
}

func GetLinks() (*Links, error) {
	resp, err := wykopAPIClient.Get(promotedEndpoint)

	if err != nil {
		println(err)
		return nil, errors.New("Error requesting Wykop.pl API")
	}
	defer resp.Body.Close()

	var parsed Links
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&parsed)
	if err != nil {
		println(err)
		return nil, errors.New("Invalid JSON")
	}

	return &parsed, nil
}
