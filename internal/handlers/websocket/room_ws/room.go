package room_ws

import (
	"log"
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
