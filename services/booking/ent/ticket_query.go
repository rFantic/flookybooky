// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"flookybooky/services/booking/ent/booking"
	"flookybooky/services/booking/ent/predicate"
	"flookybooky/services/booking/ent/ticket"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketQuery is the builder for querying Ticket entities.
type TicketQuery struct {
	config
	ctx        *QueryContext
	order      []ticket.OrderOption
	inters     []Interceptor
	predicates []predicate.Ticket
	withGoing  *BookingQuery
	withReturn *BookingQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TicketQuery builder.
func (tq *TicketQuery) Where(ps ...predicate.Ticket) *TicketQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TicketQuery) Limit(limit int) *TicketQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TicketQuery) Offset(offset int) *TicketQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TicketQuery) Unique(unique bool) *TicketQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TicketQuery) Order(o ...ticket.OrderOption) *TicketQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryGoing chains the current query on the "going" edge.
func (tq *TicketQuery) QueryGoing() *BookingQuery {
	query := (&BookingClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ticket.GoingTable, ticket.GoingColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReturn chains the current query on the "return" edge.
func (tq *TicketQuery) QueryReturn() *BookingQuery {
	query := (&BookingClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ticket.ReturnTable, ticket.ReturnColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Ticket entity from the query.
// Returns a *NotFoundError when no Ticket was found.
func (tq *TicketQuery) First(ctx context.Context) (*Ticket, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ticket.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TicketQuery) FirstX(ctx context.Context) *Ticket {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Ticket ID from the query.
// Returns a *NotFoundError when no Ticket ID was found.
func (tq *TicketQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ticket.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TicketQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Ticket entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Ticket entity is found.
// Returns a *NotFoundError when no Ticket entities are found.
func (tq *TicketQuery) Only(ctx context.Context) (*Ticket, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ticket.Label}
	default:
		return nil, &NotSingularError{ticket.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TicketQuery) OnlyX(ctx context.Context) *Ticket {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Ticket ID in the query.
// Returns a *NotSingularError when more than one Ticket ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TicketQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ticket.Label}
	default:
		err = &NotSingularError{ticket.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TicketQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tickets.
func (tq *TicketQuery) All(ctx context.Context) ([]*Ticket, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Ticket, *TicketQuery]()
	return withInterceptors[[]*Ticket](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TicketQuery) AllX(ctx context.Context) []*Ticket {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Ticket IDs.
func (tq *TicketQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(ticket.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TicketQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TicketQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TicketQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TicketQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TicketQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TicketQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TicketQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TicketQuery) Clone() *TicketQuery {
	if tq == nil {
		return nil
	}
	return &TicketQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]ticket.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Ticket{}, tq.predicates...),
		withGoing:  tq.withGoing.Clone(),
		withReturn: tq.withReturn.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithGoing tells the query-builder to eager-load the nodes that are connected to
// the "going" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TicketQuery) WithGoing(opts ...func(*BookingQuery)) *TicketQuery {
	query := (&BookingClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withGoing = query
	return tq
}

// WithReturn tells the query-builder to eager-load the nodes that are connected to
// the "return" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TicketQuery) WithReturn(opts ...func(*BookingQuery)) *TicketQuery {
	query := (&BookingClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withReturn = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FlightID uuid.UUID `json:"flight_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Ticket.Query().
//		GroupBy(ticket.FieldFlightID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TicketQuery) GroupBy(field string, fields ...string) *TicketGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TicketGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = ticket.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FlightID uuid.UUID `json:"flight_id,omitempty"`
//	}
//
//	client.Ticket.Query().
//		Select(ticket.FieldFlightID).
//		Scan(ctx, &v)
func (tq *TicketQuery) Select(fields ...string) *TicketSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TicketSelect{TicketQuery: tq}
	sbuild.label = ticket.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TicketSelect configured with the given aggregations.
func (tq *TicketQuery) Aggregate(fns ...AggregateFunc) *TicketSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TicketQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !ticket.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TicketQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Ticket, error) {
	var (
		nodes       = []*Ticket{}
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withGoing != nil,
			tq.withReturn != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Ticket).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Ticket{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withGoing; query != nil {
		if err := tq.loadGoing(ctx, query, nodes,
			func(n *Ticket) { n.Edges.Going = []*Booking{} },
			func(n *Ticket, e *Booking) { n.Edges.Going = append(n.Edges.Going, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withReturn; query != nil {
		if err := tq.loadReturn(ctx, query, nodes,
			func(n *Ticket) { n.Edges.Return = []*Booking{} },
			func(n *Ticket, e *Booking) { n.Edges.Return = append(n.Edges.Return, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TicketQuery) loadGoing(ctx context.Context, query *BookingQuery, nodes []*Ticket, init func(*Ticket), assign func(*Ticket, *Booking)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Ticket)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(booking.FieldGoingTicketID)
	}
	query.Where(predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(ticket.GoingColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.GoingTicketID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "going_ticket_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TicketQuery) loadReturn(ctx context.Context, query *BookingQuery, nodes []*Ticket, init func(*Ticket), assign func(*Ticket, *Booking)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Ticket)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(booking.FieldReturnTicketID)
	}
	query.Where(predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(ticket.ReturnColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReturnTicketID
		if fk == nil {
			return fmt.Errorf(`foreign-key "return_ticket_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "return_ticket_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TicketQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TicketQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ticket.Table, ticket.Columns, sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ticket.FieldID)
		for i := range fields {
			if fields[i] != ticket.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TicketQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(ticket.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = ticket.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TicketGroupBy is the group-by builder for Ticket entities.
type TicketGroupBy struct {
	selector
	build *TicketQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TicketGroupBy) Aggregate(fns ...AggregateFunc) *TicketGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TicketGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TicketQuery, *TicketGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TicketGroupBy) sqlScan(ctx context.Context, root *TicketQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TicketSelect is the builder for selecting fields of Ticket entities.
type TicketSelect struct {
	*TicketQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TicketSelect) Aggregate(fns ...AggregateFunc) *TicketSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TicketSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TicketQuery, *TicketSelect](ctx, ts.TicketQuery, ts, ts.inters, v)
}

func (ts *TicketSelect) sqlScan(ctx context.Context, root *TicketQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
