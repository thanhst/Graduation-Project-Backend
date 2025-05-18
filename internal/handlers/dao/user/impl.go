package userdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type userDAOImpl struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &userDAOImpl{db: db}
}

func (dao *userDAOImpl) GetByID(userId string) (*model.User, error) {
	var user model.User
	if err := dao.db.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDAOImpl) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := dao.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDAOImpl) GetClassromsOfUser(userId string, numberInPage int, offset int) (*model.User, error) {
	var user model.User
	err := dao.db.
		Preload("Classrooms", func(db *gorm.DB) *gorm.DB {
			return db.Limit(numberInPage).Offset(offset)
		}).
		Where("user_id = ?", userId).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (dao *userDAOImpl) GetNotificationsOfUser(userId string, numberInPage int, offset int) (*model.User, error) {
	var user model.User
	err := dao.db.
		Preload("Notifications", func(db *gorm.DB) *gorm.DB {
			return db.Limit(numberInPage).Offset(offset)
		}).
		Where("user_id = ?", userId).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDAOImpl) GetAccountsOfUser(userId string) (*model.User, error) {
	var user model.User
	err := dao.db.
		Preload("Accounts").
		Where("user_id = ?", userId).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDAOImpl) Create(user *model.User) error {
	return dao.db.Create(user).Error
}

func (dao *userDAOImpl) Update(user *model.User) error {
	return dao.db.Save(user).Error
}

func (dao *userDAOImpl) Delete(userId string) error {
	return dao.db.Delete(&model.User{}, userId).Error
}
