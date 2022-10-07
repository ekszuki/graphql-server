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
	model, err := r.VideoRepo.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// DeleteVideo is the resolver for the deleteVideo field.
func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (*model.Video, error) {
	if id == "" {
		return nil, fmt.Errorf("invalid id")
	}

	err := r.VideoRepo.Delete(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cannot delete document %s: %v", id, err)
	}

	return nil, nil
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	videos, err := r.VideoRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
