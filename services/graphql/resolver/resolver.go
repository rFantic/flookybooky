package resolver

import (
	"context"
	"flookybooky/internal/util"
	"flookybooky/services/graphql/gql_generated"
	"flookybooky/services/graphql/model"
	pb "flookybooky/services/graphql/proto"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Client struct {
	UserClient pb.UserServiceClient
}

type Resolver struct{ client Client }

func NewSchema(client Client) graphql.ExecutableSchema {
	var d = gql_generated.DirectiveRoot{
		HasRole: func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
			ctxRole := ctx.Value(util.ReqKey)
			if role != ctxRole {
				return nil, fmt.Errorf("access denied")
			}
			return next(ctx)
		},
	}

	return gql_generated.NewExecutableSchema(gql_generated.Config{
		Resolvers:  &Resolver{client},
		Directives: d,
	})
}
