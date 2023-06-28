// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/cogs/ent/history"
)

// HistoryCreate is the builder for creating a History entity.
type HistoryCreate struct {
	config
	mutation *HistoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (hc *HistoryCreate) SetCreatedAt(t time.Time) *HistoryCreate {
	hc.mutation.SetCreatedAt(t)
	return hc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableCreatedAt(t *time.Time) *HistoryCreate {
	if t != nil {
		hc.SetCreatedAt(*t)
	}
	return hc
}

// SetUpdatedAt sets the "updated_at" field.
func (hc *HistoryCreate) SetUpdatedAt(t time.Time) *HistoryCreate {
	hc.mutation.SetUpdatedAt(t)
	return hc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableUpdatedAt(t *time.Time) *HistoryCreate {
	if t != nil {
		hc.SetUpdatedAt(*t)
	}
	return hc
}

// SetRunID sets the "run_id" field.
func (hc *HistoryCreate) SetRunID(s string) *HistoryCreate {
	hc.mutation.SetRunID(s)
	return hc
}

// SetSuccess sets the "success" field.
func (hc *HistoryCreate) SetSuccess(b bool) *HistoryCreate {
	hc.mutation.SetSuccess(b)
	return hc
}

// SetExitCode sets the "exit_code" field.
func (hc *HistoryCreate) SetExitCode(i int) *HistoryCreate {
	hc.mutation.SetExitCode(i)
	return hc
}

// SetDuration sets the "duration" field.
func (hc *HistoryCreate) SetDuration(i int) *HistoryCreate {
	hc.mutation.SetDuration(i)
	return hc
}

// SetTrigger sets the "trigger" field.
func (hc *HistoryCreate) SetTrigger(s string) *HistoryCreate {
	hc.mutation.SetTrigger(s)
	return hc
}

// SetOutput sets the "output" field.
func (hc *HistoryCreate) SetOutput(s string) *HistoryCreate {
	hc.mutation.SetOutput(s)
	return hc
}

// SetTriggeredBy sets the "triggered_by" field.
func (hc *HistoryCreate) SetTriggeredBy(s string) *HistoryCreate {
	hc.mutation.SetTriggeredBy(s)
	return hc
}

// SetScriptID sets the "script_id" field.
func (hc *HistoryCreate) SetScriptID(i int) *HistoryCreate {
	hc.mutation.SetScriptID(i)
	return hc
}

// SetArguments sets the "arguments" field.
func (hc *HistoryCreate) SetArguments(m map[string]string) *HistoryCreate {
	hc.mutation.SetArguments(m)
	return hc
}

// SetStatus sets the "status" field.
func (hc *HistoryCreate) SetStatus(s string) *HistoryCreate {
	hc.mutation.SetStatus(s)
	return hc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableStatus(s *string) *HistoryCreate {
	if s != nil {
		hc.SetStatus(*s)
	}
	return hc
}

// Mutation returns the HistoryMutation object of the builder.
func (hc *HistoryCreate) Mutation() *HistoryMutation {
	return hc.mutation
}

// Save creates the History in the database.
func (hc *HistoryCreate) Save(ctx context.Context) (*History, error) {
	hc.defaults()
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HistoryCreate) SaveX(ctx context.Context) *History {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HistoryCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HistoryCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HistoryCreate) defaults() {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		v := history.DefaultCreatedAt()
		hc.mutation.SetCreatedAt(v)
	}
	if _, ok := hc.mutation.UpdatedAt(); !ok {
		v := history.DefaultUpdatedAt()
		hc.mutation.SetUpdatedAt(v)
	}
	if _, ok := hc.mutation.Status(); !ok {
		v := history.DefaultStatus
		hc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HistoryCreate) check() error {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "History.created_at"`)}
	}
	if _, ok := hc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "History.updated_at"`)}
	}
	if _, ok := hc.mutation.RunID(); !ok {
		return &ValidationError{Name: "run_id", err: errors.New(`ent: missing required field "History.run_id"`)}
	}
	if _, ok := hc.mutation.Success(); !ok {
		return &ValidationError{Name: "success", err: errors.New(`ent: missing required field "History.success"`)}
	}
	if _, ok := hc.mutation.ExitCode(); !ok {
		return &ValidationError{Name: "exit_code", err: errors.New(`ent: missing required field "History.exit_code"`)}
	}
	if _, ok := hc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "History.duration"`)}
	}
	if _, ok := hc.mutation.Trigger(); !ok {
		return &ValidationError{Name: "trigger", err: errors.New(`ent: missing required field "History.trigger"`)}
	}
	if _, ok := hc.mutation.Output(); !ok {
		return &ValidationError{Name: "output", err: errors.New(`ent: missing required field "History.output"`)}
	}
	if _, ok := hc.mutation.TriggeredBy(); !ok {
		return &ValidationError{Name: "triggered_by", err: errors.New(`ent: missing required field "History.triggered_by"`)}
	}
	if _, ok := hc.mutation.ScriptID(); !ok {
		return &ValidationError{Name: "script_id", err: errors.New(`ent: missing required field "History.script_id"`)}
	}
	if _, ok := hc.mutation.Arguments(); !ok {
		return &ValidationError{Name: "arguments", err: errors.New(`ent: missing required field "History.arguments"`)}
	}
	if _, ok := hc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "History.status"`)}
	}
	return nil
}

func (hc *HistoryCreate) sqlSave(ctx context.Context) (*History, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HistoryCreate) createSpec() (*History, *sqlgraph.CreateSpec) {
	var (
		_node = &History{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(history.Table, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	)
	if value, ok := hc.mutation.CreatedAt(); ok {
		_spec.SetField(history.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hc.mutation.UpdatedAt(); ok {
		_spec.SetField(history.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := hc.mutation.RunID(); ok {
		_spec.SetField(history.FieldRunID, field.TypeString, value)
		_node.RunID = value
	}
	if value, ok := hc.mutation.Success(); ok {
		_spec.SetField(history.FieldSuccess, field.TypeBool, value)
		_node.Success = value
	}
	if value, ok := hc.mutation.ExitCode(); ok {
		_spec.SetField(history.FieldExitCode, field.TypeInt, value)
		_node.ExitCode = value
	}
	if value, ok := hc.mutation.Duration(); ok {
		_spec.SetField(history.FieldDuration, field.TypeInt, value)
		_node.Duration = value
	}
	if value, ok := hc.mutation.Trigger(); ok {
		_spec.SetField(history.FieldTrigger, field.TypeString, value)
		_node.Trigger = value
	}
	if value, ok := hc.mutation.Output(); ok {
		_spec.SetField(history.FieldOutput, field.TypeString, value)
		_node.Output = value
	}
	if value, ok := hc.mutation.TriggeredBy(); ok {
		_spec.SetField(history.FieldTriggeredBy, field.TypeString, value)
		_node.TriggeredBy = value
	}
	if value, ok := hc.mutation.ScriptID(); ok {
		_spec.SetField(history.FieldScriptID, field.TypeInt, value)
		_node.ScriptID = value
	}
	if value, ok := hc.mutation.Arguments(); ok {
		_spec.SetField(history.FieldArguments, field.TypeJSON, value)
		_node.Arguments = value
	}
	if value, ok := hc.mutation.Status(); ok {
		_spec.SetField(history.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	return _node, _spec
}

// HistoryCreateBulk is the builder for creating many History entities in bulk.
type HistoryCreateBulk struct {
	config
	builders []*HistoryCreate
}

// Save creates the History entities in the database.
func (hcb *HistoryCreateBulk) Save(ctx context.Context) ([]*History, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*History, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HistoryCreateBulk) SaveX(ctx context.Context) []*History {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HistoryCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
