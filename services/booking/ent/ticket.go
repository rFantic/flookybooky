// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flookybooky/services/booking/ent/ticket"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Ticket is the model entity for the Ticket schema.
type Ticket struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ticket) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ticket.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ticket fields.
func (t *Ticket) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ticket.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Ticket.
// This includes values selected through modifiers, order, etc.
func (t *Ticket) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Ticket.
// Note that you need to call Ticket.Unwrap() before calling this method if this Ticket
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Ticket) Update() *TicketUpdateOne {
	return NewTicketClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Ticket entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Ticket) Unwrap() *Ticket {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ticket is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Ticket) String() string {
	var builder strings.Builder
	builder.WriteString("Ticket(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Tickets is a parsable slice of Ticket.
type Tickets []*Ticket
