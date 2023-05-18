// Code generated by ent, DO NOT EDIT.

package booking

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the booking type in the database.
	Label = "booking"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldGoingTicketID holds the string denoting the going_ticket_id field in the database.
	FieldGoingTicketID = "going_ticket_id"
	// FieldReturnTicketID holds the string denoting the return_ticket_id field in the database.
	FieldReturnTicketID = "return_ticket_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeTicket holds the string denoting the ticket edge name in mutations.
	EdgeTicket = "ticket"
	// Table holds the table name of the booking in the database.
	Table = "bookings"
	// TicketTable is the table that holds the ticket relation/edge. The primary key declared below.
	TicketTable = "booking_ticket"
	// TicketInverseTable is the table name for the Ticket entity.
	// It exists in this package in order to avoid circular dependency with the "ticket" package.
	TicketInverseTable = "tickets"
)

// Columns holds all SQL columns for booking fields.
var Columns = []string{
	FieldID,
	FieldCustomerID,
	FieldGoingTicketID,
	FieldReturnTicketID,
	FieldStatus,
	FieldCreatedAt,
}

var (
	// TicketPrimaryKey and TicketColumn2 are the table columns denoting the
	// primary key for the ticket relation (M2M).
	TicketPrimaryKey = []string{"booking_id", "ticket_id"}
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusCanceled  Status = "Canceled"
	StatusScheduled Status = "Scheduled"
	StatusDeparted  Status = "Departed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusCanceled, StatusScheduled, StatusDeparted:
		return nil
	default:
		return fmt.Errorf("booking: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Booking queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// ByGoingTicketID orders the results by the going_ticket_id field.
func ByGoingTicketID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGoingTicketID, opts...).ToFunc()
}

// ByReturnTicketID orders the results by the return_ticket_id field.
func ByReturnTicketID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReturnTicketID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByTicketCount orders the results by ticket count.
func ByTicketCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTicketStep(), opts...)
	}
}

// ByTicket orders the results by ticket terms.
func ByTicket(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTicketStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTicketStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TicketInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TicketTable, TicketPrimaryKey...),
	)
}
