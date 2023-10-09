package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TableInfo struct {
	Catalog string `json:"catalog"`
	Schema  string `json:"schema"`
	Table   string `json:"table"`
}

// -- name: GetTable :one
const getTable = `
SELECT Table_catalog, Table_schema, Table_name FROM information_schema.tables
`

func scanRowTable(row *sql.Row) (TableInfo, error) {
	var i TableInfo
	err := row.Scan(
		&i.Catalog,
		&i.Schema,
		&i.Table,
	)
	return i, err
}

func scanRowsTable(rows *sql.Rows) ([]TableInfo, error) {
	items := []TableInfo{}
	for rows.Next() {
		var i TableInfo
		if err := rows.Scan(
			&i.Catalog,
			&i.Schema,
			&i.Table,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *QueriesLocal) GetTable(ctx context.Context, schema string) (TableInfo, error) {
	sql := fmt.Sprintf("%s WHERE Table_schema = $1", getTable)
	row := q.db.QueryRowContext(ctx, sql, schema)
	return scanRowTable(row)
}

type ListTableParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) ListTable(ctx context.Context, schema string) ([]TableInfo, error) {
	sql := fmt.Sprintf("%s WHERE Table_schema = $1", getTable)

	rows, err := q.db.QueryContext(ctx, sql, schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsTable(rows)
}
