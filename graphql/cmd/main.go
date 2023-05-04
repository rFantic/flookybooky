//go:generate go run github.com/99designs/gqlgen
package main

import (
	"errors"
	"flookybooky/graphql/model"
	pb "flookybooky/graphql/proto"
	"flookybooky/graphql/resolver"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()
}

var users = []model.User{
	{ID: "1", Role: "user"},
	{ID: "2", Role: "admin"},
}

func main() {
	// c := gql_generated.Config{Resolvers: &resolver.Resolver{}}
	// c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
	// 	if !gql_generated.getCurrentUser(ctx).HasRole(role) {
	// 		// block calling the next resolver
	// 		return nil, fmt.Errorf("Access denied")
	// 	}

	// 	// or let it pass through
	// 	return next(ctx)
	// }

	client := dbConn()
	http.Handle("/graphql", handler.GraphQL(resolver.NewSchema(client)))
	http.Handle("/", handler.Playground("App", "/graphql"))

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func dbConn() resolver.Client {
	userConn, err := grpc.Dial("user:2220", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	customerConn, err := grpc.Dial("customer:2220", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := pb.NewUserServiceClient(userConn)
	customerClient := pb.NewCustomerServiceClient(customerConn)
	return resolver.Client{CustomerClient: customerClient, UserClient: userClient}
}

// Define a function for generating JWT tokens
var secretKey = []byte("my_secret_key")

func generateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(secretKey)
}

func verifyToken(tokenString string) (*model.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := model.User{
			ID:   claims["id"].(string),
			Role: claims["role"].(string),
		}
		return &user, nil
	}
	return nil, errors.New("invalid token")
}
