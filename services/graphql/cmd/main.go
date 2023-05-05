//go:generate go run github.com/99designs/gqlgen
package main

import (
	pb "flookybooky/services/graphql/proto"
	"flookybooky/services/graphql/resolver"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()
}

func main() {
	client := dbConn()
	http.Handle("/graphql", handler.GraphQL(resolver.NewSchema(client)))
	http.Handle("/", handler.Playground("App", "/graphql"))

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func dbConn() resolver.Client {
	userConn, err := grpc.Dial("user:2220", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	petConn, err := grpc.Dial("pet:2220", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient := pb.NewUserServiceClient(userConn)
	petClient := pb.NewPetServiceClient(petConn)
	return resolver.Client{PetClient: petClient, UserClient: userClient}
}
