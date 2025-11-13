package security

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("SECRET_STRING")

type Claims struct {
	UserID int
	jwt.RegisteredClaims
}

func (c *Claims) Valid() error {
	return c.Valid()
}

func NewClaims(userId int) *Claims {
	expirationTime := time.Now().Add(24 * time.Hour)
	return &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
}

func GenerateJWT(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(userId))
	value, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return value, nil
}

func ValidateJWT(c *gin.Context) {
	rawToken, err := c.Cookie("access_token")
	if rawToken == "" || err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(rawToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil && !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims := token.Claims.(*Claims)
	c.Set("UserId", claims.UserID)
	c.Next()
}
