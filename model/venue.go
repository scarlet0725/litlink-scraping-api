package model

type Venue struct {
	ID          int      `json:"-" gorm:"primary_key;unique;not null"`
	VenueID     string   `json:"venue_id" gorm:"unique;not null"`
	Name        string   `json:"name" gorm:"not null"`
	Description string   `json:"description"`
	WebSite     string   `json:"website"`
	Postcode    string   `json:"post_code"`
	Prefecture  string   `json:"prefecture"`
	City        string   `json:"city"`
	Street      string   `json:"street"`
	IsOpen      bool     `json:"is_open" gorm:"not null;default:true"`
	Events      []*Event `json:"-"  gorm:"foreignkey:VenueID"`
}
