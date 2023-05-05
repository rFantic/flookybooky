// Code generated by ent, DO NOT EDIT.

package flight

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the flight type in the database.
	Label = "flight"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFromID holds the string denoting the from_id field in the database.
	FieldFromID = "from_id"
	// FieldToID holds the string denoting the to_id field in the database.
	FieldToID = "to_id"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// FieldAvailableSlots holds the string denoting the available_slots field in the database.
	FieldAvailableSlots = "available_slots"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeSeats holds the string denoting the seats edge name in mutations.
	EdgeSeats = "seats"
	// Table holds the table name of the flight in the database.
	Table = "flights"
	// SeatsTable is the table that holds the seats relation/edge.
	SeatsTable = "seats"
	// SeatsInverseTable is the table name for the Seat entity.
	// It exists in this package in order to avoid circular dependency with the "seat" package.
	SeatsInverseTable = "seats"
	// SeatsColumn is the table column denoting the seats relation/edge.
	SeatsColumn = "flight_seats"
)

// Columns holds all SQL columns for flight fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFromID,
	FieldToID,
	FieldStart,
	FieldEnd,
	FieldAvailableSlots,
	FieldCreatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Flight queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFromID orders the results by the from_id field.
func ByFromID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromID, opts...).ToFunc()
}

// ByToID orders the results by the to_id field.
func ByToID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToID, opts...).ToFunc()
}

// ByStart orders the results by the start field.
func ByStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStart, opts...).ToFunc()
}

// ByEnd orders the results by the end field.
func ByEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnd, opts...).ToFunc()
}

// ByAvailableSlots orders the results by the available_slots field.
func ByAvailableSlots(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailableSlots, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// BySeatsCount orders the results by seats count.
func BySeatsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSeatsStep(), opts...)
	}
}

// BySeats orders the results by seats terms.
func BySeats(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSeatsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSeatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SeatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SeatsTable, SeatsColumn),
	)
}