package model

type Artist struct {
	ID              int              `json:"-"`
	ArtistID        string           `json:"artist_id"`
	Name            string           `json:"name"`
	URL             string           `json:"url"`
	Events          []Event          `json:"events,omitempty"`
	RyzmCrawlConfig *RyzmCrawlConfig `json:"-"`
}

type RyzmCrawlConfig struct {
	ID             int
	ArtistID       int
	RyzmHost       string
	CrawlTargetURL string
	CrawlSiteType  string
	Artist         *Artist
}
