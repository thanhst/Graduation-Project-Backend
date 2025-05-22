package accountdao

import model "server/internal/models"

type AccountDAO interface {
	GetByID(accountID string) (*model.Account, error)
	GetByEmail(email string) ([]*model.Account, error)
	GetByEmailAndMethod(email string, method string) (*model.Account, error)
	GetByUserId(userId string) ([]*model.Account, error)
	Create(account *model.Account) error
	Update(account *model.Account) error
	Delete(accountID string) error
}
