package routes

import (
	"example/todoapp/controllers"
	"github.com/gin-gonic/gin"
)

// InitializeServer initializes the Gin router and sets up routes
func InitializeServer() {
	// Initialize the Gin router
	router := gin.Default()

	router.HTMLRender = &TemplRender{}

	// Your Gin routes and handlers go here
	Routes(router)

	// Run the Gin server
	if err := router.Run(":8080"); err != nil {
		// Handle the error
	}
}

// Routes sets up routes
func Routes(router *gin.Engine) {
	// Your routes go here
	router.GET("/welcome", controllers.Welcome)
	router.GET("/users", controllers.GetUsers)
	router.GET("/user/add", controllers.AddUser)
	router.GET("/user/:id", controllers.GetUser)
	router.GET("/user/update/:id", controllers.UpdateUser)
	router.GET("/user/delete/:id", controllers.DeleteUser)

}
