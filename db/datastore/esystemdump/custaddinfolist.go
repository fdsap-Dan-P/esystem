package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createCustAddInfoList = `-- name: CreateCustAddInfoList: one
INSERT INTO esystemdump.CustAddInfoList(
   ModCtr, BrCode, ModAction, InfoCode, InfoOrder, Title, InfoType, InfoLen, InfoFormat, InputType, InfoSource )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (brCode, infoCode, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	InfoOrder =  EXCLUDED.InfoOrder,
	Title =  EXCLUDED.Title,
	InfoType =  EXCLUDED.InfoType,
	InfoLen =  EXCLUDED.InfoLen,
	InfoFormat =  EXCLUDED.InfoFormat,
	InputType =  EXCLUDED.InputType,
	InfoSource =  EXCLUDED.InfoSource
`

func (q *QueriesDump) CreateCustAddInfoList(ctx context.Context, arg model.CustAddInfoList) error {
	_, err := q.db.ExecContext(ctx, createCustAddInfoList,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.InfoCode,
		arg.InfoOrder,
		arg.Title,
		arg.InfoType,
		arg.InfoLen,
		arg.InfoFormat,
		arg.InputType,
		arg.InfoSource,
	)
	return err
}

const deleteCustAddInfoList = `-- name: DeleteCustAddInfoList :exec
DELETE FROM esystemdump.CustAddInfoList WHERE BrCode = $1 and InfoCode = $2
`

func (q *QueriesDump) DeleteCustAddInfoList(ctx context.Context, brCode string, infoCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustAddInfoList, brCode, infoCode)
	return err
}

const getCustAddInfoList = `-- name: GetCustAddInfoList :one
SELECT
ModCtr, BrCode, ModAction, InfoCode, InfoOrder, Title, InfoType, InfoLen, InfoFormat, InputType, InfoSource
FROM esystemdump.CustAddInfoList
`

func scanRowCustAddInfoList(row *sql.Row) (model.CustAddInfoList, error) {
	var i model.CustAddInfoList
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.InfoCode,
		&i.InfoOrder,
		&i.Title,
		&i.InfoType,
		&i.InfoLen,
		&i.InfoFormat,
		&i.InputType,
		&i.InfoSource,
	)
	return i, err
}

func scanRowsCustAddInfoList(rows *sql.Rows) ([]model.CustAddInfoList, error) {
	items := []model.CustAddInfoList{}
	for rows.Next() {
		var i model.CustAddInfoList
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.InfoCode,
			&i.InfoOrder,
			&i.Title,
			&i.InfoType,
			&i.InfoLen,
			&i.InfoFormat,
			&i.InputType,
			&i.InfoSource,
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

func (q *QueriesDump) GetCustAddInfoList(ctx context.Context, brCode string, infoCode int64) (model.CustAddInfoList, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and InfoCode = $2", getCustAddInfoList)
	row := q.db.QueryRowContext(ctx, sql, brCode, infoCode)
	return scanRowCustAddInfoList(row)
}

type ListCustAddInfoListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListCustAddInfoList(ctx context.Context, lastModCtr int64) ([]model.CustAddInfoList, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getCustAddInfoList)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustAddInfoList(rows)
}

const updateCustAddInfoList = `-- name: UpdateCustAddInfoList :one
UPDATE esystemdump.CustAddInfoList SET 
	ModCtr = $1,
	InfoOrder = $5,
	Title = $6,
	InfoType = $7,
	InfoLen = $8,
	InfoFormat = $9,
	InputType = $10,
	InfoSource = $11
WHERE BrCode = $2 and InfoCode = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateCustAddInfoList(ctx context.Context, arg model.CustAddInfoList) error {
	_, err := q.db.ExecContext(ctx, updateCustAddInfoList,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.InfoCode,
		arg.InfoOrder,
		arg.Title,
		arg.InfoType,
		arg.InfoLen,
		arg.InfoFormat,
		arg.InputType,
		arg.InfoSource,
	)
	return err
}
