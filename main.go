package main

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize Database
	connection_string := database.GetPostgresConnectionString()
	database.Connect(connection_string)
	database.Migrate()

	// Initialize Router
	router := routes.InitRouter()
	router.Run(":8000")
}
