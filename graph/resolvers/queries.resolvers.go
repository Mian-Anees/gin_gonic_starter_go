package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"todoGoGo/graph/generated"
	"todoGoGo/graph/model"
)

func (r *queryResolver) GProduct(ctx context.Context) ([]*model.PType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProducts(ctx context.Context) ([]*model.ProductType, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
