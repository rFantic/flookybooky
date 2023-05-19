// Code generated by ent, DO NOT EDIT.

package booking

import (
	"flookybooky/services/booking/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldID, id))
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCustomerID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCreatedAt, v))
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCustomerID, v))
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldCustomerID, v))
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldCustomerID, vs...))
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldCustomerID, vs...))
}

// CustomerIDGT applies the GT predicate on the "customer_id" field.
func CustomerIDGT(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldCustomerID, v))
}

// CustomerIDGTE applies the GTE predicate on the "customer_id" field.
func CustomerIDGTE(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldCustomerID, v))
}

// CustomerIDLT applies the LT predicate on the "customer_id" field.
func CustomerIDLT(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldCustomerID, v))
}

// CustomerIDLTE applies the LTE predicate on the "customer_id" field.
func CustomerIDLTE(v uuid.UUID) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldCustomerID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldCreatedAt, v))
}

// HasTicket applies the HasEdge predicate on the "ticket" edge.
func HasTicket() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TicketTable, TicketColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTicketWith applies the HasEdge predicate on the "ticket" edge with a given conditions (other predicates).
func HasTicketWith(preds ...predicate.Ticket) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newTicketStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
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
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		p(s.Not())
	})
}
