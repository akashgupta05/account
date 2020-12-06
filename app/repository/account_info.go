package repository

import (
	"fmt"

	"github.com/akashgupta05/account/app/models"
	"github.com/jinzhu/gorm"
)

// AccountInfo for fetching account info of a user
func (ar *AccountsRepository) AccountInfo(userID string) (*models.Account, error) {
	account := &models.Account{}

	err := ar.Db.Where("accounts.user_id = ?", userID).Find(account).Error
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("No account exists for userId : %v", userID)
	}

	ar.reduceExpiredCreditAmount(account)

	return account, err
}

func (ar *AccountsRepository) reduceExpiredCreditAmount(account *models.Account) {
	credits := []*models.Credit{}
	err := ar.Db.Where("credits.expiry < extract(epoch from NOW()) and credits.exausted = false").Find(&credits).Error
	if err != nil || len(credits) == 0 {
		return
	}

	var expiredAmount int64
	txn := ar.Db.Begin()
	for _, credit := range credits {
		expiredAmount += credit.AvailableAmount
		credit.Exausted = true
		if err := ar.Db.Save(credit).Error; err != nil {
			txn.Rollback()
		}
	}

	account.Balance -= expiredAmount
	if err := ar.updateAccount(account); err != nil {
		txn.Rollback()
	}

	txn.Commit()
}
