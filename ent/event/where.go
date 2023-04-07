// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/scarlet0725/prism-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldName, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDate, v))
}

// OpenTime applies equality check predicate on the "open_time" field. It's identical to OpenTimeEQ.
func OpenTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldOpenTime, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEndTime, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldURL, v))
}

// TicketURL applies equality check predicate on the "ticket_url" field. It's identical to TicketURLEQ.
func TicketURL(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTicketURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDeletedAt, v))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldEventID, vs...))
}

// EventIDGT applies the GT predicate on the "event_id" field.
func EventIDGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldEventID, v))
}

// EventIDGTE applies the GTE predicate on the "event_id" field.
func EventIDGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldEventID, v))
}

// EventIDLT applies the LT predicate on the "event_id" field.
func EventIDLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldEventID, v))
}

// EventIDLTE applies the LTE predicate on the "event_id" field.
func EventIDLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldEventID, v))
}

// EventIDContains applies the Contains predicate on the "event_id" field.
func EventIDContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldEventID, v))
}

// EventIDHasPrefix applies the HasPrefix predicate on the "event_id" field.
func EventIDHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldEventID, v))
}

// EventIDHasSuffix applies the HasSuffix predicate on the "event_id" field.
func EventIDHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldEventID, v))
}

// EventIDEqualFold applies the EqualFold predicate on the "event_id" field.
func EventIDEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldEventID, v))
}

// EventIDContainsFold applies the ContainsFold predicate on the "event_id" field.
func EventIDContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldEventID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldName, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDate, v))
}

// DateIsNil applies the IsNil predicate on the "date" field.
func DateIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDate))
}

// DateNotNil applies the NotNil predicate on the "date" field.
func DateNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDate))
}

// OpenTimeEQ applies the EQ predicate on the "open_time" field.
func OpenTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldOpenTime, v))
}

// OpenTimeNEQ applies the NEQ predicate on the "open_time" field.
func OpenTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldOpenTime, v))
}

// OpenTimeIn applies the In predicate on the "open_time" field.
func OpenTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldOpenTime, vs...))
}

// OpenTimeNotIn applies the NotIn predicate on the "open_time" field.
func OpenTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldOpenTime, vs...))
}

// OpenTimeGT applies the GT predicate on the "open_time" field.
func OpenTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldOpenTime, v))
}

// OpenTimeGTE applies the GTE predicate on the "open_time" field.
func OpenTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldOpenTime, v))
}

// OpenTimeLT applies the LT predicate on the "open_time" field.
func OpenTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldOpenTime, v))
}

// OpenTimeLTE applies the LTE predicate on the "open_time" field.
func OpenTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldOpenTime, v))
}

// OpenTimeIsNil applies the IsNil predicate on the "open_time" field.
func OpenTimeIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldOpenTime))
}

// OpenTimeNotNil applies the NotNil predicate on the "open_time" field.
func OpenTimeNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldOpenTime))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeIsNil applies the IsNil predicate on the "start_time" field.
func StartTimeIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldStartTime))
}

// StartTimeNotNil applies the NotNil predicate on the "start_time" field.
func StartTimeNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldStartTime))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldEndTime, v))
}

// EndTimeIsNil applies the IsNil predicate on the "end_time" field.
func EndTimeIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldEndTime))
}

// EndTimeNotNil applies the NotNil predicate on the "end_time" field.
func EndTimeNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldEndTime))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDescription, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldURL, v))
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldURL))
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldURL))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldURL, v))
}

// TicketURLEQ applies the EQ predicate on the "ticket_url" field.
func TicketURLEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTicketURL, v))
}

// TicketURLNEQ applies the NEQ predicate on the "ticket_url" field.
func TicketURLNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldTicketURL, v))
}

// TicketURLIn applies the In predicate on the "ticket_url" field.
func TicketURLIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldTicketURL, vs...))
}

// TicketURLNotIn applies the NotIn predicate on the "ticket_url" field.
func TicketURLNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldTicketURL, vs...))
}

// TicketURLGT applies the GT predicate on the "ticket_url" field.
func TicketURLGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldTicketURL, v))
}

// TicketURLGTE applies the GTE predicate on the "ticket_url" field.
func TicketURLGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldTicketURL, v))
}

// TicketURLLT applies the LT predicate on the "ticket_url" field.
func TicketURLLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldTicketURL, v))
}

// TicketURLLTE applies the LTE predicate on the "ticket_url" field.
func TicketURLLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldTicketURL, v))
}

// TicketURLContains applies the Contains predicate on the "ticket_url" field.
func TicketURLContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldTicketURL, v))
}

// TicketURLHasPrefix applies the HasPrefix predicate on the "ticket_url" field.
func TicketURLHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldTicketURL, v))
}

// TicketURLHasSuffix applies the HasSuffix predicate on the "ticket_url" field.
func TicketURLHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldTicketURL, v))
}

// TicketURLIsNil applies the IsNil predicate on the "ticket_url" field.
func TicketURLIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldTicketURL))
}

// TicketURLNotNil applies the NotNil predicate on the "ticket_url" field.
func TicketURLNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldTicketURL))
}

// TicketURLEqualFold applies the EqualFold predicate on the "ticket_url" field.
func TicketURLEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldTicketURL, v))
}

// TicketURLContainsFold applies the ContainsFold predicate on the "ticket_url" field.
func TicketURLContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldTicketURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDeletedAt))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArtists applies the HasEdge predicate on the "artists" edge.
func HasArtists() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ArtistsTable, ArtistsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtistsWith applies the HasEdge predicate on the "artists" edge with a given conditions (other predicates).
func HasArtistsWith(preds ...predicate.Artist) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArtistsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ArtistsTable, ArtistsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVenue applies the HasEdge predicate on the "venue" edge.
func HasVenue() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, VenueTable, VenueColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVenueWith applies the HasEdge predicate on the "venue" edge with a given conditions (other predicates).
func HasVenueWith(preds ...predicate.Venue) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VenueInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, VenueTable, VenueColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		p(s.Not())
	})
}
