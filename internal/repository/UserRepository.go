package repository

import (
	"go_gorm/internal/domain/model"
	"gorm.io/gorm"
)

const BatchSize = 100

type UserRepository interface {
	FindAll(db *gorm.DB) []model.User
	FindByID(db *gorm.DB, ID string) (model.User, error)
	FindByIDs(db *gorm.DB, IDs []string) ([]model.User, error)
	Create(db *gorm.DB, user model.User) error
	CreateMany(db *gorm.DB, users []model.User) error
	Update(db *gorm.DB, ID string, user model.User)
	UpdateMany(db *gorm.DB, conditions map[string]interface{}, users model.User)
	Delete(db *gorm.DB, ID string)
	DeleteMany(db *gorm.DB, IDs []string)
}

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) FindByIDs(db *gorm.DB, IDs []string) ([]model.User, error) {
	var users []model.User
	err := db.Find(&users, IDs).Error
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindAll(db *gorm.DB) []model.User {
	var users []model.User
	db.Find(&users)
	return users
}

func (u *UserRepositoryImpl) FindByID(db *gorm.DB, ID string) (model.User, error) {
	user := model.User{}
	err := db.Take(&user, "id = ?", ID).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) Create(db *gorm.DB, user model.User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) CreateMany(db *gorm.DB, users []model.User) error {
	var err error
	if len(users) > 1000 {
		err = db.CreateInBatches(&users, BatchSize).Error
	} else {
		err = db.Create(&users).Error
	}

	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) Update(db *gorm.DB, ID string, user model.User) {
	db.Where("id = ? ", ID).Updates(&user)
}

func (u *UserRepositoryImpl) UpdateMany(db *gorm.DB, conditions map[string]interface{}, users model.User) {
	db.Where(conditions).Updates(&users)
}

func (u *UserRepositoryImpl) Delete(db *gorm.DB, ID string) {
	db.Where("id = ?", ID).Delete(&model.User{})
}

func (u *UserRepositoryImpl) DeleteMany(db *gorm.DB, IDs []string) {
	db.Where("id IN ?", IDs).Delete(&model.User{})
}
