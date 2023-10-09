package db

import (
	"database/sql"
	// customer "simplebank/db/datastore/customer"
	// identity "simplebank/db/datastore/identity"
	// reference "simplebank/db/datastore/reference"
	// customer "simplebank/db/datastore/customer"
)

// Store defines all functions to execute db queriesCustomer and customers
type StoreCustomer interface {
	QuerierCustomer
	// TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides all functions to execute SQL queriesCustomer and customers
type SQLStoreCustomer struct {
	*QueriesCustomer
	db *sql.DB
	// *account.QueriesAccount
}

// NewStore creates a new store
func NewStore(db *sql.DB) StoreCustomer {
	return &SQLStoreCustomer{
		db:              db,
		QueriesCustomer: New(db),
	}
}

var _ StoreCustomer = (*SQLStoreCustomer)(nil)

// // ExecTx executes a function within a database customer
// func (store *SQLStore) execTx(ctx context.Context, fn func(*QueriesCustomer, *account.QueriesAccount) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	a := account.New(tx)
// 	err = fn(q, a)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }
