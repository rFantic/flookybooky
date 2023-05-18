package handler

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/booking/ent"
	"flookybooky/services/booking/ent/booking"
	"flookybooky/services/booking/ent/ticket"
	"flookybooky/services/booking/internal"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	client ent.Client
}

func NewBookingHandler(client ent.Client) (*BookingHandler, error) {
	return &BookingHandler{
		client: client,
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

func (h *BookingHandler) GetBookings(ctx context.Context, req *emptypb.Empty) (*pb.Bookings, error) {
	query := h.client.Booking.Query()
	// query = query.Offset(int(req.Offset)).Limit(int(req.Limit))
	bookingsRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseBookingsEntToPb(bookingsRes), nil
}

func (h *BookingHandler) GetTickets(ctx context.Context, req *emptypb.Empty) (*pb.Tickets, error) {
	query := h.client.Ticket.Query()
	// query = query.Offset(int(req.Offset)).Limit(int(req.Limit))
	ticketsRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseTicketsEntToPb(ticketsRes), nil
}

func (h *BookingHandler) PostBooking(ctx context.Context, req *pb.BookingInput) (*pb.Booking, error) {
	_customerId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return nil, err
	}
	query := h.client.Booking.Create().
		SetCustomerID(_customerId).
		SetStatus(booking.Status(req.Status))

	if req.GoingTicket != nil {
		goingTicketRes, err := h.PostTicket(ctx, req.GoingTicket)
		if err != nil {
			return nil, err
		}
		_goingTicketId, err := uuid.Parse(goingTicketRes.Id)
		if err != nil {
			return nil, err
		}
		query.SetGoingTicketID(_goingTicketId)
	}

	if req.ReturnTicket != nil {
		returnTicketRes, err := h.PostTicket(ctx, req.ReturnTicket)
		if err != nil {
			return nil, err
		}
		_returnTicketId, err := uuid.Parse(returnTicketRes.Id)
		if err != nil {
			return nil, err
		}
		query.SetReturnTicketID(_returnTicketId)
	}

	bookingRes, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseBookingEntToPb(bookingRes), err
}

func (h *BookingHandler) PostTicket(ctx context.Context, req *pb.TicketInput) (*pb.Ticket, error) {
	_flight_id, err := uuid.Parse(req.FlightId)
	if err != nil {
		return nil, err
	}
	ticketRes, err := h.client.Ticket.Create().SetFlightID(_flight_id).
		SetPassengerEmail(req.PassengerEmail).
		SetPassengerLicenseID(req.PassengerLicenseId).
		SetPassengerName(req.PassengerName).
		SetClass(ticket.Class(req.Class)).
		SetSeatNumber(req.SeatNumber).
		SetStatus(ticket.Status(req.Status)).Save(ctx)
	return internal.ParseTicketEntToPb(ticketRes), err

}
