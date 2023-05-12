package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Seat holds the schema definition for the Seat entity.
type Seat struct {
	ent.Schema
}

// Fields of the Seat.
func (Seat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("seat_number"),
	}
}

// Edges of the Seat.
func (Seat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flight", Flight.Type).Ref("seats").Unique(),
	}
}
