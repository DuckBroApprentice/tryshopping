package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConnect gorm.DB

func ConnectToMySQL() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/Demo?charset=utf8mb4&parseTime=True&loc=Local"
	DBConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(DBConnect)
}
