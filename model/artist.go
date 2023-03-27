package model

type Artist struct {
	ID              uint             `json:"-"`
	ArtistID        string           `json:"artist_id"`
	Name            string           `json:"name"`
	URL             string           `json:"url"`
	Events          []Event          `json:"events,omitempty"`
	RyzmCrawlConfig *RyzmCrawlConfig `json:"-"`
}

type RyzmCrawlConfig struct {
	ID             uint
	ArtistID       *uint
	RyzmHost       string
	CrawlTargetURL string
	CrawlSiteType  string
	Artist         *Artist
}
