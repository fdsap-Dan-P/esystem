package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

const createTrnMaster = `-- name: CreateTrnMaster: one
INSERT INTO TrnMaster (
	Acc, TrnDate, Trn, TrnType, orno, Prin, IntR, WaivedInt, RefNo, UserName, Particulars, TERMID, CANCEL
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, '', 0);
`

type TrnMasterRequest struct {
	Acc        string          `json:"acc"`
	TrnDate    time.Time       `json:"trnDate"`
	Trn        int64           `json:"trn"`
	TrnType    sql.NullInt64   `json:"trnType"`
	OrNo       sql.NullInt64   `json:"orNo"`
	Prin       decimal.Decimal `json:"prin"`
	IntR       decimal.Decimal `json:"intR"`
	WaivedInt  decimal.Decimal `json:"waivedInt"`
	RefNo      sql.NullString  `json:"refNo"`
	UserName   sql.NullString  `json:"userName"`
	Particular sql.NullString  `json:"particular"`
}

func (q *QueriesLocal) CreateTrnMaster(ctx context.Context, arg TrnMasterRequest) error {
	log.Printf("CreateTrnMaster: %v", arg)
	_, err := q.db.ExecContext(ctx, createTrnMaster,
		arg.Acc,
		arg.TrnDate,
		arg.Trn,
		arg.TrnType,
		arg.OrNo,
		arg.Prin,
		arg.IntR,
		arg.WaivedInt,
		arg.RefNo,
		arg.UserName,
		arg.Particular,
	)
	return err
}

const deleteTrnMaster = `-- name: DeleteTrnMaster :exec
DELETE FROM TrnMaster WHERE TrnDate = $1 and Trn = $2
`

func (q *QueriesLocal) DeleteTrnMaster(ctx context.Context, TrnDate time.Time, trn int64) error {
	_, err := q.db.ExecContext(ctx, deleteTrnMaster, TrnDate, trn)
	return err
}

type TrnMasterInfo struct {
	ModCtr     int64           `json:"modCtr"`
	BrCode     string          `json:"brCode"`
	ModAction  string          `json:"modAction"`
	Acc        string          `json:"acc"`
	TrnDate    time.Time       `json:"trnDate"`
	Trn        int64           `json:"trn"`
	TrnType    sql.NullInt64   `json:"trnType"`
	OrNo       sql.NullInt64   `json:"orNo"`
	Prin       decimal.Decimal `json:"prin"`
	IntR       decimal.Decimal `json:"intR"`
	WaivedInt  decimal.Decimal `json:"waivedInt"`
	RefNo      sql.NullString  `json:"refNo"`
	UserName   sql.NullString  `json:"userName"`
	Particular sql.NullString  `json:"particular"`
}

// -- name: GetTrnMaster :one
const getTrnMaster = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Acc, TrnDate, Trn, TrnType, 
  orno, Prin, IntR, WaivedInt, RefNo, UserName, Particulars
FROM OrgParms, TrnMaster d
INNER JOIN Modified m on m.UniqueKeyDate = d.TrnDate and m.UniqueKeyInt1 = d.Trn
`

func scanRowTrnMaster(row *sql.Row) (TrnMasterInfo, error) {
	var i TrnMasterInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.Acc,
		&i.TrnDate,
		&i.Trn,
		&i.TrnType,
		&i.OrNo,
		&i.Prin,
		&i.IntR,
		&i.WaivedInt,
		&i.RefNo,
		&i.UserName,
		&i.Particular,
	)
	return i, err
}

func scanRowsTrnMaster(rows *sql.Rows) ([]TrnMasterInfo, error) {
	items := []TrnMasterInfo{}
	for rows.Next() {
		var i TrnMasterInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.TrnDate,
			&i.Trn,
			&i.TrnType,
			&i.OrNo,
			&i.Prin,
			&i.IntR,
			&i.WaivedInt,
			&i.RefNo,
			&i.UserName,
			&i.Particular,
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

func (q *QueriesLocal) GetTrnMaster(ctx context.Context, trnDate time.Time, trn int64) (TrnMasterInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'TrnMaster' AND Uploaded = 0 and TrnDate = $1 and Trn = $2", getTrnMaster)
	row := q.db.QueryRowContext(ctx, sql, trnDate, trn)
	return scanRowTrnMaster(row)
}

type ListTrnMasterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) TrnMasterCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Acc, TrnDate, Trn, TrnType, 
  orno, Prin, IntR, WaivedInt, RefNo, UserName, Particulars
FROM OrgParms, TrnMaster d
`, filenamePath)
}

func (q *QueriesLocal) ListTrnMaster(ctx context.Context) ([]TrnMasterInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'TrnMaster' AND Uploaded = 0`,
		getTrnMaster)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsTrnMaster(rows)
}

// -- name: UpdateTrnMaster :one
const updateTrnMaster = `
UPDATE TrnMaster SET 
  Acc = $3,
  TrnType = $4,
  orno = $5,
  Prin = $6,
  IntR = $7,
  WaivedInt = $8,
  RefNo = $9,
  UserName = $10,
  Particulars = $11
WHERE TrnDate = $1 and Trn = $2`

func (q *QueriesLocal) UpdateTrnMaster(ctx context.Context, arg TrnMasterRequest) error {
	_, err := q.db.ExecContext(ctx, updateTrnMaster,
		arg.TrnDate,
		arg.Trn,
		arg.Acc,
		arg.TrnType,
		arg.OrNo,
		arg.Prin,
		arg.IntR,
		arg.WaivedInt,
		arg.RefNo,
		arg.UserName,
		arg.Particular,
	)
	return err
}
