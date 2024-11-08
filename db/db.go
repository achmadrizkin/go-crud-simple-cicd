package db

import (
	"fmt"
	"go-crud-simple-cicd/config"
	"go-crud-simple-cicd/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	db.AutoMigrate(&model.Book{})
	return db
}
