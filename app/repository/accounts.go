package repository

import (
	"github.com/akashgupta05/account/app/models"
	"github.com/akashgupta05/account/config/db"
	"github.com/jinzhu/gorm"
)

// AccountsRepositoryInterface holds exported methods of AccountsRepository
type AccountsRepositoryInterface interface {
	AccountInfo(string) (*models.Account, error)
	Credit(*models.Credit) error
	Debit(*models.Debit) error
	CreditHistory(string) ([]*models.Credit, error)
}

type AccountsRepository struct {
	Db *gorm.DB
}

// NewAccountsRepository returns instance of AccountsRepository
func NewAccountsRepository() *AccountsRepository {
	return &AccountsRepository{Db: db.Get()}
}

func (ar *AccountsRepository) updateAccount(account *models.Account) error {
	if err := ar.Db.Save(account).Error; err != nil {
		return err
	}

	return nil
}
