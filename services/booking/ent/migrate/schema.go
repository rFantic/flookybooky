// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BookingsColumns holds the columns for the "bookings" table.
	BookingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "customer_id", Type: field.TypeUUID},
		{Name: "flight_id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
	}
	// BookingsTable holds the schema information for the "bookings" table.
	BookingsTable = &schema.Table{
		Name:       "bookings",
		Columns:    BookingsColumns,
		PrimaryKey: []*schema.Column{BookingsColumns[0]},
	}
	// TicketsColumns holds the columns for the "tickets" table.
	TicketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "seat_id", Type: field.TypeUUID},
		{Name: "license_id", Type: field.TypeString},
		{Name: "booking_id", Type: field.TypeUUID},
	}
	// TicketsTable holds the schema information for the "tickets" table.
	TicketsTable = &schema.Table{
		Name:       "tickets",
		Columns:    TicketsColumns,
		PrimaryKey: []*schema.Column{TicketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tickets_bookings_ticket",
				Columns:    []*schema.Column{TicketsColumns[3]},
				RefColumns: []*schema.Column{BookingsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BookingsTable,
		TicketsTable,
	}
)

func init() {
	TicketsTable.ForeignKeys[0].RefTable = BookingsTable
}
