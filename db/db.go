package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn = "host=localhost user=root password=secret dbname=simple_bank port=5432 sslmode=disable TimeZone=Asia/Bangkok"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func TestDBInit() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	DB = db
	return DB
}
