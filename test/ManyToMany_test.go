package test

import (
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"go_gorm/pkg/common/json"
	"gorm.io/gorm/clause"
	"testing"
	"time"
)

var userLikeDB, _ = db.NewPostgresStore(config.DatabaseConf).Connection()

func TestCreateToMany(t *testing.T) {
	product := model.Product{
		ID:    2,
		Name:  "Luwak White Coffe",
		Price: 1500,
	}
	userLikeDB.Save(&product)

	userLikeDB.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "d6a17d3b-5756-4e6e-8bbc-bf5194650a42",
		"product_id": 1,
	})
}

func TestPreloadManyToMany(t *testing.T) {
	var products []model.Product
	userLikeDB.Preload("LikeByUser").Where("price = ? ", 1500).Find(&products)
	t.Log(json.JSONParse(products))
}

func TestAssosiationModeFindAll(t *testing.T) {
	var product model.Product
	userLikeDB.Where("id = ?", 1).Take(&product)
	user := []model.User{}
	userLikeDB.Model(&product).Where("country = ?", "ID").Association("LikeByUser").Find(&user)

	t.Log(json.JSONParse(product))
	t.Log(json.JSONParse(user))
}

func TestAssosiationModeAdd(t *testing.T) {
	id := "d40faa4b-7198-4e39-b7e5-c6f65a87716c"
	dateString, layout := "2023-12-31", "2006-01-02"
	dueDate, _ := time.Parse(layout, dateString)
	user := model.User{
		ID:       id,
		Username: "moen",
		Password: "123456789",
		Email:    "moenx@gmail.com",
		Address: &model.Address{
			Street:     "ST.Jalan ParangTritis 14KM",
			City:       "DIY",
			PostalCode: 5571,
			Country:    "ID",
		},
		TodoList: []model.TodoList{
			{UserID: id, TaskName: "Task 1.2", Description: "Testing Task 1.2", Completed: true, DueDate: dueDate},
		},
	}

	var product model.Product
	userLikeDB.Where("id = ?", 2).Take(&product)
	userLikeDB.Save(&user)
	userLikeDB.Model(&product).Association("LikeByUser").Append(&user)
}

func TestAssosiationDelete(t *testing.T) {
	var product model.Product
	userLikeDB.Where("id = ?", 2).Take(&product)
	var user model.User
	userLikeDB.Where("id = ?", "d40faa4b-7198-4e39-b7e5-c6f65a87716c").Take(&user)
	userLikeDB.Model(&user).Association("LikeProducts").Delete(&product)
}

func TestAssosiationModeClear(t *testing.T) {
	var product model.Product
	userLikeDB.Where("id = ?", 2).Take(&product)
	userLikeDB.Model(&product).Association("LikeByUser").Clear()
}

func TestPreloadingInlineCond(t *testing.T) {
	var user model.User
	userLikeDB.Preload("TodoList", "completed = ?", true).Where("id = ?", "d6a17d3b-5756-4e6e-8bbc-bf5194650a42").Take(&user)
	t.Log(json.JSONParse(user))
}

func TestPreloadingNested(t *testing.T) {
	var wallet model.Wallet
	userLikeDB.Preload("User.TodoList").Where("id = ?", 9).Take(&wallet)
	t.Log(json.JSONParse(wallet))
}

func TestPreloadALL(t *testing.T) {
	var user model.User
	userLikeDB.Preload(clause.Associations).Where("id = ?", "d6a17d3b-5756-4e6e-8bbc-bf5194650a42").Take(&user)
	t.Log(json.JSONParse(user))
}

func TestJoinQueryAndCond(t *testing.T) {
	var users []model.User
	userLikeDB.Joins("join todo_lists on todo_lists.user_id = users.id AND todo_lists.completed = ?", false).Find(&users)
	t.Log(json.JSONParse(users))
}
