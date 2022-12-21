package model

type Artist struct {
	ID              uint             `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	ArtistID        string           `json:"artist_id" gorm:"unique;not null"`
	Name            string           `json:"name" gorm:"not null"`
	URL             string           `json:"url"`
	Events          []Event          `json:"events,omitempty" gorm:"many2many:events_artists"`
	RyzmCrawlConfig *RyzmCrawlConfig `json:"-" gorm:"foreignkey:ArtistID"`
}

type RyzmCrawlConfig struct {
	ID             uint `gorm:"primary_key;unique;not null;auto_increment"`
	ArtistID       *uint
	RyzmHost       string `json:"-"`
	CrawlTargetURL string `json:"-"`
	CrawlSiteType  string `json:"-"`
	Artist         *Artist
}
