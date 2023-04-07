// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldOpenTime holds the string denoting the open_time field in the database.
	FieldOpenTime = "open_time"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldTicketURL holds the string denoting the ticket_url field in the database.
	FieldTicketURL = "ticket_url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeArtists holds the string denoting the artists edge name in mutations.
	EdgeArtists = "artists"
	// EdgeVenue holds the string denoting the venue edge name in mutations.
	EdgeVenue = "venue"
	// Table holds the table name of the event in the database.
	Table = "events"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_events"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// ArtistsTable is the table that holds the artists relation/edge. The primary key declared below.
	ArtistsTable = "event_artists"
	// ArtistsInverseTable is the table name for the Artist entity.
	// It exists in this package in order to avoid circular dependency with the "artist" package.
	ArtistsInverseTable = "artists"
	// VenueTable is the table that holds the venue relation/edge.
	VenueTable = "events"
	// VenueInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenueInverseTable = "venues"
	// VenueColumn is the table column denoting the venue relation/edge.
	VenueColumn = "event_venue"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldEventID,
	FieldName,
	FieldDate,
	FieldOpenTime,
	FieldStartTime,
	FieldEndTime,
	FieldDescription,
	FieldURL,
	FieldTicketURL,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_venue",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "event_id"}
	// ArtistsPrimaryKey and ArtistsColumn2 are the table columns denoting the
	// primary key for the artists relation (M2M).
	ArtistsPrimaryKey = []string{"event_id", "artist_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// EventIDValidator is a validator for the "event_id" field. It is called by the builders before save.
	EventIDValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
