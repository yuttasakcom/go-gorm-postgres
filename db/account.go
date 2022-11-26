package db

type Accounts struct {
	Owner   string `json:"owner"`
	Balance int64  `json:"balance"`
}

func (q *Queries) CreateAccount(arg Accounts) *Accounts {
	account := &Accounts{
		Owner:   arg.Owner,
		Balance: arg.Balance,
	}
	q.db.Create(account)

	return account
}
