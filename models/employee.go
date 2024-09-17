package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Employee struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Username string `json:"username"`
	Password 	 int    `json:"password"`
}

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("employee.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Employee{})

	if err != nil {
		panic("Failed to migrate database schema!")
	}

	DB = database
}
