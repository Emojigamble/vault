package dao

import "time"

type Opponent struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarURL"`
}

type Game struct {
	Type      string    `json:"type"`
	Id        string    `json:"id"`
	StartedAt time.Time `json:"startedAt"`
	Opponent  Opponent  `json:"opponent"`
	YourTurn  bool      `json:"yourTurn"`
	State     string    `json:"state"`
	Data      string    `json:"data"`
}

type GameAction struct {
	Type   string `json:"type"`
	GameId string `json:"gameId"`
	Data   string `json:"data"`
}
