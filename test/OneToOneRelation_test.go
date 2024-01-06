package test

import (
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"go_gorm/internal/repository"
	"go_gorm/pkg/common/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

var repositoryUser = repository.NewUserRepositoryImpl()
var userdb, _ = db.NewPostgresStore(config.DatabaseConf).Connection()

func TestCreateWallet(t *testing.T) {
	userID := "ee96fa7a-a55d-42be-adbb-ea57c9b24c1c"
	wallet := model.Wallet{
		UserID:   userID,
		Balance:  10000000000,
		Currency: "RP",
	}
	userdb.Transaction(func(tx *gorm.DB) error {
		tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&wallet)
		return nil
	})
}
func TestJoinWithPreload(t *testing.T) {
	users := []model.User{}
	userdb.Transaction(func(tx *gorm.DB) error {
		tx.Model(&model.User{}).Preload("Wallet").Find(&users)
		return nil
	})

	t.Log(json.JSONParse(users))
}
func TestJoinsOneToOne(t *testing.T) {
	users := []model.User{}
	userdb.Transaction(func(tx *gorm.DB) error {
		tx.Model(&model.User{}).Joins("Wallet").Find(&users)
		return nil
	})
	t.Log(json.JSONParse(users))
}

func TestUpsertOneToOne(t *testing.T) {
	id := "a699287f-59fe-4946-8556-661b227daf58"
	user := model.User{
		ID:       id,
		Username: "footest",
		Password: "123456789",
		Email:    "footest@gmail.com",
		Address: &model.Address{
			Street:     "ST.Jalan ParangTritis 14KM",
			City:       "DIY",
			PostalCode: 5571,
			Country:    "ID",
		},
		Wallet: model.Wallet{
			ID:       3,
			UserID:   id,
			Balance:  10000000,
			Currency: "RP",
		},
	}
	userdb.Transaction(func(tx *gorm.DB) error {
		tx.Save(&user)
		return nil
	})
}

func TestBelongsToOne(t *testing.T) {
	var wallets []model.Wallet
	userdb.Joins("User").Preload("User.TodoList").Find(&wallets)
	t.Log(json.JSONParse(wallets))
}
