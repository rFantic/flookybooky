package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("customer_id").NotEmpty().Unique().Immutable().Optional(),
		field.String("username").Immutable().Unique().NotEmpty(),
		field.String("password").NotEmpty().Sensitive(),
		field.String("email").NotEmpty().Unique(),
		field.Enum("role").Values("admin", "user"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
