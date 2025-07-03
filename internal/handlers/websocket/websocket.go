package websocket

import (
	"server/internal/handlers/websocket/hub_ws"
	"server/internal/handlers/websocket/room_ws"

	"github.com/gorilla/mux"
)

func SetupWebsocket(r *mux.Router) {
	r.HandleFunc("/ws/room", room_ws.RoomWsHandler)
	r.HandleFunc("/ws/hub", hub_ws.SignalingWSHandler)
}
