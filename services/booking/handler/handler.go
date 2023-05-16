package handler

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/booking/ent"
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

func (h *BookingHandler) GetBookings(ctx context.Context, req *emptypb.Empty) (*pb.Bookings, error) {
	query := h.client.Booking.Query()
	// query = query.Offset(int(req.Offset)).Limit(int(req.Limit))
	bookingsRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseBookingsEntToPb(bookingsRes), nil
}

func (h *BookingHandler) PostBooking(ctx context.Context, req *pb.BookingInput) (*pb.Booking, error) {
	_customerId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return nil, err
	}
	_goingFlightId, err := uuid.Parse(req.GoingFlightId)
	if err != nil {
		return nil, err
	}
	query := h.client.Booking.Create().
		SetCustomerID(_customerId).
		SetGoingFlightID(_goingFlightId)

	if req.ReturnFlightId != nil {
		_returnFlightId, err := uuid.Parse(*req.ReturnFlightId)
		if err != nil {
			return nil, err
		}
		query.SetReturnFlightID(_returnFlightId)
	}
	bookingRes, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseBookingEntToPb(bookingRes), err
}
