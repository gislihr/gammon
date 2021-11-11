package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gislihr/gammon/graph/generated"
	"github.com/gislihr/gammon/graph/model"
	"github.com/gislihr/gammon/pkg/gammon/db"
	internalModel "github.com/gislihr/gammon/pkg/gammon/model"
)

func (r *gameResolver) Loser(ctx context.Context, obj *internalModel.Game) (*model.Player, error) {
	return r.store.GetPlayerByID(obj.LoserId)
}

func (r *gameResolver) Winner(ctx context.Context, obj *internalModel.Game) (*model.Player, error) {
	return r.store.GetPlayerByID(obj.WinnerId)
}

func (r *gameResolver) Tournament(ctx context.Context, obj *internalModel.Game) (*model.Tournament, error) {
	return r.store.GetTournament(obj.TournamentId)
}

func (r *mutationResolver) AddPlayer(ctx context.Context, name string) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Player(ctx context.Context, id int) (*model.Player, error) {
	return r.store.GetPlayerByID(id)
}

func (r *queryResolver) Players(ctx context.Context, limit int, offset int) ([]*model.Player, error) {
	return r.store.GetPlayers(db.PlayerRequest{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *queryResolver) Games(ctx context.Context, request *model.GameRequest) ([]*internalModel.Game, error) {
	return r.store.GetGames(gqlReguestToDb(request))
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
