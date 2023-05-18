package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"flookybooky/services/graphql/internal"
	"flookybooky/services/graphql/model"
)

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	customerRes, err := r.client.CustomerClient.PostCustomer(ctx,
		internal.ParseCustomerInputGraphqlToPb(&input),
	)
	if err != nil {
		return nil, err
	}
	return internal.ParseCustomerPbToGraphql(customerRes), nil
}

// UpdateCustomer is the resolver for the updateCustomer field.
func (r *mutationResolver) UpdateCustomer(ctx context.Context, input model.CustomerUpdateInput) (bool, error) {
	_, err := r.client.CustomerClient.UpdateCustomer(ctx,
		internal.ParseCustomerUpdateInputGraphqlToPb(&input))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context, input *model.Pagination) ([]*model.Customer, error) {
	customersRes, err := r.client.CustomerClient.GetCustomers(ctx,
		internal.ParsePaginationGraphqlToPb(input))
	return internal.ParseCustomersPbToGraphql(customersRes), err
}
