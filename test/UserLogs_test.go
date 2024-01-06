package test

import (
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"gorm.io/gorm"
	"testing"
)

var userDB, _ = db.NewPostgresStore(config.DatabaseConf).Connection()

func TestCreateUserLogs(t *testing.T) {
	userLogs := model.UserLogs{
		UserID: "117fcbe7-0688-4a7b-a8eb-f60ae0f87fac",
		Title:  "Testing UserLogs",
		Action: "Doing Create Logs",
	}
	err := userDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&userLogs).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Fatal(err.Error())
	}
}
