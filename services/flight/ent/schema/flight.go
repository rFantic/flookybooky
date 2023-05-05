package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("from_id"),
		field.String("to_id"),
		field.Time("start"),
		field.Time("end"),
		field.Int("available_slots"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("seats", Seat.Type),
	}
}
