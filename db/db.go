package db

import (
	"gorm.io/gorm"
)

const (
	dsn = "host=localhost user=root password=secret dbname=simple_bank port=5432 sslmode=disable TimeZone=Asia/Bangkok"
)

func New(db *gorm.DB) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db *gorm.DB
}
