package handler

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/flight/ent"
	"flookybooky/services/flight/internal"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FlightHandler struct {
	pb.UnimplementedFlightServiceServer
	client ent.Client
}

func NewFlightHandler(client ent.Client) (*FlightHandler, error) {
	return &FlightHandler{
		client: client,
	}, nil
}

func (h *FlightHandler) PostAirport(ctx context.Context, req *pb.Airport) (*pb.Airport, error) {
	query := h.client.Airport.Create().
		SetAddress(req.Address).
		SetName(req.Name)
	airportRes, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseAirportEntToPb(airportRes), err
}

func (h *FlightHandler) GetAirport(ctx context.Context, req *pb.UUID) (*pb.Airport, error) {
	_uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	airportRes, err := h.client.Airport.Get(ctx, _uuid)
	return internal.ParseAirportEntToPb(airportRes), err
}

func (h *FlightHandler) GetAirports(ctx context.Context, req *emptypb.Empty) (*pb.Airports, error) {
	query := h.client.Airport.Query()
	airportsRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseAirportsEntToPb(airportsRes), nil
}
