// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/cogs/ent/predicate"
	"github.com/lbrictson/cogs/ent/scriptstats"
)

// ScriptStatsQuery is the builder for querying ScriptStats entities.
type ScriptStatsQuery struct {
	config
	ctx        *QueryContext
	order      []scriptstats.OrderOption
	inters     []Interceptor
	predicates []predicate.ScriptStats
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScriptStatsQuery builder.
func (ssq *ScriptStatsQuery) Where(ps ...predicate.ScriptStats) *ScriptStatsQuery {
	ssq.predicates = append(ssq.predicates, ps...)
	return ssq
}

// Limit the number of records to be returned by this query.
func (ssq *ScriptStatsQuery) Limit(limit int) *ScriptStatsQuery {
	ssq.ctx.Limit = &limit
	return ssq
}

// Offset to start from.
func (ssq *ScriptStatsQuery) Offset(offset int) *ScriptStatsQuery {
	ssq.ctx.Offset = &offset
	return ssq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ssq *ScriptStatsQuery) Unique(unique bool) *ScriptStatsQuery {
	ssq.ctx.Unique = &unique
	return ssq
}

// Order specifies how the records should be ordered.
func (ssq *ScriptStatsQuery) Order(o ...scriptstats.OrderOption) *ScriptStatsQuery {
	ssq.order = append(ssq.order, o...)
	return ssq
}

// First returns the first ScriptStats entity from the query.
// Returns a *NotFoundError when no ScriptStats was found.
func (ssq *ScriptStatsQuery) First(ctx context.Context) (*ScriptStats, error) {
	nodes, err := ssq.Limit(1).All(setContextOp(ctx, ssq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{scriptstats.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ssq *ScriptStatsQuery) FirstX(ctx context.Context) *ScriptStats {
	node, err := ssq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ScriptStats ID from the query.
// Returns a *NotFoundError when no ScriptStats ID was found.
func (ssq *ScriptStatsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ssq.Limit(1).IDs(setContextOp(ctx, ssq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scriptstats.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ssq *ScriptStatsQuery) FirstIDX(ctx context.Context) int {
	id, err := ssq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ScriptStats entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ScriptStats entity is found.
// Returns a *NotFoundError when no ScriptStats entities are found.
func (ssq *ScriptStatsQuery) Only(ctx context.Context) (*ScriptStats, error) {
	nodes, err := ssq.Limit(2).All(setContextOp(ctx, ssq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{scriptstats.Label}
	default:
		return nil, &NotSingularError{scriptstats.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ssq *ScriptStatsQuery) OnlyX(ctx context.Context) *ScriptStats {
	node, err := ssq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ScriptStats ID in the query.
// Returns a *NotSingularError when more than one ScriptStats ID is found.
// Returns a *NotFoundError when no entities are found.
func (ssq *ScriptStatsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ssq.Limit(2).IDs(setContextOp(ctx, ssq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scriptstats.Label}
	default:
		err = &NotSingularError{scriptstats.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ssq *ScriptStatsQuery) OnlyIDX(ctx context.Context) int {
	id, err := ssq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ScriptStatsSlice.
func (ssq *ScriptStatsQuery) All(ctx context.Context) ([]*ScriptStats, error) {
	ctx = setContextOp(ctx, ssq.ctx, "All")
	if err := ssq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ScriptStats, *ScriptStatsQuery]()
	return withInterceptors[[]*ScriptStats](ctx, ssq, qr, ssq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ssq *ScriptStatsQuery) AllX(ctx context.Context) []*ScriptStats {
	nodes, err := ssq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ScriptStats IDs.
func (ssq *ScriptStatsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ssq.ctx.Unique == nil && ssq.path != nil {
		ssq.Unique(true)
	}
	ctx = setContextOp(ctx, ssq.ctx, "IDs")
	if err = ssq.Select(scriptstats.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ssq *ScriptStatsQuery) IDsX(ctx context.Context) []int {
	ids, err := ssq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ssq *ScriptStatsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ssq.ctx, "Count")
	if err := ssq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ssq, querierCount[*ScriptStatsQuery](), ssq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ssq *ScriptStatsQuery) CountX(ctx context.Context) int {
	count, err := ssq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ssq *ScriptStatsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ssq.ctx, "Exist")
	switch _, err := ssq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ssq *ScriptStatsQuery) ExistX(ctx context.Context) bool {
	exist, err := ssq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScriptStatsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ssq *ScriptStatsQuery) Clone() *ScriptStatsQuery {
	if ssq == nil {
		return nil
	}
	return &ScriptStatsQuery{
		config:     ssq.config,
		ctx:        ssq.ctx.Clone(),
		order:      append([]scriptstats.OrderOption{}, ssq.order...),
		inters:     append([]Interceptor{}, ssq.inters...),
		predicates: append([]predicate.ScriptStats{}, ssq.predicates...),
		// clone intermediate query.
		sql:  ssq.sql.Clone(),
		path: ssq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ScriptID int `json:"script_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ScriptStats.Query().
//		GroupBy(scriptstats.FieldScriptID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ssq *ScriptStatsQuery) GroupBy(field string, fields ...string) *ScriptStatsGroupBy {
	ssq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ScriptStatsGroupBy{build: ssq}
	grbuild.flds = &ssq.ctx.Fields
	grbuild.label = scriptstats.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ScriptID int `json:"script_id,omitempty"`
//	}
//
//	client.ScriptStats.Query().
//		Select(scriptstats.FieldScriptID).
//		Scan(ctx, &v)
func (ssq *ScriptStatsQuery) Select(fields ...string) *ScriptStatsSelect {
	ssq.ctx.Fields = append(ssq.ctx.Fields, fields...)
	sbuild := &ScriptStatsSelect{ScriptStatsQuery: ssq}
	sbuild.label = scriptstats.Label
	sbuild.flds, sbuild.scan = &ssq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ScriptStatsSelect configured with the given aggregations.
func (ssq *ScriptStatsQuery) Aggregate(fns ...AggregateFunc) *ScriptStatsSelect {
	return ssq.Select().Aggregate(fns...)
}

func (ssq *ScriptStatsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ssq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ssq); err != nil {
				return err
			}
		}
	}
	for _, f := range ssq.ctx.Fields {
		if !scriptstats.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ssq.path != nil {
		prev, err := ssq.path(ctx)
		if err != nil {
			return err
		}
		ssq.sql = prev
	}
	return nil
}

func (ssq *ScriptStatsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ScriptStats, error) {
	var (
		nodes = []*ScriptStats{}
		_spec = ssq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ScriptStats).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ScriptStats{config: ssq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ssq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ssq *ScriptStatsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ssq.querySpec()
	_spec.Node.Columns = ssq.ctx.Fields
	if len(ssq.ctx.Fields) > 0 {
		_spec.Unique = ssq.ctx.Unique != nil && *ssq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ssq.driver, _spec)
}

func (ssq *ScriptStatsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(scriptstats.Table, scriptstats.Columns, sqlgraph.NewFieldSpec(scriptstats.FieldID, field.TypeInt))
	_spec.From = ssq.sql
	if unique := ssq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ssq.path != nil {
		_spec.Unique = true
	}
	if fields := ssq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scriptstats.FieldID)
		for i := range fields {
			if fields[i] != scriptstats.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ssq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ssq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ssq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ssq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ssq *ScriptStatsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ssq.driver.Dialect())
	t1 := builder.Table(scriptstats.Table)
	columns := ssq.ctx.Fields
	if len(columns) == 0 {
		columns = scriptstats.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ssq.sql != nil {
		selector = ssq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ssq.ctx.Unique != nil && *ssq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ssq.predicates {
		p(selector)
	}
	for _, p := range ssq.order {
		p(selector)
	}
	if offset := ssq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ssq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScriptStatsGroupBy is the group-by builder for ScriptStats entities.
type ScriptStatsGroupBy struct {
	selector
	build *ScriptStatsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ssgb *ScriptStatsGroupBy) Aggregate(fns ...AggregateFunc) *ScriptStatsGroupBy {
	ssgb.fns = append(ssgb.fns, fns...)
	return ssgb
}

// Scan applies the selector query and scans the result into the given value.
func (ssgb *ScriptStatsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ssgb.build.ctx, "GroupBy")
	if err := ssgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScriptStatsQuery, *ScriptStatsGroupBy](ctx, ssgb.build, ssgb, ssgb.build.inters, v)
}

func (ssgb *ScriptStatsGroupBy) sqlScan(ctx context.Context, root *ScriptStatsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ssgb.fns))
	for _, fn := range ssgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ssgb.flds)+len(ssgb.fns))
		for _, f := range *ssgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ssgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ssgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ScriptStatsSelect is the builder for selecting fields of ScriptStats entities.
type ScriptStatsSelect struct {
	*ScriptStatsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sss *ScriptStatsSelect) Aggregate(fns ...AggregateFunc) *ScriptStatsSelect {
	sss.fns = append(sss.fns, fns...)
	return sss
}

// Scan applies the selector query and scans the result into the given value.
func (sss *ScriptStatsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sss.ctx, "Select")
	if err := sss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScriptStatsQuery, *ScriptStatsSelect](ctx, sss.ScriptStatsQuery, sss, sss.inters, v)
}

func (sss *ScriptStatsSelect) sqlScan(ctx context.Context, root *ScriptStatsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sss.fns))
	for _, fn := range sss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
