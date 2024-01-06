package test

import (
	"github.com/google/uuid"
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"go_gorm/pkg/common/json"
	"testing"
	"time"
)

var psstore, _ = db.NewPostgresStore(config.DatabaseConf).Connection()

func TestUpsertOneTOMany(t *testing.T) {
	dateString, layout := "2023-12-31", "2006-01-02"
	dueDate, _ := time.Parse(layout, dateString)
	id := uuid.NewString()
	user := model.User{
		ID:       id,
		Username: "oneToMany",
		Password: "12345678",
		Email:    "onetomany@gmail.com",
		Address: &model.Address{
			Street:     "St ParangTritis 14Mil",
			City:       "DIY",
			PostalCode: 5571,
			Country:    "ID",
		},
		Wallet: model.Wallet{
			UserID:   id,
			Balance:  10000000000,
			Currency: "US",
		},
		TodoList: []model.TodoList{
			{
				UserID:      id,
				TaskName:    "Task 1",
				Description: "Testing Task 1",
				Completed:   false,
				DueDate:     dueDate,
			}, {
				UserID:      id,
				TaskName:    "Task 2",
				Description: "Testing Task 2",
				Completed:   false,
				DueDate:     dueDate,
			}, {
				UserID:      id,
				TaskName:    "Task 3",
				Description: "Testing Task 3",
				Completed:   true,
				DueDate:     dueDate,
			},
		},
	}
	err := psstore.Create(&user).Error
	if err != nil {
		t.Fatal(err.Error())
		return
	}
}

func TestFindAllOneTOMany(t *testing.T) {
	var users []model.User

	psstore.Preload("TodoList").Joins("Wallet").Find(&users)
	t.Log(json.JSONParse(users))
}

func TestBelongsTo(t *testing.T) {
	var todos []model.TodoList
	psstore.Joins("User").Find(&todos)
	t.Log(json.JSONParse(todos))

	var todo model.TodoList
	psstore.Preload("User").Joins("User.Wallet").Where("completed = ?", true).Take(&todo)
	t.Log(json.JSONParse(todo))
}
