package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type GoogleOauthToken struct {
	ent.Schema
}

func (GoogleOauthToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "google_oauth_tokens",
			Charset: "utf8mb4",
		},
	}
}

func (GoogleOauthToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("refresh_token").
			SchemaType(
				map[string]string{
					dialect.MySQL: "longtext",
				},
			),
		field.String("access_token").
			SchemaType(
				map[string]string{
					dialect.MySQL: "longtext",
				},
			),
		field.Time("expiry").
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime(3)",
				},
			),
	}
}

func (GoogleOauthToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("google_oauth_tokens").
			Unique().
			Required(),
	}
}
