package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Access holds the schema definition for the Access entity.
type Access struct {
	ent.Schema
}

func (Access) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Access.
func (Access) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("project_id"),
		field.Enum("role").Values("admin", "user").Default("user"),
	}
}

// Edges of the Access.
func (Access) Edges() []ent.Edge {
	return nil
}
