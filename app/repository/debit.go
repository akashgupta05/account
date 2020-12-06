package repository

import (
	"errors"
	"fmt"

	"github.com/akashgupta05/account/app/models"
	"github.com/jinzhu/gorm"
)

// Debit for creating transaction of debit type
func (ar *AccountsRepository) Debit(debit *models.Debit) error {

	account := &models.Account{}
	err := ar.Db.Where("accounts.user_id = ?", debit.UserID).Find(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("no user registered with given user id")
		}
		return errors.New("Failed while searching for user account")
	}

	if account.Balance < debit.Amount {
		return errors.New("insufficient balance")
	}
	debit.AccountID = account.ID

	eligibleCredits, err := ar.fetchEligibleCredits(debit)
	if err != nil {
		return err
	}

	if len(eligibleCredits) == 0 {
		return errors.New("couldn't debit with given amount and credits")
	}

	return ar.processDebit(eligibleCredits, debit, account)
}

func (ar *AccountsRepository) fetchEligibleCredits(debit *models.Debit) ([]*models.Credit, error) {
	credits := []*models.Credit{}
	err := ar.Db.Where("credits.exausted = false and credits.account_id = ?", debit.AccountID).Order("priority desc, expiry").Find(&credits).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no user registered with given user id")
		}

		return credits, fmt.Errorf("Failed while searching for credits %v", err)
	}

	for i := range credits {
		if i+debit.UsedCredits > len(credits) {
			break
		}

		var sum int64
		for j := 0; j < debit.UsedCredits; j++ {
			sum += credits[i : i+debit.UsedCredits][j].AvailableAmount
		}

		if sum >= debit.Amount {
			return credits[i : i+debit.UsedCredits], nil
		}
	}

	return credits, errors.New("couldn't find credits with given amount and credit count")
}

func (ar *AccountsRepository) processDebit(credits []*models.Credit, debit *models.Debit, account *models.Account) error {
	txn := ar.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			txn.Rollback()
		}
	}()

	debitAmount := debit.Amount
	for i, credit := range credits {
		offset := int64(debit.UsedCredits) - int64(i+1)
		if credit.AvailableAmount > debitAmount-offset {
			credit.AvailableAmount -= (debitAmount - offset)
			debitAmount = offset
		} else {
			debitAmount -= credit.AvailableAmount
			credit.AvailableAmount = 0
			credit.Exausted = true
		}

		debit.UsedCreditIDs = append(debit.UsedCreditIDs, credit.ID)
		err := txn.Save(credit).Error
		if err != nil {
			txn.Rollback()
			return fmt.Errorf("error while updating credit :%v", err)
		}
	}

	if err := txn.Create(debit).Error; err != nil {
		txn.Rollback()
		return err
	}

	account.Balance -= debit.Amount
	if err := txn.Save(account).Error; err != nil {
		txn.Rollback()
		return err
	}

	return txn.Commit().Error
}
