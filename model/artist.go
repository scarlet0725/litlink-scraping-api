package model

type Artist struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	RyzmHost       string `json:"host"`
	CrawlTargetURL string `json:"-"`
	CrawlSiteType  string `json:"-"`
}
