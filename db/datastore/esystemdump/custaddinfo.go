package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	model "simplebank/db/datastore/esystemlocal"
)

const createCustAddInfo = `-- name: CreateCustAddInfo: one
INSERT INTO esystemdump.CustAddInfo(
   ModCtr, BrCode, ModAction, CID, InfoDate, InfoCode, Info, InfoValue)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (brCode, cID, infoDate, infoCode, ModAction)
DO UPDATE SET
  ModCtr =  EXCLUDED.ModCtr,
  Info =  EXCLUDED.Info,
  InfoValue =  EXCLUDED.InfoValue
`

func (q *QueriesDump) CreateCustAddInfo(ctx context.Context, arg model.CustAddInfo) error {
	_, err := q.db.ExecContext(ctx, createCustAddInfo,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.InfoDate,
		arg.InfoCode,
		arg.Info,
		arg.InfoValue,
	)
	return err
}

func (q *QueriesDump) BulkInsertCustAddInfo(ctx context.Context, rows []model.CustAddInfo) error {
	colCtr := 8
	valueStrings := make([]string, 0, len(rows))
	valueArgs := make([]interface{}, 0, len(rows)*colCtr)
	i := 0
	for _, post := range rows {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*colCtr+1, i*colCtr+2, i*colCtr+3, i*colCtr+4, i*colCtr+5, i*colCtr+6, i*colCtr+7, i*colCtr+8))
		valueArgs = append(valueArgs, post.ModCtr)
		valueArgs = append(valueArgs, post.BrCode)
		valueArgs = append(valueArgs, post.ModAction)
		valueArgs = append(valueArgs, post.CID)
		valueArgs = append(valueArgs, post.InfoDate)
		valueArgs = append(valueArgs, post.InfoCode)
		valueArgs = append(valueArgs, post.Info)
		valueArgs = append(valueArgs, post.InfoValue)
		i++
	}
	stmt := fmt.Sprintf(`
	INSERT INTO esystemdump.CustAddInfo(
		ModCtr, BrCode, ModAction, CID, InfoDate, InfoCode, Info, InfoValue )
	 VALUES %s
	 ON CONFLICT (brCode, cID, infoDate, infoCode, ModAction)
	 DO UPDATE SET
	   ModCtr =  EXCLUDED.ModCtr,
	   Info =  EXCLUDED.Info,
	   InfoValue =  EXCLUDED.InfoValue`, strings.Join(valueStrings, ","))

	// log.Println(stmt)
	_, err := q.db.ExecContext(ctx, stmt, valueArgs...)

	return err
}

const deleteCustAddInfo = `-- name: DeleteCustAddInfo :exec
DELETE FROM esystemdump.CustAddInfo WHERE BrCode = $1 and CID = $2 and InfoDate = $3 and InfoCode = $4
`

func (q *QueriesDump) DeleteCustAddInfo(ctx context.Context, brCode string, cID int64, infoDate time.Time, infoCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustAddInfo, brCode, cID, infoDate, infoCode)
	return err
}

const getCustAddInfo = `-- name: GetCustAddInfo :one
SELECT
ModCtr, BrCode, ModAction, CID, InfoDate, InfoCode, Info, InfoValue
FROM esystemdump.CustAddInfo
`

func scanRowCustAddInfo(row *sql.Row) (model.CustAddInfo, error) {
	var i model.CustAddInfo
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.CID,
		&i.InfoDate,
		&i.InfoCode,
		&i.Info,
		&i.InfoValue,
	)
	return i, err
}

func scanRowsCustAddInfo(rows *sql.Rows) ([]model.CustAddInfo, error) {
	items := []model.CustAddInfo{}
	for rows.Next() {
		var i model.CustAddInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.CID,
			&i.InfoDate,
			&i.InfoCode,
			&i.Info,
			&i.InfoValue,
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

func (q *QueriesDump) GetCustAddInfo(ctx context.Context, brCode string, cID int64, infoDate time.Time, infoCode int64) (model.CustAddInfo, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and CID = $2 and InfoDate = $3 and InfoCode = $4", getCustAddInfo)
	row := q.db.QueryRowContext(ctx, sql, brCode, cID, infoDate, infoCode)
	return scanRowCustAddInfo(row)
}

type ListCustAddInfoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListCustAddInfo(ctx context.Context, lastModCtr int64) ([]model.CustAddInfo, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getCustAddInfo)
	// log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustAddInfo(rows)
}

const updateCustAddInfo = `-- name: UpdateCustAddInfo :one
UPDATE esystemdump.CustAddInfo SET 
  ModCtr = $1,
  Info = $7,
  InfoValue = $8
WHERE BrCode = $2 and CID = $4 and InfoDate = $5 and InfoCode = $6 and ModAction = $3
`

func (q *QueriesDump) UpdateCustAddInfo(ctx context.Context, arg model.CustAddInfo) error {
	_, err := q.db.ExecContext(ctx, updateCustAddInfo,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.InfoDate,
		arg.InfoCode,
		arg.Info,
		arg.InfoValue,
	)
	return err
}
