package internal

import (
	"flookybooky/pb"
	"flookybooky/services/booking/ent"

	"github.com/jinzhu/copier"
)

func ParseTicketEntToPb(in *ent.Ticket) (out *pb.Ticket) {
	if in == nil {
		return nil
	}
	out = &pb.Ticket{
		Id:                 in.ID.String(),
		FlightId:           in.FlightID.String(),
		PassengerLicenseId: in.PassengerLicenseID,
	}
	copier.Copy(&out, in)
	return out
}

func ParseTicketsEntToPb(in []*ent.Ticket) (out *pb.Tickets) {
	if in == nil {
		return nil
	}
	out = &pb.Tickets{
		Tickets: make([]*pb.Ticket, len(in)),
	}
	copier.Copy(&out, in)
	for i, a := range in {
		out.Tickets[i] = ParseTicketEntToPb(a)
	}
	return out
}

func ParseBookingEntToPb(in *ent.Booking) (out *pb.Booking) {
	if in == nil {
		return nil
	}
	out = &pb.Booking{
		Id: in.ID.String(),
	}
	copier.Copy(&out, in)
	out.Customer = &pb.Customer{
		Id: in.CustomerID.String(),
	}
	out.GoingTicket = &pb.Ticket{
		Id: in.GoingTicketID.String(),
	}
	if in.ReturnTicketID != nil {
		out.ReturnTicket = &pb.Ticket{
			Id: in.ReturnTicketID.String(),
		}
	}
	return out
}

func ParseBookingsEntToPb(in []*ent.Booking) (out *pb.Bookings) {
	out = &pb.Bookings{
		Bookings: make([]*pb.Booking, len(in)),
	}
	for i, a := range in {
		out.Bookings[i] = ParseBookingEntToPb(a)
	}
	return out
}
