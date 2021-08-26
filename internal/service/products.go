package service

import (
	"context"
	"todoGoGo/graph/model"
)


func CreateProduct(ctx context.Context, input *model.ProductInput) (*model.ProductType, error) {
	return &model.ProductType{
		ID:        new(string),
		Name:      &input.Name,
		Type:      &input.Type,
		Available: &input.Available,
	},nil
}