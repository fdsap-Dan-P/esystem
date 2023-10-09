package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

const createSaTrnMaster = `-- name: CreateSaTrnMaster: one
INSERT INTO SaTrnMaster (
	Acc, TrnDate, Trn, TrnType, orno, TrnAmt, RefNo, Particulars, UserName, TermId, PendApprove, TRNMNEM_CD
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 1)
`

type SaTrnMasterRequest struct {
	Acc         string              `json:"acc"`
	TrnDate     time.Time           `json:"trnDate"`
	Trn         int64               `json:"trn"`
	TrnType     sql.NullInt64       `json:"trnType"`
	OrNo        sql.NullInt64       `json:"orNo"`
	TrnAmt      decimal.NullDecimal `json:"trnAmt"`
	RefNo       sql.NullString      `json:"refNo"`
	Particular  string              `json:"particular"`
	TermId      string              `json:"termId"`
	UserName    string              `json:"userName"`
	PendApprove string              `json:"pendApprove"`
}

func (q *QueriesLocal) CreateSaTrnMaster(ctx context.Context, arg SaTrnMasterRequest) error {
	_, err := q.db.ExecContext(ctx, createSaTrnMaster,
		arg.Acc,
		arg.TrnDate,
		arg.Trn,
		arg.TrnType,
		arg.OrNo,
		arg.TrnAmt,
		arg.RefNo,
		arg.Particular,
		arg.UserName,
		arg.TermId,
		arg.PendApprove,
	)
	return err
}

const deleteSaTrnMaster = `-- name: DeleteSaTrnMaster :exec
DELETE FROM SaTrnMaster WHERE trnDate = $1 and trn = $2
`

func (q *QueriesLocal) DeleteSaTrnMaster(ctx context.Context, trnDate time.Time, trn int64) error {
	_, err := q.db.ExecContext(ctx, deleteSaTrnMaster, trnDate, trn)
	return err
}

type SaTrnMasterInfo struct {
	ModCtr      int64               `json:"modCtr"`
	BrCode      string              `json:"brCode"`
	ModAction   string              `json:"modAction"`
	Acc         string              `json:"acc"`
	TrnDate     time.Time           `json:"trnDate"`
	Trn         int64               `json:"trn"`
	TrnType     sql.NullInt64       `json:"trnType"`
	OrNo        sql.NullInt64       `json:"orNo"`
	TrnAmt      decimal.NullDecimal `json:"trnAmt"`
	RefNo       sql.NullString      `json:"refNo"`
	Particular  string              `json:"particular"`
	TermId      string              `json:"termId"`
	UserName    string              `json:"userName"`
	PendApprove string              `json:"pendApprove"`
}

// -- name: GetSaTrnMaster :one
const getSaTrnMaster = `
SELECT 
  Top 100 m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Acc, TrnDate, Trn, 
  TrnType, orno, TrnAmt, RefNo, Particulars, TermId, UserName, PendApprove
FROM OrgParms, SaTrnMaster d
INNER JOIN Modified m on m.UniqueKeyDate = d.trnDate and m.UniqueKeyInt1 = d.trn 
`

func scanRowSaTrnMaster(row *sql.Row) (SaTrnMasterInfo, error) {
	var i SaTrnMasterInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.Acc,
		&i.TrnDate,
		&i.Trn,
		&i.TrnType,
		&i.OrNo,
		&i.TrnAmt,
		&i.RefNo,
		&i.Particular,
		&i.TermId,
		&i.UserName,
		&i.PendApprove,
	)
	return i, err
}

func scanRowsSaTrnMaster(rows *sql.Rows) ([]SaTrnMasterInfo, error) {
	items := []SaTrnMasterInfo{}
	for rows.Next() {
		var i SaTrnMasterInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.TrnDate,
			&i.Trn,
			&i.TrnType,
			&i.OrNo,
			&i.TrnAmt,
			&i.RefNo,
			&i.Particular,
			&i.TermId,
			&i.UserName,
			&i.PendApprove,
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

func (q *QueriesLocal) GetSaTrnMaster(ctx context.Context, trnDate time.Time, trn int64) (SaTrnMasterInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'SaTrnMaster' AND Uploaded = 0 and trnDate = $1 and trn = $2", getSaTrnMaster)
	row := q.db.QueryRowContext(ctx, sql, trnDate, trn)
	return scanRowSaTrnMaster(row)
}

type ListSaTrnMasterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) SaTrnMasterCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Acc, TrnDate, Trn, TrnType, orno, TrnAmt, RefNo, Particulars, TermId, UserName, PendApprove
FROM OrgParms, SaTrnMaster d
`, filenamePath)
}

func (q *QueriesLocal) ListSaTrnMaster(ctx context.Context) ([]SaTrnMasterInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'SaTrnMaster' AND Uploaded = 0`,
		getSaTrnMaster)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsSaTrnMaster(rows)
}

// -- name: UpdateSaTrnMaster :one
const updateSaTrnMaster = `
UPDATE SaTrnMaster SET 
  Acc = $3,
  TrnType = $4,
  orno = $5,
  TrnAmt = $6,
  RefNo = $7,
  Particulars = $8,
  TermID = $8,
  UserName = $9,
  PendApprove = $10
WHERE trnDate = $1 and trn = $2`

func (q *QueriesLocal) UpdateSaTrnMaster(ctx context.Context, arg SaTrnMasterRequest) error {
	_, err := q.db.ExecContext(ctx, updateSaTrnMaster,
		arg.TrnDate,
		arg.Trn,
		arg.Acc,
		arg.TrnType,
		arg.OrNo,
		arg.TrnAmt,
		arg.RefNo,
		arg.Particular,
		arg.UserName,
		arg.PendApprove,
	)
	return err
}
