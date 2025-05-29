package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type SchedulerHub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

var SchedulerManager = SchedulerHub{
	clients: make(map[*websocket.Conn]bool),
}

func SchedulerWsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	SchedulerManager.mu.Lock()
	SchedulerManager.clients[conn] = true
	SchedulerManager.mu.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Scheduler message: %s\n", msg)
	}

	SchedulerManager.mu.Lock()
	delete(SchedulerManager.clients, conn)
	SchedulerManager.mu.Unlock()
}
