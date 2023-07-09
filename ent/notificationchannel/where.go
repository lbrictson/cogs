// Code generated by ent, DO NOT EDIT.

package notificationchannel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/cogs/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldType, v))
}

// Enabled applies equality check predicate on the "enabled" field. It's identical to EnabledEQ.
func Enabled(v bool) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldEnabled, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldContainsFold(FieldType, v))
}

// SlackConfigIsNil applies the IsNil predicate on the "slack_config" field.
func SlackConfigIsNil() predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldIsNull(FieldSlackConfig))
}

// SlackConfigNotNil applies the NotNil predicate on the "slack_config" field.
func SlackConfigNotNil() predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNotNull(FieldSlackConfig))
}

// EmailConfigIsNil applies the IsNil predicate on the "email_config" field.
func EmailConfigIsNil() predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldIsNull(FieldEmailConfig))
}

// EmailConfigNotNil applies the NotNil predicate on the "email_config" field.
func EmailConfigNotNil() predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNotNull(FieldEmailConfig))
}

// WebhookConfigIsNil applies the IsNil predicate on the "webhook_config" field.
func WebhookConfigIsNil() predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldIsNull(FieldWebhookConfig))
}

// WebhookConfigNotNil applies the NotNil predicate on the "webhook_config" field.
func WebhookConfigNotNil() predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNotNull(FieldWebhookConfig))
}

// EnabledEQ applies the EQ predicate on the "enabled" field.
func EnabledEQ(v bool) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldEQ(FieldEnabled, v))
}

// EnabledNEQ applies the NEQ predicate on the "enabled" field.
func EnabledNEQ(v bool) predicate.NotificationChannel {
	return predicate.NotificationChannel(sql.FieldNEQ(FieldEnabled, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotificationChannel) predicate.NotificationChannel {
	return predicate.NotificationChannel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotificationChannel) predicate.NotificationChannel {
	return predicate.NotificationChannel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NotificationChannel) predicate.NotificationChannel {
	return predicate.NotificationChannel(func(s *sql.Selector) {
		p(s.Not())
	})
}
