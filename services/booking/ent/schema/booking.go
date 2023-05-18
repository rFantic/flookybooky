package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("customer_id", uuid.UUID{}),
		field.UUID("going_ticket_id", uuid.UUID{}),
		field.UUID("return_ticket_id", uuid.UUID{}).Optional().Nillable(),
		field.Enum("status").Values("Canceled", "Scheduled", "Departed"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("going_ticket", Ticket.Type).Ref("going").Unique().Field("going_ticket_id").Required(),
		edge.From("return_ticket", Ticket.Type).Ref("return").Unique().Field("return_ticket_id"),
	}
}
