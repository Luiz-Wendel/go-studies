package db

import (
	"context"
	"testing"
	"time"

	"github.com/Luiz-Wendel/go-studies/04-backend-master-class/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, id int64) Entry {
	arg := CreateEntryParams{
		AccountID: id,
		Amount:    util.RandomMoney(),
	}

	createdEntry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdEntry)

	require.Equal(t, arg.AccountID, createdEntry.AccountID)
	require.Equal(t, arg.Amount, createdEntry.Amount)

	require.NotZero(t, createdEntry.ID)
	require.NotZero(t, createdEntry.CreatedAt)

	return createdEntry
}

func TestCreateEntry(t *testing.T) {
	createdAccount := CreateRandomAccount(t)

	createRandomEntry(t, createdAccount.ID)
}

func TestGetEntry(t *testing.T) {
	createdAccount := CreateRandomAccount(t)

	createdEntry := createRandomEntry(t, createdAccount.ID)

	recoveredEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, recoveredEntry)

	require.Equal(t, createdEntry.ID, recoveredEntry.ID)
	require.Equal(t, createdEntry.AccountID, recoveredEntry.AccountID)
	require.Equal(t, createdEntry.Amount, recoveredEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, recoveredEntry.CreatedAt, time.Second)
}
