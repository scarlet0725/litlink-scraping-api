package model

type Artist struct {
	ID             uint    `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	ArtistID       string  `json:"artist_id" gorm:"unique;not null"`
	Name           string  `json:"name" gorm:"not null"`
	URL            string  `json:"url"`
	RyzmHost       string  `json:"-"`
	CrawlTargetURL string  `json:"-"`
	CrawlSiteType  string  `json:"-"`
	Events         []Event `json:"events" gorm:"many2many:events_artists"`
}
