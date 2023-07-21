// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/cogs/ent/notificationchannel"
	"github.com/lbrictson/cogs/ent/schema"
)

// NotificationChannelCreate is the builder for creating a NotificationChannel entity.
type NotificationChannelCreate struct {
	config
	mutation *NotificationChannelMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ncc *NotificationChannelCreate) SetCreatedAt(t time.Time) *NotificationChannelCreate {
	ncc.mutation.SetCreatedAt(t)
	return ncc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ncc *NotificationChannelCreate) SetNillableCreatedAt(t *time.Time) *NotificationChannelCreate {
	if t != nil {
		ncc.SetCreatedAt(*t)
	}
	return ncc
}

// SetUpdatedAt sets the "updated_at" field.
func (ncc *NotificationChannelCreate) SetUpdatedAt(t time.Time) *NotificationChannelCreate {
	ncc.mutation.SetUpdatedAt(t)
	return ncc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ncc *NotificationChannelCreate) SetNillableUpdatedAt(t *time.Time) *NotificationChannelCreate {
	if t != nil {
		ncc.SetUpdatedAt(*t)
	}
	return ncc
}

// SetName sets the "name" field.
func (ncc *NotificationChannelCreate) SetName(s string) *NotificationChannelCreate {
	ncc.mutation.SetName(s)
	return ncc
}

// SetType sets the "type" field.
func (ncc *NotificationChannelCreate) SetType(s string) *NotificationChannelCreate {
	ncc.mutation.SetType(s)
	return ncc
}

// SetSlackConfig sets the "slack_config" field.
func (ncc *NotificationChannelCreate) SetSlackConfig(sc schema.SlackConfig) *NotificationChannelCreate {
	ncc.mutation.SetSlackConfig(sc)
	return ncc
}

// SetNillableSlackConfig sets the "slack_config" field if the given value is not nil.
func (ncc *NotificationChannelCreate) SetNillableSlackConfig(sc *schema.SlackConfig) *NotificationChannelCreate {
	if sc != nil {
		ncc.SetSlackConfig(*sc)
	}
	return ncc
}

// SetEmailConfig sets the "email_config" field.
func (ncc *NotificationChannelCreate) SetEmailConfig(sc schema.EmailConfig) *NotificationChannelCreate {
	ncc.mutation.SetEmailConfig(sc)
	return ncc
}

// SetNillableEmailConfig sets the "email_config" field if the given value is not nil.
func (ncc *NotificationChannelCreate) SetNillableEmailConfig(sc *schema.EmailConfig) *NotificationChannelCreate {
	if sc != nil {
		ncc.SetEmailConfig(*sc)
	}
	return ncc
}

// SetWebhookConfig sets the "webhook_config" field.
func (ncc *NotificationChannelCreate) SetWebhookConfig(sc schema.WebhookConfig) *NotificationChannelCreate {
	ncc.mutation.SetWebhookConfig(sc)
	return ncc
}

// SetNillableWebhookConfig sets the "webhook_config" field if the given value is not nil.
func (ncc *NotificationChannelCreate) SetNillableWebhookConfig(sc *schema.WebhookConfig) *NotificationChannelCreate {
	if sc != nil {
		ncc.SetWebhookConfig(*sc)
	}
	return ncc
}

// SetEnabled sets the "enabled" field.
func (ncc *NotificationChannelCreate) SetEnabled(b bool) *NotificationChannelCreate {
	ncc.mutation.SetEnabled(b)
	return ncc
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (ncc *NotificationChannelCreate) SetNillableEnabled(b *bool) *NotificationChannelCreate {
	if b != nil {
		ncc.SetEnabled(*b)
	}
	return ncc
}

// Mutation returns the NotificationChannelMutation object of the builder.
func (ncc *NotificationChannelCreate) Mutation() *NotificationChannelMutation {
	return ncc.mutation
}

// Save creates the NotificationChannel in the database.
func (ncc *NotificationChannelCreate) Save(ctx context.Context) (*NotificationChannel, error) {
	ncc.defaults()
	return withHooks(ctx, ncc.sqlSave, ncc.mutation, ncc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ncc *NotificationChannelCreate) SaveX(ctx context.Context) *NotificationChannel {
	v, err := ncc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncc *NotificationChannelCreate) Exec(ctx context.Context) error {
	_, err := ncc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncc *NotificationChannelCreate) ExecX(ctx context.Context) {
	if err := ncc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncc *NotificationChannelCreate) defaults() {
	if _, ok := ncc.mutation.CreatedAt(); !ok {
		v := notificationchannel.DefaultCreatedAt()
		ncc.mutation.SetCreatedAt(v)
	}
	if _, ok := ncc.mutation.UpdatedAt(); !ok {
		v := notificationchannel.DefaultUpdatedAt()
		ncc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ncc.mutation.Enabled(); !ok {
		v := notificationchannel.DefaultEnabled
		ncc.mutation.SetEnabled(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ncc *NotificationChannelCreate) check() error {
	if _, ok := ncc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NotificationChannel.created_at"`)}
	}
	if _, ok := ncc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "NotificationChannel.updated_at"`)}
	}
	if _, ok := ncc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "NotificationChannel.name"`)}
	}
	if _, ok := ncc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "NotificationChannel.type"`)}
	}
	if _, ok := ncc.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New(`ent: missing required field "NotificationChannel.enabled"`)}
	}
	return nil
}

func (ncc *NotificationChannelCreate) sqlSave(ctx context.Context) (*NotificationChannel, error) {
	if err := ncc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ncc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ncc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ncc.mutation.id = &_node.ID
	ncc.mutation.done = true
	return _node, nil
}

func (ncc *NotificationChannelCreate) createSpec() (*NotificationChannel, *sqlgraph.CreateSpec) {
	var (
		_node = &NotificationChannel{config: ncc.config}
		_spec = sqlgraph.NewCreateSpec(notificationchannel.Table, sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeInt))
	)
	if value, ok := ncc.mutation.CreatedAt(); ok {
		_spec.SetField(notificationchannel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ncc.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationchannel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ncc.mutation.Name(); ok {
		_spec.SetField(notificationchannel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ncc.mutation.GetType(); ok {
		_spec.SetField(notificationchannel.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := ncc.mutation.SlackConfig(); ok {
		_spec.SetField(notificationchannel.FieldSlackConfig, field.TypeJSON, value)
		_node.SlackConfig = value
	}
	if value, ok := ncc.mutation.EmailConfig(); ok {
		_spec.SetField(notificationchannel.FieldEmailConfig, field.TypeJSON, value)
		_node.EmailConfig = value
	}
	if value, ok := ncc.mutation.WebhookConfig(); ok {
		_spec.SetField(notificationchannel.FieldWebhookConfig, field.TypeJSON, value)
		_node.WebhookConfig = value
	}
	if value, ok := ncc.mutation.Enabled(); ok {
		_spec.SetField(notificationchannel.FieldEnabled, field.TypeBool, value)
		_node.Enabled = value
	}
	return _node, _spec
}

// NotificationChannelCreateBulk is the builder for creating many NotificationChannel entities in bulk.
type NotificationChannelCreateBulk struct {
	config
	builders []*NotificationChannelCreate
}

// Save creates the NotificationChannel entities in the database.
func (nccb *NotificationChannelCreateBulk) Save(ctx context.Context) ([]*NotificationChannel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nccb.builders))
	nodes := make([]*NotificationChannel, len(nccb.builders))
	mutators := make([]Mutator, len(nccb.builders))
	for i := range nccb.builders {
		func(i int, root context.Context) {
			builder := nccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationChannelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, nccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, nccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nccb *NotificationChannelCreateBulk) SaveX(ctx context.Context) []*NotificationChannel {
	v, err := nccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nccb *NotificationChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := nccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nccb *NotificationChannelCreateBulk) ExecX(ctx context.Context) {
	if err := nccb.Exec(ctx); err != nil {
		panic(err)
	}
}