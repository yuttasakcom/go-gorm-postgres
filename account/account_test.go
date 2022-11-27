package account

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuttasakcom/go-gorm-postgres/db"
)

func TestCreateAccount(t *testing.T) {
	db.TestDBInit()
	balance := int64(100)
	account := CreateAccount(Accounts{
		Owner:   "yuttasak",
		Balance: balance,
	})

	require.NotEmpty(t, account)
	require.Equal(t, balance, account.Balance)
}
