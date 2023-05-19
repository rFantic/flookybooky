// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flookybooky/services/booking/ent/booking"
	"flookybooky/services/booking/ent/ticket"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketCreate is the builder for creating a Ticket entity.
type TicketCreate struct {
	config
	mutation *TicketMutation
	hooks    []Hook
}

// SetBookingID sets the "booking_id" field.
func (tc *TicketCreate) SetBookingID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetBookingID(u)
	return tc
}

// SetGoingFlightID sets the "going_flight_id" field.
func (tc *TicketCreate) SetGoingFlightID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetGoingFlightID(u)
	return tc
}

// SetReturnFlightID sets the "return_flight_id" field.
func (tc *TicketCreate) SetReturnFlightID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetReturnFlightID(u)
	return tc
}

// SetNillableReturnFlightID sets the "return_flight_id" field if the given value is not nil.
func (tc *TicketCreate) SetNillableReturnFlightID(u *uuid.UUID) *TicketCreate {
	if u != nil {
		tc.SetReturnFlightID(*u)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TicketCreate) SetStatus(t ticket.Status) *TicketCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetPassengerName sets the "passenger_name" field.
func (tc *TicketCreate) SetPassengerName(s string) *TicketCreate {
	tc.mutation.SetPassengerName(s)
	return tc
}

// SetPassengerLicenseID sets the "passenger_license_id" field.
func (tc *TicketCreate) SetPassengerLicenseID(s string) *TicketCreate {
	tc.mutation.SetPassengerLicenseID(s)
	return tc
}

// SetPassengerEmail sets the "passenger_email" field.
func (tc *TicketCreate) SetPassengerEmail(s string) *TicketCreate {
	tc.mutation.SetPassengerEmail(s)
	return tc
}

// SetSeatNumber sets the "seat_number" field.
func (tc *TicketCreate) SetSeatNumber(s string) *TicketCreate {
	tc.mutation.SetSeatNumber(s)
	return tc
}

// SetClass sets the "class" field.
func (tc *TicketCreate) SetClass(t ticket.Class) *TicketCreate {
	tc.mutation.SetClass(t)
	return tc
}

// SetID sets the "id" field.
func (tc *TicketCreate) SetID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TicketCreate) SetNillableID(u *uuid.UUID) *TicketCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// SetBooking sets the "booking" edge to the Booking entity.
func (tc *TicketCreate) SetBooking(b *Booking) *TicketCreate {
	return tc.SetBookingID(b.ID)
}

// Mutation returns the TicketMutation object of the builder.
func (tc *TicketCreate) Mutation() *TicketMutation {
	return tc.mutation
}

// Save creates the Ticket in the database.
func (tc *TicketCreate) Save(ctx context.Context) (*Ticket, error) {
	tc.defaults()
	return withHooks[*Ticket, TicketMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TicketCreate) SaveX(ctx context.Context) *Ticket {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TicketCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TicketCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TicketCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := ticket.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TicketCreate) check() error {
	if _, ok := tc.mutation.BookingID(); !ok {
		return &ValidationError{Name: "booking_id", err: errors.New(`ent: missing required field "Ticket.booking_id"`)}
	}
	if _, ok := tc.mutation.GoingFlightID(); !ok {
		return &ValidationError{Name: "going_flight_id", err: errors.New(`ent: missing required field "Ticket.going_flight_id"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Ticket.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := ticket.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Ticket.status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.PassengerName(); !ok {
		return &ValidationError{Name: "passenger_name", err: errors.New(`ent: missing required field "Ticket.passenger_name"`)}
	}
	if _, ok := tc.mutation.PassengerLicenseID(); !ok {
		return &ValidationError{Name: "passenger_license_id", err: errors.New(`ent: missing required field "Ticket.passenger_license_id"`)}
	}
	if _, ok := tc.mutation.PassengerEmail(); !ok {
		return &ValidationError{Name: "passenger_email", err: errors.New(`ent: missing required field "Ticket.passenger_email"`)}
	}
	if _, ok := tc.mutation.SeatNumber(); !ok {
		return &ValidationError{Name: "seat_number", err: errors.New(`ent: missing required field "Ticket.seat_number"`)}
	}
	if _, ok := tc.mutation.Class(); !ok {
		return &ValidationError{Name: "class", err: errors.New(`ent: missing required field "Ticket.class"`)}
	}
	if v, ok := tc.mutation.Class(); ok {
		if err := ticket.ClassValidator(v); err != nil {
			return &ValidationError{Name: "class", err: fmt.Errorf(`ent: validator failed for field "Ticket.class": %w`, err)}
		}
	}
	if _, ok := tc.mutation.BookingID(); !ok {
		return &ValidationError{Name: "booking", err: errors.New(`ent: missing required edge "Ticket.booking"`)}
	}
	return nil
}

func (tc *TicketCreate) sqlSave(ctx context.Context) (*Ticket, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TicketCreate) createSpec() (*Ticket, *sqlgraph.CreateSpec) {
	var (
		_node = &Ticket{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(ticket.Table, sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.GoingFlightID(); ok {
		_spec.SetField(ticket.FieldGoingFlightID, field.TypeUUID, value)
		_node.GoingFlightID = value
	}
	if value, ok := tc.mutation.ReturnFlightID(); ok {
		_spec.SetField(ticket.FieldReturnFlightID, field.TypeUUID, value)
		_node.ReturnFlightID = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(ticket.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.PassengerName(); ok {
		_spec.SetField(ticket.FieldPassengerName, field.TypeString, value)
		_node.PassengerName = value
	}
	if value, ok := tc.mutation.PassengerLicenseID(); ok {
		_spec.SetField(ticket.FieldPassengerLicenseID, field.TypeString, value)
		_node.PassengerLicenseID = value
	}
	if value, ok := tc.mutation.PassengerEmail(); ok {
		_spec.SetField(ticket.FieldPassengerEmail, field.TypeString, value)
		_node.PassengerEmail = value
	}
	if value, ok := tc.mutation.SeatNumber(); ok {
		_spec.SetField(ticket.FieldSeatNumber, field.TypeString, value)
		_node.SeatNumber = value
	}
	if value, ok := tc.mutation.Class(); ok {
		_spec.SetField(ticket.FieldClass, field.TypeEnum, value)
		_node.Class = value
	}
	if nodes := tc.mutation.BookingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.BookingTable,
			Columns: []string{ticket.BookingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BookingID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TicketCreateBulk is the builder for creating many Ticket entities in bulk.
type TicketCreateBulk struct {
	config
	builders []*TicketCreate
}

// Save creates the Ticket entities in the database.
func (tcb *TicketCreateBulk) Save(ctx context.Context) ([]*Ticket, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Ticket, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TicketMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TicketCreateBulk) SaveX(ctx context.Context) []*Ticket {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TicketCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TicketCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
