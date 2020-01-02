package auth

import (
	"fmt"
	"net/http"
	"pages/auth/jwt"
	"pages/model/casbin"

	"github.com/gin-gonic/gin"
)

func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*jwt.CustomClaims)
		role := claims.Role
		fmt.Println("role", role)
		fmt.Println("path", c.Request.URL.Path)
		fmt.Println("method", c.Request.Method)
		e := casbin.Enforcer()
		//check perm
		res, err := e.EnforceSafe(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  -1,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		if res != true {
			c.JSON(http.StatusOK, gin.H{
				"status":  0,
				"message": "No permission",
			})
			c.Abort()
			return
		}
	}
}
