package main

import (
	"fmt"
	"github.com/Emojigamble/utility/logger"
	"github.com/Emojigamble/utility/middleware"
	"github.com/Emojigamble/utility/setup"
	"github.com/Emojigamble/vault/routes"
	socketio "github.com/googollee/go-socket.io"
	"golang.org/x/net/context"
	"net/http"
	"os"
)

func main() {
	// define logger
	log := logger.EmojigambleLogger{
		ActiveLogLevels: logger.AllLogLevels(),
		LogToDatabase:   false,
	}
	// get environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	server, _ := socketio.NewServer(nil)

	// setup firebase
	client, err := setup.FirebaseAuthClient(context.Background(), "emojigamble-key.json")
	if err != nil {
		log.Log(fmt.Sprint(err), logger.ErrorLogLevel)
	}

	// register listeners
	routes.RegisterConnectionListener(server, client, context.Background(), log)
	// TODO: game listeners
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Log(fmt.Sprint("Encountered socket.io error event: ", e), logger.WarnLogLevel)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	corsMiddleware := middleware.CorsMiddleware{
		AllowedOrigins: []string{"localhost:3000"},
		AllowedHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
	}

	go server.Serve()
	defer server.Close()
	http.Handle("/socket.io/", corsMiddleware.Cors(server))

	log.Log(fmt.Sprintf("Serving at %v", port), logger.BaseLogLevel)
	log.Log(fmt.Sprint(http.ListenAndServe(port, nil)), logger.WarnLogLevel)
}


