package model

import (
	"gorm.io/gorm"
	"time"
)

type TodoList struct {
	ID          int            `gorm:"primaryKey;column:id;<-create" json:"id"`
	UserID      string         `gorm:"column:user_id;" json:"userID"`
	TaskName    string         `json:"task_name" gorm:"type:varchar(255);column:task_name"`
	Description string         `json:"description" gorm:"type:text;column:description"`
	Completed   bool           `json:"completed" gorm:"type:boolean;column:completed;default:false"`
	DueDate     time.Time      `json:"due_date" gorm:"type:date;column:due_date"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;"`
	User        User           `gorm:"foreignKey:user_id;references:id"`
}

func (t *TodoList) TableName() string {
	return "todo_lists"
}
