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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "users",
			Charset: "utf8mb4",
		},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "username", "email").
			Unique(),
		index.Fields("deleted_at", "api_key"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			Unique().
			NotEmpty().
			SchemaType(
				map[string]string{
					dialect.MySQL: "varchar(191)",
				},
			).
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_0900_ai_ci",
				},
			),
		field.String("username").
			Unique().
			NotEmpty().
			Annotations(
				entsql.Annotation{
					Collation: "utf8mb4_0900_ai_ci",
				},
			),
		field.String("email").
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
		field.Bytes("password").
			NotEmpty().
			SchemaType(
				map[string]string{
					dialect.MySQL: "longblob",
				},
			),
		field.String("first_name").
			Default("").
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
		field.String("last_name").
			Default("").
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
		field.Bool("is_admin_verified").
			Default(false),
		field.Bool("delete_protected").
			Default(false),
		field.String("api_key").
			Default("").
			Unique().
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
			).Default(""),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime",
				},
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime",
				},
			),
		field.Time("deleted_at").
			SchemaType(
				map[string]string{
					dialect.MySQL: "datetime",
				},
			).
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("google_oauth_tokens", GoogleOauthToken.Type).
			Unique(),
		edge.To("google_oauth_states", GoogleOauthState.Type).
			Unique(),
	}
}
