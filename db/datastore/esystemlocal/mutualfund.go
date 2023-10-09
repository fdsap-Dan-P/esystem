package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

const createMutualFund = `-- name: CreateMutualFund: one
INSERT INTO Mutual_Fund (
	MFCID, MFORNUMBER, MFDATE, MFTEMPFIELD, MFAmount, MFUID
) 
VALUES ($1, $2, $3, $4, $5, $6)
`

type MutualFundRequest struct {
	CID      int64           `json:"CID"`
	OrNo     sql.NullInt64   `json:"orNo"`
	TrnDate  time.Time       `json:"trnDate"`
	TrnType  sql.NullString  `json:"trnType"`
	TrnAmt   decimal.Decimal `json:"trnAmt"`
	UserName sql.NullString  `json:"userName"`
}

func (q *QueriesLocal) CreateMutualFund(ctx context.Context, arg MutualFundRequest) error {
	_, err := q.db.ExecContext(ctx, createMutualFund,
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
DELETE FROM Mutual_Fund WHERE MFCID = $1 and MFORNumber = $2
`

func (q *QueriesLocal) DeleteMutualFund(ctx context.Context, cid int64, orno int64, trnDate time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteMutualFund, cid, orno)
	return err
}

type MutualFundInfo struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	CID       int64           `json:"CID"`
	OrNo      sql.NullInt64   `json:"orNo"`
	TrnDate   time.Time       `json:"trnDate"`
	TrnType   sql.NullString  `json:"trnType"`
	TrnAmt    decimal.Decimal `json:"trnAmt"`
	UserName  sql.NullString  `json:"userName"`
}

// -- name: GetMutualFund :one
const getMutualFund = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, MFCID, MFORNUMBER, MFDATE, MFTEMPFIELD, MFAmount, MFUID
FROM OrgParms, 
 (SELECT MFCID, MFORNUMBER, MFDATE, Max(MFTEMPFIELD) MFTEMPFIELD, Sum(MFAmount) MFAmount, Max(MFUID) MFUID 
  FROM Mutual_Fund d
  GROUP BY MFCID, MFORNUMBER, MFDATE) d
INNER JOIN Modified m 
   on m.UniqueKeyInt1 = d.MFCID 
  and m.UniqueKeyInt2 = d.MFORNumber 
  and m.UniqueKeyDate = d.MFDATE
`

func scanRowMutualFund(row *sql.Row) (MutualFundInfo, error) {
	var i MutualFundInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.CID,
		&i.OrNo,
		&i.TrnDate,
		&i.TrnType,
		&i.TrnAmt,
		&i.UserName,
	)
	return i, err
}

func scanRowsMutualFund(rows *sql.Rows) ([]MutualFundInfo, error) {
	items := []MutualFundInfo{}
	for rows.Next() {
		var i MutualFundInfo
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

func (q *QueriesLocal) GetMutualFund(ctx context.Context, cid int64, orno int64, trnDate time.Time) (MutualFundInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'Mutual_Fund' AND Uploaded = 0 and MFCID = $1 and MFORNumber = $2 and MFDate = $3", getMutualFund)
	row := q.db.QueryRowContext(ctx, sql, cid, orno, trnDate)
	return scanRowMutualFund(row)
}

type ListMutualFundParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) MutualFundCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, MFCID, MFORNUMBER, 
  MFDATE, MFTEMPFIELD, MFAmount, MFUID
FROM OrgParms, 
 (SELECT MFCID, MFORNUMBER, MFDATE, Max(MFTEMPFIELD) MFTEMPFIELD, Sum(MFAmount) MFAmount, Max(MFUID) MFUID 
  FROM Mutual_Fund d
  GROUP BY MFCID, MFORNUMBER, MFDATE) d
`, filenamePath)
}

func (q *QueriesLocal) ListMutualFund(ctx context.Context) ([]MutualFundInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'Mutual_Fund' AND Uploaded = 0`,
		getMutualFund)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsMutualFund(rows)
}

// -- name: UpdateMutualFund :one
const updateMutualFund = `
UPDATE Mutual_Fund SET 
	MFTEMPFIELD = $4,
	MFAmount = $5,
	MFUID = $6
WHERE MFCID = $1 and MFORNumber = $2 and MFDATE = $3`

func (q *QueriesLocal) UpdateMutualFund(ctx context.Context, arg MutualFundRequest) error {
	_, err := q.db.ExecContext(ctx, updateMutualFund,
		arg.CID,
		arg.OrNo,
		arg.TrnDate,
		arg.TrnType,
		arg.TrnAmt,
		arg.UserName,
	)
	return err
}
