package routes

import (
	"github.com/frozenshield/first_go/controllers"
	"github.com/frozenshield/first_go/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {

	posts := r.Group("/api/posts")
	posts.Use(middleware.JWTAuthMiddleware())
	{

		posts.GET("/", controllers.GetPosts)
		//posts.GET("/:id", controllers.GetPost)
		posts.POST("/", controllers.CreatePost)
		posts.PUT("/:id", controllers.UpdatePost)
		//posts.DELETE("/:id", controllers.DeletePost)
	}
}
