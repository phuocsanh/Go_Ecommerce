package middlewares

import (
	"go_ecommerce/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware () gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token"{
			response.ErrResponse(c, response.ErrInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}