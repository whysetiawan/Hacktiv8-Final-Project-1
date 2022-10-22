package config

import (
	"final-project-1/httpserver/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=postgres dbname=final_project port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.TodoModel{})
	return db, err
}
