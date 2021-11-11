package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/gislihr/gammon/graph/model"
	internalModel "github.com/gislihr/gammon/pkg/gammon/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

type PlayerRequest struct {
	Offset int
	Limit  int
}

type GameRequest struct {
	Offset   int
	Limit    int
	WinnerId *int
	LoserId  *int
}

type player struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Elo        float64 `db:"elo"`
	ShortName  *string `db:"short_name"`
	Email      *string `db:"email"`
	Experience int     `db:"experience"`
}

var playerFields = []string{
	"id", "name", "elo", "short_name", "email", "experience",
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
	Id           int    `db:"id"`
	WinnerId     int    `db:"winner_id"`
	LoserId      int    `db:"loser_id"`
	Length       int    `db:"length"`
	Round        int    `db:"round"`
	Created      string `db:"created"`
	WinnerScore  *int   `db:"score_winner"`
	LoserScore   *int   `db:"score_loser"`
	TournamentId int    `db:"tournament_id"`
}

var gameFields = []string{
	"id", "winner_id", "loser_id", "length", "round", "created", "score_winner", "score_loser", "tournament_id",
}

type tournament struct {
	Id       int    `db:"id"`
	Date     string `db:"date"`
	Location string `db:"location"`
	Name     string `db:"name"`
	Open     bool   `db:"open"`
}

var tournamentFields = []string{
	"id", "date", "location", "name", "open",
}

func (t tournament) toModel() *model.Tournament {
	return &model.Tournament{
		ID:       t.Id,
		Date:     t.Date,
		Location: t.Location,
		Name:     t.Name,
		Open:     t.Open,
	}
}

func (g game) toModel() *internalModel.Game {
	return &internalModel.Game{
		Id:           g.Id,
		WinnerId:     g.WinnerId,
		LoserId:      g.LoserId,
		Length:       g.Length,
		Round:        g.Round,
		Created:      g.Created,
		WinnerScore:  g.WinnerScore,
		LoserScore:   g.LoserScore,
		TournamentId: g.TournamentId,
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
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, args, err :=
		psql.Select(playerFields...).From("player").Limit(uint64(pr.Limit)).Offset(uint64(pr.Offset)).OrderBy("elo desc").ToSql()
	if err != nil {
		return nil, err
	}

	err = s.db.Select(&res, query, args...)
	return playersToModels(res), err
}

func (s *Store) GetGames(gr GameRequest) ([]*internalModel.Game, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	queryBuilder := psql.Select(gameFields...).From("game").Limit(uint64(gr.Limit)).Offset(uint64(gr.Offset)).OrderBy("created desc")

	if gr.WinnerId != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"winner_id": gr.WinnerId})
	}

	if gr.LoserId != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"loser_id": gr.LoserId})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var res []game
	err = s.db.Select(&res, query, args...)
	return gamesToModels(res), err
}

func (s Store) GetTournament(id int) (*model.Tournament, error) {
	var res tournament
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, args, err := psql.Select(tournamentFields...).From("tournament").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}

	err = s.db.Get(&res, query, args...)
	return res.toModel(), err
}
