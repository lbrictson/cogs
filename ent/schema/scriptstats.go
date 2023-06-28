package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ScriptStats holds the schema definition for the ScriptStats entity.
type ScriptStats struct {
	ent.Schema
}

// Fields of the ScriptStats.
func (ScriptStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("script_id"),
		field.Int("project_id"),
		field.Int("total_runs"),
		field.Int("total_errors"),
		field.Int("total_success"),
		field.Int("average_runtime"),
		field.Int("min_runtime"),
		field.Int("max_runtime"),
		field.Time("last_run"),
		field.Int("total_runtime"),
		field.Int("success_rate"),
	}
}

// Edges of the ScriptStats.
func (ScriptStats) Edges() []ent.Edge {
	return nil
}
