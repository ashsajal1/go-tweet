package auth

import (
	"github.com/gin-gonic/gin"
)
// SetupAuthRouter sets up the authentication routes
func SetupAuthRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", Login)
		auth.POST("/signup", Signup)
	}
}
