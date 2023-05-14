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
	for i, a := range in.Airports {
		out[i] = ParseAirportPbToGraphql(a)
	}
	return out
}
