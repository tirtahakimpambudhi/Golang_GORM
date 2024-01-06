package test

import (
	"go_gorm/internal/config"
	"go_gorm/internal/db"
	"go_gorm/internal/domain/model"
	"go_gorm/pkg/common/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
	"time"
)

type userResponse struct {
	Username string
	Email    string
	Address  model.Address `gorm:"embedded"`
}

var database, _ = db.NewPostgresStore(config.DatabaseConf).Connection()

func TestLimitAndOffset(t *testing.T) {
	var users []model.User
	current, limit := 2, 10
	offset := (current - 1) * limit
	err := database.Where("username LIKE ? OR email LIKE ?", "%testing%", "%testing%").Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	t.Log(json.JSONParse(users))
}

func TestNOTORperation(t *testing.T) {
	var user model.User
	cond := map[string]interface{}{
		"username": "testing",
	}
	err := database.Not(cond).Or(cond).Take(&user).Error
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	t.Log(json.JSONParse(user))
}

func TestQueryNONModel(t *testing.T) {
	var response []userResponse
	err := database.Model(&model.User{}).Find(&response).Error
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	t.Log(json.JSONParse(response))
}

func TestQuerySelect(t *testing.T) {
	user := map[string]interface{}{
		"username": "",
		"email":    "",
	}
	err := database.Model(&model.User{}).Select("username", "email").Take(&user).Error

	if err != nil {
		t.Fatal(err.Error())
		return
	}
	t.Log(json.JSONParse(user))
}

func TestDeleteMany(t *testing.T) {
	err := database.Where("username LIKE ?", "%testing%").Delete(&model.User{}).Error
	if err != nil {
		t.Fatal(err.Error())
		return
	}
}

func TestUpsertWithAutoIncrement(t *testing.T) {
	user := model.User{
		ID:       "117fcbe7-0688-4a7b-a8eb-f60ae0f87fac",
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
	//		ID: uuid.NewString(),
	err := database.Transaction(func(tx *gorm.DB) error {
		//the id not exist gorm try insert
		//the id exist gorm try update
		err := tx.Save(&user).Error //insert because id is default value and autofill because hooks and update because the id exist
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
func TestUpsertWithoutAI(t *testing.T) {
	dateString := "2023-12-31"
	layout := "2006-01-02"
	dueDate, _ := time.Parse(layout, dateString)
	todos := model.TodoList{
		TaskName:    "Testing Upsert 100",
		Description: "testing try upsert (update or insert)",
		DueDate:     dueDate,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	err := database.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&todos).Error
		//the id not exist gorm try update
		//the id exist gorm try insert
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

func TestUpsertWithClauses(t *testing.T) {
	user := model.User{
		ID:       "117fcbe7-0688-4a7b-a8eb-f60ae0f87fac",
		Username: "foobar",
		Password: "foobuzz12345",
		Email:    "foobar@gmail.com",
		Address: &model.Address{
			City:       "Bandung",
			Street:     "ST.Bandung Lautan Api",
			Country:    "ID",
			PostalCode: 56789,
		},
	}
	err := database.Transaction(func(tx *gorm.DB) error {
		err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user).Error
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

func TestQueryAggregationCount(t *testing.T) {
	var totalUser int64
	userdb.Model(&model.User{}).Count(&totalUser)
	t.Log(totalUser)
}
func TestQueryAggregationCountTodos(t *testing.T) {
	var totalCompletedTodos int64

	err := userdb.Model(&model.User{}).
		Joins("JOIN todo_lists ON users.id = todo_lists.user_id").
		Where("todo_lists.completed = ?", true).
		Count(&totalCompletedTodos).Error

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	t.Log(totalCompletedTodos)
}

func TestQueryAggregationAnother(t *testing.T) {
	type aggregationResult struct {
		MaxBalance   int64   `json:"maxBalance"`
		MinBalance   int64   `json:"minBalance"`
		TotalBalance int64   `json:"totalBalance"`
		AvgBalance   float64 `json:"avgBalance"`
	}
	var result []aggregationResult
	userdb.Model(model.Wallet{}).Select("sum(balance) as total_balance", "avg(balance) as avg_balance", "min(balance) as min_balance", "max(balance) as max_balance").
		Joins("User").Group("User.id").Having("sum(balance) > ?", 10000).Find(&result)
	t.Log(json.JSONParse(result))
}

func CompletedTodos(db *gorm.DB) *gorm.DB {
	return db.Where("completed = ?", true)
}

func TestScopes(t *testing.T) {
	//cara ke 1
	var todos []model.TodoList
	userdb.Scopes(CompletedTodos).Find(&todos)
	t.Log(json.JSONParse(todos))
	//cara ke 2
	userdb.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("completed = ?", false)
	}).Find(&todos)
	t.Log(json.JSONParse(todos))

}
