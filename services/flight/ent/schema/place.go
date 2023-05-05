package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Place holds the schema definition for the Place entity.
type Place struct {
	ent.Schema
}

// Fields of the Place.
func (Place) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return nil
}
