package repository

import (
	userdao "server/internal/handlers/dao/user"
	model "server/internal/models"
)

var userDAO *userdao.UserDAO

type UserRepository interface {
	GetUserByID(userid string) (*model.User, error)
}

type userRepository struct {
	dao userdao.UserDAO
}

func NewUserRepository(dao userdao.UserDAO) UserRepository {
	return &userRepository{dao}
}

func (r *userRepository) GetUserByID(userId string) (*model.User, error) {
	return r.dao.GetByID(userId)
}
