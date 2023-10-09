package db

import (
	"database/sql"
	common "simplebank/db/common"
)

func New(db common.DBTX) *QueriesSocialMedia {
	return &QueriesSocialMedia{db: db}
}

type QueriesSocialMedia struct {
	db common.DBTX
}

func (q *QueriesSocialMedia) WithTx(tx *sql.Tx) *QueriesSocialMedia {
	return &QueriesSocialMedia{
		db: tx,
	}
}
