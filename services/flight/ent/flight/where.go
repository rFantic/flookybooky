// Code generated by ent, DO NOT EDIT.

package flight

import (
	"flookybooky/services/flight/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldName, v))
}

// OriginID applies equality check predicate on the "origin_id" field. It's identical to OriginIDEQ.
func OriginID(v uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldOriginID, v))
}

// DestinartionID applies equality check predicate on the "destinartion_id" field. It's identical to DestinartionIDEQ.
func DestinartionID(v uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDestinartionID, v))
}

// DepartureTime applies equality check predicate on the "departure_time" field. It's identical to DepartureTimeEQ.
func DepartureTime(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDepartureTime, v))
}

// ArrivalTime applies equality check predicate on the "arrival_time" field. It's identical to ArrivalTimeEQ.
func ArrivalTime(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldArrivalTime, v))
}

// AvailableSlots applies equality check predicate on the "available_slots" field. It's identical to AvailableSlotsEQ.
func AvailableSlots(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAvailableSlots, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContainsFold(FieldName, v))
}

// OriginIDEQ applies the EQ predicate on the "origin_id" field.
func OriginIDEQ(v uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldOriginID, v))
}

// OriginIDNEQ applies the NEQ predicate on the "origin_id" field.
func OriginIDNEQ(v uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldOriginID, v))
}

// OriginIDIn applies the In predicate on the "origin_id" field.
func OriginIDIn(vs ...uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldOriginID, vs...))
}

// OriginIDNotIn applies the NotIn predicate on the "origin_id" field.
func OriginIDNotIn(vs ...uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldOriginID, vs...))
}

// DestinartionIDEQ applies the EQ predicate on the "destinartion_id" field.
func DestinartionIDEQ(v uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDestinartionID, v))
}

// DestinartionIDNEQ applies the NEQ predicate on the "destinartion_id" field.
func DestinartionIDNEQ(v uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldDestinartionID, v))
}

// DestinartionIDIn applies the In predicate on the "destinartion_id" field.
func DestinartionIDIn(vs ...uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldDestinartionID, vs...))
}

// DestinartionIDNotIn applies the NotIn predicate on the "destinartion_id" field.
func DestinartionIDNotIn(vs ...uuid.UUID) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldDestinartionID, vs...))
}

// DepartureTimeEQ applies the EQ predicate on the "departure_time" field.
func DepartureTimeEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDepartureTime, v))
}

// DepartureTimeNEQ applies the NEQ predicate on the "departure_time" field.
func DepartureTimeNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldDepartureTime, v))
}

// DepartureTimeIn applies the In predicate on the "departure_time" field.
func DepartureTimeIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldDepartureTime, vs...))
}

// DepartureTimeNotIn applies the NotIn predicate on the "departure_time" field.
func DepartureTimeNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldDepartureTime, vs...))
}

// DepartureTimeGT applies the GT predicate on the "departure_time" field.
func DepartureTimeGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldDepartureTime, v))
}

// DepartureTimeGTE applies the GTE predicate on the "departure_time" field.
func DepartureTimeGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldDepartureTime, v))
}

// DepartureTimeLT applies the LT predicate on the "departure_time" field.
func DepartureTimeLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldDepartureTime, v))
}

// DepartureTimeLTE applies the LTE predicate on the "departure_time" field.
func DepartureTimeLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldDepartureTime, v))
}

// ArrivalTimeEQ applies the EQ predicate on the "arrival_time" field.
func ArrivalTimeEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldArrivalTime, v))
}

// ArrivalTimeNEQ applies the NEQ predicate on the "arrival_time" field.
func ArrivalTimeNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldArrivalTime, v))
}

// ArrivalTimeIn applies the In predicate on the "arrival_time" field.
func ArrivalTimeIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldArrivalTime, vs...))
}

// ArrivalTimeNotIn applies the NotIn predicate on the "arrival_time" field.
func ArrivalTimeNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldArrivalTime, vs...))
}

// ArrivalTimeGT applies the GT predicate on the "arrival_time" field.
func ArrivalTimeGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldArrivalTime, v))
}

// ArrivalTimeGTE applies the GTE predicate on the "arrival_time" field.
func ArrivalTimeGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldArrivalTime, v))
}

// ArrivalTimeLT applies the LT predicate on the "arrival_time" field.
func ArrivalTimeLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldArrivalTime, v))
}

// ArrivalTimeLTE applies the LTE predicate on the "arrival_time" field.
func ArrivalTimeLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldArrivalTime, v))
}

// AvailableSlotsEQ applies the EQ predicate on the "available_slots" field.
func AvailableSlotsEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAvailableSlots, v))
}

// AvailableSlotsNEQ applies the NEQ predicate on the "available_slots" field.
func AvailableSlotsNEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldAvailableSlots, v))
}

// AvailableSlotsIn applies the In predicate on the "available_slots" field.
func AvailableSlotsIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldAvailableSlots, vs...))
}

// AvailableSlotsNotIn applies the NotIn predicate on the "available_slots" field.
func AvailableSlotsNotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldAvailableSlots, vs...))
}

// AvailableSlotsGT applies the GT predicate on the "available_slots" field.
func AvailableSlotsGT(v int) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldAvailableSlots, v))
}

// AvailableSlotsGTE applies the GTE predicate on the "available_slots" field.
func AvailableSlotsGTE(v int) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldAvailableSlots, v))
}

// AvailableSlotsLT applies the LT predicate on the "available_slots" field.
func AvailableSlotsLT(v int) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldAvailableSlots, v))
}

// AvailableSlotsLTE applies the LTE predicate on the "available_slots" field.
func AvailableSlotsLTE(v int) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldAvailableSlots, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldCreatedAt, v))
}

// HasSeats applies the HasEdge predicate on the "seats" edge.
func HasSeats() predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SeatsTable, SeatsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSeatsWith applies the HasEdge predicate on the "seats" edge with a given conditions (other predicates).
func HasSeatsWith(preds ...predicate.Seat) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := newSeatsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrigin applies the HasEdge predicate on the "origin" edge.
func HasOrigin() predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OriginTable, OriginColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOriginWith applies the HasEdge predicate on the "origin" edge with a given conditions (other predicates).
func HasOriginWith(preds ...predicate.Airport) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := newOriginStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDestination applies the HasEdge predicate on the "destination" edge.
func HasDestination() predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DestinationTable, DestinationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDestinationWith applies the HasEdge predicate on the "destination" edge with a given conditions (other predicates).
func HasDestinationWith(preds ...predicate.Airport) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := newDestinationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		p(s.Not())
	})
}
