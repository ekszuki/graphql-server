package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ekszuki/graphhql-server/graph/generated"
	"github.com/ekszuki/graphhql-server/graph/model"
)

// CreateVideo is the resolver for the createVideo field.
func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	panic(fmt.Errorf("not implemented: CreateVideo - createVideo"))
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	panic(fmt.Errorf("not implemented: Videos - videos"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
