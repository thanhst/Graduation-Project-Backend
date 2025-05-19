package service

import (
	"errors"
	userdao "server/internal/handlers/dao/user"
	model "server/internal/models"
)

type UserService interface {
	GetUserByID(userID string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(userID string) error
}

type userServiceImpl struct {
	userRepo userdao.UserDAO
}

func NewUserService(userRepo userdao.UserDAO) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl) GetUserByID(userID string) (*model.User, error) {
	if userID == "" {
		return nil, errors.New("user ID is empty")
	}
	return s.userRepo.GetByID(userID)
}

func (s *userServiceImpl) CreateUser(user *model.User) error {
	if user == nil {
		return errors.New("invalid user data")
	}
	return s.userRepo.Create(user)
}

func (s *userServiceImpl) UpdateUser(user *model.User) error {
	if user == nil || user.UserId == "" {
		return errors.New("invalid user for update")
	}
	return s.userRepo.Update(user)
}

func (s *userServiceImpl) DeleteUser(userID string) error {
	if userID == "" {
		return errors.New("user ID is empty")
	}
	return s.userRepo.Delete(userID)
}
