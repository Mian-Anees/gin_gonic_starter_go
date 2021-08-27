package test

import (
	"fmt"
	"testing"
	"todoGoGo/internal/service"
	"todoGoGo/models"
	// "github.com/stretchr/testify/assert"
)

func TestTodo(t *testing.T) {
	// var todo *models.Todo;
	 obj := models.Todo{ID:0 ,Title:"new" ,Description:"old"}
	 _,_ = service.CreateATodo(&obj)
	fmt.Print("+++++")
	// if err != nil {
	// 	assert.Error(t, err, "something went wrong")
	// 	return
	// }
	
	// assert.Nil(t, t, nil)
	// } else if string(res) == `{"id":"","name":"testcase","type":"1","available":false}` {
	// 	fmt.Println(string(res))
	// }
}
