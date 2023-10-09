package db

import (
	"database/sql"
	// customer "simplebank/db/datastore/customer"
	// identity "simplebank/db/datastore/identity"
	// reference "simplebank/db/datastore/reference"
	// user "simplebank/db/datastore/user"
)

// Store defines all functions to execute db queriesUser and users
type StoreUser interface {
	QuerierUser
	// TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides all functions to execute SQL queriesUser and users
type SQLStoreUser struct {
	*QueriesUser
	db *sql.DB
	// *account.QueriesAccount
}

// NewStore creates a new store
func NewStore(db *sql.DB) StoreUser {
	return &SQLStoreUser{
		db:          db,
		QueriesUser: New(db),
	}
}

// // ExecTx executes a function within a database user
// func (store *SQLStore) execTx(ctx context.Context, fn func(*QueriesUser, *account.QueriesAccount) error) error {
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
