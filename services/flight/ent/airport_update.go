// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flookybooky/services/flight/ent/airport"
	"flookybooky/services/flight/ent/flight"
	"flookybooky/services/flight/ent/predicate"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AirportUpdate is the builder for updating Airport entities.
type AirportUpdate struct {
	config
	hooks    []Hook
	mutation *AirportMutation
}

// Where appends a list predicates to the AirportUpdate builder.
func (au *AirportUpdate) Where(ps ...predicate.Airport) *AirportUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AirportUpdate) SetName(s string) *AirportUpdate {
	au.mutation.SetName(s)
	return au
}

// SetAddress sets the "address" field.
func (au *AirportUpdate) SetAddress(s string) *AirportUpdate {
	au.mutation.SetAddress(s)
	return au
}

// AddOriginIDs adds the "origin" edge to the Flight entity by IDs.
func (au *AirportUpdate) AddOriginIDs(ids ...uuid.UUID) *AirportUpdate {
	au.mutation.AddOriginIDs(ids...)
	return au
}

// AddOrigin adds the "origin" edges to the Flight entity.
func (au *AirportUpdate) AddOrigin(f ...*Flight) *AirportUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.AddOriginIDs(ids...)
}

// AddDestinationIDs adds the "destination" edge to the Flight entity by IDs.
func (au *AirportUpdate) AddDestinationIDs(ids ...uuid.UUID) *AirportUpdate {
	au.mutation.AddDestinationIDs(ids...)
	return au
}

// AddDestination adds the "destination" edges to the Flight entity.
func (au *AirportUpdate) AddDestination(f ...*Flight) *AirportUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.AddDestinationIDs(ids...)
}

// Mutation returns the AirportMutation object of the builder.
func (au *AirportUpdate) Mutation() *AirportMutation {
	return au.mutation
}

// ClearOrigin clears all "origin" edges to the Flight entity.
func (au *AirportUpdate) ClearOrigin() *AirportUpdate {
	au.mutation.ClearOrigin()
	return au
}

// RemoveOriginIDs removes the "origin" edge to Flight entities by IDs.
func (au *AirportUpdate) RemoveOriginIDs(ids ...uuid.UUID) *AirportUpdate {
	au.mutation.RemoveOriginIDs(ids...)
	return au
}

// RemoveOrigin removes "origin" edges to Flight entities.
func (au *AirportUpdate) RemoveOrigin(f ...*Flight) *AirportUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.RemoveOriginIDs(ids...)
}

// ClearDestination clears all "destination" edges to the Flight entity.
func (au *AirportUpdate) ClearDestination() *AirportUpdate {
	au.mutation.ClearDestination()
	return au
}

// RemoveDestinationIDs removes the "destination" edge to Flight entities by IDs.
func (au *AirportUpdate) RemoveDestinationIDs(ids ...uuid.UUID) *AirportUpdate {
	au.mutation.RemoveDestinationIDs(ids...)
	return au
}

// RemoveDestination removes "destination" edges to Flight entities.
func (au *AirportUpdate) RemoveDestination(f ...*Flight) *AirportUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.RemoveDestinationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AirportUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AirportMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AirportUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AirportUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AirportUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AirportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Address(); ok {
		_spec.SetField(airport.FieldAddress, field.TypeString, value)
	}
	if au.mutation.OriginCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.OriginTable,
			Columns: []string{airport.OriginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedOriginIDs(); len(nodes) > 0 && !au.mutation.OriginCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.OriginTable,
			Columns: []string{airport.OriginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.OriginIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.OriginTable,
			Columns: []string{airport.OriginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.DestinationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestinationTable,
			Columns: []string{airport.DestinationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedDestinationIDs(); len(nodes) > 0 && !au.mutation.DestinationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestinationTable,
			Columns: []string{airport.DestinationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.DestinationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestinationTable,
			Columns: []string{airport.DestinationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AirportUpdateOne is the builder for updating a single Airport entity.
type AirportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AirportMutation
}

// SetName sets the "name" field.
func (auo *AirportUpdateOne) SetName(s string) *AirportUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetAddress sets the "address" field.
func (auo *AirportUpdateOne) SetAddress(s string) *AirportUpdateOne {
	auo.mutation.SetAddress(s)
	return auo
}

// AddOriginIDs adds the "origin" edge to the Flight entity by IDs.
func (auo *AirportUpdateOne) AddOriginIDs(ids ...uuid.UUID) *AirportUpdateOne {
	auo.mutation.AddOriginIDs(ids...)
	return auo
}

// AddOrigin adds the "origin" edges to the Flight entity.
func (auo *AirportUpdateOne) AddOrigin(f ...*Flight) *AirportUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.AddOriginIDs(ids...)
}

// AddDestinationIDs adds the "destination" edge to the Flight entity by IDs.
func (auo *AirportUpdateOne) AddDestinationIDs(ids ...uuid.UUID) *AirportUpdateOne {
	auo.mutation.AddDestinationIDs(ids...)
	return auo
}

// AddDestination adds the "destination" edges to the Flight entity.
func (auo *AirportUpdateOne) AddDestination(f ...*Flight) *AirportUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.AddDestinationIDs(ids...)
}

// Mutation returns the AirportMutation object of the builder.
func (auo *AirportUpdateOne) Mutation() *AirportMutation {
	return auo.mutation
}

// ClearOrigin clears all "origin" edges to the Flight entity.
func (auo *AirportUpdateOne) ClearOrigin() *AirportUpdateOne {
	auo.mutation.ClearOrigin()
	return auo
}

// RemoveOriginIDs removes the "origin" edge to Flight entities by IDs.
func (auo *AirportUpdateOne) RemoveOriginIDs(ids ...uuid.UUID) *AirportUpdateOne {
	auo.mutation.RemoveOriginIDs(ids...)
	return auo
}

// RemoveOrigin removes "origin" edges to Flight entities.
func (auo *AirportUpdateOne) RemoveOrigin(f ...*Flight) *AirportUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.RemoveOriginIDs(ids...)
}

// ClearDestination clears all "destination" edges to the Flight entity.
func (auo *AirportUpdateOne) ClearDestination() *AirportUpdateOne {
	auo.mutation.ClearDestination()
	return auo
}

// RemoveDestinationIDs removes the "destination" edge to Flight entities by IDs.
func (auo *AirportUpdateOne) RemoveDestinationIDs(ids ...uuid.UUID) *AirportUpdateOne {
	auo.mutation.RemoveDestinationIDs(ids...)
	return auo
}

// RemoveDestination removes "destination" edges to Flight entities.
func (auo *AirportUpdateOne) RemoveDestination(f ...*Flight) *AirportUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.RemoveDestinationIDs(ids...)
}

// Where appends a list predicates to the AirportUpdate builder.
func (auo *AirportUpdateOne) Where(ps ...predicate.Airport) *AirportUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AirportUpdateOne) Select(field string, fields ...string) *AirportUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Airport entity.
func (auo *AirportUpdateOne) Save(ctx context.Context) (*Airport, error) {
	return withHooks[*Airport, AirportMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AirportUpdateOne) SaveX(ctx context.Context) *Airport {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AirportUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AirportUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AirportUpdateOne) sqlSave(ctx context.Context) (_node *Airport, err error) {
	_spec := sqlgraph.NewUpdateSpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Airport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, airport.FieldID)
		for _, f := range fields {
			if !airport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != airport.FieldID {
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
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Address(); ok {
		_spec.SetField(airport.FieldAddress, field.TypeString, value)
	}
	if auo.mutation.OriginCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.OriginTable,
			Columns: []string{airport.OriginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedOriginIDs(); len(nodes) > 0 && !auo.mutation.OriginCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.OriginTable,
			Columns: []string{airport.OriginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.OriginIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.OriginTable,
			Columns: []string{airport.OriginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.DestinationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestinationTable,
			Columns: []string{airport.DestinationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedDestinationIDs(); len(nodes) > 0 && !auo.mutation.DestinationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestinationTable,
			Columns: []string{airport.DestinationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.DestinationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestinationTable,
			Columns: []string{airport.DestinationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Airport{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
