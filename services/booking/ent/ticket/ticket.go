// Code generated by ent, DO NOT EDIT.

package ticket

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ticket type in the database.
	Label = "ticket"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGoingFlightID holds the string denoting the going_flight_id field in the database.
	FieldGoingFlightID = "going_flight_id"
	// FieldReturnFlightID holds the string denoting the return_flight_id field in the database.
	FieldReturnFlightID = "return_flight_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPassengerName holds the string denoting the passenger_name field in the database.
	FieldPassengerName = "passenger_name"
	// FieldPassengerLicenseID holds the string denoting the passenger_license_id field in the database.
	FieldPassengerLicenseID = "passenger_license_id"
	// FieldPassengerEmail holds the string denoting the passenger_email field in the database.
	FieldPassengerEmail = "passenger_email"
	// FieldSeatNumber holds the string denoting the seat_number field in the database.
	FieldSeatNumber = "seat_number"
	// FieldClass holds the string denoting the class field in the database.
	FieldClass = "class"
	// EdgeBooking holds the string denoting the booking edge name in mutations.
	EdgeBooking = "booking"
	// Table holds the table name of the ticket in the database.
	Table = "tickets"
	// BookingTable is the table that holds the booking relation/edge. The primary key declared below.
	BookingTable = "booking_ticket"
	// BookingInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingInverseTable = "bookings"
)

// Columns holds all SQL columns for ticket fields.
var Columns = []string{
	FieldID,
	FieldGoingFlightID,
	FieldReturnFlightID,
	FieldStatus,
	FieldPassengerName,
	FieldPassengerLicenseID,
	FieldPassengerEmail,
	FieldSeatNumber,
	FieldClass,
}

var (
	// BookingPrimaryKey and BookingColumn2 are the table columns denoting the
	// primary key for the booking relation (M2M).
	BookingPrimaryKey = []string{"booking_id", "ticket_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusCanceled  Status = "Canceled"
	StatusDeparted  Status = "Departed"
	StatusScheduled Status = "Scheduled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusCanceled, StatusDeparted, StatusScheduled:
		return nil
	default:
		return fmt.Errorf("ticket: invalid enum value for status field: %q", s)
	}
}

// Class defines the type for the "class" enum field.
type Class string

// Class values.
const (
	ClassFirstClass    Class = "FirstClass"
	ClassBusinessClass Class = "BusinessClass"
	ClassEconomyClass  Class = "EconomyClass"
)

func (c Class) String() string {
	return string(c)
}

// ClassValidator is a validator for the "class" field enum values. It is called by the builders before save.
func ClassValidator(c Class) error {
	switch c {
	case ClassFirstClass, ClassBusinessClass, ClassEconomyClass:
		return nil
	default:
		return fmt.Errorf("ticket: invalid enum value for class field: %q", c)
	}
}

// OrderOption defines the ordering options for the Ticket queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGoingFlightID orders the results by the going_flight_id field.
func ByGoingFlightID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGoingFlightID, opts...).ToFunc()
}

// ByReturnFlightID orders the results by the return_flight_id field.
func ByReturnFlightID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReturnFlightID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPassengerName orders the results by the passenger_name field.
func ByPassengerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassengerName, opts...).ToFunc()
}

// ByPassengerLicenseID orders the results by the passenger_license_id field.
func ByPassengerLicenseID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassengerLicenseID, opts...).ToFunc()
}

// ByPassengerEmail orders the results by the passenger_email field.
func ByPassengerEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassengerEmail, opts...).ToFunc()
}

// BySeatNumber orders the results by the seat_number field.
func BySeatNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeatNumber, opts...).ToFunc()
}

// ByClass orders the results by the class field.
func ByClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClass, opts...).ToFunc()
}

// ByBookingCount orders the results by booking count.
func ByBookingCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBookingStep(), opts...)
	}
}

// ByBooking orders the results by booking terms.
func ByBooking(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookingStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBookingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, BookingTable, BookingPrimaryKey...),
	)
}
