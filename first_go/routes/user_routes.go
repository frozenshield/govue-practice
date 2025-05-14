package routes

import (
	"github.com/frozenshield/first_go/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	users := r.Group("/api/users")

	users.POST("/login", controllers.LoginUser)
	users.POST("/register", controllers.CreateUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)

}
