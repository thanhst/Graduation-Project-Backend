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
func (RoomService *RoomService) GetByHost(userId string, offset int, limit int) ([]model.Room, error) {
	return RoomService.RoomRepo.GetByHost(userId, limit, offset)
}
func (RoomService *RoomService) CountRooms(userId string) (int64, error) {
	return RoomService.RoomRepo.CountRooms(userId)
}

// func (RoomService *RoomService) Update(room *model.Room) error {
// 	return RoomService.RoomRepo.Update(room)
// }
// func (RoomService *RoomService) Delete(roomId string) error {
// 	return RoomService.RoomRepo.Delete(roomId)
// }
