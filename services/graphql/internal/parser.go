package internal

import (
	"flookybooky/pb"
	"flookybooky/services/graphql/model"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ParseAirportInputGraphqlToPb(in *model.AirportInput) (out *pb.Airport) {
	out = &pb.Airport{}
	copier.Copy(&out, in)
	return out
}

func ParseAirportPbToGraphql(in *pb.Airport) (out *model.Airport) {
	out = &model.Airport{}
	copier.Copy(&out, in)
	out.ID = in.Id
	return out
}

func ParseAirportsPbToGraphql(in *pb.Airports) (out []*model.Airport) {
	out = make([]*model.Airport, len(in.Airports))
	for i, a := range in.Airports {
		out[i] = ParseAirportPbToGraphql(a)
	}
	return out
}

func ParseUserInputGraphqlToPb(in *model.UserInput) (out *pb.User) {
	out = &pb.User{}
	copier.Copy(&out, in)
	if out.Customer != nil {
		out.Customer.LicenseId = in.Customer.LicenseID
	}
	return out
}

func ParseUserPbToGraphql(in *pb.User) (out *model.User) {
	out = &model.User{}
	copier.Copy(&out, in)
	out.ID = in.Id
	if out.Customer != nil {
		out.Customer.ID = in.Customer.Id
		out.Customer.LicenseID = in.Customer.LicenseId
	}
	return out
}

func ParseUsersPbToGraphql(in *pb.Users) (out []*model.User) {
	out = make([]*model.User, len(in.Users))
	for i, a := range in.Users {
		out[i] = ParseUserPbToGraphql(a)
	}
	return out
}

func ParseCustomerInputGraphqlToPb(in *model.CustomerInput) (out *pb.CustomerInput) {
	out = &pb.CustomerInput{}
	copier.Copy(&out, in)
	out.LicenseId = in.LicenseID
	return out
}

func ParseCustomerPbToGraphql(in *pb.Customer) (out *model.Customer) {
	out = &model.Customer{}
	copier.Copy(&out, in)
	out.ID = in.Id
	out.LicenseID = in.LicenseId
	return out
}

func ParseCustomersPbToGraphql(in *pb.Customers) (out []*model.Customer) {
	out = make([]*model.Customer, len(in.Customers))
	for i, a := range in.Customers {
		out[i] = ParseCustomerPbToGraphql(a)
	}
	return out
}

func ParseFlightInputGraphqlToPb(in *model.FlightInput) (out *pb.Flight, err error) {
	if in == nil {
		return nil, fmt.Errorf("nil flight input")
	}
	_departureTime, err := time.Parse("2006-01-02", in.DepartureTime)
	if err != nil {
		return nil, err
	}
	_arrivalTime, err := time.Parse("2006-01-02", in.ArrivalTime)
	if err != nil {
		return nil, err
	}
	out = &pb.Flight{
		DepartureTime: timestamppb.New(_departureTime),
		ArrivalTime:   timestamppb.New(_arrivalTime),
	}
	copier.Copy(&out, in)
	out.Origin = &pb.Airport{
		Id: in.OriginID,
	}
	out.Destination = &pb.Airport{
		Id: in.DestinationID,
	}
	return out, err
}

func ParseFlightPbToGraphql(in *pb.Flight) (out *model.Flight) {
	if in == nil {
		return nil
	}
	out = &model.Flight{
		ID:            in.GetId(),
		DepartureTime: in.DepartureTime.AsTime().String(),
		ArrivalTime:   in.ArrivalTime.AsTime().String(),
	}
	copier.Copy(&out, in)
	if in.Origin != nil {
		out.Origin = &model.Airport{
			ID: in.Origin.Id,
		}
		copier.Copy(&out.Origin, in.Origin)
	}
	if in.Destination != nil {
		out.Destination = &model.Airport{
			ID: in.Destination.Id,
		}
		copier.Copy(&out.Destination, in.Destination)
	}
	return out
}

func ParseFlightsPbToGraphql(in *pb.Flights) (out []*model.Flight) {
	out = make([]*model.Flight, len(in.Flights))
	for i, a := range in.Flights {
		out[i] = ParseFlightPbToGraphql(a)
	}
	return out
}

// func ParseCustomersPbToGraphql(in *pb.Customers) (out []*model.Customer) {
// 	out = make([]*model.Customer, len(in.Customers))
// 	for i, a := range in.Customers {
// 		out[i] = ParseCustomerPbToGraphql(a)\

// 	return out
// }

func ParseBookingInputGraphqlToPb(in *model.BookingInput) (out *pb.BookingInput) {
	if in == nil {
		return nil
	}
	out = &pb.BookingInput{
		CustomerId:    in.CustomerID,
		GoingFlightId: in.GoingFlightID,
	}
	if in.ReturnFlightID != nil {
		out.ReturnFlightId = in.ReturnFlightID
	}
	copier.Copy(&out, in)
	return out
}

func ParseBookingPbToGraphql(in *pb.Booking) (out *model.Booking) {
	if in == nil {
		return nil
	}
	out = &model.Booking{
		ID: in.GetId(),
	}
	if in.Customer != nil {
		out.Customer = &model.Customer{
			ID:        in.Customer.GetId(),
			LicenseID: in.Customer.GetLicenseId(),
		}
		copier.Copy(&out.Customer, in.Customer)
	}
	if in.GoingFlight != nil {
		out.GoingFlight = &model.Flight{
			ID: in.GoingFlight.Id,
		}
		copier.Copy(&out.GoingFlight, in.GoingFlight)
	}
	if in.ReturnFlight != nil {
		out.ReturnFlight = &model.Flight{
			ID: in.ReturnFlight.Id,
		}
		copier.Copy(&out.ReturnFlight, in.ReturnFlight)
	}
	return out
}

func ParseBookingsPbToGraphql(in *pb.Bookings) (out []*model.Booking) {
	if in == nil {
		return nil
	}
	out = make([]*model.Booking, len(in.GetBookings()))
	for i, a := range in.Bookings {
		out[i] = &model.Booking{
			ID: a.Id,
		}
		copier.Copy(&out[i], a)
		if a.Customer != nil {
			out[i].Customer = &model.Customer{
				ID:        a.Customer.Id,
				LicenseID: a.Customer.LicenseId,
			}
			copier.Copy(&out[i].Customer, a.Customer)
		}
		if a.GoingFlight != nil {
			out[i].GoingFlight = &model.Flight{
				ID: a.GoingFlight.Id,
			}
			copier.Copy(&out[i].GoingFlight, a.GoingFlight)
		}
		if a.ReturnFlight != nil {
			out[i].ReturnFlight = &model.Flight{
				ID: a.ReturnFlight.Id,
			}
			copier.Copy(&out[i].ReturnFlight, a.ReturnFlight)
		}
	}
	return out
}
