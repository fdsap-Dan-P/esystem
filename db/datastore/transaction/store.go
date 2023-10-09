package db

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	dsAcc "simplebank/db/datastore/account"
// 	"simplebank/model"

// 	"github.com/shopspring/decimal"
// )

// // Store defines all functions to execute db queries and transactions
// type Store interface {
// 	QuerierTransaction
// 	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
// }

// // SQLStoreTransaction provides all functions to execute SQL queries and transactions
// type SQLStoreTransaction struct {
// 	db *sql.DB
// 	*QueriesTransaction
// }

// // NewStore creates a new store
// func NewStore(db *sql.DB) Store {
// 	return &SQLStoreTransaction{
// 		db:                 db,
// 		QueriesTransaction: New(db),
// 	}
// }

// // ExecTx executes a function within a database transaction
// func (store *SQLStoreTransaction) execTx(ctx context.Context, fn func(*dsAcc.QueriesAccount, *QueriesTransaction) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	qAcc := dsAcc.New(tx)
// 	aTrn := New(tx)
// 	err = fn(qAcc, aTrn)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }

// // TransferTxParams contains the input parameters of the transfer transaction
// type TransferTxParams struct {
// 	FromAccountId int64           `json:"from_account_id"`
// 	ToAccountId   int64           `json:"to_account_id"`
// 	Amount        decimal.Decimal `json:"amount"`
// }

// // TransferTxResult is the result of the transfer transaction
// type TransferTxResult struct {
// 	Transfer    model.Transfer `json:"transfer"`
// 	FromAccount model.Account  `json:"from_account"`
// 	ToAccount   model.Account  `json:"to_account"`
// 	FromEntry   model.Entry    `json:"from_entry"`
// 	ToEntry     model.Entry    `json:"to_entry"`
// }

// // TransferTx performs a money transfer from one account to the other.
// // It creates the transfer, add account entries, and update accounts' balance within a database transaction
// func (store *SQLStoreTransaction) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
// 	var result TransferTxResult

// 	err := store.execTx(ctx, func(qacc *dsAcc.QueriesAccount, qtrn *QueriesTransaction) error {
// 		var err error

// 		result.Transfer, err = qtrn.CreateTransfer(ctx, CreateTransferParams{
// 			FromAccountId: arg.FromAccountId,
// 			ToAccountId:   arg.ToAccountId,
// 			Amount:        arg.Amount,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		result.FromEntry, err = qtrn.CreateEntry(ctx, CreateEntryParams{
// 			AccountId: arg.FromAccountId,
// 			Amount:    arg.Amount.Neg(),
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		result.ToEntry, err = qtrn.CreateEntry(ctx, CreateEntryParams{
// 			AccountId: arg.ToAccountId,
// 			Amount:    arg.Amount,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		if arg.FromAccountId < arg.ToAccountId {
// 			result.FromAccount, result.ToAccount, err = addMoney(ctx, qacc, arg.FromAccountId, arg.Amount.Neg(), arg.ToAccountId, arg.Amount)
// 		} else {
// 			result.ToAccount, result.FromAccount, err = addMoney(ctx, qacc, arg.ToAccountId, arg.Amount, arg.FromAccountId, arg.Amount.Neg())
// 		}

// 		return err
// 	})

// 	return result, err
// }

// func addMoney(
// 	ctx context.Context,
// 	acc *dsAcc.QueriesAccount,
// 	accountId1 int64,
// 	amount1 decimal.Decimal,
// 	accountId2 int64,
// 	amount2 decimal.Decimal,
// ) (account1 model.Account, account2 model.Account, err error) {
// 	account1, err = acc.AddAccountBalance(ctx, dsAcc.AddAccountBalanceParams{
// 		Id:     accountId1,
// 		Amount: amount1,
// 	})
// 	if err != nil {
// 		return
// 	}

// 	account2, err = acc.AddAccountBalance(ctx, dsAcc.AddAccountBalanceParams{
// 		Id:     accountId2,
// 		Amount: amount2,
// 	})
// 	return
// }
