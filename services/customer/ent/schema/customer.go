package schema

import (
	"flookybooky/internal/string"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("address"),
		field.String("license_id").
			Validate(string.IsNumeric),
		field.String("phone_number").
			Validate(string.IsNumeric),
		field.Time("timestamp").Immutable().Default(time.Now()),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
