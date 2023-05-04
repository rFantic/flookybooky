//go:generate go run github.com/99designs/gqlgen
package main

import (
	"context"
	"errors"
	"flookybooky/graphql/model"
	pb "flookybooky/graphql/proto"
	"flookybooky/graphql/resolver"
	"fmt"
	"log"
	"net/http"
	"strings"
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
	http.Handle("/graphql",
		authMiddleware(
			handler.GraphQL(resolver.NewSchema(client)),
		),
	)
	http.Handle("/",
		authMiddleware(
			handler.Playground("App", "/graphql"),
		),
	)
	// http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "POST" {
	// 		var user model.User
	// 		err := json.NewDecoder(r.Body).Decode(&user)
	// 		if err != nil {
	// 			http.Error(w, "Bad Request", http.StatusBadRequest)
	// 			return
	// 		}
	// 		for _, u := range users {
	// 			if u.Role == user.Role {
	// 				token, err := generateToken(u)
	// 				if err != nil {
	// 					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 					return
	// 				}
	// 				w.Write([]byte(token))
	// 				return
	// 			} else {
	// 				http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 				return
	// 			}
	// 		}
	// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 		return
	// 	}
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// })

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func dbConn() resolver.Client {
	userConn, err := grpc.Dial("user:2220", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	petConn, err := grpc.Dial("pet:2220", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient := pb.NewUserServiceClient(userConn)
	petClient := pb.NewPetServiceClient(petConn)
	return resolver.Client{PetClient: petClient, UserClient: userClient}
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

// Define a middleware function for verifying JWT tokens
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		user, err := verifyToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
