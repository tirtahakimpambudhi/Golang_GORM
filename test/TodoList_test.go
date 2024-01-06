package test

import (
	"fmt"
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"go_gorm/pkg/common/json"
	"gorm.io/gorm"
	"testing"
	"time"
)

var todosDB, _ = db.NewPostgresStore(config.DatabaseConf).Connection()

func TestFindAllTodos(t *testing.T) {
	todos := []model.TodoList{}
	err := todosDB.Transaction(func(tx *gorm.DB) error {
		tx.Find(&todos)
		return nil
	})

	if err != nil {
		t.Fatal(err.Error())
		return
	}
	t.Log(json.JSONParse(todos))
}
func TestRestoreAllDeleteAtSoftTodos(t *testing.T) {
	err := todosDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Unscoped().Model(&model.TodoList{}).Not("deleted_at", nil).Update("deleted_at", nil).Error

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Fatal(err.Error())
		return
	}
}
func TestDeleteSoftByIDTodo(t *testing.T) {
	userId := "117fcbe7-0688-4a7b-a8eb-f60ae0f87fac"
	err := todosDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("user_id = ?", userId).Delete(&model.TodoList{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Fatal(err.Error())
	}

}
func TestCreatesManyTodos(t *testing.T) {
	dateString, layout := "2023-12-31", "2006-01-02"
	dueDate, _ := time.Parse(layout, dateString)
	todos, totalTodos := []model.TodoList{}, 100
	for i := 1; i <= totalTodos; i++ {
		todos = append(todos, model.TodoList{
			UserID:      "117fcbe7-0688-4a7b-a8eb-f60ae0f87fac",
			TaskName:    fmt.Sprintf("Tugas Ke - %v", i),
			Description: fmt.Sprintf("Ini Hanya Mencoba Saja Ke - %v", i),
			Completed:   false,
			DueDate:     dueDate,
		})
	}

	err := todosDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&todos).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	TestFindAllTodos(t)
	TestDeleteSoftByIDTodo(t)
	TestFindAllTodos(t)
	TestRestoreAllDeleteAtSoftTodos(t)
	TestFindAllTodos(t)
}
