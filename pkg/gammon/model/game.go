package model

type Game struct {
	Id       int `json:"id"`
	LoserId  int `json:"loser"`
	WinnerId int `json:"winner"`
	Length   int `json:"length"`
}
