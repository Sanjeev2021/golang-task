package database

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"task/model"

)

var gormDB *gorm.DB
var db *sql.DB

func Create_DB() {
	dsn := "root:root@tcp(localhost:3306)/task?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	gormDB = db

	err = gormDB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Error migrating the database:", err)
	}

	err = gormDB.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatal("Error migrating the database:", err)
	}

	log.Println("Database connected and migrated successfully.")
}

func Get_GormDB() *gorm.DB {
	if gormDB == nil {
		log.Fatal("Database connection is not initialized. Please call Create_DB first.")
	}
	return gormDB
}
