package repository

import (
	accountdao "server/internal/handlers/dao/account"
	model "server/internal/models"
)

type AccountRepository interface {
	GetByID(accountID string) (*model.Account, error)
	GetByEmail(email string) (*model.Account, error)
	Create(account *model.Account) error
	Update(account *model.Account) error
	Delete(accountID string) error
}

type accountRepository struct {
	accountDAO accountdao.AccountDAO
}

func NewAccountRepository(dao accountdao.AccountDAO) AccountRepository {
	return &accountRepository{accountDAO: dao}
}

func (r *accountRepository) GetByID(accountID string) (*model.Account, error) {
	return r.accountDAO.GetByID(accountID)
}

func (r *accountRepository) GetByEmail(email string) (*model.Account, error) {
	return r.accountDAO.GetByEmail(email)
}

func (r *accountRepository) Create(account *model.Account) error {
	return r.accountDAO.Create(account)
}

func (r *accountRepository) Update(account *model.Account) error {
	return r.accountDAO.Update(account)
}

func (r *accountRepository) Delete(accountID string) error {
	return r.accountDAO.Delete(accountID)
}
