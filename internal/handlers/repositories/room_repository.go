package repository

import (
	roomdao "server/internal/handlers/dao/room"
	model "server/internal/models"
)

type RoomRepository interface {
	GetByID(roomID string) (*model.Room, error)
	GetByHost(userID string, limit int, offset int) ([]model.Room, error)
	GetActiveRooms() ([]model.Room, error)
	Create(room *model.Room) error
	Update(room *model.Room) error
	CloseRoom(roomID string) error
	Delete(roomID string) error
}
type roomRepository struct {
	dao roomdao.RoomDAO
}

func NewRoomRepository(dao roomdao.RoomDAO) RoomRepository {
	return &roomRepository{dao: dao}
}

func (r *roomRepository) GetByID(roomID string) (*model.Room, error) {
	return r.dao.GetByID(roomID)
}

func (r *roomRepository) GetByHost(userID string, limit int, offset int) ([]model.Room, error) {
	return r.dao.GetByHost(userID, limit, offset)
}

func (r *roomRepository) GetActiveRooms() ([]model.Room, error) {
	return r.dao.GetActiveRooms()
}

func (r *roomRepository) Create(room *model.Room) error {
	return r.dao.Create(room)
}

func (r *roomRepository) Update(room *model.Room) error {
	return r.dao.Update(room)
}

func (r *roomRepository) CloseRoom(roomID string) error {
	return r.dao.CloseRoom(roomID)
}

func (r *roomRepository) Delete(roomID string) error {
	return r.dao.Delete(roomID)
}
