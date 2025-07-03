package hub_ws

import (
	"log"
	"sync"
)

type Hub struct {
	Clients map[string]*Client
	Mu      sync.Mutex
}

var InstanceHub *Hub = NewHub()

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[string]*Client),
	}
}

func (h *Hub) broadcast(msg *Message) {
	for _, client := range InstanceHub.Clients {
		if client.UserID == msg.UserID {
			continue
		}
		func(c *Client) {
			defer func() {
				if recover := recover(); recover != nil {
					log.Println("Recover error!")
				}
			}()
			client.SafeSend(*msg)
		}(client)
	}
}
