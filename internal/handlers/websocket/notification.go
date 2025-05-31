package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/handlers/dto"
	model "server/internal/models"
	"sync"
)

type NotificationHub struct {
	clients map[string]*Client
	mu      sync.Mutex
}

var NotificationManager = NotificationHub{
	clients: make(map[string]*Client),
}

func NotificationsWsHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.WSConnectPayloadNotification
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	userID := payload.UserId
	role := payload.Role
	// classID := payload.ClassId
	if userID == "" {
		http.Error(w, "Missing userID", http.StatusBadRequest)
		return
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	client := &Client{
		UserID: userID,
		Role:   role,
		Conn:   conn,
	}

	NotificationManager.mu.Lock()
	NotificationManager.clients[userID] = client
	NotificationManager.mu.Unlock()

	log.Printf("User %s connected to notification WS\n", userID)

	go sendPastNotifications(userID)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error from %s: %v\n", userID, err)
			break
		}
	}

	NotificationManager.mu.Lock()
	delete(NotificationManager.clients, userID)
	NotificationManager.mu.Unlock()
	_ = conn.Close()
}

func sendPastNotifications(userID string) {
	// notifs := service.GetUserNotifications(userID) // gọi service lấy từ DB
	// for _, n := range notifs {
	// 	SendNotificationToUser(userID, n)
	// }
}

func SendNotificationToUser(userID string, noti *model.Notification) {
	NotificationManager.mu.Lock()
	client, ok := NotificationManager.clients[userID]
	NotificationManager.mu.Unlock()

	if ok {
		_ = client.SendJSON(noti)
	}
}
