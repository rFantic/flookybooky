// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flookybooky/services/booking/ent/booking"
	"flookybooky/services/booking/ent/predicate"
	"flookybooky/services/booking/ent/ticket"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// Where appends a list predicates to the BookingUpdate builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetCustomerID sets the "customer_id" field.
func (bu *BookingUpdate) SetCustomerID(u uuid.UUID) *BookingUpdate {
	bu.mutation.SetCustomerID(u)
	return bu
}

// SetGoingFlightID sets the "going_flight_id" field.
func (bu *BookingUpdate) SetGoingFlightID(u uuid.UUID) *BookingUpdate {
	bu.mutation.SetGoingFlightID(u)
	return bu
}

// SetReturnFlightID sets the "return_flight_id" field.
func (bu *BookingUpdate) SetReturnFlightID(u uuid.UUID) *BookingUpdate {
	bu.mutation.SetReturnFlightID(u)
	return bu
}

// SetNillableReturnFlightID sets the "return_flight_id" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableReturnFlightID(u *uuid.UUID) *BookingUpdate {
	if u != nil {
		bu.SetReturnFlightID(*u)
	}
	return bu
}

// ClearReturnFlightID clears the value of the "return_flight_id" field.
func (bu *BookingUpdate) ClearReturnFlightID() *BookingUpdate {
	bu.mutation.ClearReturnFlightID()
	return bu
}

// SetStatus sets the "status" field.
func (bu *BookingUpdate) SetStatus(b booking.Status) *BookingUpdate {
	bu.mutation.SetStatus(b)
	return bu
}

// AddTicketIDs adds the "ticket" edge to the Ticket entity by IDs.
func (bu *BookingUpdate) AddTicketIDs(ids ...uuid.UUID) *BookingUpdate {
	bu.mutation.AddTicketIDs(ids...)
	return bu
}

// AddTicket adds the "ticket" edges to the Ticket entity.
func (bu *BookingUpdate) AddTicket(t ...*Ticket) *BookingUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTicketIDs(ids...)
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// ClearTicket clears all "ticket" edges to the Ticket entity.
func (bu *BookingUpdate) ClearTicket() *BookingUpdate {
	bu.mutation.ClearTicket()
	return bu
}

// RemoveTicketIDs removes the "ticket" edge to Ticket entities by IDs.
func (bu *BookingUpdate) RemoveTicketIDs(ids ...uuid.UUID) *BookingUpdate {
	bu.mutation.RemoveTicketIDs(ids...)
	return bu
}

// RemoveTicket removes "ticket" edges to Ticket entities.
func (bu *BookingUpdate) RemoveTicket(t ...*Ticket) *BookingUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTicketIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BookingMutation](ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookingUpdate) check() error {
	if v, ok := bu.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	return nil
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeUUID))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.CustomerID(); ok {
		_spec.SetField(booking.FieldCustomerID, field.TypeUUID, value)
	}
	if value, ok := bu.mutation.GoingFlightID(); ok {
		_spec.SetField(booking.FieldGoingFlightID, field.TypeUUID, value)
	}
	if value, ok := bu.mutation.ReturnFlightID(); ok {
		_spec.SetField(booking.FieldReturnFlightID, field.TypeUUID, value)
	}
	if bu.mutation.ReturnFlightIDCleared() {
		_spec.ClearField(booking.FieldReturnFlightID, field.TypeUUID)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
	}
	if bu.mutation.TicketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.TicketTable,
			Columns: []string{booking.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTicketIDs(); len(nodes) > 0 && !bu.mutation.TicketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.TicketTable,
			Columns: []string{booking.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TicketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.TicketTable,
			Columns: []string{booking.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookingMutation
}

// SetCustomerID sets the "customer_id" field.
func (buo *BookingUpdateOne) SetCustomerID(u uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetCustomerID(u)
	return buo
}

// SetGoingFlightID sets the "going_flight_id" field.
func (buo *BookingUpdateOne) SetGoingFlightID(u uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetGoingFlightID(u)
	return buo
}

// SetReturnFlightID sets the "return_flight_id" field.
func (buo *BookingUpdateOne) SetReturnFlightID(u uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetReturnFlightID(u)
	return buo
}

// SetNillableReturnFlightID sets the "return_flight_id" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableReturnFlightID(u *uuid.UUID) *BookingUpdateOne {
	if u != nil {
		buo.SetReturnFlightID(*u)
	}
	return buo
}

// ClearReturnFlightID clears the value of the "return_flight_id" field.
func (buo *BookingUpdateOne) ClearReturnFlightID() *BookingUpdateOne {
	buo.mutation.ClearReturnFlightID()
	return buo
}

// SetStatus sets the "status" field.
func (buo *BookingUpdateOne) SetStatus(b booking.Status) *BookingUpdateOne {
	buo.mutation.SetStatus(b)
	return buo
}

// AddTicketIDs adds the "ticket" edge to the Ticket entity by IDs.
func (buo *BookingUpdateOne) AddTicketIDs(ids ...uuid.UUID) *BookingUpdateOne {
	buo.mutation.AddTicketIDs(ids...)
	return buo
}

// AddTicket adds the "ticket" edges to the Ticket entity.
func (buo *BookingUpdateOne) AddTicket(t ...*Ticket) *BookingUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTicketIDs(ids...)
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// ClearTicket clears all "ticket" edges to the Ticket entity.
func (buo *BookingUpdateOne) ClearTicket() *BookingUpdateOne {
	buo.mutation.ClearTicket()
	return buo
}

// RemoveTicketIDs removes the "ticket" edge to Ticket entities by IDs.
func (buo *BookingUpdateOne) RemoveTicketIDs(ids ...uuid.UUID) *BookingUpdateOne {
	buo.mutation.RemoveTicketIDs(ids...)
	return buo
}

// RemoveTicket removes "ticket" edges to Ticket entities.
func (buo *BookingUpdateOne) RemoveTicket(t ...*Ticket) *BookingUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTicketIDs(ids...)
}

// Where appends a list predicates to the BookingUpdate builder.
func (buo *BookingUpdateOne) Where(ps ...predicate.Booking) *BookingUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookingUpdateOne) Select(field string, fields ...string) *BookingUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Booking entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	return withHooks[*Booking, BookingMutation](ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookingUpdateOne) check() error {
	if v, ok := buo.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	return nil
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (_node *Booking, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeUUID))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Booking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for _, f := range fields {
			if !booking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.CustomerID(); ok {
		_spec.SetField(booking.FieldCustomerID, field.TypeUUID, value)
	}
	if value, ok := buo.mutation.GoingFlightID(); ok {
		_spec.SetField(booking.FieldGoingFlightID, field.TypeUUID, value)
	}
	if value, ok := buo.mutation.ReturnFlightID(); ok {
		_spec.SetField(booking.FieldReturnFlightID, field.TypeUUID, value)
	}
	if buo.mutation.ReturnFlightIDCleared() {
		_spec.ClearField(booking.FieldReturnFlightID, field.TypeUUID)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
	}
	if buo.mutation.TicketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.TicketTable,
			Columns: []string{booking.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTicketIDs(); len(nodes) > 0 && !buo.mutation.TicketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.TicketTable,
			Columns: []string{booking.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TicketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.TicketTable,
			Columns: []string{booking.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Booking{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
