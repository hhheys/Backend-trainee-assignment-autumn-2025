package security

import (
	"AvitoPRService/internal/config"
	"AvitoPRService/internal/db"
	errorResponse "AvitoPRService/internal/response/error_response"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("SECRET_STRING"))

type Claims struct {
	UserID uint
	jwt.RegisteredClaims
}

func (c *Claims) Valid() error {
	return c.Valid()
}

func NewClaims(userId uint) *Claims {
	expirationTime := time.Now().Add(24 * time.Hour)
	return &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
}

func GenerateJWT(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(userId))
	value, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return value, nil
}

func AdminAuthReqired(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || authHeader != "Bearer "+config.AccessToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse.NewErrorResponse(errorResponse.NOT_FOUND, db.ErrUserNotFound.Error()))
			return
		}
		c.Next()
	}
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
