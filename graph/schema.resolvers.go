package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gislihr/gammon/graph/generated"
	"github.com/gislihr/gammon/graph/model"
)

func (r *mutationResolver) AddPlayer(ctx context.Context, name string) (*model.Player, error) {
	return r.store.InsertPlayer(name)
}

func (r *queryResolver) Player(ctx context.Context, id int) (*model.Player, error) {
	return r.store.GetPlayerByID(id)
}

func (r *queryResolver) Players(ctx context.Context) ([]*model.Player, error) {
	return r.store.GetAllPlayers()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
