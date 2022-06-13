package db

import (
	"context"
	"testing"
	"time"

	"github.com/Luiz-Wendel/go-studies/04-backend-master-class/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	recoveredAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, recoveredAccount)

	require.Equal(t, createdAccount.ID, recoveredAccount.ID)
	require.Equal(t, createdAccount.Owner, recoveredAccount.Owner)
	require.Equal(t, createdAccount.Balance, recoveredAccount.Balance)
	require.Equal(t, createdAccount.Currency, recoveredAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, recoveredAccount.CreatedAt, time.Second)
}
