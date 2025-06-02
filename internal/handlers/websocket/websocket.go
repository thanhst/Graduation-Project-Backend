package websocket

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:4200"
	},
}

func SetupWebsocket(r *mux.Router) {
	r.HandleFunc("/ws/notifications", NotificationsWsHandler)
	r.HandleFunc("/ws/scheduler", SchedulerWsHandler)
	r.HandleFunc("/ws/room", RoomWsHandler)
}
