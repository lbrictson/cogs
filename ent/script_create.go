// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/cogs/ent/schema"
	"github.com/lbrictson/cogs/ent/script"
)

// ScriptCreate is the builder for creating a Script entity.
type ScriptCreate struct {
	config
	mutation *ScriptMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *ScriptCreate) SetCreatedAt(t time.Time) *ScriptCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ScriptCreate) SetNillableCreatedAt(t *time.Time) *ScriptCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ScriptCreate) SetUpdatedAt(t time.Time) *ScriptCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ScriptCreate) SetNillableUpdatedAt(t *time.Time) *ScriptCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *ScriptCreate) SetName(s string) *ScriptCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *ScriptCreate) SetDescription(s string) *ScriptCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *ScriptCreate) SetNillableDescription(s *string) *ScriptCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetScript sets the "script" field.
func (sc *ScriptCreate) SetScript(s string) *ScriptCreate {
	sc.mutation.SetScript(s)
	return sc
}

// SetTimeoutSeconds sets the "timeout_seconds" field.
func (sc *ScriptCreate) SetTimeoutSeconds(i int) *ScriptCreate {
	sc.mutation.SetTimeoutSeconds(i)
	return sc
}

// SetNillableTimeoutSeconds sets the "timeout_seconds" field if the given value is not nil.
func (sc *ScriptCreate) SetNillableTimeoutSeconds(i *int) *ScriptCreate {
	if i != nil {
		sc.SetTimeoutSeconds(*i)
	}
	return sc
}

// SetProjectID sets the "project_id" field.
func (sc *ScriptCreate) SetProjectID(i int) *ScriptCreate {
	sc.mutation.SetProjectID(i)
	return sc
}

// SetParameters sets the "parameters" field.
func (sc *ScriptCreate) SetParameters(sio []schema.ScriptInputOptions) *ScriptCreate {
	sc.mutation.SetParameters(sio)
	return sc
}

// SetScheduleEnabled sets the "schedule_enabled" field.
func (sc *ScriptCreate) SetScheduleEnabled(b bool) *ScriptCreate {
	sc.mutation.SetScheduleEnabled(b)
	return sc
}

// SetNillableScheduleEnabled sets the "schedule_enabled" field if the given value is not nil.
func (sc *ScriptCreate) SetNillableScheduleEnabled(b *bool) *ScriptCreate {
	if b != nil {
		sc.SetScheduleEnabled(*b)
	}
	return sc
}

// SetScheduleCron sets the "schedule_cron" field.
func (sc *ScriptCreate) SetScheduleCron(s string) *ScriptCreate {
	sc.mutation.SetScheduleCron(s)
	return sc
}

// SetNillableScheduleCron sets the "schedule_cron" field if the given value is not nil.
func (sc *ScriptCreate) SetNillableScheduleCron(s *string) *ScriptCreate {
	if s != nil {
		sc.SetScheduleCron(*s)
	}
	return sc
}

// SetSuccessNotificationChannelID sets the "success_notification_channel_id" field.
func (sc *ScriptCreate) SetSuccessNotificationChannelID(i int) *ScriptCreate {
	sc.mutation.SetSuccessNotificationChannelID(i)
	return sc
}

// SetNillableSuccessNotificationChannelID sets the "success_notification_channel_id" field if the given value is not nil.
func (sc *ScriptCreate) SetNillableSuccessNotificationChannelID(i *int) *ScriptCreate {
	if i != nil {
		sc.SetSuccessNotificationChannelID(*i)
	}
	return sc
}

// SetFailureNotificationChannelID sets the "failure_notification_channel_id" field.
func (sc *ScriptCreate) SetFailureNotificationChannelID(i int) *ScriptCreate {
	sc.mutation.SetFailureNotificationChannelID(i)
	return sc
}

// SetNillableFailureNotificationChannelID sets the "failure_notification_channel_id" field if the given value is not nil.
func (sc *ScriptCreate) SetNillableFailureNotificationChannelID(i *int) *ScriptCreate {
	if i != nil {
		sc.SetFailureNotificationChannelID(*i)
	}
	return sc
}

// Mutation returns the ScriptMutation object of the builder.
func (sc *ScriptCreate) Mutation() *ScriptMutation {
	return sc.mutation
}

// Save creates the Script in the database.
func (sc *ScriptCreate) Save(ctx context.Context) (*Script, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScriptCreate) SaveX(ctx context.Context) *Script {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScriptCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScriptCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ScriptCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := script.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := script.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.TimeoutSeconds(); !ok {
		v := script.DefaultTimeoutSeconds
		sc.mutation.SetTimeoutSeconds(v)
	}
	if _, ok := sc.mutation.ScheduleEnabled(); !ok {
		v := script.DefaultScheduleEnabled
		sc.mutation.SetScheduleEnabled(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScriptCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Script.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Script.updated_at"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Script.name"`)}
	}
	if _, ok := sc.mutation.Script(); !ok {
		return &ValidationError{Name: "script", err: errors.New(`ent: missing required field "Script.script"`)}
	}
	if _, ok := sc.mutation.TimeoutSeconds(); !ok {
		return &ValidationError{Name: "timeout_seconds", err: errors.New(`ent: missing required field "Script.timeout_seconds"`)}
	}
	if _, ok := sc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project_id", err: errors.New(`ent: missing required field "Script.project_id"`)}
	}
	if _, ok := sc.mutation.ScheduleEnabled(); !ok {
		return &ValidationError{Name: "schedule_enabled", err: errors.New(`ent: missing required field "Script.schedule_enabled"`)}
	}
	return nil
}

func (sc *ScriptCreate) sqlSave(ctx context.Context) (*Script, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScriptCreate) createSpec() (*Script, *sqlgraph.CreateSpec) {
	var (
		_node = &Script{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(script.Table, sqlgraph.NewFieldSpec(script.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(script.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(script.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(script.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(script.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.Script(); ok {
		_spec.SetField(script.FieldScript, field.TypeString, value)
		_node.Script = value
	}
	if value, ok := sc.mutation.TimeoutSeconds(); ok {
		_spec.SetField(script.FieldTimeoutSeconds, field.TypeInt, value)
		_node.TimeoutSeconds = value
	}
	if value, ok := sc.mutation.ProjectID(); ok {
		_spec.SetField(script.FieldProjectID, field.TypeInt, value)
		_node.ProjectID = value
	}
	if value, ok := sc.mutation.Parameters(); ok {
		_spec.SetField(script.FieldParameters, field.TypeJSON, value)
		_node.Parameters = value
	}
	if value, ok := sc.mutation.ScheduleEnabled(); ok {
		_spec.SetField(script.FieldScheduleEnabled, field.TypeBool, value)
		_node.ScheduleEnabled = value
	}
	if value, ok := sc.mutation.ScheduleCron(); ok {
		_spec.SetField(script.FieldScheduleCron, field.TypeString, value)
		_node.ScheduleCron = value
	}
	if value, ok := sc.mutation.SuccessNotificationChannelID(); ok {
		_spec.SetField(script.FieldSuccessNotificationChannelID, field.TypeInt, value)
		_node.SuccessNotificationChannelID = &value
	}
	if value, ok := sc.mutation.FailureNotificationChannelID(); ok {
		_spec.SetField(script.FieldFailureNotificationChannelID, field.TypeInt, value)
		_node.FailureNotificationChannelID = &value
	}
	return _node, _spec
}

// ScriptCreateBulk is the builder for creating many Script entities in bulk.
type ScriptCreateBulk struct {
	config
	builders []*ScriptCreate
}

// Save creates the Script entities in the database.
func (scb *ScriptCreateBulk) Save(ctx context.Context) ([]*Script, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Script, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScriptMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScriptCreateBulk) SaveX(ctx context.Context) []*Script {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScriptCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScriptCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
