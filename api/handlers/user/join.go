package user

import (
	"encoding/json"
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/joseph56-coder/weekend-game/api"
)

type userJoinReq struct {
	Username string `json:"username"`
}

func joinEvent(game *api.Game) interface{} {
	return func(c socketio.Conn, msg string) string {
		log.Println("Users:", game.Users)
		req := new(userJoinReq)
		err := json.Unmarshal([]byte(msg), req)
		if err != nil {
			log.Println(err)
			return "Could not join"
		}
		_, found := game.Users.FindByUsername(req.Username)
		if found {
			return "Username in use"
		}
		game.Users.AddUser(api.NewUser(req.Username))
		return "Done!"
	}
}
