package routes

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/Emojigamble/utility/logger"
	"github.com/Emojigamble/vault/dao"
	socketio "github.com/googollee/go-socket.io"
)

// Basic game request type.
type GeneralGameRequest struct {
	Type string `json:"type"`
}

type TokenizedRequest struct {
	Token string `json:"token"`
}

// Payload a client should send when searching a public game.
// Token won't be sent in the response.
type PublicGameSearchRequest struct {
	*GeneralGameRequest
	*TokenizedRequest
}

// Payload a client should send when searching a private game.
// JoinCode and Token won't be sent in the response.
type PrivateGameJoinRequest struct {
	*GeneralGameRequest
	JoinCode string `json:"joinCode"`
	*TokenizedRequest
}

// Response after a client requested to join a game.
type GameJoinResponse struct {
	Request GeneralGameRequest `json:"request"`
	Error dao.CommonError `json:"error"`
	Game dao.Game `json:"game"`
}

