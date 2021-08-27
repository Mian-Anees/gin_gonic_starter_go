package test

import (
	"context"
	"testing"
	"todoGoGo/graph/model"
	"todoGoGo/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	result, err := service.CreateProduct(context.Background(), &model.ProductInput{Name: "testcase", Type: "1", Available: false})
	if err != nil {
		assert.Error(t, err, "test failed")
	}
	assert.Contains(t, *result.Name, "testcase")
	// } else if string(res) == `{"id":"","name":"testcase","type":"1","available":false}` {
	// 	fmt.Println(string(res))
	// }
}
