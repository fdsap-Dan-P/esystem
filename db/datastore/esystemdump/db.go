package db

import (
	"database/sql"
	common "simplebank/db/common"
)

func New(db common.DBTX) *QueriesDump {
	return &QueriesDump{db: db}
}

type QueriesDump struct {
	db common.DBTX
}

func (q *QueriesDump) WithTx(tx *sql.Tx) *QueriesDump {
	return &QueriesDump{
		db: tx,
	}
}
