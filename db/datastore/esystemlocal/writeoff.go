package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

const createWriteoff = `-- name: CreateWriteoff: one
INSERT INTO Writeoff (
	Acc, DisbDate, Principal, Interest, BalPrin, BalInt, TrnDate, AcctType, [Print], PostedBy, VerifiedBy
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
`

type WriteoffRequest struct {
	Acc        string          `json:"acc"`
	DisbDate   time.Time       `json:"disbDate"`
	Principal  decimal.Decimal `json:"principal"`
	Interest   decimal.Decimal `json:"interest"`
	BalPrin    decimal.Decimal `json:"balPrin"`
	BalInt     decimal.Decimal `json:"balInt"`
	TrnDate    time.Time       `json:"trnDate"`
	AcctType   string          `json:"acctType"`
	Print      sql.NullString  `json:"print"`
	PostedBy   sql.NullString  `json:"postedBy"`
	VerifiedBy sql.NullString  `json:"verifiedBy"`
}

func (q *QueriesLocal) CreateWriteoff(ctx context.Context, arg WriteoffRequest) error {
	_, err := q.db.ExecContext(ctx, createWriteoff,
		arg.Acc,
		arg.DisbDate,
		arg.Principal,
		arg.Interest,
		arg.BalPrin,
		arg.BalInt,
		arg.TrnDate,
		arg.AcctType,
		arg.Print,
		arg.PostedBy,
		arg.VerifiedBy,
	)
	return err
}

const deleteWriteoff = `-- name: DeleteWriteoff :exec
DELETE FROM Writeoff WHERE Acc = $1
`

func (q *QueriesLocal) DeleteWriteoff(ctx context.Context, Acc string) error {
	_, err := q.db.ExecContext(ctx, deleteWriteoff, Acc)
	return err
}

type WriteoffInfo struct {
	ModCtr     int64           `json:"modCtr"`
	BrCode     string          `json:"brCode"`
	ModAction  string          `json:"modAction"`
	Acc        string          `json:"acc"`
	DisbDate   time.Time       `json:"disbDate"`
	Principal  decimal.Decimal `json:"principal"`
	Interest   decimal.Decimal `json:"interest"`
	BalPrin    decimal.Decimal `json:"balPrin"`
	BalInt     decimal.Decimal `json:"balInt"`
	TrnDate    time.Time       `json:"trnDate"`
	AcctType   string          `json:"acctType"`
	Print      sql.NullString  `json:"print"`
	PostedBy   sql.NullString  `json:"postedBy"`
	VerifiedBy sql.NullString  `json:"verifiedBy"`
}

// -- name: GetWriteoff :one
const getWriteoff = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Acc, DisbDate, Principal, Interest, BalPrin, BalInt, TrnDate, AcctType, [Print], PostedBy, VerifiedBy
FROM OrgParms, Writeoff d
INNER JOIN Modified m on m.UniqueKeyString1 = d.Acc
`

func scanRowWriteoff(row *sql.Row) (WriteoffInfo, error) {
	var i WriteoffInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.Acc,
		&i.DisbDate,
		&i.Principal,
		&i.Interest,
		&i.BalPrin,
		&i.BalInt,
		&i.TrnDate,
		&i.AcctType,
		&i.Print,
		&i.PostedBy,
		&i.VerifiedBy,
	)
	return i, err
}

func scanRowsWriteoff(rows *sql.Rows) ([]WriteoffInfo, error) {
	items := []WriteoffInfo{}
	for rows.Next() {
		var i WriteoffInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.DisbDate,
			&i.Principal,
			&i.Interest,
			&i.BalPrin,
			&i.BalInt,
			&i.TrnDate,
			&i.AcctType,
			&i.Print,
			&i.PostedBy,
			&i.VerifiedBy,
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

func (q *QueriesLocal) GetWriteoff(ctx context.Context, Acc string) (WriteoffInfo, error) {
	sql := fmt.Sprintf("%s WHERE Uploaded = 0 and Acc = $1", getWriteoff)
	row := q.db.QueryRowContext(ctx, sql, Acc)
	return scanRowWriteoff(row)
}

type ListWriteoffParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) WriteoffCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Acc, DisbDate, Principal, 
  Interest, BalPrin, BalInt, TrnDate, AcctType, [Print], PostedBy, VerifiedBy
FROM OrgParms, Writeoff d
`, filenamePath)
}

func (q *QueriesLocal) ListWriteoff(ctx context.Context) ([]WriteoffInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'Writeoff' and Uploaded = 0`,
		getWriteoff)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsWriteoff(rows)
}

// -- name: UpdateWriteoff :one
const updateWriteoff = `
UPDATE Writeoff SET 
	DisbDate = $2,
	Principal = $3,
	Interest = $4,
	BalPrin = $5,
	BalInt = $6,
	TrnDate = $7,
	AcctType = $8,
	[Print] = $9,
	PostedBy = $10,
	VerifiedBy = $11
WHERE Acc = $1`

func (q *QueriesLocal) UpdateWriteoff(ctx context.Context, arg WriteoffRequest) error {
	_, err := q.db.ExecContext(ctx, updateWriteoff,
		arg.Acc,
		arg.DisbDate,
		arg.Principal,
		arg.Interest,
		arg.BalPrin,
		arg.BalInt,
		arg.TrnDate,
		arg.AcctType,
		arg.Print,
		arg.PostedBy,
		arg.VerifiedBy,
	)
	return err
}
