package dto

import "time"

type WSConnectPayload struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

type WSConnectPayloadScheduler struct {
	UserId       string    `json:"user_id"`
	Role         string    `json:"role"`
	SelectedDate time.Time `json:"date"`
}
type WSConnectPayloadNotification struct {
	UserId  string `json:"user_id"`
	Role    string `json:"role"`
	ClassId string `json:"class_id"`
}
