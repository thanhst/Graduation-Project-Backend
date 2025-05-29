package websocket

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func SetupWebsocket(r *mux.Router) {
	r.HandleFunc("/ws/notifications", NotificationsWsHandler)
	r.HandleFunc("/ws/scheduler", SchedulerWsHandler)
}
