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

type Venue struct {
	ent.Schema
}

func (Venue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "venues",
			Charset: "utf8mb4",
		},
	}
}

func (Venue) Fields() []ent.Field {
	return []ent.Field{
		field.String("venue_id").
			Unique().
			NotEmpty().
			SchemaType(
				map[string]string{
					dialect.MySQL: "varchar(32)",
				},
			).
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_0900_ai_ci",
				},
			),
		field.String("name").
			NotEmpty().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_ja_0900_as_cs_ks",
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
		field.Bytes("web_site").
			Optional().
			SchemaType(
				map[string]string{
					dialect.MySQL: "longtext",
				},
			),
		field.String("postcode").
			Optional().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_ja_0900_as_cs_ks",
				},
			).
			SchemaType(
				map[string]string{
					dialect.MySQL: "varchar(8)",
				},
			),
		field.String("prefecture").
			Optional().
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
		field.String("city").
			Optional().
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
		field.String("street").
			Optional().
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
		field.Bool("is_open").
			Default(true),
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

func (Venue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).
			Ref("venue"),
	}
}
