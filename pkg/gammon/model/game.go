package model

type Game struct {
	ID       string `json:"id"`
	LoserId  string `json:"loser"`
	WinnerId string `json:"winner"`
	Length   int    `json:"length"`
}
