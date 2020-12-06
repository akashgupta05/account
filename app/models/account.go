package models

import (
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// Account holds account table model
type Account struct {
	ID        string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID    string    `json:"user_id" gorm:"not null;column:user_id;default:null"`
	Balance   int64     `json:"balance" gorm:"not null;column:balance;default:null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Credit holds transaction table model
type Credit struct {
	ID              string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AccountID       string    `json:"account_id" gorm:"not null;column:account_id;default:null"`
	UserID          string    `json:"user_id,omitempty" gorm:"-"`
	Type            string    `json:"type" gorm:"not null;column:type;default:null"`
	CreditAmount    int64     `json:"credit_amount" gorm:"not null;column:credit_amount;default:null"`
	AvailableAmount int64     `json:"available_amount" gorm:"not null;column:available_amount;default:null"`
	Exausted        bool      `json:"exausted" gorm:"not null;column:exausted;default:false"`
	Priority        int       `json:"priority" gorm:"column:priority;default:null"`
	Expiry          int64     `json:"expiry" gorm:"column:expiry;default:null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Debit holds transaction table model
type Debit struct {
	ID            string         `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AccountID     string         `json:"account_id" gorm:"not null;column:account_id;default:null"`
	UserID        string         `json:"user_id,omitempty" gorm:"-"`
	UsedCredits   int            `json:"used_credits" gorm:"not null;column:used_credits;default:0"`
	UsedCreditIDs pq.StringArray `json:"used_credit_ids" gorm:"not null;column:used_credit_ids;default:{}"`
	Amount        int64          `json:"amount" gorm:"not null;column:amount;default:null"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}
