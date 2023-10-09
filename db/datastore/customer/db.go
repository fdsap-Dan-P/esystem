package db

import (
	"database/sql"
	common "simplebank/db/common"

)

func New(db common.DBTX) *QueriesCustomer {
	return &QueriesCustomer{db: db}
}

type QueriesCustomer struct {
	db common.DBTX
}

func (q *QueriesCustomer) WithTx(tx *sql.Tx) *QueriesCustomer {
	return &QueriesCustomer{
		db: tx,
	}
}
