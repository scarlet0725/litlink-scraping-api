package model

type Venue struct {
	ID          uint     `json:"-"`
	VenueID     string   `json:"venue_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	WebSite     string   `json:"website"`
	Postcode    string   `json:"post_code"`
	Prefecture  string   `json:"prefecture"`
	City        string   `json:"city"`
	Street      string   `json:"street"`
	IsOpen      bool     `json:"is_open"`
	Events      []*Event `json:"-"`
}
