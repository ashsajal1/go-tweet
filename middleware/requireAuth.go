package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// requireAuth is a middleware that checks if the user is authenticated before executing the next handler
func RequireAuth(c *gin.Context) {
	// Get the JWT token from the Authorization header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		c.Abort()
		return
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signature of the token using the secret key
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// Check if the token is valid
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// Get the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Get the user ID from the token claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Set the user ID in the Gin context
	c.Set("user_id", userID)

	// Call the next handler
	c.Next()
}
