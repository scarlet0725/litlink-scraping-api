package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type GoogleOauthState struct {
	ent.Schema
}

func (GoogleOauthState) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "google_oauth_states",
			Charset: "utf8mb4",
		},
	}
}

func (GoogleOauthState) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("state").
			Unique(),
	}
}

func (GoogleOauthState) Fields() []ent.Field {
	return []ent.Field{
		field.String("state").
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
	}
}

func (GoogleOauthState) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("google_oauth_states").
			Unique().
			Required(),
	}
}
