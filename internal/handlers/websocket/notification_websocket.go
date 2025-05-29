package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type NotificationHub struct {
    clients map[*websocket.Conn]bool
    mu      sync.Mutex
}

var NotificationManager = NotificationHub{
    clients: make(map[*websocket.Conn]bool),
}

func NotificationsWsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := Upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    NotificationManager.mu.Lock()
    NotificationManager.clients[conn] = true
    NotificationManager.mu.Unlock()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            break
        }
        log.Printf("Notification message: %s\n", msg)
        // Xử lý message, gửi cho client khác hoặc gì đó
    }

    NotificationManager.mu.Lock()
    delete(NotificationManager.clients, conn)
    NotificationManager.mu.Unlock()
}
