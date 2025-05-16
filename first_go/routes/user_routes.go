package routes

import (
	"github.com/frozenshield/first_go/controllers"
	"github.com/frozenshield/first_go/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	users := r.Group("/api/users")
	users.Use(middleware.JWTAuthMiddleware())

	{
		users.POST("/login", controllers.LoginUser)
		users.POST("/register", controllers.CreateUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}

}
