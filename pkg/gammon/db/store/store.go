package store

import (
	"github.com/gislihr/gammon/graph/model"
	internalModel "github.com/gislihr/gammon/pkg/gammon/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) (*Store, error) {
	return &Store{
		db: db,
	}, nil
}

type PlayerRequest struct {
	Offset int
	Limit  int
}

type GameRequest struct {
	Offset int
	Limit  int
}

type player struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Elo        float64 `db:"elo"`
	ShortName  *string `db:"short_name"`
	Email      *string `db:"email"`
	Experience int     `db:"experience"`
}

func (p player) toModel() *model.Player {
	return &model.Player{
		ID:          p.Id,
		Name:        p.Name,
		Elo:         p.Elo,
		ShortName:   *p.ShortName,
		Experiennce: p.Experience,
		Email:       p.Email,
	}
}

func playersToModels(ps []player) []*model.Player {
	res := make([]*model.Player, len(ps))
	for i, p := range ps {
		res[i] = p.toModel()
	}
	return res
}

type game struct {
	Id       int `db:"id"`
	WinnerId int `db:"winner_id"`
	LoserId  int `db:"loser_id"`
	Length   int `db:"length"`
}

func (g game) toModel() *internalModel.Game {
	return &internalModel.Game{
		Id:       g.Id,
		WinnerId: g.WinnerId,
		LoserId:  g.LoserId,
		Length:   g.Length,
	}
}

func gamesToModels(gs []game) []*internalModel.Game {
	res := make([]*internalModel.Game, len(gs))
	for i, p := range gs {
		res[i] = p.toModel()
	}
	return res
}

func (s *Store) GetPlayerByID(id int) (*model.Player, error) {
	p := player{}
	err := s.db.Get(&p,
		"select id, name, elo, short_name, email, experience from player where id = $1", id)
	return p.toModel(), err
}

func (s *Store) GetPlayers(pr PlayerRequest) ([]*model.Player, error) {
	var res []player
	err := s.db.Select(&res,
		"select id, name, elo, short_name, email, experience from player limit $1 offset $2",
		pr.Limit, pr.Offset)

	return playersToModels(res), err
}

func (s *Store) GetGames(gr GameRequest) ([]*internalModel.Game, error) {
	var res []game
	err := s.db.Select(&res, "select id, winner_id, loser_id, length from game limit $1 offset $2", gr.Limit, gr.Offset)
	return gamesToModels(res), err
}
