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

func RegisterGameJoinListeners(server *socketio.Server, client *auth.Client, context context.Context, log logger.EmojigambleLogger) {
	server.OnEvent("/", "searchPublicGame", func(s socketio.Conn, data string) {
		response := GameJoinResponse{
			Request: nil,
			Error:   nil,
			Game:    nil,
		}
		defer marshalAndSendInterface(s, "gameJoinResponse", response)

		publicGameRequest := PublicGameSearchRequest{}
		err := json.Unmarshal([]byte(data), &publicGameRequest)
		if err != nil {
			log.Log(fmt.Sprint(err, "while parsing 'searchPublicGame' Request payload!"), logger.MatchMakingLogLevel)

			response.Error = dao.PayloadParsingError
			return
		}

		response.Request = *publicGameRequest.GeneralGameRequest

		// TODO: verify user and get userdata
		// TODO: check if user has game access
		// TODO: create new game or assign existing
	})
}

func marshalAndSendInterface(s socketio.Conn, event string, data interface{}) {
	response, _ := json.Marshal(data)
	s.Emit(event, response)
}
