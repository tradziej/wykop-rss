package endpoints

import (
	"encoding/xml"
	"html/template"
	"net/http"
	"path"

	"github.com/tradziej/wykop-rss/api"
	"github.com/tradziej/wykop-rss/config"
	"github.com/tradziej/wykop-rss/rss"
	"github.com/tradziej/wykop-rss/utils"
)

type page struct {
	Title string
}

// Index renders index.html page.
func Index(w http.ResponseWriter, r *http.Request) {
	p := page{Title: "wykop-rss"}
	fp := path.Join("html", "index.html")

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Promoted renders links from promoted endpoint.
func Promoted(w http.ResponseWriter, r *http.Request) {
	params := rss.Params{
		ChannelTitle: "Wykop.pl - Strona Główna",
	}
	handler(w, r, &params)
}

// Upcoming renders links from upcoming endpoint.
func Upcoming(w http.ResponseWriter, r *http.Request) {
	params := rss.Params{
		ChannelTitle: "Wykop.pl - Wykopalisko",
	}
	handler(w, r, &params)
}

// Popular renders links from popular endpoint.
func Popular(w http.ResponseWriter, r *http.Request) {
	params := rss.Params{
		ChannelTitle: "Wykop.pl - Ostatnio Popularne",
	}
	handler(w, r, &params)
}

func handler(w http.ResponseWriter, r *http.Request, p *rss.Params) {
	requestURL := r.URL.String()
	p.AtomLink = config.Get().AppURL + requestURL

	links, err := api.GetLinks(requestURL[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	if len(links.Data) > 0 {
		link := links.Data[0]
		w.Header().Set("Last-Modified", utils.StringToDate("http", link.GetCreatedAt()))
	}

	rss := rss.Generate(links, p)

	x, err := xml.MarshalIndent(rss, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
