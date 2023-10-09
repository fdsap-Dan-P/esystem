package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/util"
)

type ModifiedTable struct {
	LocalTableName string `json:"localTableName"`
	ModCtr         int64  `json:"modCtr"`
	// MaxModCtr      int64  `json:"maxModCtr"`
}

const getModifiedTable = `-- name: GetModifiedTable:one
SELECT
  TableName LocalTableName, Max(ModCtr) ModCtr
FROM Modified WITH(NOLOCK)
WHERE Uploaded = 0
GROUP BY TableName
`

func scanRowsModifiedTables(rows *sql.Rows) ([]ModifiedTable, error) {
	items := []ModifiedTable{}
	for rows.Next() {
		var i ModifiedTable
		if err := rows.Scan(
			&i.LocalTableName,
			// &i.MinModCtr,
			&i.ModCtr,
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

func (q *QueriesLocal) ListModifiedTable(ctx context.Context) ([]ModifiedTable, error) {
	rows, err := q.db.QueryContext(ctx, getModifiedTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsModifiedTables(rows)
}

type ModifiedTableInfo struct {
	ModCtr           int64          `json:"modCtr"`
	TableName        string         `json:"localTableName"`
	UniqueKeyInt1    sql.NullInt64  `json:"uniqueKeyInt1"`
	UniqueKeyInt2    sql.NullInt64  `json:"uniqueKeyInt2"`
	UniqueKeyInt3    sql.NullInt64  `json:"uniqueKeyInt3"`
	UniqueKeyDate    sql.NullTime   `json:"uniqueKeyDate"`
	UniqueKeyString1 sql.NullString `json:"uniqueKeyString1"`
	UniqueKeyString2 sql.NullString `json:"uniqueKeyString2"`
	Uploaded         bool           `json:"uploaded"`
	// MaxModCtr      int64  `json:"maxModCtr"`
}

func scanRowsModifiedTable(row *sql.Row) (ModifiedTableInfo, error) {
	var i ModifiedTableInfo
	err := row.Scan(
		&i.ModCtr,
		&i.TableName,
		&i.UniqueKeyInt1,
		&i.UniqueKeyInt2,
		&i.UniqueKeyInt3,
		&i.UniqueKeyDate,
		&i.UniqueKeyString1,
		&i.UniqueKeyString2,
		&i.Uploaded,
	)
	return i, err
}

func (q *QueriesLocal) GetModifiedTable(ctx context.Context, modCtr int64) (ModifiedTableInfo, error) {
	sql := `-- name: GetModifiedTable:one
	SELECT
	  ModCtr, TableName, UniqueKeyInt1, UniqueKeyInt2, UniqueKeyInt3, 
	  UniqueKeyDate, UniqueKeyString1, UniqueKeyString2, Uploaded
	FROM Modified WITH(NOLOCK)
	WHERE modCtr = $1`
	row := q.db.QueryRowContext(ctx, sql, modCtr)
	return scanRowsModifiedTable(row)
}

func (q *QueriesLocal) UpdateModifiedTableUploaded(ctx context.Context, modCtrList []int64, updated bool) error {
	_, err := q.db.ExecContext(ctx,
		fmt.Sprintf(`UPDATE Modified SET Uploaded = $1 WHERE ModCtr in (%v)`, util.Int64List2Comma(modCtrList)), updated)
	return err
}
