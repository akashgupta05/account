package repository

import (
	"fmt"

	"github.com/akashgupta05/account/app/models"
	"github.com/jinzhu/gorm"
)

// CreditHistory returns all the credit activity of a user
func (ar *AccountsRepository) CreditHistory(userID string) ([]*models.Credit, error) {
	account := &models.Account{}
	if err := ar.Db.Where("accounts.user_id = ?", userID).Find(account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("No account exists for userId : %v", userID)
		}

		return nil, err
	}

	credits := []*models.Credit{}
	if err := ar.Db.Where("credits.account_id = ?", account.ID).Order("created_at").Find(&credits).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("No credit activity for userId : %v", userID)
		}
		return nil, err
	}

	return credits, nil
}
