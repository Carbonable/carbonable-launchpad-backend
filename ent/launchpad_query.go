// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carbonable/carbonable-launchpad-backend/ent/launchpad"
	"github.com/carbonable/carbonable-launchpad-backend/ent/predicate"
	"github.com/carbonable/carbonable-launchpad-backend/ent/project"
)

// LaunchpadQuery is the builder for querying Launchpad entities.
type LaunchpadQuery struct {
	config
	ctx         *QueryContext
	order       []launchpad.OrderOption
	inters      []Interceptor
	predicates  []predicate.Launchpad
	withProject *ProjectQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LaunchpadQuery builder.
func (lq *LaunchpadQuery) Where(ps ...predicate.Launchpad) *LaunchpadQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LaunchpadQuery) Limit(limit int) *LaunchpadQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LaunchpadQuery) Offset(offset int) *LaunchpadQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LaunchpadQuery) Unique(unique bool) *LaunchpadQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LaunchpadQuery) Order(o ...launchpad.OrderOption) *LaunchpadQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryProject chains the current query on the "project" edge.
func (lq *LaunchpadQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(launchpad.Table, launchpad.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, launchpad.ProjectTable, launchpad.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Launchpad entity from the query.
// Returns a *NotFoundError when no Launchpad was found.
func (lq *LaunchpadQuery) First(ctx context.Context) (*Launchpad, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{launchpad.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LaunchpadQuery) FirstX(ctx context.Context) *Launchpad {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Launchpad ID from the query.
// Returns a *NotFoundError when no Launchpad ID was found.
func (lq *LaunchpadQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{launchpad.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LaunchpadQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Launchpad entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Launchpad entity is found.
// Returns a *NotFoundError when no Launchpad entities are found.
func (lq *LaunchpadQuery) Only(ctx context.Context) (*Launchpad, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{launchpad.Label}
	default:
		return nil, &NotSingularError{launchpad.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LaunchpadQuery) OnlyX(ctx context.Context) *Launchpad {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Launchpad ID in the query.
// Returns a *NotSingularError when more than one Launchpad ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LaunchpadQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{launchpad.Label}
	default:
		err = &NotSingularError{launchpad.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LaunchpadQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Launchpads.
func (lq *LaunchpadQuery) All(ctx context.Context) ([]*Launchpad, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Launchpad, *LaunchpadQuery]()
	return withInterceptors[[]*Launchpad](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LaunchpadQuery) AllX(ctx context.Context) []*Launchpad {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Launchpad IDs.
func (lq *LaunchpadQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(launchpad.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LaunchpadQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LaunchpadQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LaunchpadQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LaunchpadQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LaunchpadQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LaunchpadQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LaunchpadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LaunchpadQuery) Clone() *LaunchpadQuery {
	if lq == nil {
		return nil
	}
	return &LaunchpadQuery{
		config:      lq.config,
		ctx:         lq.ctx.Clone(),
		order:       append([]launchpad.OrderOption{}, lq.order...),
		inters:      append([]Interceptor{}, lq.inters...),
		predicates:  append([]predicate.Launchpad{}, lq.predicates...),
		withProject: lq.withProject.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LaunchpadQuery) WithProject(opts ...func(*ProjectQuery)) *LaunchpadQuery {
	query := (&ProjectClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withProject = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IsReady bool `json:"is_ready,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Launchpad.Query().
//		GroupBy(launchpad.FieldIsReady).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LaunchpadQuery) GroupBy(field string, fields ...string) *LaunchpadGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LaunchpadGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = launchpad.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IsReady bool `json:"is_ready,omitempty"`
//	}
//
//	client.Launchpad.Query().
//		Select(launchpad.FieldIsReady).
//		Scan(ctx, &v)
func (lq *LaunchpadQuery) Select(fields ...string) *LaunchpadSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LaunchpadSelect{LaunchpadQuery: lq}
	sbuild.label = launchpad.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LaunchpadSelect configured with the given aggregations.
func (lq *LaunchpadQuery) Aggregate(fns ...AggregateFunc) *LaunchpadSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LaunchpadQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !launchpad.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LaunchpadQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Launchpad, error) {
	var (
		nodes       = []*Launchpad{}
		withFKs     = lq.withFKs
		_spec       = lq.querySpec()
		loadedTypes = [1]bool{
			lq.withProject != nil,
		}
	)
	if lq.withProject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, launchpad.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Launchpad).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Launchpad{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withProject; query != nil {
		if err := lq.loadProject(ctx, query, nodes, nil,
			func(n *Launchpad, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LaunchpadQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*Launchpad, init func(*Launchpad), assign func(*Launchpad, *Project)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Launchpad)
	for i := range nodes {
		if nodes[i].project_launchpad == nil {
			continue
		}
		fk := *nodes[i].project_launchpad
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_launchpad" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lq *LaunchpadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LaunchpadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(launchpad.Table, launchpad.Columns, sqlgraph.NewFieldSpec(launchpad.FieldID, field.TypeInt))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, launchpad.FieldID)
		for i := range fields {
			if fields[i] != launchpad.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LaunchpadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(launchpad.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = launchpad.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LaunchpadGroupBy is the group-by builder for Launchpad entities.
type LaunchpadGroupBy struct {
	selector
	build *LaunchpadQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LaunchpadGroupBy) Aggregate(fns ...AggregateFunc) *LaunchpadGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LaunchpadGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LaunchpadQuery, *LaunchpadGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LaunchpadGroupBy) sqlScan(ctx context.Context, root *LaunchpadQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LaunchpadSelect is the builder for selecting fields of Launchpad entities.
type LaunchpadSelect struct {
	*LaunchpadQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LaunchpadSelect) Aggregate(fns ...AggregateFunc) *LaunchpadSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LaunchpadSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LaunchpadQuery, *LaunchpadSelect](ctx, ls.LaunchpadQuery, ls, ls.inters, v)
}

func (ls *LaunchpadSelect) sqlScan(ctx context.Context, root *LaunchpadQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}