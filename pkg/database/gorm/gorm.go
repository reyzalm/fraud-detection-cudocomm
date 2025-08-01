package gorm

import (
	"fmt"

	"github.com/CudoCommunication/cudocomm/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.Env.DBHost,
		config.Env.DBUser,
		config.Env.DBPass,
		config.Env.DBName,
		config.Env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
