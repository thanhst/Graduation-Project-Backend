package roomdao

import model "server/internal/models"

type RoomDAO interface {
	GetByID(roomID string) (*model.Room, error)
	GetByHost(userID string, limit int, offset int) ([]model.Room, error)
	GetActiveRooms() ([]model.Room, error)
	Create(room *model.Room) error
	Update(room *model.Room) error
	CloseRoom(roomID string) error
	Delete(roomID string) error
}
