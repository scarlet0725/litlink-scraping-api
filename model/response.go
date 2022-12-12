package model

type UserResponse struct {
	OK   bool   `json:"ok"`
	User *User  `json:"user,omitempty"`
	Err  string `json:"err,omitempty"`
}

type EventResponse struct {
}

type RegistrationEventResponse struct {
	CalenderAdded   bool   `json:"calender_added"`
	EventRegistered bool   `json:"event_registered"`
	EventID         string `json:"event_id,omitempty"`
}
