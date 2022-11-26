package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	balance := int64(100)
	account := testQueries.CreateAccount(Accounts{
		Owner:   "yea",
		Balance: balance,
	})

	require.NotEmpty(t, account)
	require.Equal(t, balance, account.Balance)
}
