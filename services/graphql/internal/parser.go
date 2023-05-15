package internal

import (
	"flookybooky/pb"
	"flookybooky/services/graphql/model"

	"github.com/jinzhu/copier"
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

func ParseCustomerInputGraphqlToPb(in *model.CustomerInput) (out *pb.Customer) {
	out = &pb.Customer{}
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

func ParseFlightInputGraphqlToPb(in *model.FlightInput) (out *pb.Flight) {
	out = &pb.Flight{}
	copier.Copy(&out, in)
	if in != nil {
		out.Origin = &pb.Airport{
			Id: in.OriginID,
		}
		out.Destination = &pb.Airport{
			Id: in.DestinationID,
		}
	}
	return out
}

func ParseFlightPbToGraphql(in *pb.Flight) (out *model.Flight) {
	out = &model.Flight{}
	copier.Copy(&out, in)
	if in.Origin != nil {
		out.Origin = &model.Airport{
			ID: in.Origin.Id,
		}
	}
	if in.Destination != nil {
		out.Destination = &model.Airport{
			ID: in.Destination.Id,
		}
	}
	return out
}

// func ParseCustomersPbToGraphql(in *pb.Customers) (out []*model.Customer) {
// 	out = make([]*model.Customer, len(in.Customers))
// 	for i, a := range in.Customers {
// 		out[i] = ParseCustomerPbToGraphql(a)
// 	}
// 	return out
// }
