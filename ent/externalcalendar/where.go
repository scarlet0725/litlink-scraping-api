// Code generated by ent, DO NOT EDIT.

package externalcalendar

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/scarlet0725/prism-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldDescription, v))
}

// CalendarID applies equality check predicate on the "calendar_id" field. It's identical to CalendarIDEQ.
func CalendarID(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldCalendarID, v))
}

// SourceType applies equality check predicate on the "source_type" field. It's identical to SourceTypeEQ.
func SourceType(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldSourceType, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldUserID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldDeletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldContainsFold(FieldDescription, v))
}

// CalendarIDEQ applies the EQ predicate on the "calendar_id" field.
func CalendarIDEQ(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldCalendarID, v))
}

// CalendarIDNEQ applies the NEQ predicate on the "calendar_id" field.
func CalendarIDNEQ(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldCalendarID, v))
}

// CalendarIDIn applies the In predicate on the "calendar_id" field.
func CalendarIDIn(vs ...string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldCalendarID, vs...))
}

// CalendarIDNotIn applies the NotIn predicate on the "calendar_id" field.
func CalendarIDNotIn(vs ...string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldCalendarID, vs...))
}

// CalendarIDGT applies the GT predicate on the "calendar_id" field.
func CalendarIDGT(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldCalendarID, v))
}

// CalendarIDGTE applies the GTE predicate on the "calendar_id" field.
func CalendarIDGTE(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldCalendarID, v))
}

// CalendarIDLT applies the LT predicate on the "calendar_id" field.
func CalendarIDLT(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldCalendarID, v))
}

// CalendarIDLTE applies the LTE predicate on the "calendar_id" field.
func CalendarIDLTE(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldCalendarID, v))
}

// CalendarIDContains applies the Contains predicate on the "calendar_id" field.
func CalendarIDContains(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldContains(FieldCalendarID, v))
}

// CalendarIDHasPrefix applies the HasPrefix predicate on the "calendar_id" field.
func CalendarIDHasPrefix(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldHasPrefix(FieldCalendarID, v))
}

// CalendarIDHasSuffix applies the HasSuffix predicate on the "calendar_id" field.
func CalendarIDHasSuffix(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldHasSuffix(FieldCalendarID, v))
}

// CalendarIDEqualFold applies the EqualFold predicate on the "calendar_id" field.
func CalendarIDEqualFold(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEqualFold(FieldCalendarID, v))
}

// CalendarIDContainsFold applies the ContainsFold predicate on the "calendar_id" field.
func CalendarIDContainsFold(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldContainsFold(FieldCalendarID, v))
}

// SourceTypeEQ applies the EQ predicate on the "source_type" field.
func SourceTypeEQ(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldSourceType, v))
}

// SourceTypeNEQ applies the NEQ predicate on the "source_type" field.
func SourceTypeNEQ(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldSourceType, v))
}

// SourceTypeIn applies the In predicate on the "source_type" field.
func SourceTypeIn(vs ...string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldSourceType, vs...))
}

// SourceTypeNotIn applies the NotIn predicate on the "source_type" field.
func SourceTypeNotIn(vs ...string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldSourceType, vs...))
}

// SourceTypeGT applies the GT predicate on the "source_type" field.
func SourceTypeGT(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldSourceType, v))
}

// SourceTypeGTE applies the GTE predicate on the "source_type" field.
func SourceTypeGTE(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldSourceType, v))
}

// SourceTypeLT applies the LT predicate on the "source_type" field.
func SourceTypeLT(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldSourceType, v))
}

// SourceTypeLTE applies the LTE predicate on the "source_type" field.
func SourceTypeLTE(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldSourceType, v))
}

// SourceTypeContains applies the Contains predicate on the "source_type" field.
func SourceTypeContains(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldContains(FieldSourceType, v))
}

// SourceTypeHasPrefix applies the HasPrefix predicate on the "source_type" field.
func SourceTypeHasPrefix(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldHasPrefix(FieldSourceType, v))
}

// SourceTypeHasSuffix applies the HasSuffix predicate on the "source_type" field.
func SourceTypeHasSuffix(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldHasSuffix(FieldSourceType, v))
}

// SourceTypeEqualFold applies the EqualFold predicate on the "source_type" field.
func SourceTypeEqualFold(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEqualFold(FieldSourceType, v))
}

// SourceTypeContainsFold applies the ContainsFold predicate on the "source_type" field.
func SourceTypeContainsFold(v string) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldContainsFold(FieldSourceType, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldUserID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.ExternalCalendar {
	return predicate.ExternalCalendar(sql.FieldNotNull(FieldDeletedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExternalCalendar) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExternalCalendar) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(func(s *sql.Selector) {
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
func Not(p predicate.ExternalCalendar) predicate.ExternalCalendar {
	return predicate.ExternalCalendar(func(s *sql.Selector) {
		p(s.Not())
	})
}
