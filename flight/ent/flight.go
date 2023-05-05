// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flookybooky/flight/ent/flight"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Flight is the model entity for the Flight schema.
type Flight struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// FromID holds the value of the "from_id" field.
	FromID string `json:"from_id,omitempty"`
	// ToID holds the value of the "to_id" field.
	ToID string `json:"to_id,omitempty"`
	// Start holds the value of the "start" field.
	Start time.Time `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End time.Time `json:"end,omitempty"`
	// AvailableSlots holds the value of the "available_slots" field.
	AvailableSlots int `json:"available_slots,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlightQuery when eager-loading is set.
	Edges        FlightEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FlightEdges holds the relations/edges for other nodes in the graph.
type FlightEdges struct {
	// Seats holds the value of the seats edge.
	Seats []*Seat `json:"seats,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SeatsOrErr returns the Seats value or an error if the edge
// was not loaded in eager-loading.
func (e FlightEdges) SeatsOrErr() ([]*Seat, error) {
	if e.loadedTypes[0] {
		return e.Seats, nil
	}
	return nil, &NotLoadedError{edge: "seats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Flight) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case flight.FieldID, flight.FieldAvailableSlots:
			values[i] = new(sql.NullInt64)
		case flight.FieldName, flight.FieldFromID, flight.FieldToID:
			values[i] = new(sql.NullString)
		case flight.FieldStart, flight.FieldEnd, flight.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Flight fields.
func (f *Flight) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flight.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case flight.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case flight.FieldFromID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from_id", values[i])
			} else if value.Valid {
				f.FromID = value.String
			}
		case flight.FieldToID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to_id", values[i])
			} else if value.Valid {
				f.ToID = value.String
			}
		case flight.FieldStart:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				f.Start = value.Time
			}
		case flight.FieldEnd:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				f.End = value.Time
			}
		case flight.FieldAvailableSlots:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field available_slots", values[i])
			} else if value.Valid {
				f.AvailableSlots = int(value.Int64)
			}
		case flight.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Flight.
// This includes values selected through modifiers, order, etc.
func (f *Flight) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QuerySeats queries the "seats" edge of the Flight entity.
func (f *Flight) QuerySeats() *SeatQuery {
	return NewFlightClient(f.config).QuerySeats(f)
}

// Update returns a builder for updating this Flight.
// Note that you need to call Flight.Unwrap() before calling this method if this Flight
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Flight) Update() *FlightUpdateOne {
	return NewFlightClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Flight entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Flight) Unwrap() *Flight {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Flight is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Flight) String() string {
	var builder strings.Builder
	builder.WriteString("Flight(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("from_id=")
	builder.WriteString(f.FromID)
	builder.WriteString(", ")
	builder.WriteString("to_id=")
	builder.WriteString(f.ToID)
	builder.WriteString(", ")
	builder.WriteString("start=")
	builder.WriteString(f.Start.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(f.End.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("available_slots=")
	builder.WriteString(fmt.Sprintf("%v", f.AvailableSlots))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Flights is a parsable slice of Flight.
type Flights []*Flight
