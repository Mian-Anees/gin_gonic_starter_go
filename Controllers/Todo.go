package Controllers

import (
	"fmt"
	"net/http"
	"todoGoGo/internal/service"
	"github.com/gin-gonic/gin"
	"todoGoGo/Models"
)

//List all todos
func GetTodos(c *gin.Context){

	var todo []Models.Todo
	err := service.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Create a Todo
func CreateATodo(c *gin.Context){
	var todo Models.Todo
	fmt.Printf("%p is controller",c)
	c.BindJSON(&todo)
	err := service.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Get a particular Todo with id
func GetATodo(c *gin.Context){
	id := c.Params.ByName("id")
	var todo Models.Todo
	err := service.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Update an existing Todo
func UpdateATodo(c *gin.Context){
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := service.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = service.UpdateATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Delete a Todo
func DeleteATodo(c *gin.Context){
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := service.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
