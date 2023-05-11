package handler

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/customer/ent"
	"flookybooky/services/customer/ent/customer"

	"github.com/google/uuid"
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
	res.Id = c.ID.String()
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
		res.Customers[i].Id = c.ID.String()
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *CustomerHandler) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.Customer, error) {
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

	customerOut := &pb.Customer{}
	copier.Copy(&customerOut, customerRes)
	customerOut.Id = customerRes.ID.String()
	customerOut.LicenseId = customerRes.LicenseID
	return customerOut, nil
}
