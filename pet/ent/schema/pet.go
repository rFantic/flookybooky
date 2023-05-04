package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("species"),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return nil
}
