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

	return account, err
}
