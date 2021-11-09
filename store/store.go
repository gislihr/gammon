package store

import (
	"fmt"

	"github.com/gislihr/gammon/graph/model"
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

type player struct {
	Id         int32   `db:"id"`
	Name       string  `db:"name"`
	Elo        float64 `db:"elo"`
	ShortName  *string `db:"short_name"`
	Email      *string `db:"email"`
	Experience int     `db:"experience"`
}

func (p player) toModel() *model.Player {
	return &model.Player{
		ID:          fmt.Sprintf("%d", p.Id),
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

func (s *Store) InsertPlayer(name string) (*model.Player, error) {
	p := player{}
	err := s.db.Get(&p, "insert into gammon.player (name, elo) values ($1, 0) returning id, name, elo", name)
	return p.toModel(), err
}

func (s *Store) GetPlayerByID(id int) (*model.Player, error) {
	p := player{}
	err := s.db.Get(&p, "select id, name, elo, short_name, email, experience from player where id = $1", id)
	return p.toModel(), err
}

func (s *Store) GetPlayers(pr PlayerRequest) ([]*model.Player, error) {
	var res []player
	err := s.db.Select(&res, "select id, name, elo, short_name, email, experience from player limit $1 offset $2", pr.Limit, pr.Offset)

	return playersToModels(res), err
}
