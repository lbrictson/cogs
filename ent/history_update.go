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
	"github.com/lbrictson/cogs/ent/history"
	"github.com/lbrictson/cogs/ent/predicate"
)

// HistoryUpdate is the builder for updating History entities.
type HistoryUpdate struct {
	config
	hooks    []Hook
	mutation *HistoryMutation
}

// Where appends a list predicates to the HistoryUpdate builder.
func (hu *HistoryUpdate) Where(ps ...predicate.History) *HistoryUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetUpdatedAt sets the "updated_at" field.
func (hu *HistoryUpdate) SetUpdatedAt(t time.Time) *HistoryUpdate {
	hu.mutation.SetUpdatedAt(t)
	return hu
}

// SetRunID sets the "run_id" field.
func (hu *HistoryUpdate) SetRunID(s string) *HistoryUpdate {
	hu.mutation.SetRunID(s)
	return hu
}

// SetSuccess sets the "success" field.
func (hu *HistoryUpdate) SetSuccess(b bool) *HistoryUpdate {
	hu.mutation.SetSuccess(b)
	return hu
}

// SetExitCode sets the "exit_code" field.
func (hu *HistoryUpdate) SetExitCode(i int) *HistoryUpdate {
	hu.mutation.ResetExitCode()
	hu.mutation.SetExitCode(i)
	return hu
}

// AddExitCode adds i to the "exit_code" field.
func (hu *HistoryUpdate) AddExitCode(i int) *HistoryUpdate {
	hu.mutation.AddExitCode(i)
	return hu
}

// SetDuration sets the "duration" field.
func (hu *HistoryUpdate) SetDuration(i int) *HistoryUpdate {
	hu.mutation.ResetDuration()
	hu.mutation.SetDuration(i)
	return hu
}

// AddDuration adds i to the "duration" field.
func (hu *HistoryUpdate) AddDuration(i int) *HistoryUpdate {
	hu.mutation.AddDuration(i)
	return hu
}

// SetTrigger sets the "trigger" field.
func (hu *HistoryUpdate) SetTrigger(s string) *HistoryUpdate {
	hu.mutation.SetTrigger(s)
	return hu
}

// SetOutput sets the "output" field.
func (hu *HistoryUpdate) SetOutput(s string) *HistoryUpdate {
	hu.mutation.SetOutput(s)
	return hu
}

// SetTriggeredBy sets the "triggered_by" field.
func (hu *HistoryUpdate) SetTriggeredBy(s string) *HistoryUpdate {
	hu.mutation.SetTriggeredBy(s)
	return hu
}

// SetScriptID sets the "script_id" field.
func (hu *HistoryUpdate) SetScriptID(i int) *HistoryUpdate {
	hu.mutation.ResetScriptID()
	hu.mutation.SetScriptID(i)
	return hu
}

// AddScriptID adds i to the "script_id" field.
func (hu *HistoryUpdate) AddScriptID(i int) *HistoryUpdate {
	hu.mutation.AddScriptID(i)
	return hu
}

// SetArguments sets the "arguments" field.
func (hu *HistoryUpdate) SetArguments(m map[string]string) *HistoryUpdate {
	hu.mutation.SetArguments(m)
	return hu
}

// SetStatus sets the "status" field.
func (hu *HistoryUpdate) SetStatus(s string) *HistoryUpdate {
	hu.mutation.SetStatus(s)
	return hu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableStatus(s *string) *HistoryUpdate {
	if s != nil {
		hu.SetStatus(*s)
	}
	return hu
}

// Mutation returns the HistoryMutation object of the builder.
func (hu *HistoryUpdate) Mutation() *HistoryMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HistoryUpdate) Save(ctx context.Context) (int, error) {
	hu.defaults()
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HistoryUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HistoryUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hu *HistoryUpdate) defaults() {
	if _, ok := hu.mutation.UpdatedAt(); !ok {
		v := history.UpdateDefaultUpdatedAt()
		hu.mutation.SetUpdatedAt(v)
	}
}

func (hu *HistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.UpdatedAt(); ok {
		_spec.SetField(history.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hu.mutation.RunID(); ok {
		_spec.SetField(history.FieldRunID, field.TypeString, value)
	}
	if value, ok := hu.mutation.Success(); ok {
		_spec.SetField(history.FieldSuccess, field.TypeBool, value)
	}
	if value, ok := hu.mutation.ExitCode(); ok {
		_spec.SetField(history.FieldExitCode, field.TypeInt, value)
	}
	if value, ok := hu.mutation.AddedExitCode(); ok {
		_spec.AddField(history.FieldExitCode, field.TypeInt, value)
	}
	if value, ok := hu.mutation.Duration(); ok {
		_spec.SetField(history.FieldDuration, field.TypeInt, value)
	}
	if value, ok := hu.mutation.AddedDuration(); ok {
		_spec.AddField(history.FieldDuration, field.TypeInt, value)
	}
	if value, ok := hu.mutation.Trigger(); ok {
		_spec.SetField(history.FieldTrigger, field.TypeString, value)
	}
	if value, ok := hu.mutation.Output(); ok {
		_spec.SetField(history.FieldOutput, field.TypeString, value)
	}
	if value, ok := hu.mutation.TriggeredBy(); ok {
		_spec.SetField(history.FieldTriggeredBy, field.TypeString, value)
	}
	if value, ok := hu.mutation.ScriptID(); ok {
		_spec.SetField(history.FieldScriptID, field.TypeInt, value)
	}
	if value, ok := hu.mutation.AddedScriptID(); ok {
		_spec.AddField(history.FieldScriptID, field.TypeInt, value)
	}
	if value, ok := hu.mutation.Arguments(); ok {
		_spec.SetField(history.FieldArguments, field.TypeJSON, value)
	}
	if value, ok := hu.mutation.Status(); ok {
		_spec.SetField(history.FieldStatus, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{history.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HistoryUpdateOne is the builder for updating a single History entity.
type HistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (huo *HistoryUpdateOne) SetUpdatedAt(t time.Time) *HistoryUpdateOne {
	huo.mutation.SetUpdatedAt(t)
	return huo
}

// SetRunID sets the "run_id" field.
func (huo *HistoryUpdateOne) SetRunID(s string) *HistoryUpdateOne {
	huo.mutation.SetRunID(s)
	return huo
}

// SetSuccess sets the "success" field.
func (huo *HistoryUpdateOne) SetSuccess(b bool) *HistoryUpdateOne {
	huo.mutation.SetSuccess(b)
	return huo
}

// SetExitCode sets the "exit_code" field.
func (huo *HistoryUpdateOne) SetExitCode(i int) *HistoryUpdateOne {
	huo.mutation.ResetExitCode()
	huo.mutation.SetExitCode(i)
	return huo
}

// AddExitCode adds i to the "exit_code" field.
func (huo *HistoryUpdateOne) AddExitCode(i int) *HistoryUpdateOne {
	huo.mutation.AddExitCode(i)
	return huo
}

// SetDuration sets the "duration" field.
func (huo *HistoryUpdateOne) SetDuration(i int) *HistoryUpdateOne {
	huo.mutation.ResetDuration()
	huo.mutation.SetDuration(i)
	return huo
}

// AddDuration adds i to the "duration" field.
func (huo *HistoryUpdateOne) AddDuration(i int) *HistoryUpdateOne {
	huo.mutation.AddDuration(i)
	return huo
}

// SetTrigger sets the "trigger" field.
func (huo *HistoryUpdateOne) SetTrigger(s string) *HistoryUpdateOne {
	huo.mutation.SetTrigger(s)
	return huo
}

// SetOutput sets the "output" field.
func (huo *HistoryUpdateOne) SetOutput(s string) *HistoryUpdateOne {
	huo.mutation.SetOutput(s)
	return huo
}

// SetTriggeredBy sets the "triggered_by" field.
func (huo *HistoryUpdateOne) SetTriggeredBy(s string) *HistoryUpdateOne {
	huo.mutation.SetTriggeredBy(s)
	return huo
}

// SetScriptID sets the "script_id" field.
func (huo *HistoryUpdateOne) SetScriptID(i int) *HistoryUpdateOne {
	huo.mutation.ResetScriptID()
	huo.mutation.SetScriptID(i)
	return huo
}

// AddScriptID adds i to the "script_id" field.
func (huo *HistoryUpdateOne) AddScriptID(i int) *HistoryUpdateOne {
	huo.mutation.AddScriptID(i)
	return huo
}

// SetArguments sets the "arguments" field.
func (huo *HistoryUpdateOne) SetArguments(m map[string]string) *HistoryUpdateOne {
	huo.mutation.SetArguments(m)
	return huo
}

// SetStatus sets the "status" field.
func (huo *HistoryUpdateOne) SetStatus(s string) *HistoryUpdateOne {
	huo.mutation.SetStatus(s)
	return huo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableStatus(s *string) *HistoryUpdateOne {
	if s != nil {
		huo.SetStatus(*s)
	}
	return huo
}

// Mutation returns the HistoryMutation object of the builder.
func (huo *HistoryUpdateOne) Mutation() *HistoryMutation {
	return huo.mutation
}

// Where appends a list predicates to the HistoryUpdate builder.
func (huo *HistoryUpdateOne) Where(ps ...predicate.History) *HistoryUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HistoryUpdateOne) Select(field string, fields ...string) *HistoryUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated History entity.
func (huo *HistoryUpdateOne) Save(ctx context.Context) (*History, error) {
	huo.defaults()
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HistoryUpdateOne) SaveX(ctx context.Context) *History {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HistoryUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (huo *HistoryUpdateOne) defaults() {
	if _, ok := huo.mutation.UpdatedAt(); !ok {
		v := history.UpdateDefaultUpdatedAt()
		huo.mutation.SetUpdatedAt(v)
	}
}

func (huo *HistoryUpdateOne) sqlSave(ctx context.Context) (_node *History, err error) {
	_spec := sqlgraph.NewUpdateSpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "History.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, history.FieldID)
		for _, f := range fields {
			if !history.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != history.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.UpdatedAt(); ok {
		_spec.SetField(history.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := huo.mutation.RunID(); ok {
		_spec.SetField(history.FieldRunID, field.TypeString, value)
	}
	if value, ok := huo.mutation.Success(); ok {
		_spec.SetField(history.FieldSuccess, field.TypeBool, value)
	}
	if value, ok := huo.mutation.ExitCode(); ok {
		_spec.SetField(history.FieldExitCode, field.TypeInt, value)
	}
	if value, ok := huo.mutation.AddedExitCode(); ok {
		_spec.AddField(history.FieldExitCode, field.TypeInt, value)
	}
	if value, ok := huo.mutation.Duration(); ok {
		_spec.SetField(history.FieldDuration, field.TypeInt, value)
	}
	if value, ok := huo.mutation.AddedDuration(); ok {
		_spec.AddField(history.FieldDuration, field.TypeInt, value)
	}
	if value, ok := huo.mutation.Trigger(); ok {
		_spec.SetField(history.FieldTrigger, field.TypeString, value)
	}
	if value, ok := huo.mutation.Output(); ok {
		_spec.SetField(history.FieldOutput, field.TypeString, value)
	}
	if value, ok := huo.mutation.TriggeredBy(); ok {
		_spec.SetField(history.FieldTriggeredBy, field.TypeString, value)
	}
	if value, ok := huo.mutation.ScriptID(); ok {
		_spec.SetField(history.FieldScriptID, field.TypeInt, value)
	}
	if value, ok := huo.mutation.AddedScriptID(); ok {
		_spec.AddField(history.FieldScriptID, field.TypeInt, value)
	}
	if value, ok := huo.mutation.Arguments(); ok {
		_spec.SetField(history.FieldArguments, field.TypeJSON, value)
	}
	if value, ok := huo.mutation.Status(); ok {
		_spec.SetField(history.FieldStatus, field.TypeString, value)
	}
	_node = &History{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{history.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}