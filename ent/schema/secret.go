package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Secret holds the schema definition for the Secret entity.
type Secret struct {
	ent.Schema
}

func (Secret) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Secret.
func (Secret) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("value"),
		field.Int("project_id"),
		field.String("created_by"),
	}
}

// Edges of the Secret.
func (Secret) Edges() []ent.Edge {
	return nil
}
