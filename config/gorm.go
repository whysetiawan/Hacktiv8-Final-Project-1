package config

import (
	"final-project-1/httpserver/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.TodoModel{})
	return db, err
}
