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
	"github.com/lbrictson/cogs/ent/access"
	"github.com/lbrictson/cogs/ent/predicate"
)

// AccessUpdate is the builder for updating Access entities.
type AccessUpdate struct {
	config
	hooks    []Hook
	mutation *AccessMutation
}

// Where appends a list predicates to the AccessUpdate builder.
func (au *AccessUpdate) Where(ps ...predicate.Access) *AccessUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccessUpdate) SetUpdatedAt(t time.Time) *AccessUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetUserID sets the "user_id" field.
func (au *AccessUpdate) SetUserID(i int) *AccessUpdate {
	au.mutation.ResetUserID()
	au.mutation.SetUserID(i)
	return au
}

// AddUserID adds i to the "user_id" field.
func (au *AccessUpdate) AddUserID(i int) *AccessUpdate {
	au.mutation.AddUserID(i)
	return au
}

// SetProjectID sets the "project_id" field.
func (au *AccessUpdate) SetProjectID(i int) *AccessUpdate {
	au.mutation.ResetProjectID()
	au.mutation.SetProjectID(i)
	return au
}

// AddProjectID adds i to the "project_id" field.
func (au *AccessUpdate) AddProjectID(i int) *AccessUpdate {
	au.mutation.AddProjectID(i)
	return au
}

// SetRole sets the "role" field.
func (au *AccessUpdate) SetRole(a access.Role) *AccessUpdate {
	au.mutation.SetRole(a)
	return au
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (au *AccessUpdate) SetNillableRole(a *access.Role) *AccessUpdate {
	if a != nil {
		au.SetRole(*a)
	}
	return au
}

// Mutation returns the AccessMutation object of the builder.
func (au *AccessUpdate) Mutation() *AccessMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccessUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccessUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccessUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccessUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccessUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := access.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccessUpdate) check() error {
	if v, ok := au.mutation.Role(); ok {
		if err := access.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Access.role": %w`, err)}
		}
	}
	return nil
}

func (au *AccessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(access.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UserID(); ok {
		_spec.SetField(access.FieldUserID, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedUserID(); ok {
		_spec.AddField(access.FieldUserID, field.TypeInt, value)
	}
	if value, ok := au.mutation.ProjectID(); ok {
		_spec.SetField(access.FieldProjectID, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedProjectID(); ok {
		_spec.AddField(access.FieldProjectID, field.TypeInt, value)
	}
	if value, ok := au.mutation.Role(); ok {
		_spec.SetField(access.FieldRole, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{access.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccessUpdateOne is the builder for updating a single Access entity.
type AccessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccessUpdateOne) SetUpdatedAt(t time.Time) *AccessUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetUserID sets the "user_id" field.
func (auo *AccessUpdateOne) SetUserID(i int) *AccessUpdateOne {
	auo.mutation.ResetUserID()
	auo.mutation.SetUserID(i)
	return auo
}

// AddUserID adds i to the "user_id" field.
func (auo *AccessUpdateOne) AddUserID(i int) *AccessUpdateOne {
	auo.mutation.AddUserID(i)
	return auo
}

// SetProjectID sets the "project_id" field.
func (auo *AccessUpdateOne) SetProjectID(i int) *AccessUpdateOne {
	auo.mutation.ResetProjectID()
	auo.mutation.SetProjectID(i)
	return auo
}

// AddProjectID adds i to the "project_id" field.
func (auo *AccessUpdateOne) AddProjectID(i int) *AccessUpdateOne {
	auo.mutation.AddProjectID(i)
	return auo
}

// SetRole sets the "role" field.
func (auo *AccessUpdateOne) SetRole(a access.Role) *AccessUpdateOne {
	auo.mutation.SetRole(a)
	return auo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableRole(a *access.Role) *AccessUpdateOne {
	if a != nil {
		auo.SetRole(*a)
	}
	return auo
}

// Mutation returns the AccessMutation object of the builder.
func (auo *AccessUpdateOne) Mutation() *AccessMutation {
	return auo.mutation
}

// Where appends a list predicates to the AccessUpdate builder.
func (auo *AccessUpdateOne) Where(ps ...predicate.Access) *AccessUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccessUpdateOne) Select(field string, fields ...string) *AccessUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Access entity.
func (auo *AccessUpdateOne) Save(ctx context.Context) (*Access, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccessUpdateOne) SaveX(ctx context.Context) *Access {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccessUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccessUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccessUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := access.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccessUpdateOne) check() error {
	if v, ok := auo.mutation.Role(); ok {
		if err := access.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Access.role": %w`, err)}
		}
	}
	return nil
}

func (auo *AccessUpdateOne) sqlSave(ctx context.Context) (_node *Access, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Access.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, access.FieldID)
		for _, f := range fields {
			if !access.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != access.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(access.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UserID(); ok {
		_spec.SetField(access.FieldUserID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedUserID(); ok {
		_spec.AddField(access.FieldUserID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.ProjectID(); ok {
		_spec.SetField(access.FieldProjectID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedProjectID(); ok {
		_spec.AddField(access.FieldProjectID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.Role(); ok {
		_spec.SetField(access.FieldRole, field.TypeEnum, value)
	}
	_node = &Access{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{access.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}