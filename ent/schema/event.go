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

type Event struct {
	ent.Schema
}

func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "events",
			Charset: "utf8mb4",
		},
	}
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_id").
			Unique(),
		index.Fields("date", "open_time", "start_time", "end_time"), //TODO: ElasticSearchに切り替えたらインデックスを外す
	}
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id").
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
		field.Time("date").
			Optional().
			Nillable().
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime(3)",
				},
			),
		field.Time("open_time").
			Optional().
			Nillable().
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime(3)",
				},
			),
		field.Time("start_time").
			Optional().
			Nillable().
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime(3)",
				},
			),
		field.Time("end_time").
			Optional().
			Nillable().
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime(3)",
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
		field.String("url").
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
		field.String("ticket_url").
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

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("events"),
		edge.To("artists", Artist.Type),
	}
}
