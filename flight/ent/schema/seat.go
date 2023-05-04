package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Seat holds the schema definition for the Seat entity.
type Seat struct {
	ent.Schema
}

// Fields of the Seat.
func (Seat) Fields() []ent.Field {
	return []ent.Field{
		field.String("flight_id"),
		field.String("seat_number"),
	}
}

// Edges of the Seat.
func (Seat) Edges() []ent.Edge {
	return nil
}
