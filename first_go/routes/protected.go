package routes

import (
	"net/http"

	"github.com/frozenshield/first_go/middleware"
	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(r *gin.Engine) {
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Your are authorized!"})
	})
}
