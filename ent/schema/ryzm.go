package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type RyzmEvent struct {
	ent.Schema
}

func (RyzmEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "ryzm_events",
			Charset: "utf8mb4",
		},
	}
}

func (RyzmEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique().Immutable(),
	}
}

func (RyzmEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("related_ryzm_events").
			Unique().
			Required(),
	}
}

type UnStructuredEventInformation struct {
	ent.Schema
}

func (UnStructuredEventInformation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "un_structured_event_informations",
			Charset: "utf8mb4",
		},
	}
}

func (UnStructuredEventInformation) Fields() []ent.Field {
	return []ent.Field{
		field.String("ryzmuuid"),
		field.String("venue_name"),
		field.String("artist_name"),
		field.String("price"),
	}
}

func (UnStructuredEventInformation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("un_structured_event_informations").
			Unique().
			Required(),
	}
}
