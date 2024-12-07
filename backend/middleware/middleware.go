package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

// Middleware authentication method that validates JWT Token
func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Step 1: Extraction. Extract token from the Authorization Header.
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort() // Stop current middleware and prevent further handler function calls.
			return
		}

		// Check if raw authHeader is valid.
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error: ": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Step 2: Parse and valiidate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next() // Allow access route handler if token is valid

	}
}
