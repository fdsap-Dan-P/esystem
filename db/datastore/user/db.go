package db

import (
	"database/sql"
	common "simplebank/db/common"
)

func New(db common.DBTX) *QueriesUser {
	return &QueriesUser{db: db}
}

type QueriesUser struct {
	db common.DBTX
}

func (q *QueriesUser) WithTx(tx *sql.Tx) *QueriesUser {
	return &QueriesUser{
		db: tx,
	}
}
