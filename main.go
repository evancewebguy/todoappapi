package main

import (
	conf "example/todoapp/config"
	"example/todoapp/routes"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	//Loads env file
	LoadEnvFile()

	// Connect to the database
	conf.ConnectDatabase()

	// Initialize the Gin router and setup routes
	routes.InitializeServer()
}

func LoadEnvFile() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
		return
	}
}
