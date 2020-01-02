package auth

import (
	"log"
	"net/http"
	"pages/auth/jwt"
	"pages/constants"

	"github.com/gin-gonic/gin"
)

// JWT auth middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) > 0 {
			log.Print("get token: ", token)

			j := jwt.New(constants.JWTSigningKey)
			claims, err := j.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  -1,
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			c.Set("claims", claims)
		}
	}
}
