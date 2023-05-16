package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/graphql/gql_generated"
	"flookybooky/services/graphql/internal"
	"flookybooky/services/graphql/model"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Origin is the resolver for the origin field.
func (r *flightResolver) Origin(ctx context.Context, obj *model.Flight) (*model.Airport, error) {
	var out *model.Airport
	if obj.Origin != nil {
		originRes, err := r.client.FlightClient.GetAirport(ctx, &pb.UUID{Id: obj.Origin.ID})
		if err != nil {
			return nil, err
		}
		out = internal.ParseAirportPbToGraphql(originRes)
	}
	return out, nil
}

// Destination is the resolver for the destination field.
func (r *flightResolver) Destination(ctx context.Context, obj *model.Flight) (*model.Airport, error) {
	var out *model.Airport
	if obj.Destination != nil {
		destinationRes, err := r.client.FlightClient.GetAirport(ctx, &pb.UUID{Id: obj.Destination.ID})
		if err != nil {
			return nil, err
		}
		out = internal.ParseAirportPbToGraphql(destinationRes)
	}
	return out, nil
}

// CreateFlight is the resolver for the createFlight field.
func (r *mutationResolver) CreateFlight(ctx context.Context, input model.FlightInput) (*model.Flight, error) {
	flightReq, err := internal.ParseFlightInputGraphqlToPb(&input)
	if err != nil {
		return nil, err
	}
	flightRes, err := r.client.FlightClient.PostFlight(ctx, flightReq)
	return internal.ParseFlightPbToGraphql(flightRes), err
}

// Flight is the resolver for the flight field.
func (r *queryResolver) Flight(ctx context.Context) ([]*model.Flight, error) {
	flightsRes, err := r.client.FlightClient.GetFlights(ctx, &emptypb.Empty{})
	return internal.ParseFlightsPbToGraphql(flightsRes), err
}

// Flight returns gql_generated.FlightResolver implementation.
func (r *Resolver) Flight() gql_generated.FlightResolver { return &flightResolver{r} }

type flightResolver struct{ *Resolver }