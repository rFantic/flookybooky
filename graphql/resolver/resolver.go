package resolver

import (
	"context"
	"flookybooky/graphql/gql_generated"
	"flookybooky/graphql/model"
	pb "flookybooky/graphql/proto"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Client struct {
	PetClient  pb.PetServiceClient
	UserClient pb.UserServiceClient
}

type Resolver struct{ client Client }

func NewSchema(client Client) graphql.ExecutableSchema {
	var d = gql_generated.DirectiveRoot{
		HasRole: func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
			ctxUserID := ctx.Value(key)
			if ctxUserID != nil {
				return next(ctx)
			}
			return nil, ctx.Err()
		},
	}

	return gql_generated.NewExecutableSchema(gql_generated.Config{
		Resolvers: &Resolver{client},
	})
}
