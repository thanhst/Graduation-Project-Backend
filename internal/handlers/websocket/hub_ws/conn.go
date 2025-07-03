package hub_ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Event   string
	UserID  string
	Payload map[string]interface{}
}

type Client struct {
	Conn      *websocket.Conn
	UserID    string
	Send      chan Message
	Read      chan Message
	CloseOnce sync.Once
}

func ClientConnect(connection *websocket.Conn, userId string) *Client {
	return &Client{
		Conn:   connection,
		UserID: userId,
		Send:   make(chan Message),
		Read:   make(chan Message, 1000),
	}
}


func writePump(user *Client) {
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

func readPump(user *Client) {
	defer func() {
		if r := recover(); r != nil {
		}
		InstanceHub.Mu.Lock()
		delete(InstanceHub.Clients, user.UserID)
		InstanceHub.Mu.Unlock()
		user.CloseOnce.Do(func() {
			close(user.Send)
			close(user.Read)
			user.Conn.Close()
		})
	}()
	log.Println("Ready to read msg of user: ", user.UserID)
	for {
		var msg Message
		if err := user.Conn.ReadJSON(&msg); err != nil {
			return
		}
		func(user *Client) {
			defer func() {
				if recover := recover(); recover != nil {
					log.Println("Recover panic!")
				}
			}()
			user.Read <- msg
		}(user)
	}
}

func (user *Client) SafeSend(msg Message) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from send panic:", r)
		}
	}()
	user.Send <- msg
}


