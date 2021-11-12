package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/gislihr/gammon/graph/model"
	"github.com/gislihr/gammon/pkg/gammon/db"
)

func newLoaders(store *db.Store) *Loaders {
	return &Loaders{
		PlayerById: NewPlayerLoader(PlayerLoaderConfig{
			Fetch: func(keys []int) ([]*model.Player, []error) {
				players, err := store.GetPlayersByIds(keys)

				if err != nil {
					return players, []error{err}
				}

				playersByIdMap := make(map[int]*model.Player)
				for _, player := range players {
					playersByIdMap[player.ID] = player
				}

				playersOrdered := make([]*model.Player, len(keys))
				for i, key := range keys {
					playersOrdered[i] = playersByIdMap[key]
				}
				return playersOrdered, nil
			},
			MaxBatch: 100,
			Wait:     1 * time.Millisecond,
		}),
		TournamentById: NewTournamentLoader(TournamentLoaderConfig{
			Fetch: func(keys []int) ([]*model.Tournament, []error) {
				tournaments, err := store.GetTournaments(keys)

				if err != nil {
					return tournaments, []error{err}
				}

				tournamentsByIdMap := make(map[int]*model.Tournament)
				for _, tournament := range tournaments {
					tournamentsByIdMap[tournament.ID] = tournament
				}

				tournamentsOrdered := make([]*model.Tournament, len(keys))
				for i, key := range keys {
					tournamentsOrdered[i] = tournamentsByIdMap[key]
				}
				return tournamentsOrdered, nil
			},
			MaxBatch: 100,
			Wait:     1 * time.Millisecond,
		}),
	}
}

type contextKey string

func (c contextKey) String() string {
	return "dataloader context key " + string(c)
}

const loadersKey = contextKey("dataLoaders")

type Loaders struct {
	PlayerById     *PlayerLoader
	TournamentById *TournamentLoader
}

func Middleware(store db.Store, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, newLoaders(&store))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
