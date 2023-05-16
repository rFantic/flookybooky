package internal

import (
	"flookybooky/pb"
	"flookybooky/services/booking/ent"

	"github.com/jinzhu/copier"
)

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
	out.GoingFlight = &pb.Flight{
		Id: in.GoingFlightID.String(),
	}
	out.ReturnFlight = &pb.Flight{
		Id: in.ReturnFlightID.String(),
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
