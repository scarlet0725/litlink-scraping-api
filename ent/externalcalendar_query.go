// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scarlet0725/prism-api/ent/externalcalendar"
	"github.com/scarlet0725/prism-api/ent/predicate"
	"github.com/scarlet0725/prism-api/ent/user"
)

// ExternalCalendarQuery is the builder for querying ExternalCalendar entities.
type ExternalCalendarQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.ExternalCalendar
	withUser   *UserQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExternalCalendarQuery builder.
func (ecq *ExternalCalendarQuery) Where(ps ...predicate.ExternalCalendar) *ExternalCalendarQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *ExternalCalendarQuery) Limit(limit int) *ExternalCalendarQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *ExternalCalendarQuery) Offset(offset int) *ExternalCalendarQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *ExternalCalendarQuery) Unique(unique bool) *ExternalCalendarQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *ExternalCalendarQuery) Order(o ...OrderFunc) *ExternalCalendarQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// QueryUser chains the current query on the "user" edge.
func (ecq *ExternalCalendarQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(externalcalendar.Table, externalcalendar.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, externalcalendar.UserTable, externalcalendar.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ExternalCalendar entity from the query.
// Returns a *NotFoundError when no ExternalCalendar was found.
func (ecq *ExternalCalendarQuery) First(ctx context.Context) (*ExternalCalendar, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{externalcalendar.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *ExternalCalendarQuery) FirstX(ctx context.Context) *ExternalCalendar {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExternalCalendar ID from the query.
// Returns a *NotFoundError when no ExternalCalendar ID was found.
func (ecq *ExternalCalendarQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{externalcalendar.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *ExternalCalendarQuery) FirstIDX(ctx context.Context) int {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExternalCalendar entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExternalCalendar entity is found.
// Returns a *NotFoundError when no ExternalCalendar entities are found.
func (ecq *ExternalCalendarQuery) Only(ctx context.Context) (*ExternalCalendar, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{externalcalendar.Label}
	default:
		return nil, &NotSingularError{externalcalendar.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *ExternalCalendarQuery) OnlyX(ctx context.Context) *ExternalCalendar {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExternalCalendar ID in the query.
// Returns a *NotSingularError when more than one ExternalCalendar ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *ExternalCalendarQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{externalcalendar.Label}
	default:
		err = &NotSingularError{externalcalendar.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *ExternalCalendarQuery) OnlyIDX(ctx context.Context) int {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExternalCalendars.
func (ecq *ExternalCalendarQuery) All(ctx context.Context) ([]*ExternalCalendar, error) {
	ctx = setContextOp(ctx, ecq.ctx, "All")
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExternalCalendar, *ExternalCalendarQuery]()
	return withInterceptors[[]*ExternalCalendar](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *ExternalCalendarQuery) AllX(ctx context.Context) []*ExternalCalendar {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExternalCalendar IDs.
func (ecq *ExternalCalendarQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, "IDs")
	if err = ecq.Select(externalcalendar.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *ExternalCalendarQuery) IDsX(ctx context.Context) []int {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *ExternalCalendarQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Count")
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*ExternalCalendarQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *ExternalCalendarQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *ExternalCalendarQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Exist")
	switch _, err := ecq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *ExternalCalendarQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExternalCalendarQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *ExternalCalendarQuery) Clone() *ExternalCalendarQuery {
	if ecq == nil {
		return nil
	}
	return &ExternalCalendarQuery{
		config:     ecq.config,
		ctx:        ecq.ctx.Clone(),
		order:      append([]OrderFunc{}, ecq.order...),
		inters:     append([]Interceptor{}, ecq.inters...),
		predicates: append([]predicate.ExternalCalendar{}, ecq.predicates...),
		withUser:   ecq.withUser.Clone(),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *ExternalCalendarQuery) WithUser(opts ...func(*UserQuery)) *ExternalCalendarQuery {
	query := (&UserClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withUser = query
	return ecq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExternalCalendar.Query().
//		GroupBy(externalcalendar.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *ExternalCalendarQuery) GroupBy(field string, fields ...string) *ExternalCalendarGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExternalCalendarGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = externalcalendar.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ExternalCalendar.Query().
//		Select(externalcalendar.FieldName).
//		Scan(ctx, &v)
func (ecq *ExternalCalendarQuery) Select(fields ...string) *ExternalCalendarSelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &ExternalCalendarSelect{ExternalCalendarQuery: ecq}
	sbuild.label = externalcalendar.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExternalCalendarSelect configured with the given aggregations.
func (ecq *ExternalCalendarQuery) Aggregate(fns ...AggregateFunc) *ExternalCalendarSelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *ExternalCalendarQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ecq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ecq); err != nil {
				return err
			}
		}
	}
	for _, f := range ecq.ctx.Fields {
		if !externalcalendar.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	return nil
}

func (ecq *ExternalCalendarQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExternalCalendar, error) {
	var (
		nodes       = []*ExternalCalendar{}
		withFKs     = ecq.withFKs
		_spec       = ecq.querySpec()
		loadedTypes = [1]bool{
			ecq.withUser != nil,
		}
	)
	if ecq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, externalcalendar.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExternalCalendar).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExternalCalendar{config: ecq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ecq.modifiers) > 0 {
		_spec.Modifiers = ecq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ecq.withUser; query != nil {
		if err := ecq.loadUser(ctx, query, nodes, nil,
			func(n *ExternalCalendar, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ecq *ExternalCalendarQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ExternalCalendar, init func(*ExternalCalendar), assign func(*ExternalCalendar, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ExternalCalendar)
	for i := range nodes {
		if nodes[i].user_id == nil {
			continue
		}
		fk := *nodes[i].user_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ecq *ExternalCalendarQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	if len(ecq.modifiers) > 0 {
		_spec.Modifiers = ecq.modifiers
	}
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *ExternalCalendarQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(externalcalendar.Table, externalcalendar.Columns, sqlgraph.NewFieldSpec(externalcalendar.FieldID, field.TypeInt))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, externalcalendar.FieldID)
		for i := range fields {
			if fields[i] != externalcalendar.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *ExternalCalendarQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(externalcalendar.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = externalcalendar.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.ctx.Unique != nil && *ecq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ecq.modifiers {
		m(selector)
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ecq *ExternalCalendarQuery) Modify(modifiers ...func(s *sql.Selector)) *ExternalCalendarSelect {
	ecq.modifiers = append(ecq.modifiers, modifiers...)
	return ecq.Select()
}

// ExternalCalendarGroupBy is the group-by builder for ExternalCalendar entities.
type ExternalCalendarGroupBy struct {
	selector
	build *ExternalCalendarQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *ExternalCalendarGroupBy) Aggregate(fns ...AggregateFunc) *ExternalCalendarGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *ExternalCalendarGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, "GroupBy")
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExternalCalendarQuery, *ExternalCalendarGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *ExternalCalendarGroupBy) sqlScan(ctx context.Context, root *ExternalCalendarQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ecgb.flds)+len(ecgb.fns))
		for _, f := range *ecgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ecgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExternalCalendarSelect is the builder for selecting fields of ExternalCalendar entities.
type ExternalCalendarSelect struct {
	*ExternalCalendarQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *ExternalCalendarSelect) Aggregate(fns ...AggregateFunc) *ExternalCalendarSelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *ExternalCalendarSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, "Select")
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExternalCalendarQuery, *ExternalCalendarSelect](ctx, ecs.ExternalCalendarQuery, ecs, ecs.inters, v)
}

func (ecs *ExternalCalendarSelect) sqlScan(ctx context.Context, root *ExternalCalendarQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ecs.fns))
	for _, fn := range ecs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ecs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ecs *ExternalCalendarSelect) Modify(modifiers ...func(s *sql.Selector)) *ExternalCalendarSelect {
	ecs.modifiers = append(ecs.modifiers, modifiers...)
	return ecs
}
