package upgrader

import (
	"net/http"
	"server/internal/utils/dotenv"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == dotenv.GetDotEnv("FE_URL")+":"+dotenv.GetDotEnv("FE_PORT") || origin == dotenv.GetDotEnv("FE_URL")
		// return true
	},
}
