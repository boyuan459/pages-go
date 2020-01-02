package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT struct
type JWT struct {
	SigningKey []byte
}

// CustomClaims custom claims
type CustomClaims struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// New create an jwt
func New(key string) *JWT {
	return &JWT{[]byte(key)}
}

//CreateToken create jwt token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//ParseToken parse token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if _, ok := err.(*jwt.ValidationError); ok {
			return nil, err
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//RefreshToken refresh token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", errors.New("invalid token")
}
