package db

import (
	"fmt"
	"go_gorm/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type PostgresStore struct {
	DBConf *model.Database
}

func NewPostgresStore(DBConf *model.Database) *PostgresStore {
	return &PostgresStore{DBConf: DBConf}
}

func (p *PostgresStore) Connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=America/New_York", p.DBConf.Host, p.DBConf.User, p.DBConf.Password, p.DBConf.Name, p.DBConf.Port)
	//Best Perfomance Config GORM
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	//	Logger:                 logger.Default.LogMode(logger.Info),
	//	SkipDefaultTransaction: true,
	//	PrepareStmt:            true,
	//})

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return db, nil
}
