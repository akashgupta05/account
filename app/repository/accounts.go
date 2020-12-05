package repository

import (
	"github.com/akashgupta05/account/app/models"
	"github.com/akashgupta05/account/config/db"
	"github.com/jinzhu/gorm"
)

type AccountsRepositoryInterface interface {
	Credit(*models.Transaction) error
}

type AccountsRepository struct {
	Db *gorm.DB
}

func NewAccountsRepository() *AccountsRepository {
	return &AccountsRepository{Db: db.Get()}
}

func (ar *AccountsRepository) createTransaction(t *models.Transaction) error {
	if err := ar.Db.Create(t).Error; err != nil {
		return err
	}

	return nil
}

func (ar *AccountsRepository) updateAccount(account *models.Account) error {
	if err := ar.Db.Save(account).Error; err != nil {
		return err
	}

	return nil
}
