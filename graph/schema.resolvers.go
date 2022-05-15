package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gislihr/gammon/graph/generated"
	"github.com/gislihr/gammon/graph/model"
	"github.com/gislihr/gammon/pkg/gammon/dataloader"
	internalModel "github.com/gislihr/gammon/pkg/gammon/model"
)

func (r *gameResolver) Loser(ctx context.Context, obj *internalModel.Game) (*model.Player, error) {
	return dataloader.For(ctx).PlayerById.Load(obj.LoserId)
}

func (r *gameResolver) Winner(ctx context.Context, obj *internalModel.Game) (*model.Player, error) {
	return dataloader.For(ctx).PlayerById.Load(obj.WinnerId)
}

func (r *gameResolver) Tournament(ctx context.Context, obj *internalModel.Game) (*model.Tournament, error) {
	return dataloader.For(ctx).TournamentById.Load(obj.TournamentId)
}

func (r *queryResolver) Player(ctx context.Context, id int) (*model.Player, error) {
	return dataloader.For(ctx).PlayerById.Load(id)
}

func (r *queryResolver) Players(ctx context.Context, request model.PlayerRequest) ([]*model.Player, error) {
	return r.store.GetPlayers(playerRequestToDb(request))
}

func (r *queryResolver) Games(ctx context.Context, request model.GameRequest) ([]*internalModel.Game, error) {
	return r.store.GetGames(gameRequestToDb(request))
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type gameResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
