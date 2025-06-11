package room_ws

import (
	"log"
	"net/http"
	"server/internal/app"
	"server/internal/handlers/websocket/upgrader"
	model "server/internal/models"
	"sync"
)

type RoomMessage struct {
	Sender  *UserConn
	Event   string
	Payload map[string]interface{}
}

type Room struct {
	ID             string
	HostID         string
	Users          map[string]*UserConn
	Waiting        map[string]*UserConn
	UsersModel     map[string]*model.User
	WaitingModel   map[string]*model.User
	AllowAutoJoin  bool
	AllowAutoShare bool
	Mu             sync.Mutex
	MsgChan        chan RoomMessage
	QuitChan       chan struct{}
}

var rooms = make(map[string]*Room)
var roomsMu sync.Mutex
var wg sync.WaitGroup

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

func (r *Room) Run() {
	for {
		select {
		case msg := <-r.MsgChan:
			if msg.Payload["forward"] == "waiting" {
				r.broadcastWaitingRoom(&msg)
			} else {
				r.broadcast(&msg)
			}
		case <-r.QuitChan:
			return
		}
	}
}
func (r *Room) broadcast(msg *RoomMessage) {
	for _, u := range r.Users {
		if u == msg.Sender {
			continue
		}
		func(u *UserConn) {
			defer func() {
				if recover := recover(); recover != nil {
					log.Println("Recover error!")
				}
			}()
			u.SafeSend(RoomMessage{
				Event:   msg.Event,
				Payload: msg.Payload,
			})
		}(u)
	}
}
func (r *Room) broadcastWaitingRoom(msg *RoomMessage) {
	for _, u := range r.Waiting {
		u.SafeSend(RoomMessage{
			Event:   msg.Event,
			Payload: msg.Payload,
		})
	}
}

func createNewRoom(roomID, hostID string) *Room {
	return &Room{
		ID:             roomID,
		HostID:         hostID,
		Users:          make(map[string]*UserConn),
		Waiting:        make(map[string]*UserConn),
		UsersModel:     make(map[string]*model.User),
		WaitingModel:   make(map[string]*model.User),
		AllowAutoJoin:  false,
		AllowAutoShare: false,
		MsgChan:        make(chan RoomMessage),
		QuitChan:       make(chan struct{}),
	}
}

func createNewRoomWithHost(roomID, hostID string) *Room {
	return createNewRoom(roomID, hostID)
}
