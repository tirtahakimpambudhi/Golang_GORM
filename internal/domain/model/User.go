package model

import (
	"github.com/google/uuid"
	"go_gorm/pkg/common/bcrypts"
	"gorm.io/gorm"
)

type User struct {
	ID           string     `gorm:"primaryKey;column:id;<-create" json:"id"`
	Username     string     `json:"username" gorm:"unique;type:varchar(255);column:username"`
	Password     string     `json:"password" gorm:"type:varchar(255);column:password"`
	Email        string     `json:"email" gorm:"unique;type:varchar(100);column:email"`
	Address      *Address   `gorm:"embedded"`
	Wallet       Wallet     `gorm:"foreignKey:user_id;references:id"`
	TodoList     []TodoList `gorm:"foreignKey:user_id;references:id"`
	LikeProducts []Product  `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id;"`
}

type Address struct {
	Street     string `json:"street" gorm:"type:varchar(255);column:street"`
	City       string `json:"city" gorm:"type:varchar(255);column:city"`
	PostalCode int    `json:"postal_code" gorm:"type:int;column:postal_code"`
	Country    string `json:"country" gorm:"type:varchar(255);column:country"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	hashPass, err := bcrypts.HashPassword(u.Password, 10)
	if err != nil {
		return err
	}
	u.Password = hashPass
	return nil
}
