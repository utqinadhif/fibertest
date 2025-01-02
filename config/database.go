package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	const dsn = "root:@tcp(127.0.0.1:33060)/simtk?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
}
