// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AirportsColumns holds the columns for the "airports" table.
	AirportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
	}
	// AirportsTable holds the schema information for the "airports" table.
	AirportsTable = &schema.Table{
		Name:       "airports",
		Columns:    AirportsColumns,
		PrimaryKey: []*schema.Column{AirportsColumns[0]},
	}
	// FlightsColumns holds the columns for the "flights" table.
	FlightsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "departure_time", Type: field.TypeTime},
		{Name: "arrival_time", Type: field.TypeTime},
		{Name: "available_slots", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "origin_id", Type: field.TypeUUID},
		{Name: "destinartion_id", Type: field.TypeUUID},
	}
	// FlightsTable holds the schema information for the "flights" table.
	FlightsTable = &schema.Table{
		Name:       "flights",
		Columns:    FlightsColumns,
		PrimaryKey: []*schema.Column{FlightsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "flights_airports_origin",
				Columns:    []*schema.Column{FlightsColumns[6]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "flights_airports_destination",
				Columns:    []*schema.Column{FlightsColumns[7]},
				RefColumns: []*schema.Column{AirportsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SeatsColumns holds the columns for the "seats" table.
	SeatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "seat_number", Type: field.TypeString},
		{Name: "flight_seats", Type: field.TypeUUID, Nullable: true},
	}
	// SeatsTable holds the schema information for the "seats" table.
	SeatsTable = &schema.Table{
		Name:       "seats",
		Columns:    SeatsColumns,
		PrimaryKey: []*schema.Column{SeatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "seats_flights_seats",
				Columns:    []*schema.Column{SeatsColumns[2]},
				RefColumns: []*schema.Column{FlightsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AirportsTable,
		FlightsTable,
		SeatsTable,
	}
)

func init() {
	FlightsTable.ForeignKeys[0].RefTable = AirportsTable
	FlightsTable.ForeignKeys[1].RefTable = AirportsTable
	SeatsTable.ForeignKeys[0].RefTable = FlightsTable
}
