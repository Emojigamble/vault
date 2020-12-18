package routes

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/adwirawien/Emojigamble/utility/logger"
	socketio "github.com/googollee/go-socket.io"
)

func registerConnectListener(server socketio.Server, client *auth.Client, context context.Context, logger logger.EmojigambleLogger) {
	server.OnConnect("/", func(s socketio.Conn) error {
		url := s.URL()
		idToken := url.Query().Get("idToken")

		token, err := client.VerifyIDToken(context, idToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			s.Close()
			return nil
		}
		s.Emit("authenticated")

		log.Printf("Verified ID token: %v\n", token.UID)

		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
}