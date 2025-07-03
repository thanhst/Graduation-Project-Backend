package hub_ws

import (
	"log"
	"server/internal/app"
)

func handleSignaling(client *Client) {
	defer func() {
		InstanceHub.Mu.Lock()
		handleDisconnect(client)
		InstanceHub.Mu.Unlock()
		client.Conn.Close()
	}()
	for msg := range client.Read {
		event := msg.Event
		switch event {
		case "new_notification":
			classId := msg.Payload["class_id"].(string)
			students_class, err := app.ServiceContainer.StudentClassService.ListByClass(classId)
			if err != nil {
				log.Print(err)
			}
			for _, std := range students_class {
				if c, exist := InstanceHub.Clients[std.UserId]; exist {
					c.SafeSend(msg)
				}
			}
		case "ping":
			client.SafeSend(Message{
				Event: "pong",
			})
		}
	}
}
func handleDisconnect(client *Client) {
	client.Conn.Close()
	delete(InstanceHub.Clients, client.UserID)
	log.Println("User ", client.UserID, " disconnect from ws-hub")
}
