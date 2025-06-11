package room_ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type UserConn struct {
	UserID    string
	RoomID    string
	Role      string
	Conn      *websocket.Conn
	Accepted  bool
	Send      chan RoomMessage
	Read      chan map[string]interface{}
	CloseOnce sync.Once
}

func readPump(user *UserConn, room *Room) {
	defer func() {
		if r := recover(); r != nil {
		}
		room.Mu.Lock()
		delete(room.Users, user.UserID)
		room.Mu.Unlock()
		user.CloseOnce.Do(func() {
			close(user.Send)
			close(user.Read)
			user.Conn.Close()
		})
	}()
	log.Println("Ready to read msg of user: ", user.UserID)
	for {
		var msg map[string]interface{}
		if err := user.Conn.ReadJSON(&msg); err != nil {
			return
		}
		func(user *UserConn) {
			defer func() {
				if recover := recover(); recover != nil {
					log.Println("Recover panic!")
				}
			}()
			user.Read <- msg
		}(user)
	}
}

func writePump(user *UserConn) {
	log.Println("Ready to write msg of user: ", user.UserID)
	for msg := range user.Send {
		err := user.Conn.WriteJSON(map[string]interface{}{
			"event": msg.Event,
			"data":  msg.Payload,
		})
		if err != nil {
			user.CloseOnce.Do(func() {
				close(user.Send)
				close(user.Read)
				user.Conn.Close()
			})
			return
		}
	}
}
func createConnection(userId string, roomId string, role string, connection *websocket.Conn) *UserConn {
	return &UserConn{
		UserID:   userId,
		RoomID:   roomId,
		Role:     role,
		Conn:     connection,
		Accepted: role == "host",
		Send:     make(chan RoomMessage, 256),
		Read:     make(chan map[string]interface{}, 256),
	}
}
func (user *UserConn) SafeSend(msg RoomMessage) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from send panic:", r)
		}
	}()
	user.Send <- msg
}
