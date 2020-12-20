package routes

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/Emojigamble/utility/logger"
	socketio "github.com/googollee/go-socket.io"
)

// Registers a socket.io connection event listener which expects a firebase id token from the user.
// When no token is found the connection will be closed otherwise an 'authenticated' event is emitted
func RegisterConnectionListener(server *socketio.Server, client *auth.Client, context context.Context, log logger.EmojigambleLogger) {
	server.OnConnect("/", func(s socketio.Conn) error {
		url := s.URL()
		idToken := url.Query().Get("idToken")

		_, err := client.VerifyIDToken(context, idToken)
		if err != nil {
			log.Log(fmt.Sprintf("Error verifying ID token while connecting: %v\n", err), logger.ConnectionLogLevel)
			_ = s.Close()
			return nil
		}
		s.Emit("authenticated")

		s.SetContext("")
		log.Log(fmt.Sprint("Connected new client: ", s.ID()), logger.ConnectionLogLevel)
		return nil
	})
}