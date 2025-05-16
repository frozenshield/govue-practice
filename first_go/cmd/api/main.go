package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/frozenshield/first_go/database"
	"github.com/frozenshield/first_go/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	ginMode := os.Getenv("GIN_MODE")

	if ginMode == gin.ReleaseMode || strings.TrimSpace(ginMode) == "" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database

	database.ConnectDB()

	r := gin.Default()

	//cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.PostRoutes(r)
	routes.UserRoutes(r)
	routes.ProtectedRoutes(r)

	r.Run(":8080")

}
