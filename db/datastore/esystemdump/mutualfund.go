package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createMutualFund = `-- name: CreateMutualFund: one
INSERT INTO esystemdump.MutualFund(
   ModCtr, BrCode, ModAction, CID, OrNo, TrnDate, TrnType, TrnAmt, UserName )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (brCode, cID, orNo, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	TrnDate =  EXCLUDED.TrnDate,
	TrnType =  EXCLUDED.TrnType,
	TrnAmt =  EXCLUDED.TrnAmt,
	UserName =  EXCLUDED.UserName
`

func (q *QueriesDump) CreateMutualFund(ctx context.Context, arg model.MutualFund) error {
	_, err := q.db.ExecContext(ctx, createMutualFund,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.OrNo,
		arg.TrnDate,
		arg.TrnType,
		arg.TrnAmt,
		arg.UserName,
	)
	return err
}

const deleteMutualFund = `-- name: DeleteMutualFund :exec
DELETE FROM esystemdump.MutualFund WHERE BrCode = $1 and CID = $2 and OrNo = $3
`

func (q *QueriesDump) DeleteMutualFund(ctx context.Context, brCode string, cID int64, orNo int64) error {
	_, err := q.db.ExecContext(ctx, deleteMutualFund, brCode, cID, orNo)
	return err
}

const getMutualFund = `-- name: GetMutualFund :one
SELECT
ModCtr, BrCode, ModAction, CID, OrNo, TrnDate, TrnType, TrnAmt, UserName
FROM esystemdump.MutualFund
`

func scanRowMutualFund(row *sql.Row) (model.MutualFund, error) {
	var i model.MutualFund
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.CID,
		&i.OrNo,
		&i.TrnDate,
		&i.TrnType,
		&i.TrnAmt,
		&i.UserName,
	)
	return i, err
}

func scanRowsMutualFund(rows *sql.Rows) ([]model.MutualFund, error) {
	items := []model.MutualFund{}
	for rows.Next() {
		var i model.MutualFund
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.CID,
			&i.OrNo,
			&i.TrnDate,
			&i.TrnType,
			&i.TrnAmt,
			&i.UserName,
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

func (q *QueriesDump) GetMutualFund(ctx context.Context, brCode string, cID int64, orNo int64) (model.MutualFund, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and CID = $2 and OrNo = $3", getMutualFund)
	row := q.db.QueryRowContext(ctx, sql, brCode, cID, orNo)
	return scanRowMutualFund(row)
}

type ListMutualFundParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListMutualFund(ctx context.Context, lastModCtr int64) ([]model.MutualFund, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getMutualFund)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsMutualFund(rows)
}

const updateMutualFund = `-- name: UpdateMutualFund :one
UPDATE esystemdump.MutualFund SET 
	ModCtr = $1,
	TrnDate = $6,
	TrnType = $7,
	TrnAmt = $8,
	UserName = $9
WHERE BrCode = $2 and CID = $4 and OrNo = $5 and ModAction = $3
`

func (q *QueriesDump) UpdateMutualFund(ctx context.Context, arg model.MutualFund) error {
	_, err := q.db.ExecContext(ctx, updateMutualFund,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.OrNo,
		arg.TrnDate,
		arg.TrnType,
		arg.TrnAmt,
		arg.UserName,
	)
	return err
}
