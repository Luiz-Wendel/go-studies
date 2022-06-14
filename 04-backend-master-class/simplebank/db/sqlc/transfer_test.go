package db

import (
	"context"
	"testing"

	"github.com/Luiz-Wendel/go-studies/04-backend-master-class/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
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
}
