package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SlackConfig struct {
	WebhookURL string `json:"webhook_url"`
}

type EmailConfig struct {
	To string `json:"to"`
}

type WebhookConfig struct {
	URL string `json:"url"`
}

// NotificationChannel holds the schema definition for the NotificationChannel entity.
type NotificationChannel struct {
	ent.Schema
}

// Fields of the NotificationChannel.
func (NotificationChannel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("type"),
		field.JSON("slack_config", SlackConfig{}).Optional(),
		field.JSON("email_config", EmailConfig{}).Optional(),
		field.JSON("webhook_config", WebhookConfig{}).Optional(),
		field.Bool("enabled").Default(true),
		field.Time("last_used").Optional().Nillable(),
		field.Bool("last_used_success").Optional().Nillable(),
	}
}

func (NotificationChannel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the NotificationChannel.
func (NotificationChannel) Edges() []ent.Edge {
	return nil
}
