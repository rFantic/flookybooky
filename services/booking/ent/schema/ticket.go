package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("going_flight_id", uuid.UUID{}),
		field.UUID("return_flight_id", uuid.UUID{}),
		field.Enum("status").Values(
			"Canceled", "Departed", "Scheduled",
		),
		field.String("passenger_name"),
		field.String("passenger_license_id"),
		field.String("passenger_email"),
		field.String("seat_number"),
		field.Enum("class").Values(
			"FirstClass", "BusinessClass", "EconomyClass",
		),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("booking", Booking.Type).Ref("ticket"),
	}
}
