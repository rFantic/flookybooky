package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"flookybooky/services/graphql/internal"
	"flookybooky/services/graphql/model"

	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateAirport is the resolver for the createAirport field.
func (r *mutationResolver) CreateAirport(ctx context.Context, input model.AirportInput) (*model.Airport, error) {
	airportRes, err := r.client.FlightClient.PostAirport(ctx, internal.ParseAirportInputGraphqlToPb(&input))
	if err != nil {
		return nil, err
	}
	return internal.ParseAirportPbToGraphql(airportRes), nil
}

// Airport is the resolver for the airport field.
func (r *queryResolver) Airport(ctx context.Context) ([]*model.Airport, error) {
	airportsRes, err := r.client.FlightClient.GetAirports(ctx, &emptypb.Empty{})
	return internal.ParseAirportsPbToGraphql(airportsRes), err
}
