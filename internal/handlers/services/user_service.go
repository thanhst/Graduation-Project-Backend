package service

import (
	"errors"
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(userID string) (*model.User, error) {
	if userID == "" {
		return nil, errors.New("user ID is empty")
	}
	return s.userRepo.GetByID(userID)
}

func (s *UserService) CreateUser(user *model.User) error {
	if user == nil {
		return errors.New("invalid user data")
	}
	return s.userRepo.Create(user)
}

func (s *UserService) UpdateUser(user *model.User) error {
	if user == nil || user.UserId == "" {
		return errors.New("invalid user for update")
	}
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(userID string) error {
	if userID == "" {
		return errors.New("user ID is empty")
	}
	return s.userRepo.Delete(userID)
}
