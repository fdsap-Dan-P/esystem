package db

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	// dsAcc "simplebank/db/datastore/account"

// 	"github.com/shopspring/decimal"
// 	"github.com/stretchr/testify/require"
// )

// func TestTransferTx(t *testing.T) {
// 	store := NewStore(testDB)

// 	account1, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
// 	account2, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")

// 	fmt.Println(">> before:", account1.Balance, account2.Balance)

// 	n := 5
// 	decn := decimal.NewFromInt(int64(n))
// 	amount := decimal.RequireFromString("10")

// 	errs := make(chan error)
// 	results := make(chan TransferTxResult)

// 	// run n concurrent transfer transaction
// 	for i := 0; i < n; i++ {
// 		go func() {
// 			result, err := store.TransferTx(context.Background(), TransferTxParams{
// 				FromAccountId: account1.Id,
// 				ToAccountId:   account2.Id,
// 				Amount:        amount,
// 			})

// 			errs <- err
// 			results <- result
// 		}()
// 	}

// 	// check results
// 	existed := make(map[int]bool)

// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)

// 		result := <-results
// 		require.NotEmpty(t, result)

// 		// check transfer
// 		transfer := result.Transfer
// 		require.NotEmpty(t, transfer)
// 		require.Equal(t, account1.Id, transfer.FromAccountId)
// 		require.Equal(t, account2.Id, transfer.ToAccountId)
// 		require.Equal(t, amount, transfer.Amount)
// 		require.NotZero(t, transfer.Id)
// 		require.NotZero(t, transfer.CreatedAt)

// 		_, err = store.GetTransfer(context.Background(), transfer.Id)
// 		require.NoError(t, err)

// 		// check entries
// 		fromEntry := result.FromEntry
// 		require.NotEmpty(t, fromEntry)
// 		require.Equal(t, account1.Id, fromEntry.AccountId)
// 		require.Equal(t, amount.Neg(), fromEntry.Amount)
// 		require.NotZero(t, fromEntry.Id)
// 		require.NotZero(t, fromEntry.CreatedAt)

// 		_, err = store.GetEntry(context.Background(), fromEntry.Id)
// 		require.NoError(t, err)

// 		toEntry := result.ToEntry
// 		require.NotEmpty(t, toEntry)
// 		require.Equal(t, account2.Id, toEntry.AccountId)
// 		require.Equal(t, amount, toEntry.Amount)
// 		require.NotZero(t, toEntry.Id)
// 		require.NotZero(t, toEntry.CreatedAt)

// 		_, err = store.GetEntry(context.Background(), toEntry.Id)
// 		require.NoError(t, err)

// 		// check accounts
// 		fromAccount := result.FromAccount
// 		require.NotEmpty(t, fromAccount)
// 		require.Equal(t, account1.Id, fromAccount.Id)

// 		toAccount := result.ToAccount
// 		require.NotEmpty(t, toAccount)
// 		require.Equal(t, account2.Id, toAccount.Id)

// 		// check balances
// 		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)

// 		diff1 := account1.Balance.Sub(fromAccount.Balance)
// 		diff2 := toAccount.Balance.Sub(account2.Balance)
// 		require.Equal(t, diff1, diff2)
// 		require.True(t, diff1.GreaterThan(decimal.Zero))
// 		require.True(t, diff1.Mod(amount).IsZero()) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

// 		k := int(diff1.Div(amount).IntPart())
// 		require.True(t, k >= 1 && k <= n)
// 		require.NotContains(t, existed, k)
// 		existed[k] = true
// 	}

// 	// check the final updated balance
// 	updatedAccount1, err := testQueriesAccount.GetAccount(context.Background(), account1.Id)
// 	require.NoError(t, err)

// 	updatedAccount2, err := testQueriesAccount.GetAccount(context.Background(), account2.Id)
// 	require.NoError(t, err)

// 	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

// 	require.Equal(t, account1.Balance.Sub(amount.Mul(decn)), updatedAccount1.Balance)
// 	require.Equal(t, account2.Balance.Add(amount.Mul(decn)), updatedAccount2.Balance)
// }

// func TestTransferTxDeadlock(t *testing.T) {
// 	store := NewStore(testDB)

// 	account1, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
// 	account2, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")

// 	fmt.Println(">> before:", account1.Balance, account2.Balance)

// 	n := 10
// 	// decn := decimal.NewFromInt(int64(n))
// 	amount := decimal.RequireFromString("10")

// 	errs := make(chan error)

// 	for i := 0; i < n; i++ {
// 		fromAccountId := account1.Id
// 		toAccountId := account2.Id

// 		if i%2 == 1 {
// 			fromAccountId = account2.Id
// 			toAccountId = account1.Id
// 		}

// 		go func() {
// 			_, err := store.TransferTx(context.Background(), TransferTxParams{
// 				FromAccountId: fromAccountId,
// 				ToAccountId:   toAccountId,
// 				Amount:        amount,
// 			})

// 			errs <- err
// 		}()
// 	}

// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)
// 	}

// 	// check the final updated balance
// 	updatedAccount1, err := testQueriesAccount.GetAccount(context.Background(), account1.Id)
// 	require.NoError(t, err)

// 	updatedAccount2, err := testQueriesAccount.GetAccount(context.Background(), account2.Id)
// 	require.NoError(t, err)

// 	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
// 	require.Equal(t, account1.Balance, updatedAccount1.Balance)
// 	require.Equal(t, account2.Balance, updatedAccount2.Balance)
// }
