package repository

import (
	"errors"

	"github.com/akashgupta05/account/app/models"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Credit for creating transaction of credit type
func (ar *AccountsRepository) Credit(credit *models.Credit) error {
	account := &models.Account{}
	if err := ar.Db.Where("accounts.user_id = ?", credit.UserID).Find(account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			account.UserID = credit.UserID
			account.Balance = credit.CreditAmount
			if err := ar.Db.Create(account).Error; err != nil {
				return err
			}

			credit.AccountID = account.ID
			return ar.createCredit(credit)
		}
		log.Error("Could not search accounts", err)
		return errors.New("Failed while searching for user account")
	}

	credit.AccountID = account.ID
	if err := ar.createCredit(credit); err != nil {
		return err
	}

	account.Balance += credit.CreditAmount
	return ar.updateAccount(account)
}

func (ar *AccountsRepository) createCredit(t *models.Credit) error {
	if err := ar.Db.Create(t).Error; err != nil {
		return err
	}

	return nil
}
