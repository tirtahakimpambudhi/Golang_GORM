package model

import "time"

type Wallet struct {
	ID        int       `gorm:"primaryKey;column:id;<-create" json:"id"`
	UserID    string    `gorm:"column:user_id;" json:"userID"`
	Balance   int64     `json:"balance" gorm:"column:balance"`
	Currency  string    `json:"currency" gorm:"currency"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      *User     `json:"user" gorm:"foreignKey:user_id;references:id"`
}

func (w *Wallet) TableName() string {
	return "wallet"
}
