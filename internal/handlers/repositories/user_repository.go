package repository

import (
	userdao "server/internal/handlers/dao/user"
	model "server/internal/models"
)

type UserRepository interface {
	GetByID(userId string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetClassroomsOfUser(userId string, numberInPage int, offset int) (*model.User, error)
	GetNotificationsOfUser(userId string, numberInPage int, offset int) (*model.User, error)
	GetAccountsOfUser(userId string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(userId string) error
}

type userRepository struct {
	userDAO userdao.UserDAO
}

func NewUserRepository(userDAO userdao.UserDAO) UserRepository {
	return &userRepository{userDAO: userDAO}
}

func (r *userRepository) GetByID(userId string) (*model.User, error) {
	return r.userDAO.GetByID(userId)
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	return r.userDAO.GetByEmail(email)
}

func (r *userRepository) GetClassroomsOfUser(userId string, numberInPage int, offset int) (*model.User, error) {
	return r.userDAO.GetClassromsOfUser(userId, numberInPage, offset)
}

func (r *userRepository) GetNotificationsOfUser(userId string, numberInPage int, offset int) (*model.User, error) {
	return r.userDAO.GetNotificationsOfUser(userId, numberInPage, offset)
}

func (r *userRepository) GetAccountsOfUser(userId string) (*model.User, error) {
	return r.userDAO.GetAccountsOfUser(userId)
}

func (r *userRepository) Create(user *model.User) error {
	return r.userDAO.Create(user)
}

func (r *userRepository) Update(user *model.User) error {
	return r.userDAO.Update(user)
}

func (r *userRepository) Delete(userId string) error {
	return r.userDAO.Delete(userId)
}
