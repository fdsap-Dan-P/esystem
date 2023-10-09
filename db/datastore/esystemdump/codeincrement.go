package db

import (
	"context"
)

func (q *QueriesDump) CodeIncrement(ctx context.Context, tableName string, keyCode string) (int64, error) {
	var trn int64
	sql := `
	DECLARE @Trn as Numeric
	exec usp_CodeIncrement $1, $2, @Trn Output
    SELECT  @Trn Trn
	`
	row := q.db.QueryRowContext(ctx, sql, tableName, keyCode)
	err := row.Scan(&trn)

	return trn, err
}
