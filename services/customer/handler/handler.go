package handler

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/customer/ent"

	"github.com/jinzhu/copier"
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

func (h *CustomerHandler) PostCustomer(ctx context.Context, customer *pb.Customer) (*emptypb.Empty, error) {
	err := h.client.Customer.Create().
		SetName(customer.Name).
		SetAddress(customer.Address).
		SetLicenseID(customer.LicenseId).
		SetName(customer.Name).
		SetPhoneNumber(customer.PhoneNumber).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *CustomerHandler) GetCustomers(ctx context.Context, req *emptypb.Empty) (*pb.GetCustomersResponse, error) {
	customers, err := h.client.Customer.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	res := &pb.GetCustomersResponse{}
	err = copier.Copy(&res.Customers, &customers)
	if err != nil {
		return nil, err
	}
	return res, nil
}
