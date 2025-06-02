package websocket

import (
	"log"
	"net/http"
	"server/internal/app"
	"sync"

	"github.com/gorilla/websocket"
)

type RoomMessage struct {
	Sender  *UserConn
	Payload map[string]interface{}
}

type UserConn struct {
	UserID   string
	RoomID   string
	Role     string
	Conn     *websocket.Conn
	Accepted bool
}

type Room struct {
	ID             string
	HostID         string
	Users          map[string]*UserConn
	Waiting        map[string]*UserConn
	AllowAutoJoin  bool
	AllowAutoShare bool
	Mu             sync.Mutex
	MsgChan        chan RoomMessage
	QuitChan       chan struct{}
}

var rooms = make(map[string]*Room)
var roomsMu sync.Mutex

func RoomWsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
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

	user := &UserConn{
		UserID:   initMsg.UserID,
		RoomID:   initMsg.RoomID,
		Role:     initMsg.Role,
		Conn:     conn,
		Accepted: initMsg.Role == "host",
	}

	userInfo, err := app.ServiceContainer.UserService.GetUserByID(user.UserID)
	if err != nil || userInfo.UserId == "" {
		log.Println("User not found")
		conn.Close()
		return
	}

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
			go room.Run()
		}
		rooms[room.ID] = room
	}
	roomsMu.Unlock()

	handleUserJoin(user, userInfo, room)
}

func createNewRoom(roomID, hostID string) *Room {
	return &Room{
		ID:             roomID,
		HostID:         hostID,
		Users:          make(map[string]*UserConn),
		Waiting:        make(map[string]*UserConn),
		AllowAutoJoin:  false,
		AllowAutoShare: false,
		MsgChan:        make(chan RoomMessage),
		QuitChan:       make(chan struct{}),
	}
}

func createNewRoomWithHost(roomID, hostID string) *Room {
	return createNewRoom(roomID, hostID)
}

func handleUserJoin(user *UserConn, userInfo any, room *Room) {
	room.Mu.Lock()
	defer room.Mu.Unlock()

	if user.Role == "host" {
		handleHostJoin(user, userInfo, room)
		return
	}

	_, hostExists := room.Users[room.HostID]
	if !hostExists {
		room.Waiting[user.UserID] = user
		notifyWaiting(user, userInfo, room)
		go handleRoomMessages(user, room)
		return
	}

	if room.AllowAutoJoin {
		user.Accepted = true
		room.Users[user.UserID] = user
		notifyAccepted(user, userInfo, room)
	} else {
		room.Waiting[user.UserID] = user
		notifyWaiting(user, userInfo, room)
		if host, ok := room.Users[room.HostID]; ok {
			host.Conn.WriteJSON(map[string]interface{}{
				"event":  "user_waiting",
				"roomId": room.ID,
				"userId": user.UserID,
				"role":   user.Role,
				"user":   userInfo,
			})
		}
	}
	go handleRoomMessages(user, room)
}

func handleHostJoin(user *UserConn, userInfo any, room *Room) {
	room.Users[user.UserID] = user
	log.Println("Host joined:", user.UserID)
	user.Conn.WriteJSON(map[string]interface{}{
		"event":  "host_check",
		"roomId": room.ID,
		"role":   user.Role,
		"user":   userInfo,
	})
	for uid, u := range room.Waiting {
		if uid != user.UserID {
			u.Conn.WriteJSON(map[string]interface{}{
				"event":   "host_joined",
				"roomId":  room.ID,
				"message": "Host has joined. Please wait for approval.",
				"user":    userInfo,
			})
		}
	}
	go handleRoomMessages(user, room)
}

func notifyWaiting(user *UserConn, userInfo any, room *Room) {
	user.Conn.WriteJSON(map[string]interface{}{
		"event":   "join_waiting_room",
		"roomId":  room.ID,
		"userId":  user.UserID,
		"role":    user.Role,
		"message": "Waiting for host to open this room!",
		"user":    userInfo,
	})
}

func notifyAccepted(user *UserConn, userInfo any, room *Room) {
	room.Users[user.UserID] = user
	user.Conn.WriteJSON(map[string]interface{}{
		"event":  "accepted",
		"roomId": room.ID,
		"role":   user.Role,
		"user":   userInfo,
	})
}

func handleRoomMessages(user *UserConn, room *Room) {
	defer func() {
		room.Mu.Lock()
		handleDisconnect(user, room)
		room.Mu.Unlock()
		user.Conn.Close()
	}()
	for {
		var msg map[string]interface{}
		err := user.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		event := msg["event"].(string)

		switch event {

		case "accept_user":
			if user.Role != "host" {
				continue
			}
			acceptedID := msg["userId"].(string)

			room.Mu.Lock()
			if guest, ok := room.Waiting[acceptedID]; ok {
				guest.Accepted = true
				room.Users[acceptedID] = guest
				delete(room.Waiting, acceptedID)

				guest.Conn.WriteJSON(map[string]interface{}{
					"event": "accepted",
				})
			}
			room.Mu.Unlock()

		case "toggle_auto_join":
			if user.Role != "host" {
				continue
			}
			value := msg["value"].(bool)
			room.Mu.Lock()
			room.AllowAutoJoin = value
			room.Mu.Unlock()

		case "set_permission":

		case "offer", "answer", "ice_candidate":
			targetID := msg["targetId"].(string)

			room.Mu.Lock()
			if target, ok := room.Users[targetID]; ok && target.Accepted {
				target.Conn.WriteJSON(msg)
			}
			room.Mu.Unlock()

		case "start_share":

		case "transfer_host":
			if user.Role != "host" {
				continue
			}
			newHostId := msg["newHostId"].(string)

			room.Mu.Lock()
			if newHost, ok := room.Users[newHostId]; ok {
				// Cập nhật HostID
				oldHostId := room.HostID
				room.HostID = newHostId

				// Cập nhật Role
				oldHost := room.Users[oldHostId]
				oldHost.Role = "guest"
				newHost.Role = "host"

				// Thông báo cho các user trong phòng
				for _, u := range room.Users {
					u.Conn.WriteJSON(map[string]interface{}{
						"event":   "host_transferred",
						"newHost": newHostId,
						"oldHost": oldHostId,
					})
				}
			}
			room.Mu.Unlock()
		}
	}
}
func handleDisconnect(user *UserConn, room *Room) {
	if user.Accepted {
		delete(room.Users, user.UserID)
	} else {
		delete(room.Waiting, user.UserID)
	}
	log.Printf("User %s disconnected from room %s", user.UserID, room.ID)
	if len(room.Users) >= 1 && user.Role == "host" {
		newHostId := ""
		for uid := range room.Users {
			newHostId = uid
		}
		if newHostId != "" {
			room.HostID = newHostId
			if newHost, ok := room.Users[newHostId]; ok {
				newHost.Role = "host"
				newHost.Accepted = true
			}

			for _, u := range room.Users {
				u.Conn.WriteJSON(map[string]interface{}{
					"event":   "host_transferred",
					"newHost": newHostId,
					"oldHost": user.UserID,
				})
			}

			log.Printf("Host transferred from %s to %s", user.UserID, newHostId)
		} else {
			log.Printf("Room %s is empty after host disconnected", room.ID)
			delete(rooms, room.ID)
		}
	}
	if len(room.Users) == 0 && user.Role == "host" {
		log.Printf("Room %s is empty after host disconnected", room.ID)
		for _, u := range room.Waiting {
			u.Conn.WriteJSON(map[string]interface{}{
				"event":   "join_waiting_room",
				"roomId":  room.ID,
				"role":    user.Role,
				"message": "Waiting for host open this room!",
			})
		}
	}
	if len(room.Users) == 0 && len(room.Waiting) == 0 {
		log.Println("Remove room: ", room.ID)
		delete(rooms, room.ID)
	}
}

func (r *Room) Run() {
	for {
		select {
		case msg := <-r.MsgChan:
			r.broadcast(msg.Event, msg.Payload)
		case <-r.QuitChan:
			log.Println("Room closed:", r.ID)
			return
		}
	}
}
func (r *Room) broadcast(event string, data map[string]interface{}) {
	for _, u := range r.Users {
		u.Conn.WriteJSON(map[string]interface{}{
			"event": event,
			"data":  data,
		})
	}
}
