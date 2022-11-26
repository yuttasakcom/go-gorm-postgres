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

type Accounts struct {
	Owner   string `json:"owner"`
	Balance int64  `json:"balance"`
}

type Queries struct {
	db *gorm.DB
}

func (q *Queries) CreateAccount(arg Accounts) *Accounts {
	account := &Accounts{
		Owner:   arg.Owner,
		Balance: arg.Balance,
	}
	q.db.Create(account)

	return account
}
