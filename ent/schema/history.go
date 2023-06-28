package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

func (History) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.String("run_id").Unique(),
		field.Bool("success"),
		field.Int("exit_code"),
		field.Int("duration"),
		field.String("trigger"),
		field.String("output"),
		field.String("triggered_by"),
		field.Int("script_id"),
		field.JSON("arguments", map[string]string{}),
		field.String("status").Default("running"),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return nil
}
