package Middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTcxNzA5NDgxNCwiaWF0IjoxNzE3MDk0ODE0fQ.1ZsPNsWkIznYaBU_rxnGHQnEPt1Cj4EhRcKLKInhcgQ")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenStr := strings.Split(authHeader, "Bearer ")[1]
		claims, err := ValidateJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorised token")
			return
		}
		c.Set("username", claims.Subject)
		c.Next()
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(userID string) (string, error) {
	// Define the expiration time for the token (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, including the user ID and expiration time
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   userID,
	}

	// Create the JWT token with the claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT validates a JWT token and returns the claims if the token is valid
func ValidateJWT(tokenString string) (*jwt.StandardClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("that's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("token is either expired or not active yet")
			} else {
				return nil, fmt.Errorf("couldn't handle this token: %v", err)
			}
		} else {
			return nil, fmt.Errorf("couldn't handle this token: %v", err)
		}
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func NoAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		return
	}
}
