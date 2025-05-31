package service

import (
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
)

type RoomService struct {
	RoomRepo repository.RoomRepository
}

func NewRoomService(RoomRepo repository.RoomRepository) *RoomService {
	return &RoomService{RoomRepo: RoomRepo}
}
func (RoomService *RoomService) Create(room *model.Room) error {
	return RoomService.RoomRepo.Create(room)
}

// func (RoomService *RoomService) Update(room *model.Room) error {
// 	return RoomService.RoomRepo.Update(room)
// }
// func (RoomService *RoomService) Delete(roomId string) error {
// 	return RoomService.RoomRepo.Delete(roomId)
// }
