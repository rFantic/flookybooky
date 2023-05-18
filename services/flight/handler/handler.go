package handler

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/flight/ent"
	"flookybooky/services/flight/ent/airport"
	"flookybooky/services/flight/ent/flight"
	"flookybooky/services/flight/internal"
	"fmt"

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

func (h *FlightHandler) GetAirports(ctx context.Context, req *pb.Pagination) (*pb.Airports, error) {
	query := h.client.Airport.Query()
	if req != nil {
		var options []airport.OrderOption
		if req.AscFields != nil {
			options = append(options, ent.Asc(req.AscFields...))
		}
		if req.DesFields != nil {
			options = append(options, ent.Desc(req.DesFields...))
		}
		query.Order(options...)
		if req.Limit != nil {
			query.Limit(int(*req.Limit))
		}
		if req.Offset != nil {
			query.Offset(int(*req.Offset))
		}
	}
	airportsRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseAirportsEntToPb(airportsRes), nil
}

func (h *FlightHandler) GetFlight(ctx context.Context, req *pb.UUID) (*pb.Flight, error) {
	_uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	flightRes, err := h.client.Flight.Get(ctx, _uuid)
	return internal.ParseFlightEntToPb(flightRes), err
}

func (h *FlightHandler) GetFlights(ctx context.Context, req *pb.Pagination) (*pb.Flights, error) {
	query := h.client.Flight.Query()
	if req != nil {
		var options []flight.OrderOption
		if req.AscFields != nil {
			options = append(options, ent.Asc(req.AscFields...))
		}
		if req.DesFields != nil {
			options = append(options, ent.Desc(req.DesFields...))
		}
		query.Order(options...)
		if req.Limit != nil {
			query.Limit(int(*req.Limit))
		}
		if req.Offset != nil {
			query.Offset(int(*req.Offset))
		}
	}
	flightsRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseFlightsEntToPb(flightsRes), nil
}

func (h *FlightHandler) PostFlight(ctx context.Context, req *pb.FlightInput) (*pb.Flight, error) {
	if req == nil {
		return nil, nil
	}
	_originReq, err := uuid.Parse(req.OriginId)
	if err != nil {
		return nil, err
	}
	_destinationReq, err := uuid.Parse(req.DestinationId)
	if err != nil {
		return nil, err
	}
	query := h.client.Flight.Create().
		SetAvailableSlots(int(req.AvailableSlots)).
		SetDepartureTime(req.DepartureTime.AsTime()).
		SetArrivalTime(req.ArrivalTime.AsTime()).
		SetName(req.Name).
		SetDestinationID(_destinationReq).
		SetOriginID(_originReq).
		SetStatus(flight.Status(req.Status))
	flightRes, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseFlightEntToPb(flightRes), err
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, req *pb.FlightUpdateInput) (*emptypb.Empty, error) {
	_flightId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	res, err := h.client.Flight.Get(ctx, _flightId)
	if err != nil {
		return nil, err
	}
	query := h.client.Flight.UpdateOneID(_flightId)
	if req.Status != nil {
		if *req.Status != "Scheduled" {
			query.SetStatus(flight.Status(*req.Status))
		}
	}
	if res.Status != "Scheduled" {
		return nil, fmt.Errorf("flight already departed")
	}
	if req.ArrivalTime != nil {
		query.SetArrivalTime(req.ArrivalTime.AsTime())
	}
	if req.DepartureTime != nil {
		query.SetDepartureTime(req.DepartureTime.AsTime())
	}
	if req.Name != nil {
		query.SetName(*req.Name)
	}
	if req.DestinationId != nil {
		_DestinationID, err := uuid.Parse(*req.DestinationId)
		if err != nil {
			return nil, err
		}
		query.SetDestinationID(_DestinationID)
	}
	if req.OriginId != nil {
		_OriginID, err := uuid.Parse(*req.OriginId)
		if err != nil {
			return nil, err
		}
		query.SetOriginID(_OriginID)
	}
	if req.AvailableSlots != nil {
		query.SetAvailableSlots(int(*req.AvailableSlots))
	}
	err = query.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
