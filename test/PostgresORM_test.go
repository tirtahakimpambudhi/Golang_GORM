package test

import (
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"testing"
	"time"
)

func TestCreateTodo(t *testing.T) {
	conf := config.DatabaseConf
	dbs := db.NewPostgresStore(conf)
	todoDB, _ := dbs.Connection()
	dateString := "2023-12-31"
	layout := "2006-01-02"
	dueDate, _ := time.Parse(layout, dateString)

	input := model.TodoList{
		TaskName:    "Testing GORM",
		UserID:      "117fcbe7-0688-4a7b-a8eb-f60ae0f87fac",
		Description: "For Test ORM POSTGRES",
		Completed:   true,
		DueDate:     dueDate,
	}
	args := model.NewArgsORM(todoDB, &input, false)

	tests := []*model.Test{
		model.NewTest("Create TodoList", args, nil),
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			argsConv := test.Args.(*model.ArgsORM)
			err := argsConv.DB.Model(&model.TodoList{}).Unscoped().Where("id = ?", 1).Update("deleted_at", nil).Error

			if err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	conf := config.DatabaseConf
	dbs := db.NewPostgresStore(conf)
	userDB, _ := dbs.Connection()

	input := model.User{
		Username: "tirtazyygamingganzz",
		Password: "tirta123",
		Email:    "tirtaz@gmail.com",
		Address: &model.Address{
			Street:     "ST.JENDRAL SOEDIRMAN",
			PostalCode: 52802,
			Country:    "ID",
			City:       "Jakarta",
		},
	}
	args := model.NewArgsORMUsers(userDB, &input, false)

	tests := []*model.Test{
		model.NewTest("Create Users", args, nil),
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			argsConv := test.Args.(*model.ArgsORMUsers)
			err := argsConv.DB.Create(argsConv.Todo).Error

			if err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}
