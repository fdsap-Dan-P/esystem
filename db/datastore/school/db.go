package db

import (
	ds "simplebank/db/common"
)

func New(db ds.DBTX) *QueriesSchool {
	return &QueriesSchool{db: db}
}

type QueriesSchool struct {
	db ds.DBTX
	// DB  *sql.DB
	// Row *sql.Row
}
