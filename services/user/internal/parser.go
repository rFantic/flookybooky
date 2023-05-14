package internal

import (
	"flookybooky/pb"
	"flookybooky/services/user/ent"

	"github.com/jinzhu/copier"
)

func ParseUserEntToPb(in *ent.User) (out *pb.User) {
	out = &pb.User{Customer: &pb.Customer{}}
	copier.Copy(&out, in)
	if in != nil {
		out.Id = in.ID.String()
		out.Customer.Id = in.CustomerID
	}
	return out
}

func ParseUsersEntToPb(in []*ent.User) (out *pb.Users) {
	out = &pb.Users{
		Users: make([]*pb.User, len(in)),
	}
	for i, a := range in {
		out.Users[i] = ParseUserEntToPb(a)
	}
	return out
}
