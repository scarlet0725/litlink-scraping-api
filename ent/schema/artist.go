package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Artist struct {
	ent.Schema
}

func (Artist) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "artists",
			Charset: "utf8mb4",
		},
	}
}

func (Artist) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("artist_id").
			Unique(),
		index.Fields("deleted_at"),
	}
}

func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.String("artist_id").
			Unique().
			NotEmpty().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_0900_ai_ci",
				},
			).
			SchemaType(
				map[string]string{
					dialect.MySQL: "varchar(191)",
				},
			),
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
		field.String("url").
			Optional().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_0900_ai_ci",
				},
			).
			SchemaType(
				map[string]string{
					dialect.MySQL: "longtext",
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

func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).
			Ref("artists"),
	}
}
