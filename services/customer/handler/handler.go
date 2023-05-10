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

func (h *CustomerHandler) PostCustomer(ctx context.Context, customer *pb.Customer) (*pb.Customer, error) {
	c, err := h.client.Customer.Create().
		SetName(customer.Name).
		SetAddress(customer.Address).
		SetLicenseID(customer.LicenseId).
		SetName(customer.Name).
		SetPhoneNumber(customer.PhoneNumber).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	res := &pb.Customer{}
	err = copier.Copy(&res, c)
	res.Id = uint32(c.ID)
	res.LicenseId = c.LicenseID
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *CustomerHandler) GetCustomers(ctx context.Context, req *emptypb.Empty) (*pb.GetCustomersResponse, error) {
	customers, err := h.client.Customer.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	res := &pb.GetCustomersResponse{}
	err = copier.Copy(&res.Customers, &customers)
	for i, c := range customers {
		res.Customers[i].Id = uint32(c.ID)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
