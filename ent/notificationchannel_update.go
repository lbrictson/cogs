// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/cogs/ent/notificationchannel"
	"github.com/lbrictson/cogs/ent/predicate"
	"github.com/lbrictson/cogs/ent/schema"
)

// NotificationChannelUpdate is the builder for updating NotificationChannel entities.
type NotificationChannelUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationChannelMutation
}

// Where appends a list predicates to the NotificationChannelUpdate builder.
func (ncu *NotificationChannelUpdate) Where(ps ...predicate.NotificationChannel) *NotificationChannelUpdate {
	ncu.mutation.Where(ps...)
	return ncu
}

// SetUpdatedAt sets the "updated_at" field.
func (ncu *NotificationChannelUpdate) SetUpdatedAt(t time.Time) *NotificationChannelUpdate {
	ncu.mutation.SetUpdatedAt(t)
	return ncu
}

// SetName sets the "name" field.
func (ncu *NotificationChannelUpdate) SetName(s string) *NotificationChannelUpdate {
	ncu.mutation.SetName(s)
	return ncu
}

// SetType sets the "type" field.
func (ncu *NotificationChannelUpdate) SetType(s string) *NotificationChannelUpdate {
	ncu.mutation.SetType(s)
	return ncu
}

// SetSlackConfig sets the "slack_config" field.
func (ncu *NotificationChannelUpdate) SetSlackConfig(sc schema.SlackConfig) *NotificationChannelUpdate {
	ncu.mutation.SetSlackConfig(sc)
	return ncu
}

// SetNillableSlackConfig sets the "slack_config" field if the given value is not nil.
func (ncu *NotificationChannelUpdate) SetNillableSlackConfig(sc *schema.SlackConfig) *NotificationChannelUpdate {
	if sc != nil {
		ncu.SetSlackConfig(*sc)
	}
	return ncu
}

// ClearSlackConfig clears the value of the "slack_config" field.
func (ncu *NotificationChannelUpdate) ClearSlackConfig() *NotificationChannelUpdate {
	ncu.mutation.ClearSlackConfig()
	return ncu
}

// SetEmailConfig sets the "email_config" field.
func (ncu *NotificationChannelUpdate) SetEmailConfig(sc schema.EmailConfig) *NotificationChannelUpdate {
	ncu.mutation.SetEmailConfig(sc)
	return ncu
}

// SetNillableEmailConfig sets the "email_config" field if the given value is not nil.
func (ncu *NotificationChannelUpdate) SetNillableEmailConfig(sc *schema.EmailConfig) *NotificationChannelUpdate {
	if sc != nil {
		ncu.SetEmailConfig(*sc)
	}
	return ncu
}

// ClearEmailConfig clears the value of the "email_config" field.
func (ncu *NotificationChannelUpdate) ClearEmailConfig() *NotificationChannelUpdate {
	ncu.mutation.ClearEmailConfig()
	return ncu
}

// SetWebhookConfig sets the "webhook_config" field.
func (ncu *NotificationChannelUpdate) SetWebhookConfig(sc schema.WebhookConfig) *NotificationChannelUpdate {
	ncu.mutation.SetWebhookConfig(sc)
	return ncu
}

// SetNillableWebhookConfig sets the "webhook_config" field if the given value is not nil.
func (ncu *NotificationChannelUpdate) SetNillableWebhookConfig(sc *schema.WebhookConfig) *NotificationChannelUpdate {
	if sc != nil {
		ncu.SetWebhookConfig(*sc)
	}
	return ncu
}

// ClearWebhookConfig clears the value of the "webhook_config" field.
func (ncu *NotificationChannelUpdate) ClearWebhookConfig() *NotificationChannelUpdate {
	ncu.mutation.ClearWebhookConfig()
	return ncu
}

// SetEnabled sets the "enabled" field.
func (ncu *NotificationChannelUpdate) SetEnabled(b bool) *NotificationChannelUpdate {
	ncu.mutation.SetEnabled(b)
	return ncu
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (ncu *NotificationChannelUpdate) SetNillableEnabled(b *bool) *NotificationChannelUpdate {
	if b != nil {
		ncu.SetEnabled(*b)
	}
	return ncu
}

// Mutation returns the NotificationChannelMutation object of the builder.
func (ncu *NotificationChannelUpdate) Mutation() *NotificationChannelMutation {
	return ncu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ncu *NotificationChannelUpdate) Save(ctx context.Context) (int, error) {
	ncu.defaults()
	return withHooks(ctx, ncu.sqlSave, ncu.mutation, ncu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ncu *NotificationChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := ncu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ncu *NotificationChannelUpdate) Exec(ctx context.Context) error {
	_, err := ncu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncu *NotificationChannelUpdate) ExecX(ctx context.Context) {
	if err := ncu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncu *NotificationChannelUpdate) defaults() {
	if _, ok := ncu.mutation.UpdatedAt(); !ok {
		v := notificationchannel.UpdateDefaultUpdatedAt()
		ncu.mutation.SetUpdatedAt(v)
	}
}

func (ncu *NotificationChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(notificationchannel.Table, notificationchannel.Columns, sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeInt))
	if ps := ncu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncu.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationchannel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ncu.mutation.Name(); ok {
		_spec.SetField(notificationchannel.FieldName, field.TypeString, value)
	}
	if value, ok := ncu.mutation.GetType(); ok {
		_spec.SetField(notificationchannel.FieldType, field.TypeString, value)
	}
	if value, ok := ncu.mutation.SlackConfig(); ok {
		_spec.SetField(notificationchannel.FieldSlackConfig, field.TypeJSON, value)
	}
	if ncu.mutation.SlackConfigCleared() {
		_spec.ClearField(notificationchannel.FieldSlackConfig, field.TypeJSON)
	}
	if value, ok := ncu.mutation.EmailConfig(); ok {
		_spec.SetField(notificationchannel.FieldEmailConfig, field.TypeJSON, value)
	}
	if ncu.mutation.EmailConfigCleared() {
		_spec.ClearField(notificationchannel.FieldEmailConfig, field.TypeJSON)
	}
	if value, ok := ncu.mutation.WebhookConfig(); ok {
		_spec.SetField(notificationchannel.FieldWebhookConfig, field.TypeJSON, value)
	}
	if ncu.mutation.WebhookConfigCleared() {
		_spec.ClearField(notificationchannel.FieldWebhookConfig, field.TypeJSON)
	}
	if value, ok := ncu.mutation.Enabled(); ok {
		_spec.SetField(notificationchannel.FieldEnabled, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ncu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationchannel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ncu.mutation.done = true
	return n, nil
}

// NotificationChannelUpdateOne is the builder for updating a single NotificationChannel entity.
type NotificationChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationChannelMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ncuo *NotificationChannelUpdateOne) SetUpdatedAt(t time.Time) *NotificationChannelUpdateOne {
	ncuo.mutation.SetUpdatedAt(t)
	return ncuo
}

// SetName sets the "name" field.
func (ncuo *NotificationChannelUpdateOne) SetName(s string) *NotificationChannelUpdateOne {
	ncuo.mutation.SetName(s)
	return ncuo
}

// SetType sets the "type" field.
func (ncuo *NotificationChannelUpdateOne) SetType(s string) *NotificationChannelUpdateOne {
	ncuo.mutation.SetType(s)
	return ncuo
}

// SetSlackConfig sets the "slack_config" field.
func (ncuo *NotificationChannelUpdateOne) SetSlackConfig(sc schema.SlackConfig) *NotificationChannelUpdateOne {
	ncuo.mutation.SetSlackConfig(sc)
	return ncuo
}

// SetNillableSlackConfig sets the "slack_config" field if the given value is not nil.
func (ncuo *NotificationChannelUpdateOne) SetNillableSlackConfig(sc *schema.SlackConfig) *NotificationChannelUpdateOne {
	if sc != nil {
		ncuo.SetSlackConfig(*sc)
	}
	return ncuo
}

// ClearSlackConfig clears the value of the "slack_config" field.
func (ncuo *NotificationChannelUpdateOne) ClearSlackConfig() *NotificationChannelUpdateOne {
	ncuo.mutation.ClearSlackConfig()
	return ncuo
}

// SetEmailConfig sets the "email_config" field.
func (ncuo *NotificationChannelUpdateOne) SetEmailConfig(sc schema.EmailConfig) *NotificationChannelUpdateOne {
	ncuo.mutation.SetEmailConfig(sc)
	return ncuo
}

// SetNillableEmailConfig sets the "email_config" field if the given value is not nil.
func (ncuo *NotificationChannelUpdateOne) SetNillableEmailConfig(sc *schema.EmailConfig) *NotificationChannelUpdateOne {
	if sc != nil {
		ncuo.SetEmailConfig(*sc)
	}
	return ncuo
}

// ClearEmailConfig clears the value of the "email_config" field.
func (ncuo *NotificationChannelUpdateOne) ClearEmailConfig() *NotificationChannelUpdateOne {
	ncuo.mutation.ClearEmailConfig()
	return ncuo
}

// SetWebhookConfig sets the "webhook_config" field.
func (ncuo *NotificationChannelUpdateOne) SetWebhookConfig(sc schema.WebhookConfig) *NotificationChannelUpdateOne {
	ncuo.mutation.SetWebhookConfig(sc)
	return ncuo
}

// SetNillableWebhookConfig sets the "webhook_config" field if the given value is not nil.
func (ncuo *NotificationChannelUpdateOne) SetNillableWebhookConfig(sc *schema.WebhookConfig) *NotificationChannelUpdateOne {
	if sc != nil {
		ncuo.SetWebhookConfig(*sc)
	}
	return ncuo
}

// ClearWebhookConfig clears the value of the "webhook_config" field.
func (ncuo *NotificationChannelUpdateOne) ClearWebhookConfig() *NotificationChannelUpdateOne {
	ncuo.mutation.ClearWebhookConfig()
	return ncuo
}

// SetEnabled sets the "enabled" field.
func (ncuo *NotificationChannelUpdateOne) SetEnabled(b bool) *NotificationChannelUpdateOne {
	ncuo.mutation.SetEnabled(b)
	return ncuo
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (ncuo *NotificationChannelUpdateOne) SetNillableEnabled(b *bool) *NotificationChannelUpdateOne {
	if b != nil {
		ncuo.SetEnabled(*b)
	}
	return ncuo
}

// Mutation returns the NotificationChannelMutation object of the builder.
func (ncuo *NotificationChannelUpdateOne) Mutation() *NotificationChannelMutation {
	return ncuo.mutation
}

// Where appends a list predicates to the NotificationChannelUpdate builder.
func (ncuo *NotificationChannelUpdateOne) Where(ps ...predicate.NotificationChannel) *NotificationChannelUpdateOne {
	ncuo.mutation.Where(ps...)
	return ncuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ncuo *NotificationChannelUpdateOne) Select(field string, fields ...string) *NotificationChannelUpdateOne {
	ncuo.fields = append([]string{field}, fields...)
	return ncuo
}

// Save executes the query and returns the updated NotificationChannel entity.
func (ncuo *NotificationChannelUpdateOne) Save(ctx context.Context) (*NotificationChannel, error) {
	ncuo.defaults()
	return withHooks(ctx, ncuo.sqlSave, ncuo.mutation, ncuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ncuo *NotificationChannelUpdateOne) SaveX(ctx context.Context) *NotificationChannel {
	node, err := ncuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ncuo *NotificationChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := ncuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncuo *NotificationChannelUpdateOne) ExecX(ctx context.Context) {
	if err := ncuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncuo *NotificationChannelUpdateOne) defaults() {
	if _, ok := ncuo.mutation.UpdatedAt(); !ok {
		v := notificationchannel.UpdateDefaultUpdatedAt()
		ncuo.mutation.SetUpdatedAt(v)
	}
}

func (ncuo *NotificationChannelUpdateOne) sqlSave(ctx context.Context) (_node *NotificationChannel, err error) {
	_spec := sqlgraph.NewUpdateSpec(notificationchannel.Table, notificationchannel.Columns, sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeInt))
	id, ok := ncuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotificationChannel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ncuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationchannel.FieldID)
		for _, f := range fields {
			if !notificationchannel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notificationchannel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ncuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationchannel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ncuo.mutation.Name(); ok {
		_spec.SetField(notificationchannel.FieldName, field.TypeString, value)
	}
	if value, ok := ncuo.mutation.GetType(); ok {
		_spec.SetField(notificationchannel.FieldType, field.TypeString, value)
	}
	if value, ok := ncuo.mutation.SlackConfig(); ok {
		_spec.SetField(notificationchannel.FieldSlackConfig, field.TypeJSON, value)
	}
	if ncuo.mutation.SlackConfigCleared() {
		_spec.ClearField(notificationchannel.FieldSlackConfig, field.TypeJSON)
	}
	if value, ok := ncuo.mutation.EmailConfig(); ok {
		_spec.SetField(notificationchannel.FieldEmailConfig, field.TypeJSON, value)
	}
	if ncuo.mutation.EmailConfigCleared() {
		_spec.ClearField(notificationchannel.FieldEmailConfig, field.TypeJSON)
	}
	if value, ok := ncuo.mutation.WebhookConfig(); ok {
		_spec.SetField(notificationchannel.FieldWebhookConfig, field.TypeJSON, value)
	}
	if ncuo.mutation.WebhookConfigCleared() {
		_spec.ClearField(notificationchannel.FieldWebhookConfig, field.TypeJSON)
	}
	if value, ok := ncuo.mutation.Enabled(); ok {
		_spec.SetField(notificationchannel.FieldEnabled, field.TypeBool, value)
	}
	_node = &NotificationChannel{config: ncuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ncuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationchannel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ncuo.mutation.done = true
	return _node, nil
}