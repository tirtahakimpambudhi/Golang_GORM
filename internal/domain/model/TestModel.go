package model

import "gorm.io/gorm"

type Test struct {
	Name     string
	Args     interface{}
	WantData interface{}
}

func NewTest(name string, args interface{}, wantData interface{}) *Test {
	return &Test{Name: name, Args: args, WantData: wantData}
}

type ArgsORM struct {
	DB      *gorm.DB
	Todo    *TodoList
	WantErr bool
}
type ArgsORMUsers struct {
	DB      *gorm.DB
	Todo    *User
	WantErr bool
}

func NewArgsORMUsers(DB *gorm.DB, todo *User, wantErr bool) *ArgsORMUsers {
	return &ArgsORMUsers{DB: DB, Todo: todo, WantErr: wantErr}
}

func NewArgsORM(DB *gorm.DB, todo *TodoList, wantErr bool) *ArgsORM {
	return &ArgsORM{DB: DB, Todo: todo, WantErr: wantErr}
}

type ArgsRawQuery struct {
	DB      *gorm.DB
	Query   string
	WantErr bool
}

type ArgsRawExec struct {
	DB    *gorm.DB
	Query string
}

func NewArgsRawExec(DB *gorm.DB, query string) *ArgsRawExec {
	return &ArgsRawExec{DB: DB, Query: query}
}

func NewArgsRawQuery(DB *gorm.DB, query string, wantErr bool) *ArgsRawQuery {
	return &ArgsRawQuery{DB: DB, Query: query, WantErr: wantErr}
}
