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

type player struct {
	Id   int32  `db:"id"`
	Name string `db:"name"`
	Elo  int    `db:"elo"`
}

func (p player) toModel() *model.Player {
	return &model.Player{
		ID:   fmt.Sprintf("%d", p.Id),
		Name: p.Name,
		Elo:  p.Elo,
	}
}

func (s *Store) InsertPlayer(name string) (*model.Player, error) {
	p := player{}
	err := s.db.Get(&p, "insert into gammon.player (name, elo) values ($1, 0) returning id, name, elo", name)
	return p.toModel(), err
}

func (s *Store) GetPlayerByID(id int) (*model.Player, error) {
	p := player{}
	err := s.db.Get(&p, "select id, name, elo from gammon.player where id = $1", id)
	return p.toModel(), err
}

func (s *Store) GetAllPlayers() ([]*model.Player, error) {
	var res []*model.Player
	err := s.db.Select(&res, "select id, name, elo from gammon.player")
	return res, err
}
