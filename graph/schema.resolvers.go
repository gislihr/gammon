package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gislihr/gammon/graph/generated"
	"github.com/gislihr/gammon/graph/model"
	internalModel "github.com/gislihr/gammon/pkg/gammon/model"
	"github.com/gislihr/gammon/store"
)

func (r *gameResolver) Loser(ctx context.Context, obj *internalModel.Game) (*model.Player, error) {
	return r.store.GetPlayerByID(obj.LoserId)
}

func (r *gameResolver) Winner(ctx context.Context, obj *internalModel.Game) (*model.Player, error) {
	return r.store.GetPlayerByID(obj.WinnerId)
}

func (r *mutationResolver) AddPlayer(ctx context.Context, name string) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Player(ctx context.Context, id int) (*model.Player, error) {
	return r.store.GetPlayerByID(id)
}

func (r *queryResolver) Players(ctx context.Context, limit int, offset int) ([]*model.Player, error) {
	return r.store.GetPlayers(store.PlayerRequest{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *queryResolver) Games(ctx context.Context, limit int, offset int) ([]*internalModel.Game, error) {
	return r.store.GetGames(store.GameRequest{
		Limit:  limit,
		Offset: offset,
	})
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type gameResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *gameResolver) ID(ctx context.Context, obj *internalModel.Game) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
