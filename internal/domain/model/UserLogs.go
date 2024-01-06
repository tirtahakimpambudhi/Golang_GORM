package model

import (
	"time"
)

type UserLogs struct {
	ID           int       `gorm:"primaryKey;column:id;<-create;AutoIncrement;type:int" json:"id"`
	UserID       string    `gorm:"column:user_id;type:varchar;" json:"user_id"`
	Title        string    `gorm:"column:title;type:varchar;" json:"title"`
	Action       string    `gorm:"column:action;type:varchar;" json:"action"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
	LastModified int64     `json:"lastModified" gorm:"column:last_modified;autoCreateTime:mili;autoUpdateTime:mili;"`
}

func (l *UserLogs) TableName() string {
	return "user_logs"
}
