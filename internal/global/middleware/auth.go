package middleware

import (
	"gin-rush-template/internal/global/errs"
	"gin-rush-template/internal/global/jwt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if payload, valid := jwt.ParseToken(token); !valid {
			errs.Fail(c, errs.ErrTokenInvalid)
			c.Abort()
			return
		} else {
			c.Set("payload", payload)
		}
		c.Next()
	}
}
