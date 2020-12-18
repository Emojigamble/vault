package routes

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/Emojigamble/utility/logger"
	socketio "github.com/googollee/go-socket.io"
)

func RegisterConnectListener(server *socketio.Server, client *auth.Client, context context.Context, log logger.EmojigambleLogger) {
	server.OnConnect("/", func(s socketio.Conn) error {
		url := s.URL()
		idToken := url.Query().Get("idToken")

		token, err := client.VerifyIDToken(context, idToken)
		if err != nil {
			log.Log(fmt.Sprintf("error verifying ID token: %v\n", err), logger.ErrorLogLevel)
			_ = s.Close()
			return nil
		}
		s.Emit("authenticated")

		log.Log(fmt.Sprintf("Verified ID token: %v\n", token.UID), logger.ConnectionLogLevel)

		s.SetContext("")
		log.Log(fmt.Sprint("connected: ", s.ID()), logger.ConnectionLogLevel)
		return nil
	})
}