package repository

import (
	"errors"

	"github.com/akashgupta05/account/app/models"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Credit for creating transaction of credit type
func (ar *AccountsRepository) Credit(t *models.Transaction) error {
	account := &models.Account{}
	err := ar.Db.Find(account).Where("accounts.user_id = ?", t.UserID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			account.UserID = t.UserID
			account.Balance = t.Amount
			if err := ar.Db.Create(account).Error; err != nil {
				return err
			}

			t.AccountID = account.ID
			return ar.createTransaction(t)
		}
		log.Error("Could not search accounts", err)
		return errors.New("Failed while searching for user account")
	}

	t.AccountID = account.ID
	if err = ar.createTransaction(t); err != nil {
		return err
	}

	account.Balance += t.Amount
	return ar.updateAccount(account)
}
