// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flookybooky/services/flight/ent/airport"
	"flookybooky/services/flight/ent/flight"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Flight is the model entity for the Flight schema.
type Flight struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DepartureTime holds the value of the "departure_time" field.
	DepartureTime time.Time `json:"departure_time,omitempty"`
	// ArrivalTime holds the value of the "arrival_time" field.
	ArrivalTime time.Time `json:"arrival_time,omitempty"`
	// AvailableSlots holds the value of the "available_slots" field.
	AvailableSlots int `json:"available_slots,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlightQuery when eager-loading is set.
	Edges               FlightEdges `json:"edges"`
	airport_origin      *uuid.UUID
	airport_destination *uuid.UUID
	selectValues        sql.SelectValues
}

// FlightEdges holds the relations/edges for other nodes in the graph.
type FlightEdges struct {
	// Seats holds the value of the seats edge.
	Seats []*Seat `json:"seats,omitempty"`
	// Origin holds the value of the origin edge.
	Origin *Airport `json:"origin,omitempty"`
	// Destination holds the value of the destination edge.
	Destination *Airport `json:"destination,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SeatsOrErr returns the Seats value or an error if the edge
// was not loaded in eager-loading.
func (e FlightEdges) SeatsOrErr() ([]*Seat, error) {
	if e.loadedTypes[0] {
		return e.Seats, nil
	}
	return nil, &NotLoadedError{edge: "seats"}
}

// OriginOrErr returns the Origin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightEdges) OriginOrErr() (*Airport, error) {
	if e.loadedTypes[1] {
		if e.Origin == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: airport.Label}
		}
		return e.Origin, nil
	}
	return nil, &NotLoadedError{edge: "origin"}
}

// DestinationOrErr returns the Destination value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FlightEdges) DestinationOrErr() (*Airport, error) {
	if e.loadedTypes[2] {
		if e.Destination == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: airport.Label}
		}
		return e.Destination, nil
	}
	return nil, &NotLoadedError{edge: "destination"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Flight) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case flight.FieldAvailableSlots:
			values[i] = new(sql.NullInt64)
		case flight.FieldName:
			values[i] = new(sql.NullString)
		case flight.FieldDepartureTime, flight.FieldArrivalTime, flight.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case flight.FieldID:
			values[i] = new(uuid.UUID)
		case flight.ForeignKeys[0]: // airport_origin
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case flight.ForeignKeys[1]: // airport_destination
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
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
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case flight.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case flight.FieldDepartureTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field departure_time", values[i])
			} else if value.Valid {
				f.DepartureTime = value.Time
			}
		case flight.FieldArrivalTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field arrival_time", values[i])
			} else if value.Valid {
				f.ArrivalTime = value.Time
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
		case flight.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field airport_origin", values[i])
			} else if value.Valid {
				f.airport_origin = new(uuid.UUID)
				*f.airport_origin = *value.S.(*uuid.UUID)
			}
		case flight.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field airport_destination", values[i])
			} else if value.Valid {
				f.airport_destination = new(uuid.UUID)
				*f.airport_destination = *value.S.(*uuid.UUID)
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

// QueryOrigin queries the "origin" edge of the Flight entity.
func (f *Flight) QueryOrigin() *AirportQuery {
	return NewFlightClient(f.config).QueryOrigin(f)
}

// QueryDestination queries the "destination" edge of the Flight entity.
func (f *Flight) QueryDestination() *AirportQuery {
	return NewFlightClient(f.config).QueryDestination(f)
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
	builder.WriteString("departure_time=")
	builder.WriteString(f.DepartureTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("arrival_time=")
	builder.WriteString(f.ArrivalTime.Format(time.ANSIC))
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
