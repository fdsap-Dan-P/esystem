package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

const createSaMaster = `-- name: CreateSaMaster: one
INSERT INTO SaMaster (
	Acc, CID, Type, Balance, doLastTrn, DoStatus, Dopen, DoMaturity, Status, DOINTEFF, CLASSIFICATION_CODE, CLASSIFICATION_TYPE
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 0, 1, 1) 
`

type SaMasterRequest struct {
	Acc        string              `json:"acc"`
	CID        int64               `json:"cID"`
	Type       int64               `json:"type"`
	Balance    decimal.NullDecimal `json:"balance"`
	DoLastTrn  sql.NullTime        `json:"doLastTrn"`
	DoStatus   sql.NullTime        `json:"doStatus"`
	Dopen      sql.NullTime        `json:"dopen"`
	DoMaturity sql.NullTime        `json:"doMaturity"`
	Status     sql.NullString      `json:"status"`
}

func (q *QueriesLocal) CreateSaMaster(ctx context.Context, arg SaMasterRequest) error {
	_, err := q.db.ExecContext(ctx, createSaMaster,
		arg.Acc,
		arg.CID,
		arg.Type,
		arg.Balance,
		arg.DoLastTrn,
		arg.DoStatus,
		arg.Dopen,
		arg.DoMaturity,
		arg.Status,
	)
	return err
}

const deleteSaMaster = `-- name: DeleteSaMaster :exec
DELETE FROM SaMaster WHERE Acc = $1
`

func (q *QueriesLocal) DeleteSaMaster(ctx context.Context, Acc string) error {
	_, err := q.db.ExecContext(ctx, deleteSaMaster, Acc)
	return err
}

type SaMasterInfo struct {
	ModCtr     int64               `json:"modCtr"`
	BrCode     string              `json:"brCode"`
	ModAction  string              `json:"modAction"`
	Acc        string              `json:"acc"`
	CID        int64               `json:"cID"`
	Type       int64               `json:"type"`
	Balance    decimal.NullDecimal `json:"balance"`
	DoLastTrn  sql.NullTime        `json:"doLastTrn"`
	DoStatus   sql.NullTime        `json:"doStatus"`
	Dopen      sql.NullTime        `json:"dopen"`
	DoMaturity sql.NullTime        `json:"doMaturity"`
	Status     sql.NullString      `json:"status"`
}

// -- name: GetSaMaster :one
const getSaMaster = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Acc, CID, Type, Balance, doLastTrn, DoStatus, Dopen, DoMaturity, Status
FROM OrgParms, SaMaster d
INNER JOIN Modified m on m.UniqueKeyString1 = d.Acc 
`

func scanRowSaMaster(row *sql.Row) (SaMasterInfo, error) {
	var i SaMasterInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.Acc,
		&i.CID,
		&i.Type,
		&i.Balance,
		&i.DoLastTrn,
		&i.DoStatus,
		&i.Dopen,
		&i.DoMaturity,
		&i.Status,
	)
	return i, err
}

func scanRowsSaMaster(rows *sql.Rows) ([]SaMasterInfo, error) {
	items := []SaMasterInfo{}
	for rows.Next() {
		var i SaMasterInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.CID,
			&i.Type,
			&i.Balance,
			&i.DoLastTrn,
			&i.DoStatus,
			&i.Dopen,
			&i.DoMaturity,
			&i.Status,
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

func (q *QueriesLocal) GetSaMaster(ctx context.Context, Acc string) (SaMasterInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'SaMaster' AND Uploaded = 0 and Acc = $1", getSaMaster)
	row := q.db.QueryRowContext(ctx, sql, Acc)
	return scanRowSaMaster(row)
}

type ListSaMasterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) SaMasterCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Acc, CID, Type, Balance, doLastTrn, DoStatus, Dopen, DoMaturity, Status
FROM OrgParms, SaMaster d
`, filenamePath)
}

func (q *QueriesLocal) ListSaMaster(ctx context.Context) ([]SaMasterInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'SaMaster' AND Uploaded = 0`,
		getSaMaster)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsSaMaster(rows)
}

// -- name: UpdateSaMaster :one
const updateSaMaster = `
UPDATE SaMaster SET 
  Acc = $1,
  CID = $2,
  Type = $3,
  Balance = $4,
  doLastTrn = $5,
  DoStatus = $6,
  Dopen = $7,
  DoMaturity = $8,
  Status = $9
WHERE Acc = $1`

func (q *QueriesLocal) UpdateSaMaster(ctx context.Context, arg SaMasterRequest) error {
	_, err := q.db.ExecContext(ctx, updateSaMaster,
		arg.Acc,
		arg.CID,
		arg.Type,
		arg.Balance,
		arg.DoLastTrn,
		arg.DoStatus,
		arg.Dopen,
		arg.DoMaturity,
		arg.Status,
	)
	return err
}
