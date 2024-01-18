package controllers

import (
	welcome "example/todoapp/views/todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(c *gin.Context) {
	//create ui using templ and htmx
	//c.IndentedJSON(http.StatusOK, map[string]interface{}{"message": "welcome all"})
	c.HTML(http.StatusOK, "", welcome.Show("welcome all"))
}
