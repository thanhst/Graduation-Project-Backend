package hub_ws

import (
	"log"
	"net/http"
	"server/internal/handlers/websocket/upgrader"
)

func SignalingWSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	var initMsg Message
	if err := conn.ReadJSON(&initMsg); err != nil {
		log.Println("Init read error:", err)
		conn.Close()
		return
	}
	client := ClientConnect(conn, initMsg.UserID)
	log.Println("User ", initMsg.UserID, " connect to hub ws!")
	InstanceHub.Mu.Lock()
	InstanceHub.Clients[initMsg.UserID] = client
	InstanceHub.Mu.Unlock()
	go readPump(client)
	go writePump(client)
	handleSignaling(client)
}
