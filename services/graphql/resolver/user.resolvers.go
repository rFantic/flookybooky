package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"flookybooky/pb"
	"flookybooky/services/graphql/gql_generated"
	"flookybooky/services/graphql/model"

	"github.com/jinzhu/copier"
)

// Customer is the resolver for the customer field.
func (r *userResolver) Customer(ctx context.Context, obj *model.User) (*model.Customer, error) {
	var out *model.Customer
	if obj.Customer != nil {
		out = &model.Customer{}
		req := &pb.GetCustomerRequest{
			Id: obj.Customer.ID,
		}
		customerRes, err := r.client.CustomerClient.GetCustomer(ctx, req)
		if err != nil {
			return nil, err
		}

		copier.Copy(&out, customerRes)
		out.ID = customerRes.Id
		out.LicenseID = customerRes.LicenseId
	}
	return out, nil
}

// User returns gql_generated.UserResolver implementation.
func (r *Resolver) User() gql_generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }