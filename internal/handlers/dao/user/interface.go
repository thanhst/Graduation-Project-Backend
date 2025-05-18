package userdao

import model "server/internal/models"

type UserDAO interface {
	GetByID(userId string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetClassromsOfUser(userId string, numberInPage int, offset int) (*model.User, error)
	GetNotificationsOfUser(userId string, numberInPage int, offset int) (*model.User, error)
	GetAccountsOfUser(userId string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(userId string) error
}
