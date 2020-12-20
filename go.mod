module github.com/Emojigamble/vault

go 1.15

require (
	cloud.google.com/go/firestore v1.4.0 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/Emojigamble/utility/logger v0.0.0-00010101000000-000000000000
	github.com/Emojigamble/utility/middleware v0.0.0-00010101000000-000000000000
	github.com/Emojigamble/utility/setup v0.0.0-00010101000000-000000000000
	github.com/googollee/go-socket.io v1.4.4
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102
	google.golang.org/api v0.36.0
)

replace github.com/Emojigamble/utility/logger => ../utility/logger

replace github.com/Emojigamble/utility/middleware => ../utility/middleware

replace github.com/Emojigamble/utility/setup => ../utility/setup
