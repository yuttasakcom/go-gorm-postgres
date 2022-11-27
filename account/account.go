package account

import "github.com/yuttasakcom/go-gorm-postgres/db"

type Accounts struct {
	Owner   string `json:"owner"`
	Balance int64  `json:"balance"`
}

func CreateAccount(arg Accounts) *Accounts {
	db := db.GetDB()
	account := &Accounts{
		Owner:   arg.Owner,
		Balance: arg.Balance,
	}
	db.Create(account)

	return account
}
