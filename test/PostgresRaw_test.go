package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	conf := config.DatabaseConf
	dbs := db.NewPostgresStore(conf)
	todoDB, err := dbs.Connection()

	if err != nil {
		assert.NotEqual(t, nil, err)
		t.Fatal(err.Error())
	}
	assert.NotEqual(t, nil, todoDB)

}

func TestFindAll1(t *testing.T) {
	conf := config.DatabaseConf
	dbs := db.NewPostgresStore(conf)
	todoDB, err := dbs.Connection()

	if err != nil {
		assert.NotEqual(t, nil, err)
		t.Fatal(err.Error())
	}

	todolist := []model.TodoList{}
	err = todoDB.Raw("SELECT * FROM public.todolist").Scan(&todolist).Error

	if err != nil {
		assert.NotEqual(t, nil, err)
		t.Fatal(err.Error())
	}

	assert.Equal(t, 3, len(todolist))
	t.Log(todolist)
}

func TestFindAll(t *testing.T) {
	result := []model.TodoList{}
	conf := config.DatabaseConf
	dbs := db.NewPostgresStore(conf)
	todoDB, err := dbs.Connection()

	if err != nil {
		assert.NotEqual(t, nil, err)
		t.Fatal(err.Error())
	}

	args := model.NewArgsRawQuery(todoDB, "SELECT * FROM public.todolist", false)
	tests := []*model.Test{
		model.NewTest("Find ALL", args, 3),
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var gotErr error
			argsConv := test.Args.(*model.ArgsRawQuery)
			argsConv.DB.Raw(argsConv.Query).Scan(&result)
			if wantdata := test.WantData.(int); len(result) != wantdata {
				gotErr = fmt.Errorf("FindAll() gotTotal = %v, want %v", len(result), test.WantData)
			}

			if (gotErr != nil) != argsConv.WantErr {
				t.Errorf("FindALL() error = %v, wantErr %v", gotErr, argsConv.WantErr)
			}

		})
	}
}

func TestUpdate(t *testing.T) {
	dateString := "2023-12-31"
	layout := "2006-01-02"
	dueData, _ := time.Parse(layout, dateString)
	conf := config.DatabaseConf
	dbs := db.NewPostgresStore(conf)
	todoDB, err := dbs.Connection()

	if err != nil {
		assert.NotEqual(t, nil, err)
		t.Fatal(err.Error())
	}
	args := model.NewArgsRawExec(todoDB, "UPDATE public.todolist  SET task_name = ?, description = ? , due_date = ? , completed = ? WHERE id = ?;")
	tests := []*model.Test{
		model.NewTest("update", args, nil),
	}

	for i, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			argsConv := test.Args.(*model.ArgsRawExec)
			err = argsConv.DB.Exec(argsConv.Query, fmt.Sprintf("Task to %v", i), fmt.Sprintf("Description for Task %v", i), dueData, true, 5).Error

			if err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}
func TestDelete(t *testing.T) {
	conf := config.DatabaseConf
	dbs := db.NewPostgresStore(conf)
	todoDB, err := dbs.Connection()

	if err != nil {
		assert.NotEqual(t, nil, err)
		t.Fatal(err.Error())
	}
	args := model.NewArgsRawExec(todoDB, "DELETE FROM public.todolist WHERE id = ?;")
	tests := []*model.Test{
		model.NewTest("update", args, nil),
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			argsConv := test.Args.(*model.ArgsRawExec)
			err = argsConv.DB.Exec(argsConv.Query, 5).Error

			if err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}
func TestInsert(t *testing.T) {
	dateString := "2023-12-31"
	layout := "2006-01-02"
	dueData, _ := time.Parse(layout, dateString)
	conf := config.DatabaseConf
	dbs := db.NewPostgresStore(conf)
	todoDB, err := dbs.Connection()

	if err != nil {
		assert.NotEqual(t, nil, err)
		t.Fatal(err.Error())
	}
	args := model.NewArgsRawExec(todoDB, "INSERT INTO public.todolist (task_name, description, due_date, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	args1 := model.NewArgsRawExec(todoDB, "INSERT INTO public.todolist (task_name, description, due_date, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	args2 := model.NewArgsRawExec(todoDB, "INSERT INTO public.todolist (task_name, description, due_date, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	args3 := model.NewArgsRawExec(todoDB, "INSERT INTO public.todolist (task_name, description, due_date, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	tests := []*model.Test{
		model.NewTest("Insert", args, nil),
		model.NewTest("Insert", args1, nil),
		model.NewTest("Insert", args2, nil),
		model.NewTest("Insert", args3, nil),
	}

	for i, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			argsConv := test.Args.(*model.ArgsRawExec)
			err = argsConv.DB.Exec(argsConv.Query, fmt.Sprintf("Task to %v", i), fmt.Sprintf("Description for Task %v", i), dueData, false, time.Now(), nil).Error

			if err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}
