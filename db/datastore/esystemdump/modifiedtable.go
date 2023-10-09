package db

import (
	"context"
	"database/sql"
)

type ModifiedTable struct {
	BrCode     string `json:"brCode"`
	DumpTable  string `json:"dumpTable"`
	LastModCtr int64  `json:"lastModCtr"`
}

const getModifiedTable = `-- name: GetModifiedTable:one
SELECT
  BrCode, DumpTable, LastModCtr
FROM esystemdump.ModifiedTable
WHERE BrCode = $1
`

func scanRowsModifiedTable(rows *sql.Rows) ([]ModifiedTable, error) {
	items := []ModifiedTable{}
	for rows.Next() {
		var i ModifiedTable
		if err := rows.Scan(
			&i.BrCode,
			&i.DumpTable,
			&i.LastModCtr,
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

func (q *QueriesDump) ListModifiedTable(ctx context.Context, brCode string) ([]ModifiedTable, error) {
	rows, err := q.db.QueryContext(ctx, getModifiedTable, brCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsModifiedTable(rows)
}
