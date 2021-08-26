package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"todoGoGo/graph/generated"
	"todoGoGo/graph/model"
	"todoGoGo/internal/service"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.ProductInput) (*model.ProductType, error) {
	result,err:=service.CreateProduct(ctx,input)
	if err != nil{
		return nil, err
	}
	return result,nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CProduct(ctx context.Context, input *model.PInput) (*model.PType, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
