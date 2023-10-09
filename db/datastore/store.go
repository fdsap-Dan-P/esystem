package db

import (
	"context"
	"database/sql"
	"fmt"
	account "simplebank/db/datastore/account"
	transaction "simplebank/db/datastore/transaction"
	"simplebank/model"

	common "simplebank/db/common"

	"github.com/shopspring/decimal"
)

// Store defines all functions to execute db queriesUser and users
type Store interface {
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	// TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides all functions to execute SQL queriesUser and users
type SQLStore struct {
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db: db,
	}
}

type Queries struct {
	db common.DBTX
}

func New(db common.DBTX) *Queries {
	return &Queries{db: db}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context,
	fn func(*transaction.QueriesTransaction, *account.QueriesAccount) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := transaction.New(tx)
	a := account.New(tx)
	err = fn(q, a)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountId int64           `json:"fromAccountId"`
	ToAccountId   int64           `json:"toAccountId"`
	Amount        decimal.Decimal `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    model.Transfer `json:"transfer"`
	FromAccount model.Account  `json:"fromAccount"`
	ToAccount   model.Account  `json:"toAccount"`
	FromEntry   model.Entry    `json:"fromEntry"`
	ToEntry     model.Entry    `json:"toEntry"`
}

// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts' balance within a database transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *transaction.QueriesTransaction, a *account.QueriesAccount) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, transaction.CreateTransferParams{
			FromAccountId: arg.FromAccountId,
			ToAccountId:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, transaction.CreateEntryParams{
			AccountId: arg.FromAccountId,
			Amount:    arg.Amount.Neg(),
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, transaction.CreateEntryParams{
			AccountId: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountId < arg.ToAccountId {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, a, arg.FromAccountId, arg.Amount.Neg(), arg.ToAccountId, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, a, arg.ToAccountId, arg.Amount, arg.FromAccountId, arg.Amount.Neg())
		}
		return err
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *account.QueriesAccount,
	accountId1 int64,
	amount1 decimal.Decimal,
	accountId2 int64,
	amount2 decimal.Decimal,
) (account1 model.Account, account2 model.Account, err error) {

	account1, err = q.AddAccountBalance(ctx, account.AddAccountBalanceParams{
		Id:     accountId1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, account.AddAccountBalanceParams{
		Id:     accountId2,
		Amount: amount2,
	})
	return
}
