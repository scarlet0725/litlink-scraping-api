// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/scarlet0725/prism-api/ent/externalcalendar"
)

// ExternalCalendar is the model entity for the ExternalCalendar schema.
type ExternalCalendar struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CalendarID holds the value of the "calendar_id" field.
	CalendarID string `json:"calendar_id,omitempty"`
	// SourceType holds the value of the "source_type" field.
	SourceType string `json:"source_type,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExternalCalendar) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case externalcalendar.FieldID, externalcalendar.FieldUserID:
			values[i] = new(sql.NullInt64)
		case externalcalendar.FieldName, externalcalendar.FieldDescription, externalcalendar.FieldCalendarID, externalcalendar.FieldSourceType:
			values[i] = new(sql.NullString)
		case externalcalendar.FieldCreatedAt, externalcalendar.FieldUpdatedAt, externalcalendar.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ExternalCalendar", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExternalCalendar fields.
func (ec *ExternalCalendar) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case externalcalendar.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int(value.Int64)
		case externalcalendar.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ec.Name = value.String
			}
		case externalcalendar.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ec.Description = value.String
			}
		case externalcalendar.FieldCalendarID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field calendar_id", values[i])
			} else if value.Valid {
				ec.CalendarID = value.String
			}
		case externalcalendar.FieldSourceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_type", values[i])
			} else if value.Valid {
				ec.SourceType = value.String
			}
		case externalcalendar.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ec.UserID = int(value.Int64)
			}
		case externalcalendar.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ec.CreatedAt = value.Time
			}
		case externalcalendar.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ec.UpdatedAt = value.Time
			}
		case externalcalendar.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ec.DeletedAt = new(time.Time)
				*ec.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ExternalCalendar.
// Note that you need to call ExternalCalendar.Unwrap() before calling this method if this ExternalCalendar
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *ExternalCalendar) Update() *ExternalCalendarUpdateOne {
	return (&ExternalCalendarClient{config: ec.config}).UpdateOne(ec)
}

// Unwrap unwraps the ExternalCalendar entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *ExternalCalendar) Unwrap() *ExternalCalendar {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExternalCalendar is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *ExternalCalendar) String() string {
	var builder strings.Builder
	builder.WriteString("ExternalCalendar(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("name=")
	builder.WriteString(ec.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ec.Description)
	builder.WriteString(", ")
	builder.WriteString("calendar_id=")
	builder.WriteString(ec.CalendarID)
	builder.WriteString(", ")
	builder.WriteString("source_type=")
	builder.WriteString(ec.SourceType)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.UserID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ec.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ec.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ec.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// ExternalCalendars is a parsable slice of ExternalCalendar.
type ExternalCalendars []*ExternalCalendar

func (ec ExternalCalendars) config(cfg config) {
	for _i := range ec {
		ec[_i].config = cfg
	}
}
