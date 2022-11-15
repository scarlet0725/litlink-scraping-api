package model

import (
	"time"
)

type RyzmAPIResponse struct {
	Data  []RyzmLiveData `json:"data"`
	Links struct {
		First string      `json:"first"`
		Last  string      `json:"last"`
		Prev  interface{} `json:"prev"`
		Next  interface{} `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int `json:"current_page"`
		From        int `json:"from"`
		LastPage    int `json:"last_page"`
		Links       []struct {
			URL    interface{} `json:"url"`
			Label  string      `json:"label"`
			Active bool        `json:"active"`
		} `json:"links"`
		Path    string `json:"path"`
		PerPage int    `json:"per_page"`
		To      int    `json:"to"`
		Total   int    `json:"total"`
	} `json:"meta"`
}

type RyzmAPI struct {
}

type RyzmLiveData struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	EventDate       string `json:"event_date"`
	EventDateStatus string `json:"event_date_status"`
	CoverImage      struct {
		ID       string      `json:"id"`
		URL      string      `json:"url"`
		MimeType string      `json:"mime_type"`
		FileName string      `json:"file_name"`
		Width    string      `json:"width"`
		Height   string      `json:"height"`
		AltText  interface{} `json:"alt_text"`
		Title    interface{} `json:"title"`
	} `json:"cover_image"`
	Category struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
		Position int    `json:"position"`
	} `json:"category"`
	Venue              string `json:"venue"`
	Title              string `json:"title"`
	Artist             string `json:"artist"`
	DoorsStartsTime    string `json:"doors_starts_time"`
	Price              string `json:"price"`
	ReservationSetting struct {
		TicketReservationType              string      `json:"ticket_reservation_type"`
		WebReservationMaxQuantity          interface{} `json:"web_reservation_max_quantity"`
		WebReservationMaxQuantityPerPerson interface{} `json:"web_reservation_max_quantity_per_person"`
		Platforms                          []struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"platforms"`
	} `json:"reservation_setting"`
	Body        string    `json:"body"`
	PublishesAt time.Time `json:"publishes_at"`
	Archived    int       `json:"archived"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
