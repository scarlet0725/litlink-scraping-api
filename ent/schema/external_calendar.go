package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ExternalCalendar struct {
	ent.Schema
}

func (ExternalCalendar) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "external_calendars",
			Charset: "utf8mb4",
		},
	}
}

func (ExternalCalendar) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_ja_0900_as_cs_ks",
				},
			).
			SchemaType(
				map[string]string{
					dialect.MySQL: "longtext",
				},
			),
		field.String("description").
			Optional().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_bin",
				},
			).
			SchemaType(
				map[string]string{
					dialect.MySQL: "longtext",
				},
			),
		field.String("calendar_id").
			NotEmpty().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_bin",
				},
			),
		field.String("type").
			NotEmpty().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_bin",
				},
			),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime(3)",
				},
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime(3)",
				},
			),
		field.Time("deleted_at").
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime(3)",
				},
			).
			Nillable().
			Optional(),
	}
}

func (ExternalCalendar) Edge() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("external_calendars").
			Unique().
			Required(),
	}
}
