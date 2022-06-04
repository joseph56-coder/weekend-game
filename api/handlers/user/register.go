package user

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/joseph56-coder/weekend-game/api"
)

func RegisterHandlers(s *socketio.Server, g *api.Game) {
	s.OnEvent("/user", "join", joinEvent(g))
}
