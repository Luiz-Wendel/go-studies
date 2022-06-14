package db

import (
	"context"
	"testing"
	"time"

	"github.com/Luiz-Wendel/go-studies/04-backend-master-class/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAccountID, toAccountID int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        util.RandomMoney(),
	}

	createdTransfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdTransfer)

	require.Equal(t, arg.FromAccountID, createdTransfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, createdTransfer.ToAccountID)
	require.Equal(t, arg.Amount, createdTransfer.Amount)
	require.NotZero(t, createdTransfer.ID)
	require.NotZero(t, createdTransfer.CreatedAt)

	return createdTransfer
}

func TestCreateTransfer(t *testing.T) {
	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)

	createRandomTransfer(t, fromAccount.ID, toAccount.ID)
}

func TestGetTransfer(t *testing.T) {
	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)
	createdTransfer := createRandomTransfer(t, fromAccount.ID, toAccount.ID)

	recoveredTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, recoveredTransfer)

	require.Equal(t, createdTransfer.ID, recoveredTransfer.ID)
	require.Equal(t, createdTransfer.FromAccountID, recoveredTransfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, recoveredTransfer.ToAccountID)
	require.Equal(t, createdTransfer.Amount, recoveredTransfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, recoveredTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)

	const SIZE = 5

	for index := 0; index < SIZE*2; index++ {
		createRandomTransfer(t, fromAccount.ID, toAccount.ID)
	}

	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         SIZE,
		Offset:        SIZE,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, SIZE)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
