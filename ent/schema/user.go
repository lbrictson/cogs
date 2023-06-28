package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("hashed_password"),
		field.String("api_key"),
		field.Enum("role").Values("admin", "user").Default("user"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
