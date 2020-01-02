package user

import (
	"fmt"
	"log"
	"net/http"
	"pages/auth/jwt"
	"pages/constants"
	"pages/model/user"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginResult struct {
	Token    string `json:"token"`
	Username string
	Role     string
}

func Create(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model user.User
		err := c.BindJSON(&model)
		if role != "" {
			model.Role = role
		}
		result := model.FindUserByUsername(model.Username)
		fmt.Println("result", result)
		fmt.Println("username", model.Username)
		if result.Username == model.Username {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "user already existed",
			})
			return
		}
		hashpwd, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln(err)
		}
		model.Password = string(hashpwd)
		err = model.Create(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "create failed",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "create success",
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model user.User
		err := c.BindJSON(&model)
		result := model.FindUserByUsername(model.Username)

		fmt.Println("username", model.Username)
		fmt.Println("result", result)
		if result.Username != model.Username {
			c.JSON(http.StatusOK, gin.H{
				"status":  -1,
				"message": "login failure",
			})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(model.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "username or password is incorrect",
			})
			return
		}

		j := jwt.New(constants.JWTSigningKey)

		claims := jwt.CustomClaims{
			result.Username,
			result.Phone,
			result.Role,
			jwtgo.StandardClaims{
				NotBefore: int64(time.Now().Unix() - 1000),
				ExpiresAt: int64(time.Now().Unix() + 3600), //expired in 1 hour
				Issuer:    constants.JWTIssuer,
			},
		}

		token, err := j.CreateToken(claims)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  -1,
				"message": err.Error(),
			})
			return
		}
		data := LoginResult{
			Token:    token,
			Username: result.Username,
			Role:     result.Role,
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "login success",
			"data":    data,
		})
	}
}
