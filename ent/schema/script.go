package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Script holds the schema definition for the Script entity.
type Script struct {
	ent.Schema
}

func (Script) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

type ScriptInputOptions struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	StrictOptions bool     `json:"strict_options"`
	Options       []string `json:"options,omitempty"`
	VariableType  string   `json:"-"`
}

// Fields of the Script.
func (Script) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.String("script"),
		field.Int("timeout_seconds").Default(300),
		field.Int("project_id"),
		field.JSON("parameters", []ScriptInputOptions{}).Optional(),
	}
}

// Edges of the Script.
func (Script) Edges() []ent.Edge {
	return nil
}
