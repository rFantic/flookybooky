package internal

import (
	"flookybooky/pb"
	"flookybooky/services/flight/ent"

	"github.com/jinzhu/copier"
)

func ParseAirportEntToPb(in *ent.Airport) (out *pb.Airport) {
	if in == nil {
		return nil
	}
	copier.Copy(&out, in)
	out.Id = in.ID.String()
	return out
}

func ParseAirportsEntToPb(in []*ent.Airport) (out *pb.Airports) {
	out = &pb.Airports{}
	for i, a := range in {
		out.Airports[i] = ParseAirportEntToPb(a)
	}
	return out
}
