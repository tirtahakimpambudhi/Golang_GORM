package test

import (
	"fmt"
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"go_gorm/internal/repository"
	"go_gorm/pkg/common/bcrypts"
	"go_gorm/pkg/common/json"
	"gorm.io/gorm"
	"testing"
)

var repo = repository.NewUserRepositoryImpl()
var dbs, _ = db.NewPostgresStore(config.DatabaseConf).Connection()

func TestFindAllUserRepository(t *testing.T) {
	dbs.Transaction(func(tx *gorm.DB) error {
		users := repo.FindAll(tx)
		t.Log(json.JSONParse(users))
		return nil
	})
}

func TestFindByID(t *testing.T) {
	id := "d976c140-b1d7-4e3c-b41c-8ed502d930de"
	errTX := dbs.Transaction(func(tx *gorm.DB) error {
		user, err := repo.FindByID(tx, id)
		if err != nil {
			t.Fatal(err.Error())
			return err
		}
		t.Log(json.JSONParse(user))
		t.Log(bcrypts.CheckPasswordHash("foobar@gmail.com", user.Password))
		return nil
	})

	if errTX != nil {
		t.Fatal(errTX.Error())
	}
}

func TestCreate(t *testing.T) {
	user := model.User{
		Username: "foobar",
		Password: "foobuzz1234",
		Email:    "foobar@gmail.com",
		Address: &model.Address{
			City:       "Bandung",
			Street:     "ST.Bandung Lautan Api",
			Country:    "ID",
			PostalCode: 56789,
		},
	}

	err := dbs.Transaction(func(tx *gorm.DB) error {
		err := repo.Create(tx, user)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCreateMany(t *testing.T) {
	var users []model.User
	var usernames []string

	for i := 1; i < 10001; i++ {
		user := model.User{
			Username: fmt.Sprintf("testing ke -%v", i),
			Password: fmt.Sprintf("testing%v", i),
			Email:    fmt.Sprintf("testing%v@gmail.com", i),
			Address: &model.Address{
				City:       "Bandung",
				Street:     "ST.Bandung Lautan Api",
				Country:    "ID",
				PostalCode: 56789,
			},
		}

		users = append(users, user)
		usernames = append(usernames, user.Username)
	}

	err := dbs.Transaction(func(tx *gorm.DB) error {
		err := repo.CreateMany(tx, users)
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

func TestUpdateUsers(t *testing.T) {
	id := "39e917e9-28fd-4212-bb6b-4644cbcfe08f"
	err := dbs.Transaction(func(tx *gorm.DB) error {
		repo.Update(tx, id, model.User{
			Username: "tirtahakimpambudhi",
			Email:    "tirtanewwhakim123@gmail.com",
		})
		return nil
	})

	if err != nil {
		t.Fatal(err.Error())
		return
	}

}

func TestDeleteUsers(t *testing.T) {
	id := "39e917e9-28fd-4212-bb6b-4644cbcfe08f"

	err := dbs.Transaction(func(tx *gorm.DB) error {
		repo.Delete(tx, id)
		return nil
	})

	if err != nil {
		t.Fatal(err.Error())
		return
	}
}
