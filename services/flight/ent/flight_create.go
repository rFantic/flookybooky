// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flookybooky/services/flight/ent/airport"
	"flookybooky/services/flight/ent/flight"
	"flookybooky/services/flight/ent/seat"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FlightCreate is the builder for creating a Flight entity.
type FlightCreate struct {
	config
	mutation *FlightMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (fc *FlightCreate) SetName(s string) *FlightCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetDepartureTime sets the "departure_time" field.
func (fc *FlightCreate) SetDepartureTime(t time.Time) *FlightCreate {
	fc.mutation.SetDepartureTime(t)
	return fc
}

// SetArrivalTime sets the "arrival_time" field.
func (fc *FlightCreate) SetArrivalTime(t time.Time) *FlightCreate {
	fc.mutation.SetArrivalTime(t)
	return fc
}

// SetAvailableSlots sets the "available_slots" field.
func (fc *FlightCreate) SetAvailableSlots(i int) *FlightCreate {
	fc.mutation.SetAvailableSlots(i)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FlightCreate) SetCreatedAt(t time.Time) *FlightCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FlightCreate) SetNillableCreatedAt(t *time.Time) *FlightCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FlightCreate) SetID(u uuid.UUID) *FlightCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FlightCreate) SetNillableID(u *uuid.UUID) *FlightCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// AddSeatIDs adds the "seats" edge to the Seat entity by IDs.
func (fc *FlightCreate) AddSeatIDs(ids ...uuid.UUID) *FlightCreate {
	fc.mutation.AddSeatIDs(ids...)
	return fc
}

// AddSeats adds the "seats" edges to the Seat entity.
func (fc *FlightCreate) AddSeats(s ...*Seat) *FlightCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return fc.AddSeatIDs(ids...)
}

// SetOriginID sets the "origin" edge to the Airport entity by ID.
func (fc *FlightCreate) SetOriginID(id uuid.UUID) *FlightCreate {
	fc.mutation.SetOriginID(id)
	return fc
}

// SetNillableOriginID sets the "origin" edge to the Airport entity by ID if the given value is not nil.
func (fc *FlightCreate) SetNillableOriginID(id *uuid.UUID) *FlightCreate {
	if id != nil {
		fc = fc.SetOriginID(*id)
	}
	return fc
}

// SetOrigin sets the "origin" edge to the Airport entity.
func (fc *FlightCreate) SetOrigin(a *Airport) *FlightCreate {
	return fc.SetOriginID(a.ID)
}

// SetDestinationID sets the "destination" edge to the Airport entity by ID.
func (fc *FlightCreate) SetDestinationID(id uuid.UUID) *FlightCreate {
	fc.mutation.SetDestinationID(id)
	return fc
}

// SetNillableDestinationID sets the "destination" edge to the Airport entity by ID if the given value is not nil.
func (fc *FlightCreate) SetNillableDestinationID(id *uuid.UUID) *FlightCreate {
	if id != nil {
		fc = fc.SetDestinationID(*id)
	}
	return fc
}

// SetDestination sets the "destination" edge to the Airport entity.
func (fc *FlightCreate) SetDestination(a *Airport) *FlightCreate {
	return fc.SetDestinationID(a.ID)
}

// Mutation returns the FlightMutation object of the builder.
func (fc *FlightCreate) Mutation() *FlightMutation {
	return fc.mutation
}

// Save creates the Flight in the database.
func (fc *FlightCreate) Save(ctx context.Context) (*Flight, error) {
	fc.defaults()
	return withHooks[*Flight, FlightMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FlightCreate) SaveX(ctx context.Context) *Flight {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FlightCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FlightCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FlightCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := flight.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := flight.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FlightCreate) check() error {
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Flight.name"`)}
	}
	if _, ok := fc.mutation.DepartureTime(); !ok {
		return &ValidationError{Name: "departure_time", err: errors.New(`ent: missing required field "Flight.departure_time"`)}
	}
	if _, ok := fc.mutation.ArrivalTime(); !ok {
		return &ValidationError{Name: "arrival_time", err: errors.New(`ent: missing required field "Flight.arrival_time"`)}
	}
	if _, ok := fc.mutation.AvailableSlots(); !ok {
		return &ValidationError{Name: "available_slots", err: errors.New(`ent: missing required field "Flight.available_slots"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Flight.created_at"`)}
	}
	return nil
}

func (fc *FlightCreate) sqlSave(ctx context.Context) (*Flight, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
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
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FlightCreate) createSpec() (*Flight, *sqlgraph.CreateSpec) {
	var (
		_node = &Flight{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(flight.Table, sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID))
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(flight.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.DepartureTime(); ok {
		_spec.SetField(flight.FieldDepartureTime, field.TypeTime, value)
		_node.DepartureTime = value
	}
	if value, ok := fc.mutation.ArrivalTime(); ok {
		_spec.SetField(flight.FieldArrivalTime, field.TypeTime, value)
		_node.ArrivalTime = value
	}
	if value, ok := fc.mutation.AvailableSlots(); ok {
		_spec.SetField(flight.FieldAvailableSlots, field.TypeInt, value)
		_node.AvailableSlots = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(flight.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := fc.mutation.SeatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flight.SeatsTable,
			Columns: []string{flight.SeatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(seat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.OriginIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flight.OriginTable,
			Columns: []string{flight.OriginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.airport_origin = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.DestinationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   flight.DestinationTable,
			Columns: []string{flight.DestinationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.airport_destination = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FlightCreateBulk is the builder for creating many Flight entities in bulk.
type FlightCreateBulk struct {
	config
	builders []*FlightCreate
}

// Save creates the Flight entities in the database.
func (fcb *FlightCreateBulk) Save(ctx context.Context) ([]*Flight, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Flight, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlightMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FlightCreateBulk) SaveX(ctx context.Context) []*Flight {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FlightCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FlightCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
