package main

import (
	"log"

	"github.com/frozenshield/first_go/database"
	"github.com/frozenshield/first_go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database

	database.ConnectDB()

	r := gin.Default()

	routes.PostRoutes(r)

	r.Run(":8080")

}
