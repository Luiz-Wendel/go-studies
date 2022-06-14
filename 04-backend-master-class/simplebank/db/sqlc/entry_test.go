package db

import (
	"context"
	"testing"

	"github.com/Luiz-Wendel/go-studies/04-backend-master-class/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	createdAccount := CreateRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: createdAccount.ID,
		Amount:    util.RandomMoney(),
	}

	createdEntry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdEntry)

	require.Equal(t, arg.AccountID, createdEntry.AccountID)
	require.Equal(t, arg.Amount, createdEntry.Amount)

	require.NotZero(t, createdEntry.ID)
	require.NotZero(t, createdEntry.CreatedAt)
}
