package controllers

import (
	"fmt"
	"net/http"
	"todoGoGo/internal/service"
	"github.com/gin-gonic/gin"
	"todoGoGo/models"
)

//List all todos
func GetTodos(c *gin.Context){
	result,err := service.GetAllTodos()
	fmt.Printf("----in get %+v",result)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

//Create a Todo
func CreateATodo(c *gin.Context){
	var todo models.Todo
	c.BindJSON(&todo)
	fmt.Printf("%+v is controller",&todo)
	_,err := service.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Get a particular Todo with id
func GetATodo(c *gin.Context){
	id := c.Params.ByName("id")
	var todo models.Todo
	err := service.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Update an existing Todo
func UpdateATodo(c *gin.Context){
	var todo models.Todo
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
	var todo models.Todo
	id := c.Params.ByName("id")
	err := service.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
