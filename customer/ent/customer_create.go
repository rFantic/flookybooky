// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flookybooky/customer/ent/customer"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CustomerCreate) SetName(s string) *CustomerCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetAddress sets the "address" field.
func (cc *CustomerCreate) SetAddress(s string) *CustomerCreate {
	cc.mutation.SetAddress(s)
	return cc
}

// SetLicenseID sets the "license_id" field.
func (cc *CustomerCreate) SetLicenseID(s string) *CustomerCreate {
	cc.mutation.SetLicenseID(s)
	return cc
}

// SetPhoneNumber sets the "phone_number" field.
func (cc *CustomerCreate) SetPhoneNumber(s string) *CustomerCreate {
	cc.mutation.SetPhoneNumber(s)
	return cc
}

// SetTimestamp sets the "timestamp" field.
func (cc *CustomerCreate) SetTimestamp(t time.Time) *CustomerCreate {
	cc.mutation.SetTimestamp(t)
	return cc
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableTimestamp(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetTimestamp(*t)
	}
	return cc
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	cc.defaults()
	return withHooks[*Customer, CustomerMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CustomerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CustomerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CustomerCreate) defaults() {
	if _, ok := cc.mutation.Timestamp(); !ok {
		v := customer.DefaultTimestamp
		cc.mutation.SetTimestamp(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Customer.name"`)}
	}
	if _, ok := cc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Customer.address"`)}
	}
	if _, ok := cc.mutation.LicenseID(); !ok {
		return &ValidationError{Name: "license_id", err: errors.New(`ent: missing required field "Customer.license_id"`)}
	}
	if v, ok := cc.mutation.LicenseID(); ok {
		if err := customer.LicenseIDValidator(v); err != nil {
			return &ValidationError{Name: "license_id", err: fmt.Errorf(`ent: validator failed for field "Customer.license_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`ent: missing required field "Customer.phone_number"`)}
	}
	if v, ok := cc.mutation.PhoneNumber(); ok {
		if err := customer.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Customer.phone_number": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Customer.timestamp"`)}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(customer.Table, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Address(); ok {
		_spec.SetField(customer.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := cc.mutation.LicenseID(); ok {
		_spec.SetField(customer.FieldLicenseID, field.TypeString, value)
		_node.LicenseID = value
	}
	if value, ok := cc.mutation.PhoneNumber(); ok {
		_spec.SetField(customer.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := cc.mutation.Timestamp(); ok {
		_spec.SetField(customer.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CustomerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
