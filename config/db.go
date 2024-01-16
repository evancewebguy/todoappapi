package config

import (
	"example/todoapp/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB // Declare a global variable to hold the database connection

// ConnectDatabase Connect to the database
func ConnectDatabase() {
	// Replace these with your actual MySQL database credentials
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")

	// Create a DSN (Data Source Name) string : MYSQL
	//dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Create a DSN (Data Source Name) string : POSTGRES
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port + ""

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to database")
	}

	// Set the global variable to the connected database
	DB = db

	// Migrate the schema // Replace YourModel with your actual model struct
	if err := db.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		logrus.WithError(err).Fatal("Failed to AutoMigrate Table to database")
	}
}
