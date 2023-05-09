package middleware

import (
	"context"
	"flookybooky/internal/util"

	"github.com/gin-gonic/gin"
)

// func RequestCtxMiddleware() func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		ctx := context.WithValue(c.Request.Context(), util.ReqKey, model.RoleAdmin)
// 		c.Request = c.Request.WithContext(ctx)
// 		c.Next()
// 	}
// }

func CookieMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), util.ContextKey{}, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// func RequireAuth() func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		tokenString, err := c.Cookie("Authentication")
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return util.Secretkey, nil
// 		})

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 				c.AbortWithStatus(http.StatusUnauthorized)
// 			}
// 		} else {
// 			fmt.Println(err)
// 		}
// 		c.Next()
// 	}
// }
