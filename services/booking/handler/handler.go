package handler

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/booking/ent"
	"flookybooky/services/booking/ent/booking"
	"flookybooky/services/booking/ent/ticket"
	"flookybooky/services/booking/internal"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	customerClient pb.CustomerServiceClient
	client         ent.Client
}

func NewBookingHandler(client ent.Client, customerClient pb.CustomerServiceClient) (*BookingHandler, error) {
	return &BookingHandler{
		client:         client,
		customerClient: customerClient,
	}, nil
}

func (h *BookingHandler) GetBooking(ctx context.Context, req *pb.UUID) (*pb.Booking, error) {
	_uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	bookingRes, err := h.client.Booking.Get(ctx, _uuid)
	return internal.ParseBookingEntToPb(bookingRes), err
}

func (h *BookingHandler) GetTicket(ctx context.Context, req *pb.UUID) (*pb.Ticket, error) {
	_uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	ticketRes, err := h.client.Ticket.Get(ctx, _uuid)
	return internal.ParseTicketEntToPb(ticketRes), err
}

func (h *BookingHandler) GetBookings(ctx context.Context, req *pb.Pagination) (*pb.Bookings, error) {
	query := h.client.Booking.Query()
	if req != nil {
		var options []booking.OrderOption
		if req.AscFields != nil {
			options = append(options, ent.Asc(req.AscFields...))
		}
		if req.DesFields != nil {
			options = append(options, ent.Desc(req.DesFields...))
		}
		query.Order(options...)
		if req.Limit != nil {
			query.Limit(int(*req.Limit))
		} else {
			query.Limit(10)
		}
		if req.Offset != nil {
			query.Offset(int(*req.Offset))
		}
	} else {
		query.Limit(10)
	}
	bookingsRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	res := internal.ParseBookingsEntToPb(bookingsRes)
	for i, b := range bookingsRes {
		ticketsRes, err := b.QueryTicket().All(ctx)
		if err != nil {
			return nil, err
		}
		res.Bookings[i].Tickets = internal.ParseTicketsEntToPb(ticketsRes)
	}
	return res, nil
}

func (h *BookingHandler) GetTickets(ctx context.Context, req *pb.Pagination) (*pb.Tickets, error) {
	query := h.client.Ticket.Query()
	if req != nil {
		var options []ticket.OrderOption
		if req.AscFields != nil {
			options = append(options, ent.Asc(req.AscFields...))
		}
		if req.DesFields != nil {
			options = append(options, ent.Desc(req.DesFields...))
		}
		query.Order(options...)
		if req.Limit != nil {
			query.Limit(int(*req.Limit))
		} else {
			query.Limit(10)
		}
		if req.Offset != nil {
			query.Offset(int(*req.Offset))
		}
	} else {
		query.Limit(10)
	}
	ticketsRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseTicketsEntToPb(ticketsRes), nil
}

func (h *BookingHandler) PostBookingForGuest(ctx context.Context, req *pb.BookingInputForGuest) (*pb.Booking, error) {
	customerRes, err := h.customerClient.PostCustomer(ctx, req.CustomerInput)
	if err != nil {
		return nil, err
	}
	bookingInput := &pb.BookingInput{}
	copier.Copy(&bookingInput, req)
	bookingInput.CustomerId = customerRes.Id
	return h.PostBooking(ctx, bookingInput)
}

func (h *BookingHandler) PostBooking(ctx context.Context, req *pb.BookingInput) (*pb.Booking, error) {
	_customerId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return nil, err
	}
	_going_flight_id, err := uuid.Parse(req.GoingFlightId)
	if err != nil {
		return nil, err
	}
	bookingQuery := h.client.Booking.Create().
		SetCustomerID(_customerId).
		SetStatus(booking.Status(req.Status)).
		SetGoingFlightID(_going_flight_id)
	if req.ReturnFlightId != nil {
		_return_flight_id, err := uuid.Parse(*req.ReturnFlightId)
		if err != nil {
			return nil, err
		}
		bookingQuery.SetReturnFlightID(_return_flight_id)
	}
	bookingRes, err := bookingQuery.Save(ctx)

	var ent_tickets []*ent.Ticket
	for _, t := range req.Tickets {
		ticketQuery := h.client.Ticket.Create().
			SetClass(ticket.Class(t.Class)).
			SetPassengerEmail(t.PassengerEmail).
			SetPassengerLicenseID(t.PassengerLicenseId).
			SetPassengerName(t.PassengerName).
			SetSeatNumber(t.SeatNumber).
			SetStatus(ticket.Status(t.Status)).
			SetBookingID(bookingRes.ID)
		ent_ticket, err := ticketQuery.Save(ctx)
		if err != nil {
			return nil, err
		}
		ent_tickets = append(ent_tickets, ent_ticket)
	}
	if err != nil {
		return nil, err
	}
	res := internal.ParseBookingEntToPb(bookingRes)
	if ent_tickets != nil {
		res.Tickets = internal.ParseTicketsEntToPb(ent_tickets)
	}
	return res, err
}

func (h *BookingHandler) PostTicket(ctx context.Context, req *pb.TicketInput) (*pb.Ticket, error) {
	query := h.client.Ticket.Create().
		SetPassengerEmail(req.PassengerEmail).
		SetPassengerLicenseID(req.PassengerLicenseId).
		SetPassengerName(req.PassengerName).
		SetClass(ticket.Class(req.Class)).
		SetSeatNumber(req.SeatNumber).
		SetStatus(ticket.Status(req.Status))
	ticketRes, err := query.Save(ctx)
	return internal.ParseTicketEntToPb(ticketRes), err

}
