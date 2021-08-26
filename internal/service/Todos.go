package service

import (
	"fmt"
	"todoGoGo/Config"
	"todoGoGo/Models"
)

//fetch all todos at once
func GetAllTodos(todo *[]Models.Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

//insert a todo
func CreateATodo(todo *Models.Todo) (err error) {
	fmt.Print(todo, "todo--------")
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//fetch one todo
func GetATodo(todo *Models.Todo, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//update a todo
func UpdateATodo(todo *Models.Todo, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

//delete a todo
func DeleteATodo(todo *Models.Todo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
