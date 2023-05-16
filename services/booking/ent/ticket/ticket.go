// Code generated by ent, DO NOT EDIT.

package ticket

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ticket type in the database.
	Label = "ticket"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBookingID holds the string denoting the booking_id field in the database.
	FieldBookingID = "booking_id"
	// FieldSeatID holds the string denoting the seat_id field in the database.
	FieldSeatID = "seat_id"
	// FieldLicenseID holds the string denoting the license_id field in the database.
	FieldLicenseID = "license_id"
	// EdgeBooking holds the string denoting the booking edge name in mutations.
	EdgeBooking = "booking"
	// Table holds the table name of the ticket in the database.
	Table = "tickets"
	// BookingTable is the table that holds the booking relation/edge.
	BookingTable = "tickets"
	// BookingInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingInverseTable = "bookings"
	// BookingColumn is the table column denoting the booking relation/edge.
	BookingColumn = "booking_id"
)

// Columns holds all SQL columns for ticket fields.
var Columns = []string{
	FieldID,
	FieldBookingID,
	FieldSeatID,
	FieldLicenseID,
}

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

// OrderOption defines the ordering options for the Ticket queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBookingID orders the results by the booking_id field.
func ByBookingID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBookingID, opts...).ToFunc()
}

// BySeatID orders the results by the seat_id field.
func BySeatID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeatID, opts...).ToFunc()
}

// ByLicenseID orders the results by the license_id field.
func ByLicenseID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseID, opts...).ToFunc()
}

// ByBookingField orders the results by booking field.
func ByBookingField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookingStep(), sql.OrderByField(field, opts...))
	}
}
func newBookingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BookingTable, BookingColumn),
	)
}
