package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createCustAddInfoGroupNeed = `-- name: CreateCustAddInfoGroupNeed: one
INSERT INTO esystemdump.CustAddInfoGroupNeed(
   ModCtr, BrCode, ModAction, InfoGroup, InfoCode, InfoProcess )
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (brCode, infoGroup, infoCode, ModAction)
DO UPDATE SET
ModCtr =  EXCLUDED.ModCtr,
InfoProcess =  EXCLUDED.InfoProcess
`

func (q *QueriesDump) CreateCustAddInfoGroupNeed(ctx context.Context, arg model.CustAddInfoGroupNeed) error {
	_, err := q.db.ExecContext(ctx, createCustAddInfoGroupNeed,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.InfoGroup,
		arg.InfoCode,
		arg.InfoProcess,
	)
	return err
}

const deleteCustAddInfoGroupNeed = `-- name: DeleteCustAddInfoGroupNeed :exec
DELETE FROM esystemdump.CustAddInfoGroupNeed WHERE BrCode = $1 and InfoGroup = $2 and InfoCode = $3
`

func (q *QueriesDump) DeleteCustAddInfoGroupNeed(ctx context.Context, brCode string, infoGroup int64, infoCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustAddInfoGroupNeed, brCode, infoGroup, infoCode)
	return err
}

const getCustAddInfoGroupNeed = `-- name: GetCustAddInfoGroupNeed :one
SELECT
ModCtr, BrCode, ModAction, InfoGroup, InfoCode, InfoProcess
FROM esystemdump.CustAddInfoGroupNeed
`

func scanRowCustAddInfoGroupNeed(row *sql.Row) (model.CustAddInfoGroupNeed, error) {
	var i model.CustAddInfoGroupNeed
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.InfoGroup,
		&i.InfoCode,
		&i.InfoProcess,
	)
	return i, err
}

func scanRowsCustAddInfoGroupNeed(rows *sql.Rows) ([]model.CustAddInfoGroupNeed, error) {
	items := []model.CustAddInfoGroupNeed{}
	for rows.Next() {
		var i model.CustAddInfoGroupNeed
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.InfoGroup,
			&i.InfoCode,
			&i.InfoProcess,
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

func (q *QueriesDump) GetCustAddInfoGroupNeed(ctx context.Context, brCode string, infoGroup int64, infoCode int64) (model.CustAddInfoGroupNeed, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and InfoGroup = $2 and InfoCode = $3", getCustAddInfoGroupNeed)
	row := q.db.QueryRowContext(ctx, sql, brCode, infoGroup, infoCode)
	return scanRowCustAddInfoGroupNeed(row)
}

type ListCustAddInfoGroupNeedParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListCustAddInfoGroupNeed(ctx context.Context, lastModCtr int64) ([]model.CustAddInfoGroupNeed, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getCustAddInfoGroupNeed)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustAddInfoGroupNeed(rows)
}

const updateCustAddInfoGroupNeed = `-- name: UpdateCustAddInfoGroupNeed :one
UPDATE esystemdump.CustAddInfoGroupNeed SET 
ModCtr = $1,
InfoProcess = $6
WHERE BrCode = $2 and InfoGroup = $4 and InfoCode = $5 and ModAction = $3
`

func (q *QueriesDump) UpdateCustAddInfoGroupNeed(ctx context.Context, arg model.CustAddInfoGroupNeed) error {
	_, err := q.db.ExecContext(ctx, updateCustAddInfoGroupNeed,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.InfoGroup,
		arg.InfoCode,
		arg.InfoProcess,
	)
	return err
}
