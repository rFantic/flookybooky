package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.Time("departure_time"),
		field.Time("arrival_time"),
		field.Int("available_slots"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("seats", Seat.Type),
		edge.From("origin", Airport.Type).Ref("origin").Unique(),
		edge.From("destination", Airport.Type).Ref("destination").Unique(),
	}
}
