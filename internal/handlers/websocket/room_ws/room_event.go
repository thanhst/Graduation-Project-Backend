package room_ws

import (
	"log"
	"server/internal/app"
	model "server/internal/models"
)

func handleUserJoin(user *UserConn, userInfo *model.User, room *Room) {
	room.Mu.Lock()
	defer room.Mu.Unlock()
	if user.Role == "host" {
		handleHostJoin(user, userInfo, room)
		return
	}

	_, hostExists := room.Users[room.HostID]
	if !hostExists {
		room.Waiting[user.UserID] = user
		notifyWaiting(user, userInfo, room)
		log.Printf("User %s join wt room!", user.UserID)
		go handleRoomMessages(user, room)
		return
	}

	if room.AllowAutoJoin {
		user.Accepted = true
		notifyAccepted(user, userInfo, room)
	} else {
		room.Waiting[user.UserID] = user
		notifyWaiting(user, userInfo, room)
	}

	go handleRoomMessages(user, room)
}

func handleHostJoin(user *UserConn, userInfo *model.User, room *Room) {
	room.Users[user.UserID] = user
	room.UsersModel[user.UserID] = userInfo
	log.Println("Host joined:", user.UserID)
	user.SafeSend(RoomMessage{
		Event:  "host_check",
		Sender: nil,
		Payload: map[string]interface{}{
			"roomId":                room.ID,
			"role":                  user.Role,
			"user":                  userInfo,
			"users_in_room":         room.UsersModel,
			"users_in_waiting_room": room.WaitingModel,
			"auto_join":             room.AllowAutoJoin,
			"share_state":           room.AllowAutoShare,
		},
	})
	for uid, u := range room.Waiting {
		if uid != user.UserID {
			u.SafeSend(RoomMessage{
				Event: "host_joined",
				Payload: map[string]interface{}{
					"roomId":  room.ID,
					"role":    user.Role,
					"message": "Host has joined. Please wait for approval.",
					"user":    userInfo,
				},
			})
		}
	}
	handleActionAdd(room, userInfo, user)
	go handleRoomMessages(user, room)
}

func notifyWaiting(user *UserConn, userInfo *model.User, room *Room) {
	user.SafeSend(RoomMessage{
		Event: "join_waiting_room",
		Payload: map[string]interface{}{
			"roomId":  room.ID,
			"userId":  user.UserID,
			"role":    user.Role,
			"message": "Waiting for host to accept you join this room!",
			"user":    userInfo,
		},
	})
	handleActionWTAdd(room, userInfo)
}

func notifyAccepted(user *UserConn, userInfo *model.User, room *Room) {
	room.Users[user.UserID] = user
	room.UsersModel[user.UserID] = userInfo
	user.SafeSend(RoomMessage{
		Event: "accepted",
		Payload: map[string]interface{}{
			"roomId":        room.ID,
			"role":          user.Role,
			"user":          userInfo,
			"users_in_room": room.UsersModel,
			"share_state":   room.AllowAutoShare,
			"host":          getUserById(room.HostID),
		},
	})
	handleActionAdd(room, userInfo, user)
}

func handleRoomMessages(user *UserConn, room *Room) {
	defer func() {
		room.Mu.Lock()
		handleDisconnect(user, room)
		room.Mu.Unlock()
		user.Conn.Close()
	}()
	for msg := range user.Read {
		event := msg["event"].(string)

		switch event {

		case "accept_user":
			if user.Role != "host" {
				continue
			}
			data, ok := msg["data"].(map[string]interface{})
			if !ok {
				log.Println("msg['data'] không phải map[string]interface{}")
				return
			}
			acceptedID := data["user_id"].(string)

			room.Mu.Lock()
			if guest, ok := room.Waiting[acceptedID]; ok {
				guest.Accepted = true
				userInfo := getUserById(guest.UserID)
				guest.SafeSend(RoomMessage{
					Event: "accepted",
					Payload: map[string]interface{}{
						"roomId":        room.ID,
						"role":          guest.Role,
						"user":          userInfo,
						"users_in_room": room.UsersModel,
						"share_state":   room.AllowAutoShare,
						"host":          getUserById(room.HostID),
					},
				})
				handleActionAdd(room, userInfo, guest)
				handleActionWTRemove(room, userInfo, guest)
			}
			room.Mu.Unlock()

		case "toggle_auto_join":
			if user.Role != "host" {
				continue
			}
			data, ok := msg["data"].(map[string]interface{})
			if !ok {
				log.Println("msg['data'] không phải map[string]interface{}")
				return
			}
			value := data["auto_join"].(bool)
			room.Mu.Lock()
			room.AllowAutoJoin = value
			if value {
				waitingUsers := room.Waiting

				for userID, waitingUser := range waitingUsers {
					waitingUser.Accepted = true
					userInfo := getUserById(userID)
					waitingUser.SafeSend(RoomMessage{
						Event: "accepted",
						Payload: map[string]interface{}{
							"roomId":        room.ID,
							"role":          waitingUser.Role,
							"user":          userInfo,
							"users_in_room": room.UsersModel,
							"share_state":   room.AllowAutoShare,
							"host":          getUserById(room.HostID),
						},
					})
					handleActionAdd(room, userInfo, waitingUser)
					handleActionWTRemove(room, userInfo, waitingUser)
				}
			}
			room.Mu.Unlock()

		case "toggle_share":
			if user.Role != "host" {
				continue
			}
			data, ok := msg["data"].(map[string]interface{})
			if !ok {
				log.Println("msg['data'] không phải map[string]interface{}")
				return
			}
			value := data["allow_share"].(bool)
			room.Mu.Lock()
			room.AllowAutoShare = value
			room.MsgChan <- RoomMessage{
				Event:  "change_share_permission",
				Sender: nil,
				Payload: map[string]interface{}{
					"share_state": room.AllowAutoShare,
				},
			}
			room.Mu.Unlock()

		case "start_share":
			if user.Role == "host" {

			} else if room.AllowAutoShare {
				// thì mới gửi lại là cho phép share
			}

		case "start_mic":

		case "start_cam":

		case "stop_mic":

		case "stop_cam":

		case "tranfer_host":
			if user.Role != "host" {
				continue
			}
			data, ok := msg["data"].(map[string]interface{})
			if !ok {
				log.Println("msg['data'] không phải map[string]interface{}")
				return
			}
			newHostId := data["newHostId"].(string)
			room.Mu.Lock()

			if newHost, ok := room.Users[newHostId]; ok {
				oldHostId := room.HostID
				room.HostID = newHostId
				newHostInfo := getUserById(newHostId)
				oldHost := room.Users[oldHostId]
				oldHost.Role = "guest"
				newHost.Role = "host"
				room.MsgChan <- RoomMessage{
					Sender: nil,
					Event:  "host_transferred",
					Payload: map[string]interface{}{
						"newHost": newHostId,
						"oldHost": oldHostId,
						"host":    newHostInfo,
					},
				}
				newHost.SafeSend(RoomMessage{
					Event:  "host_check",
					Sender: nil,
					Payload: map[string]interface{}{
						"roomId":                room.ID,
						"role":                  newHost.Role,
						"user":                  newHostInfo,
						"users_in_room":         room.UsersModel,
						"users_in_waiting_room": room.WaitingModel,
						"auto_join":             room.AllowAutoJoin,
						"share_state":           room.AllowAutoShare,
					},
				})
			}
			room.Mu.Unlock()
		case "remove_user":
			if user.Role != "host" {
				continue
			}
			data, ok := msg["data"].(map[string]interface{})
			if !ok {
				log.Println("msg['data'] không phải map[string]interface{}")
				return
			}
			acceptedID := data["user_id"].(string)

			room.Mu.Lock()
			if guest, ok := room.Waiting[acceptedID]; ok {
				userInfo := getUserById(guest.UserID)
				guest.SafeSend(RoomMessage{
					Event: "removed",
					Payload: map[string]interface{}{
						"role":   guest.Role,
						"roomId": room.ID,
					},
				})
				handleActionWTRemove(room, userInfo, guest)
			}
			room.Mu.Unlock()
		case "return_host":
			if user.Role != "host" {
				continue
			}
			room.Mu.Lock()
			userInfo := getUserById(user.UserID)
			oldHost := room.Users[room.HostID]
			oldHost.Role = "guest"
			user.Role = "host"
			room.HostID = user.UserID
			room.MsgChan <- RoomMessage{
				Sender: nil,
				Event:  "host_transferred",
				Payload: map[string]interface{}{
					"newHost": user.UserID,
					"oldHost": oldHost.UserID,
					"host":    userInfo,
				},
			}
			user.SafeSend(RoomMessage{
				Event:  "host_check",
				Sender: nil,
				Payload: map[string]interface{}{
					"roomId":                room.ID,
					"role":                  user.Role,
					"user":                  userInfo,
					"users_in_room":         room.UsersModel,
					"users_in_waiting_room": room.WaitingModel,
					"auto_join":             room.AllowAutoJoin,
					"share_state":           room.AllowAutoShare,
				},
			})
			room.Mu.Unlock()
		case "close_room":
			wg.Add(1)
			go func() {
				defer wg.Done()
				room.Mu.Lock()
				room.MsgChan <- RoomMessage{
					Event:  "room_closed",
					Sender: room.Users[room.HostID],
					Payload: map[string]interface{}{
						"roomId": room.ID,
					},
				}
				room.MsgChan <- RoomMessage{
					Event:  "room_closed",
					Sender: room.Users[room.HostID],
					Payload: map[string]interface{}{
						"roomId":  room.ID,
						"forward": "waiting",
					},
				}
				room.Mu.Unlock()
			}()
			wg.Wait()
			roomsMu.Lock()
			delete(rooms, room.ID)
			roomsMu.Unlock()
		}
	}
}
func handleDisconnect(user *UserConn, room *Room) {
	if user.Accepted {
		handleActionRemove(room, getUserById(user.UserID), user)
	} else {
		handleActionWTRemove(room, getUserById(user.UserID), user)
	}
	log.Printf("User %s disconnected from room %s", user.UserID, room.ID)
	if len(room.Users) >= 1 && user.Role == "host" {
		newHostId := ""
		for uid := range room.Users {
			newHostId = uid
		}
		if newHostId != "" {
			room.HostID = newHostId
			newHost, ok := room.Users[newHostId]
			if ok {
				newHost.Role = "host"
				newHost.Accepted = true
			}
			newHostInfo := getUserById(newHostId)
			room.MsgChan <- RoomMessage{
				Sender: nil,
				Event:  "host_transferred",
				Payload: map[string]interface{}{
					"newHost": newHostId,
					"oldHost": user.UserID,
					"host":    newHostInfo,
				},
			}
			newHost.SafeSend(RoomMessage{
				Event:  "host_check",
				Sender: nil,
				Payload: map[string]interface{}{
					"roomId":                room.ID,
					"role":                  newHost.Role,
					"user":                  newHostInfo,
					"users_in_room":         room.UsersModel,
					"users_in_waiting_room": room.WaitingModel,
					"auto_join":             room.AllowAutoJoin,
					"share_state":           room.AllowAutoShare,
				},
			})

			log.Printf("Host transferred from %s to %s", user.UserID, newHostId)
		} else {
			log.Printf("Room %s is empty after host disconnected", room.ID)
			roomsMu.Lock()
			delete(rooms, room.ID)
			roomsMu.Unlock()
		}
	}
	if len(room.Users) == 0 && user.Role == "host" {
		log.Printf("Room %s is empty after host disconnected", room.ID)
		room.MsgChan <- RoomMessage{
			Sender: nil,
			Event:  "host_leave",
			Payload: map[string]interface{}{
				"roomId":  room.ID,
				"role":    user.Role,
				"message": "Host leave this room! Waiting for host open this room!",
			},
		}
	}
	if len(room.Users) == 0 && len(room.Waiting) == 0 {
		log.Println("Remove room: ", room.ID)
		close(room.QuitChan)
		roomsMu.Lock()
		delete(rooms, room.ID)
		roomsMu.Unlock()
	}
}
func handleActionAdd(room *Room, userInfo *model.User, userConn *UserConn) {
	room.Users[userConn.UserID] = userConn
	room.UsersModel[userInfo.UserId] = userInfo
	room.MsgChan <- RoomMessage{
		Sender: nil,
		Event:  "update_user_in_room",
		Payload: map[string]interface{}{
			"action": "add",
			"user":   userInfo,
		},
	}
}
func handleActionRemove(room *Room, userInfo *model.User, userConn *UserConn) {
	delete(room.Users, userConn.UserID)
	delete(room.UsersModel, userInfo.UserId)
	room.MsgChan <- RoomMessage{
		Sender: nil,
		Event:  "update_user_in_room",
		Payload: map[string]interface{}{
			"action": "remove",
			"user":   userInfo,
		},
	}
}
func handleActionWTAdd(room *Room, userInfo *model.User) {
	room.WaitingModel[userInfo.UserId] = userInfo
	hostConn := room.Users[room.HostID]
	if hostConn != nil {
		hostConn.SafeSend(RoomMessage{
			Event: "update_user_waiting_room",
			Payload: map[string]interface{}{
				"action":  "add",
				"user":    userInfo,
				"role":    hostConn.Role,
				"roomId":  room.ID,
				"forward": "waiting",
			},
		})
	}
}
func handleActionWTRemove(room *Room, userInfo *model.User, userConn *UserConn) {
	delete(room.Waiting, userConn.UserID)
	delete(room.WaitingModel, userInfo.UserId)
	hostConn := room.Users[room.HostID]
	if hostConn != nil {
		hostConn.SafeSend(RoomMessage{
			Event: "update_user_waiting_room",
			Payload: map[string]interface{}{
				"action":  "remove",
				"user":    userInfo,
				"role":    hostConn.Role,
				"roomId":  room.ID,
				"forward": "waiting",
			},
		})
	}
}

func getUserById(userId string) *model.User {
	userInfo, err := app.ServiceContainer.UserService.GetUserByID(userId)
	if err != nil || userInfo.UserId == "" {
		log.Println("User not found")
		return nil
	}
	return userInfo
}
