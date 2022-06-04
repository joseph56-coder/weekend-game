package handlers

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/joseph56-coder/weekend-game/api"
	"github.com/joseph56-coder/weekend-game/api/handlers/user"
)

func RegisterHandlers(s *socketio.Server, g *api.Game) {
	user.RegisterHandlers(s, g)
}
