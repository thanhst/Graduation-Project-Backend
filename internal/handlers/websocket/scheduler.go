package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/handlers/dto"
	model "server/internal/models"
	"sync"
	"time"
)

type SchedulerHub struct {
	clients map[string]*Client
	mu      sync.Mutex
}

var SchedulerManager = SchedulerHub{
	clients: make(map[string]*Client),
}

func SchedulerWsHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.WSConnectPayloadScheduler
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	userID := payload.UserId
	role := payload.Role
	// selectedDate := payload.SelectedDate
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

	SchedulerManager.mu.Lock()
	SchedulerManager.clients[userID] = client
	SchedulerManager.mu.Unlock()

	log.Printf("User %s connected to scheduler WS\n", userID)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error from %s: %v\n", userID, err)
			break
		}
	}

	SchedulerManager.mu.Lock()
	delete(SchedulerManager.clients, userID)
	SchedulerManager.mu.Unlock()
	_ = conn.Close()
}
func sendPastSchedulers(userID string, selectedDate time.Time) {
	// schedulers :=
	// for _, n := range schedulers {
	// 	SendNotificationToUser(userID, n)
	// }
}
func SendSchedulerToUser(userId string, scheduler *model.Scheduler) {
	SchedulerManager.mu.Lock()
	client, ok := SchedulerManager.clients[userId]
	SchedulerManager.mu.Unlock()

	if ok {
		_ = client.SendJSON(scheduler)
	}
}
