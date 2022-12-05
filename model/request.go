package model

import "time"

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}

type CreateArtist struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CreateEvent struct {
	Name        string     `json:"name" binding:"required"`
	Date        string     `json:"date"`
	Description string     `json:"description"`
	OpenTime    *time.Time `json:"open_time"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	UUID        string     `json:"uuid"`
	VenueName   string     `json:"venue_name"`
	Url         string     `json:"url"`
	TicketURL   string     `json:"ticket_url"`
	ArtistName  string     `json:"artist_name"`
	ArtistIDs   []string   `json:"artist_ids"`
	VenueID     string     `json:"venue_id"`
}

type GetEvent struct {
	EventID string `json:"event_id" form:"event_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateVenue struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	WebSite     string `json:"website"`
	Postcode    string `json:"post_code"`
	Prefecture  string `json:"prefecture"`
	City        string `json:"city"`
	Street      string `json:"street"`
}

type CrawlerRequest struct {
	ArtistID   string `json:"artist_id"`
	ArtistName string `json:"artist_name"`
}

type MergeEvent struct {
	EventID            string `json:"event_id"`
	MergeTargetEventID string `json:"merge_target_id"`
}

type UpdateEvent struct {
	EventID     string     `json:"event_id"`
	Name        string     `json:"name,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	Description string     `json:"description,omitempty"`
	OpenTime    *time.Time `json:"open_time,omitempty"`
	StartTime   *time.Time `json:"start_time,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	VenueID     string     `json:"venue_id,omitempty"`
	Url         string     `json:"url,omitempty"`
	TicketURL   string     `json:"ticket_url,omitempty"`
	ArtistName  string     `json:"artist_name,omitempty"`
	//ArtistIDs   []string   `json:"artist_ids"`
}

type AdminVerify struct {
	UserID string `json:"user_id"`
}

type GoogleOauthCallback struct {
	Code  string `json:"code" form:"code"`
	State string `json:"state" form:"state"`
}

