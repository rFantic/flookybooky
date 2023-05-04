package resolver

import (
	"flookybooky/graphql/gql_generated"
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
	return gql_generated.NewExecutableSchema(gql_generated.Config{
		Resolvers: &Resolver{client},
	})
}
