package middleware

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/known"
	"basetable.com/pkg/token"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		kid, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, err, nil)
			c.Abort()
		}
		c.Writer.Header().Set(known.XUsernameKey, kid)
		c.Next()
	}
}
