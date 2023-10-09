package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createJnlHeaders = `-- name: CreateJnlHeaders: one
INSERT INTO esystemdump.JnlHeaders(
   ModCtr, BrCode, ModAction, Trn, TrnDate, Particulars, UserName, Code )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (brCode, trn, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	TrnDate =  EXCLUDED.TrnDate,
	Particulars =  EXCLUDED.Particulars,
	UserName =  EXCLUDED.UserName,
	Code =  EXCLUDED.Code
`

func (q *QueriesDump) CreateJnlHeaders(ctx context.Context, arg model.JnlHeaders) error {
	_, err := q.db.ExecContext(ctx, createJnlHeaders,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Trn,
		arg.TrnDate,
		arg.Particulars,
		arg.UserName,
		arg.Code,
	)
	return err
}

const deleteJnlHeaders = `-- name: DeleteJnlHeaders :exec
DELETE FROM esystemdump.JnlHeaders WHERE BrCode = $1 and Trn = $2
`

func (q *QueriesDump) DeleteJnlHeaders(ctx context.Context, brCode string, trn string) error {
	_, err := q.db.ExecContext(ctx, deleteJnlHeaders, brCode, trn)
	return err
}

const getJnlHeaders = `-- name: GetJnlHeaders :one
SELECT
ModCtr, BrCode, ModAction, Trn, TrnDate, Particulars, UserName, Code
FROM esystemdump.JnlHeaders
`

func scanRowJnlHeaders(row *sql.Row) (model.JnlHeaders, error) {
	var i model.JnlHeaders
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Trn,
		&i.TrnDate,
		&i.Particulars,
		&i.UserName,
		&i.Code,
	)
	return i, err
}

func scanRowsJnlHeaders(rows *sql.Rows) ([]model.JnlHeaders, error) {
	items := []model.JnlHeaders{}
	for rows.Next() {
		var i model.JnlHeaders
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Trn,
			&i.TrnDate,
			&i.Particulars,
			&i.UserName,
			&i.Code,
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

func (q *QueriesDump) GetJnlHeaders(ctx context.Context, brCode string, trn string) (model.JnlHeaders, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Trn = $2", getJnlHeaders)
	row := q.db.QueryRowContext(ctx, sql, brCode, trn)
	return scanRowJnlHeaders(row)
}

type ListJnlHeadersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListJnlHeaders(ctx context.Context, lastModCtr int64) ([]model.JnlHeaders, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getJnlHeaders)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsJnlHeaders(rows)
}

const updateJnlHeaders = `-- name: UpdateJnlHeaders :one
UPDATE esystemdump.JnlHeaders SET 
	ModCtr = $1,
	TrnDate = $5,
	Particulars = $6,
	UserName = $7,
	Code = $8
WHERE BrCode = $2 and Trn = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateJnlHeaders(ctx context.Context, arg model.JnlHeaders) error {
	_, err := q.db.ExecContext(ctx, updateJnlHeaders,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Trn,
		arg.TrnDate,
		arg.Particulars,
		arg.UserName,
		arg.Code,
	)
	return err
}
