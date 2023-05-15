package internal

import (
	"flookybooky/pb"
	"flookybooky/services/customer/ent"

	"github.com/jinzhu/copier"
)

func ParseCustomerEntToPb(in *ent.Customer) (out *pb.Customer) {
	out = &pb.Customer{}
	copier.Copy(&out, in)
	if in != nil {
		out.Id = in.ID.String()
		out.LicenseId = in.LicenseID
	}
	return out
}

func ParseCustomersEntToPb(in []*ent.Customer) (out *pb.Customers) {
	out = &pb.Customers{
		Customers: make([]*pb.Customer, len(in)),
	}
	for i, a := range in {
		out.Customers[i] = ParseCustomerEntToPb(a)
	}
	return out
}
