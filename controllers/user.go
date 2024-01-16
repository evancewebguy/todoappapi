package controllers

import (
	"errors"
	"example/todoapp/config"
	"example/todoapp/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUsers(c *gin.Context) {
	// Query the database to get all users
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Return the list of users as JSON
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {

	// Retrieve the user ID from the URL parameter
	userID := c.Param("id")

	var user models.User

	// Assuming db is your GORM database instance
	if err := config.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the user with the specified ID is not found
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Handle other database errors
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func AddUser(c *gin.Context) {
	var user models.User

	user.Name = "james"

	if err := config.DB.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the user with the specified ID is not found
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Handle other database errors
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)

}

func UpdateUser(c *gin.Context) {
	var user models.User

	userID := c.Param("id")

	if err := config.DB.Model(&user).Update("id", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the user with the specified ID is not found
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// Handle other database errors
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User

	userID := c.Param("id")

	if err := config.DB.Delete(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where the user with the specified ID is not found
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// Handle other database errors
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": "User has been deleted successfully"})
}
