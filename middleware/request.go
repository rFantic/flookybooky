package middleware

import (
	"context"
	"flookybooky/internal/util"
	"flookybooky/services/graphql/model"

	"github.com/gin-gonic/gin"
)

func RequestCtxMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), util.ReqKey, model.RoleAdmin)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
