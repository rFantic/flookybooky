// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flookybooky/flight/ent/seat"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Seat is the model entity for the Seat schema.
type Seat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FlightID holds the value of the "flight_id" field.
	FlightID string `json:"flight_id,omitempty"`
	// SeatNumber holds the value of the "seat_number" field.
	SeatNumber   string `json:"seat_number,omitempty"`
	flight_seats *int
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Seat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case seat.FieldID:
			values[i] = new(sql.NullInt64)
		case seat.FieldFlightID, seat.FieldSeatNumber:
			values[i] = new(sql.NullString)
		case seat.ForeignKeys[0]: // flight_seats
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Seat fields.
func (s *Seat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case seat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case seat.FieldFlightID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flight_id", values[i])
			} else if value.Valid {
				s.FlightID = value.String
			}
		case seat.FieldSeatNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seat_number", values[i])
			} else if value.Valid {
				s.SeatNumber = value.String
			}
		case seat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field flight_seats", value)
			} else if value.Valid {
				s.flight_seats = new(int)
				*s.flight_seats = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Seat.
// This includes values selected through modifiers, order, etc.
func (s *Seat) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Seat.
// Note that you need to call Seat.Unwrap() before calling this method if this Seat
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Seat) Update() *SeatUpdateOne {
	return NewSeatClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Seat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Seat) Unwrap() *Seat {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Seat is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Seat) String() string {
	var builder strings.Builder
	builder.WriteString("Seat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("flight_id=")
	builder.WriteString(s.FlightID)
	builder.WriteString(", ")
	builder.WriteString("seat_number=")
	builder.WriteString(s.SeatNumber)
	builder.WriteByte(')')
	return builder.String()
}

// Seats is a parsable slice of Seat.
type Seats []*Seat