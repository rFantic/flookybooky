package handler

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/customer/ent"
	"flookybooky/services/customer/ent/customer"
	"flookybooky/services/customer/internal"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	client ent.Client
}

func NewCustomerHandler(client ent.Client) (*CustomerHandler, error) {
	return &CustomerHandler{
		client: client,
	}, nil
}

func (h *CustomerHandler) GetCustomer(ctx context.Context, req *pb.UUID) (*pb.Customer, error) {
	neededID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	customerRes, err := h.client.Customer.Query().
		Where(customer.ID(neededID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseCustomerEntToPb(customerRes), nil
}

func (h *CustomerHandler) GetCustomers(ctx context.Context, req *emptypb.Empty) (*pb.Customers, error) {
	customersRes, err := h.client.Customer.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseCustomersEntToPb(customersRes), nil
}

func (h *CustomerHandler) PostCustomer(ctx context.Context, req *pb.CustomerInput) (*pb.Customer, error) {
	customerRes, err := h.client.Customer.Create().
		SetName(req.Name).
		SetAddress(req.Address).
		SetLicenseID(req.LicenseId).
		SetName(req.Name).
		SetPhoneNumber(req.PhoneNumber).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseCustomerEntToPb(customerRes), err
}
