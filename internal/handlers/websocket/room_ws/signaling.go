package room_ws

import (
	"log"
	"net/http"
	"server/internal/app"
	"server/internal/handlers/websocket/upgrader"
)

func RoomWsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	var initMsg struct {
		UserID string `json:"userId"`
		RoomID string `json:"roomId"`
		Role   string `json:"role"`
	}
	if err := conn.ReadJSON(&initMsg); err != nil {
		log.Println("Init read error:", err)
		conn.Close()
		return
	}

	user := createConnection(initMsg.UserID, initMsg.RoomID, initMsg.Role, conn)
	userInfo := getUserById(user.UserID)

	roomsMu.Lock()
	room, exists := rooms[initMsg.RoomID]
	if !exists {
		roomInfo, err := app.ServiceContainer.RoomService.RoomRepo.GetByID(initMsg.RoomID)
		if err != nil || roomInfo.RoomId == "" {
			room = createNewRoom(initMsg.RoomID, initMsg.UserID)
		} else {
			if roomInfo.Host != userInfo.UserId {
				user.Role = "guest"
			} else {
				user.Role = "host"
			}
			room = createNewRoomWithHost(initMsg.RoomID, roomInfo.Host)
		}
		go room.Run()
		rooms[room.ID] = room
	}
	go readPump(user, room)
	go writePump(user)
	roomsMu.Unlock()

	handleUserJoin(user, userInfo, room)
}
