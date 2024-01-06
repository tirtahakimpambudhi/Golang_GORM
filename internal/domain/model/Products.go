package model

import "time"

type Product struct {
	ID         int       `gorm:"primaryKey;column:id"`
	Name       string    `gorm:"column:name"`
	Price      int       `gorm:"column:price"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	LikeByUser []User    `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id;"`
}

func (p *Product) TableName() string {
	return "products"
}
