package accountdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type accountDAOImpl struct {
	db *gorm.DB
}

func NewAccountDAO(db *gorm.DB) AccountDAO {
	return &accountDAOImpl{db: db}
}

func (dao *accountDAOImpl) GetByID(accountID string) (*model.Account, error) {
	var account model.Account
	err := dao.db.Preload("User").First(&account, "account_id = ?", accountID).Error
	return &account, err
}

func (dao *accountDAOImpl) GetByEmail(email string) (*model.Account, error) {
	var account model.Account
	err := dao.db.Preload("User").First(&account, "email = ?", email).Error
	return &account, err
}

func (dao *accountDAOImpl) Create(account *model.Account) error {
	return dao.db.Create(account).Error
}

func (dao *accountDAOImpl) Update(account *model.Account) error {
	return dao.db.Save(account).Error
}

func (dao *accountDAOImpl) Delete(accountID string) error {
	return dao.db.Delete(&model.Account{}, "account_id = ?", accountID).Error
}
