package models

import (
	"time"
)

// Account holds account table model
type Account struct {
	ID        string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID    string    `json:"user_id" gorm:"not null;column:user_id;default:null"`
	Balance   int64     `json:"balance" gorm:"not null;column:balance;default:null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Transaction holds transaction table model
type Transaction struct {
	ID        string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AccountID string    `json:"account_id" gorm:"not null;column:account_id;default:null"`
	UserID    string    `json:"user_id" gorm:"not null;column:user_id;default:null"`
	Type      string    `json:"type" gorm:"not null;column:type;default:null"`
	Amount    int64     `json:"amount" gorm:"not null;column:amount;default:null"`
	Priority  int       `json:"priority" gorm:"column:priority;default:null"`
	Expiry    int64     `json:"expiry" gorm:"column:expiry;default:null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
