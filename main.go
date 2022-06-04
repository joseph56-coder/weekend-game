package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
	"github.com/joseph56-coder/weekend-game/api"
	"github.com/joseph56-coder/weekend-game/api/handlers"
)

func main() {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get the current filename")
	}
	root := filepath.Dir(f)

	err := godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		panic(err)
	}

	public := filepath.Join(root, "public")
	publicFs := http.FileServer(http.Dir(public))
	http.Handle("/public/", http.StripPrefix("/public/", publicFs))

	// pages := filepath.Join(root, "public", "html")
	// pagesFs := http.FileServer(http.Dir(pages))
	// http.Handle("/", pagesFs)

	server := socketio.NewServer(nil)
	server.OnConnect("/", func(c socketio.Conn) error {
		log.Println("New Connection:", c.ID())
		return nil
	})

	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		log.Println("Disconnected", c.ID(), reason)
	})

	game := &api.Game{}
	handlers.RegisterHandlers(server, game)

	go server.Serve()
	defer server.Close()

	http.Handle("/game", server)

	err = http.ListenAndServe(os.Getenv("ADDR"), nil)
	if err != http.ErrServerClosed {
		panic(err)
	}
}
