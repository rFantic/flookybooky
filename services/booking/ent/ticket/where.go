// Code generated by ent, DO NOT EDIT.

package ticket

import (
	"flookybooky/services/booking/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldID, id))
}

// BookingID applies equality check predicate on the "booking_id" field. It's identical to BookingIDEQ.
func BookingID(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldBookingID, v))
}

// GoingFlightID applies equality check predicate on the "going_flight_id" field. It's identical to GoingFlightIDEQ.
func GoingFlightID(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldGoingFlightID, v))
}

// ReturnFlightID applies equality check predicate on the "return_flight_id" field. It's identical to ReturnFlightIDEQ.
func ReturnFlightID(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldReturnFlightID, v))
}

// PassengerName applies equality check predicate on the "passenger_name" field. It's identical to PassengerNameEQ.
func PassengerName(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldPassengerName, v))
}

// PassengerLicenseID applies equality check predicate on the "passenger_license_id" field. It's identical to PassengerLicenseIDEQ.
func PassengerLicenseID(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldPassengerLicenseID, v))
}

// PassengerEmail applies equality check predicate on the "passenger_email" field. It's identical to PassengerEmailEQ.
func PassengerEmail(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldPassengerEmail, v))
}

// SeatNumber applies equality check predicate on the "seat_number" field. It's identical to SeatNumberEQ.
func SeatNumber(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldSeatNumber, v))
}

// BookingIDEQ applies the EQ predicate on the "booking_id" field.
func BookingIDEQ(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldBookingID, v))
}

// BookingIDNEQ applies the NEQ predicate on the "booking_id" field.
func BookingIDNEQ(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldBookingID, v))
}

// BookingIDIn applies the In predicate on the "booking_id" field.
func BookingIDIn(vs ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldBookingID, vs...))
}

// BookingIDNotIn applies the NotIn predicate on the "booking_id" field.
func BookingIDNotIn(vs ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldBookingID, vs...))
}

// GoingFlightIDEQ applies the EQ predicate on the "going_flight_id" field.
func GoingFlightIDEQ(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldGoingFlightID, v))
}

// GoingFlightIDNEQ applies the NEQ predicate on the "going_flight_id" field.
func GoingFlightIDNEQ(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldGoingFlightID, v))
}

// GoingFlightIDIn applies the In predicate on the "going_flight_id" field.
func GoingFlightIDIn(vs ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldGoingFlightID, vs...))
}

// GoingFlightIDNotIn applies the NotIn predicate on the "going_flight_id" field.
func GoingFlightIDNotIn(vs ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldGoingFlightID, vs...))
}

// GoingFlightIDGT applies the GT predicate on the "going_flight_id" field.
func GoingFlightIDGT(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldGoingFlightID, v))
}

// GoingFlightIDGTE applies the GTE predicate on the "going_flight_id" field.
func GoingFlightIDGTE(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldGoingFlightID, v))
}

// GoingFlightIDLT applies the LT predicate on the "going_flight_id" field.
func GoingFlightIDLT(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldGoingFlightID, v))
}

// GoingFlightIDLTE applies the LTE predicate on the "going_flight_id" field.
func GoingFlightIDLTE(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldGoingFlightID, v))
}

// ReturnFlightIDEQ applies the EQ predicate on the "return_flight_id" field.
func ReturnFlightIDEQ(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldReturnFlightID, v))
}

// ReturnFlightIDNEQ applies the NEQ predicate on the "return_flight_id" field.
func ReturnFlightIDNEQ(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldReturnFlightID, v))
}

// ReturnFlightIDIn applies the In predicate on the "return_flight_id" field.
func ReturnFlightIDIn(vs ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldReturnFlightID, vs...))
}

// ReturnFlightIDNotIn applies the NotIn predicate on the "return_flight_id" field.
func ReturnFlightIDNotIn(vs ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldReturnFlightID, vs...))
}

// ReturnFlightIDGT applies the GT predicate on the "return_flight_id" field.
func ReturnFlightIDGT(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldReturnFlightID, v))
}

// ReturnFlightIDGTE applies the GTE predicate on the "return_flight_id" field.
func ReturnFlightIDGTE(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldReturnFlightID, v))
}

// ReturnFlightIDLT applies the LT predicate on the "return_flight_id" field.
func ReturnFlightIDLT(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldReturnFlightID, v))
}

// ReturnFlightIDLTE applies the LTE predicate on the "return_flight_id" field.
func ReturnFlightIDLTE(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldReturnFlightID, v))
}

// ReturnFlightIDIsNil applies the IsNil predicate on the "return_flight_id" field.
func ReturnFlightIDIsNil() predicate.Ticket {
	return predicate.Ticket(sql.FieldIsNull(FieldReturnFlightID))
}

// ReturnFlightIDNotNil applies the NotNil predicate on the "return_flight_id" field.
func ReturnFlightIDNotNil() predicate.Ticket {
	return predicate.Ticket(sql.FieldNotNull(FieldReturnFlightID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldStatus, vs...))
}

// PassengerNameEQ applies the EQ predicate on the "passenger_name" field.
func PassengerNameEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldPassengerName, v))
}

// PassengerNameNEQ applies the NEQ predicate on the "passenger_name" field.
func PassengerNameNEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldPassengerName, v))
}

// PassengerNameIn applies the In predicate on the "passenger_name" field.
func PassengerNameIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldPassengerName, vs...))
}

// PassengerNameNotIn applies the NotIn predicate on the "passenger_name" field.
func PassengerNameNotIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldPassengerName, vs...))
}

// PassengerNameGT applies the GT predicate on the "passenger_name" field.
func PassengerNameGT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldPassengerName, v))
}

// PassengerNameGTE applies the GTE predicate on the "passenger_name" field.
func PassengerNameGTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldPassengerName, v))
}

// PassengerNameLT applies the LT predicate on the "passenger_name" field.
func PassengerNameLT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldPassengerName, v))
}

// PassengerNameLTE applies the LTE predicate on the "passenger_name" field.
func PassengerNameLTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldPassengerName, v))
}

// PassengerNameContains applies the Contains predicate on the "passenger_name" field.
func PassengerNameContains(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContains(FieldPassengerName, v))
}

// PassengerNameHasPrefix applies the HasPrefix predicate on the "passenger_name" field.
func PassengerNameHasPrefix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasPrefix(FieldPassengerName, v))
}

// PassengerNameHasSuffix applies the HasSuffix predicate on the "passenger_name" field.
func PassengerNameHasSuffix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasSuffix(FieldPassengerName, v))
}

// PassengerNameEqualFold applies the EqualFold predicate on the "passenger_name" field.
func PassengerNameEqualFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEqualFold(FieldPassengerName, v))
}

// PassengerNameContainsFold applies the ContainsFold predicate on the "passenger_name" field.
func PassengerNameContainsFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContainsFold(FieldPassengerName, v))
}

// PassengerLicenseIDEQ applies the EQ predicate on the "passenger_license_id" field.
func PassengerLicenseIDEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDNEQ applies the NEQ predicate on the "passenger_license_id" field.
func PassengerLicenseIDNEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDIn applies the In predicate on the "passenger_license_id" field.
func PassengerLicenseIDIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldPassengerLicenseID, vs...))
}

// PassengerLicenseIDNotIn applies the NotIn predicate on the "passenger_license_id" field.
func PassengerLicenseIDNotIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldPassengerLicenseID, vs...))
}

// PassengerLicenseIDGT applies the GT predicate on the "passenger_license_id" field.
func PassengerLicenseIDGT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDGTE applies the GTE predicate on the "passenger_license_id" field.
func PassengerLicenseIDGTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDLT applies the LT predicate on the "passenger_license_id" field.
func PassengerLicenseIDLT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDLTE applies the LTE predicate on the "passenger_license_id" field.
func PassengerLicenseIDLTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDContains applies the Contains predicate on the "passenger_license_id" field.
func PassengerLicenseIDContains(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContains(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDHasPrefix applies the HasPrefix predicate on the "passenger_license_id" field.
func PassengerLicenseIDHasPrefix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasPrefix(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDHasSuffix applies the HasSuffix predicate on the "passenger_license_id" field.
func PassengerLicenseIDHasSuffix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasSuffix(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDEqualFold applies the EqualFold predicate on the "passenger_license_id" field.
func PassengerLicenseIDEqualFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEqualFold(FieldPassengerLicenseID, v))
}

// PassengerLicenseIDContainsFold applies the ContainsFold predicate on the "passenger_license_id" field.
func PassengerLicenseIDContainsFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContainsFold(FieldPassengerLicenseID, v))
}

// PassengerEmailEQ applies the EQ predicate on the "passenger_email" field.
func PassengerEmailEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldPassengerEmail, v))
}

// PassengerEmailNEQ applies the NEQ predicate on the "passenger_email" field.
func PassengerEmailNEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldPassengerEmail, v))
}

// PassengerEmailIn applies the In predicate on the "passenger_email" field.
func PassengerEmailIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldPassengerEmail, vs...))
}

// PassengerEmailNotIn applies the NotIn predicate on the "passenger_email" field.
func PassengerEmailNotIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldPassengerEmail, vs...))
}

// PassengerEmailGT applies the GT predicate on the "passenger_email" field.
func PassengerEmailGT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldPassengerEmail, v))
}

// PassengerEmailGTE applies the GTE predicate on the "passenger_email" field.
func PassengerEmailGTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldPassengerEmail, v))
}

// PassengerEmailLT applies the LT predicate on the "passenger_email" field.
func PassengerEmailLT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldPassengerEmail, v))
}

// PassengerEmailLTE applies the LTE predicate on the "passenger_email" field.
func PassengerEmailLTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldPassengerEmail, v))
}

// PassengerEmailContains applies the Contains predicate on the "passenger_email" field.
func PassengerEmailContains(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContains(FieldPassengerEmail, v))
}

// PassengerEmailHasPrefix applies the HasPrefix predicate on the "passenger_email" field.
func PassengerEmailHasPrefix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasPrefix(FieldPassengerEmail, v))
}

// PassengerEmailHasSuffix applies the HasSuffix predicate on the "passenger_email" field.
func PassengerEmailHasSuffix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasSuffix(FieldPassengerEmail, v))
}

// PassengerEmailEqualFold applies the EqualFold predicate on the "passenger_email" field.
func PassengerEmailEqualFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEqualFold(FieldPassengerEmail, v))
}

// PassengerEmailContainsFold applies the ContainsFold predicate on the "passenger_email" field.
func PassengerEmailContainsFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContainsFold(FieldPassengerEmail, v))
}

// SeatNumberEQ applies the EQ predicate on the "seat_number" field.
func SeatNumberEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldSeatNumber, v))
}

// SeatNumberNEQ applies the NEQ predicate on the "seat_number" field.
func SeatNumberNEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldSeatNumber, v))
}

// SeatNumberIn applies the In predicate on the "seat_number" field.
func SeatNumberIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldSeatNumber, vs...))
}

// SeatNumberNotIn applies the NotIn predicate on the "seat_number" field.
func SeatNumberNotIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldSeatNumber, vs...))
}

// SeatNumberGT applies the GT predicate on the "seat_number" field.
func SeatNumberGT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldSeatNumber, v))
}

// SeatNumberGTE applies the GTE predicate on the "seat_number" field.
func SeatNumberGTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldSeatNumber, v))
}

// SeatNumberLT applies the LT predicate on the "seat_number" field.
func SeatNumberLT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldSeatNumber, v))
}

// SeatNumberLTE applies the LTE predicate on the "seat_number" field.
func SeatNumberLTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldSeatNumber, v))
}

// SeatNumberContains applies the Contains predicate on the "seat_number" field.
func SeatNumberContains(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContains(FieldSeatNumber, v))
}

// SeatNumberHasPrefix applies the HasPrefix predicate on the "seat_number" field.
func SeatNumberHasPrefix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasPrefix(FieldSeatNumber, v))
}

// SeatNumberHasSuffix applies the HasSuffix predicate on the "seat_number" field.
func SeatNumberHasSuffix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasSuffix(FieldSeatNumber, v))
}

// SeatNumberEqualFold applies the EqualFold predicate on the "seat_number" field.
func SeatNumberEqualFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEqualFold(FieldSeatNumber, v))
}

// SeatNumberContainsFold applies the ContainsFold predicate on the "seat_number" field.
func SeatNumberContainsFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContainsFold(FieldSeatNumber, v))
}

// ClassEQ applies the EQ predicate on the "class" field.
func ClassEQ(v Class) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldClass, v))
}

// ClassNEQ applies the NEQ predicate on the "class" field.
func ClassNEQ(v Class) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldClass, v))
}

// ClassIn applies the In predicate on the "class" field.
func ClassIn(vs ...Class) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldClass, vs...))
}

// ClassNotIn applies the NotIn predicate on the "class" field.
func ClassNotIn(vs ...Class) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldClass, vs...))
}

// HasBooking applies the HasEdge predicate on the "booking" edge.
func HasBooking() predicate.Ticket {
	return predicate.Ticket(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BookingTable, BookingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookingWith applies the HasEdge predicate on the "booking" edge with a given conditions (other predicates).
func HasBookingWith(preds ...predicate.Booking) predicate.Ticket {
	return predicate.Ticket(func(s *sql.Selector) {
		step := newBookingStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(func(s *sql.Selector) {
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
func Not(p predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(func(s *sql.Selector) {
		p(s.Not())
	})
}
