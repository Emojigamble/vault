package main

import (
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	socketio "github.com/googollee/go-socket.io"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
		allowedOrigins := []string{"localhost:3000"}

		if contains(allowedOrigins, r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		r.Header.Del("Origin")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}

func contains(s []string, search string) bool {
	contains := false
	for entry := range s {
		if strings.Contains(search, s[entry]) {
			contains = true
		}
	}
	return contains
}

type TicTacToeMove struct {
	Col     int    `json:"col"`
	Row     int    `json:"row"`
	Emoji   string `json:"emoji"`
	IdToken string `json:"idToken"`
}

func main() {
	server, _ := socketio.NewServer(nil)

	client, err := setupFirebase("emojigamble-key.json")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Sprint("vault", os.Getenv("HOSTNAME"))

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/", "tictactoe-move", func(s socketio.Conn, data string) {
		var move TicTacToeMove
		json.Unmarshal([]byte(data), &move)

		fmt.Println("submitted move in row ", move.Row, " and col ", move.Col)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", corsMiddleware(server))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func setupFirebase(keyFile string) (*auth.Client, error) {
	opt := option.WithCredentialsFile(keyFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}