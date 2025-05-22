package accountdao

import (
	"errors"
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

func (dao *accountDAOImpl) GetByEmail(email string) ([]*model.Account, error) {
	var accounts []*model.Account
	err := dao.db.Preload("User").Where("email = ?", email).Find(&accounts).Error
	return accounts, err
}

func (dao *accountDAOImpl) GetByEmailAndMethod(email string, method string) (*model.Account, error) {
	var account model.Account
	err := dao.db.Where("email = ? AND login_method = ?", email, method).First(&account).Error
	return &account, err
}
func (dao *accountDAOImpl) GetByUserId(userId string) ([]*model.Account, error) {
	var accounts []*model.Account
	err := dao.db.Where("user_id = ?", userId).Find(&accounts).Error
	return accounts, err
}
func (dao *accountDAOImpl) Create(account *model.Account) error {
	err := dao.db.Create(account).Error
	if err != nil {
		return errors.New("cannot create account")
	}
	return nil
}

func (dao *accountDAOImpl) Update(account *model.Account) error {
	return dao.db.Save(account).Error
}

func (dao *accountDAOImpl) Delete(accountID string) error {
	return dao.db.Delete(&model.Account{}, "account_id = ?", accountID).Error
}
