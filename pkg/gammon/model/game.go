package model

type Game struct {
	Id           int
	LoserId      int
	WinnerId     int
	Length       int
	Round        int
	Created      string
	WinnerScore  *int
	LoserScore   *int
	TournamentId int
}
