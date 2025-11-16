// Package security provides authentication, authorization,
// and related security utilities for the application.
package security

import (
	"AvitoPRService/internal/config"
	"AvitoPRService/internal/model/response/error_response"
	"AvitoPRService/internal/repository/postgres"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// secret is the secret key used to sign the JWT tokens.
var secret = []byte(os.Getenv("SECRET_STRING"))

// Claims represents the JWT claims.
type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

// Validate validates the Claims.
func (c *Claims) Validate() error {
	if c.UserID == "" {
		return errors.New("userID is required")
	}
	return nil
}

// NewClaims creates a new Claims instance.
func NewClaims(userID string) *Claims {
	expirationTime := time.Now().Add(24 * time.Hour)
	return &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
}

// GenerateJWT generates a JWT token for the given userID.
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(userID))
	value, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return value, nil
}

// AdminAuthRequired is a middleware that checks if the request is authorized.
func AdminAuthRequired(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || authHeader != "Bearer "+config.AccessToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse(response.NotFound, postgres.ErrUserNotFound.Error()))
			return
		}
		c.Next()
	}
}

// ValidateJWT validates the JWT token and sets the UserID in the context
func ValidateJWT(c *gin.Context) {
	rawToken, err := c.Cookie("access_token")
	if rawToken == "" || err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(rawToken, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil && !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims := token.Claims.(*Claims)
	c.Set("UserID", claims.UserID)
	c.Next()
}
