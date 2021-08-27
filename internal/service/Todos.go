package service

import (
	"errors"
	"fmt"
	"todoGoGo/config"
	"todoGoGo/models"
)

//fetch all todos at once
func GetAllTodos() (todo *[]models.Todo, err error) {
	fmt.Printf("+++++");
	
	if todo == nil {
		return nil, errors.New("something went wrong")
	}
	_ = config.DB.Find(&todo).Error; 
	fmt.Printf("+++++ %+v",todo);
	 if err != nil {
		return nil, err
	}
	return todo, nil
}

//insert a todo
func CreateATodo(todo *models.Todo) (result *models.Todo,err error) {
	fmt.Print(todo, "todo--------")
	 _ = config.DB.Create(todo);
		fmt.Print("todo----++----")
	return &models.Todo{Title:"name",Description:"long"},nil
}

//fetch one todo
func GetATodo(todo *models.Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//update a todo
func UpdateATodo(todo *models.Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

//delete a todo
func DeleteATodo(todo *models.Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
