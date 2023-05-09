package resolver

import (
	"context"
	"flookybooky/internal/util"
	"flookybooky/services/graphql/gql_generated"
	"flookybooky/services/graphql/model"
	pb "flookybooky/services/graphql/proto"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
			c, _ := ctx.Value(util.ContextKey{}).(*gin.Context)
			tokenString, err := c.Cookie("Authentication")
			if err != nil {
				// c.AbortWithStatus(http.StatusUnauthorized)
				return nil, fmt.Errorf("missing authentication cookie")
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return util.Secretkey, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if float64(time.Now().Unix()) > claims["exp"].(float64) {
					// c.AbortWithStatus(http.StatusUnauthorized)
					return nil, fmt.Errorf("token expired")
				}
				id := claims["sub"].(string)
				user := &pb.GetUserRequest{Id: id}
				res, err := client.UserClient.GetUser(ctx, user)
				if err != nil {
					return nil, fmt.Errorf("claims user not found")
				}
				ctxRole := model.Role(res.User.Role)
				if role != ctxRole {
					return nil, fmt.Errorf("current role not qualified")
				}
			} else {
				fmt.Println(err)
			}
			return next(ctx)
		},
	}

	return gql_generated.NewExecutableSchema(gql_generated.Config{
		Resolvers:  &Resolver{client},
		Directives: d,
	})
}
