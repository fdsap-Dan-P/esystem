package db

import (
	"context"
	"testing"
	"time"

	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountId1 int64, accountId2 int64) model.Transfer {
	arg := CreateTransferParams{
		FromAccountId: accountId1,
		ToAccountId:   accountId2,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueriesTransaction.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountId, transfer.FromAccountId)
	require.Equal(t, arg.ToAccountId, transfer.ToAccountId)
	require.True(t, arg.Amount.Equal(transfer.Amount))

	require.NotZero(t, transfer.Id)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	account2, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")
	createRandomTransfer(t, account1.Id, account2.Id)
}

func TestGetTransfer(t *testing.T) {
	account1, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	account2, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")

	transfer1 := createRandomTransfer(t, account1.Id, account2.Id)

	transfer2, err := testQueriesTransaction.GetTransfer(context.Background(), transfer1.Id)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.Id, transfer2.Id)
	require.Equal(t, transfer1.FromAccountId, transfer2.FromAccountId)
	require.Equal(t, transfer1.ToAccountId, transfer2.ToAccountId)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	// account1 := CreateRandomAccount(t)
	// account2 := createRandomAccount(t)

	account1, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	account2, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1.Id, account2.Id)
		createRandomTransfer(t, account2.Id, account1.Id)
	}

	arg := ListTransfersParams{
		FromAccountId: account1.Id,
		ToAccountId:   account1.Id,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueriesTransaction.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountId == account1.Id || transfer.ToAccountId == account1.Id)
	}
}
