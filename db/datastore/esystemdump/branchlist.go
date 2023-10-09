package db

import (
	"context"
	"database/sql"
	"fmt"

	"time"
)

type BranchListLight struct {
	BrCode         string    `json:"brCode"`
	EbSysDate      time.Time `json:"ebSysDate"`
	RunState       int64     `json:"runState"`
	LastConnection time.Time `json:"lastConnection"`
}

type BranchList struct {
	BrCode         string    `json:"brCode"`
	EbSysDate      time.Time `json:"ebSysDate"`
	RunState       int64     `json:"runState"`
	OrgAddress     string    `json:"orgAddress"`
	TaxInfo        string    `json:"taxInfo"`
	DefCity        string    `json:"defCity"`
	DefProvince    string    `json:"defProvince"`
	DefCountry     string    `json:"defCountry"`
	DefZip         string    `json:"defZip"`
	WaivableInt    bool      `json:"waivableInt"`
	DBVersion      string    `json:"dBVersion"`
	ESystemVer     []byte    `json:"eSystemVer"`
	NewBrCode      int64     `json:"newBrCode"`
	LastConnection time.Time `json:"lastConnection"`
}

const createBranchList = `-- name: CreateBranchList: one
INSERT INTO esystemdump.BranchList(
   BrCode, EbSysDate, RunState, OrgAddress, TaxInfo, DefCity, DefProvince, DefCountry, DefZip, WaivableInt, 
   DBVersion, ESystemVer, NewBrCode, LastConnection )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, CURRENT_TIMESTAMP)
ON CONFLICT (brCode)
DO UPDATE SET
	EbSysDate =  EXCLUDED.EbSysDate,
	RunState =  EXCLUDED.RunState,
	OrgAddress =  EXCLUDED.OrgAddress,
	TaxInfo =  EXCLUDED.TaxInfo,
	DefCity =  EXCLUDED.DefCity,
	DefProvince =  EXCLUDED.DefProvince,
	DefCountry =  EXCLUDED.DefCountry,
	DefZip =  EXCLUDED.DefZip,
	WaivableInt =  EXCLUDED.WaivableInt,
	DBVersion =  EXCLUDED.DBVersion,
	ESystemVer =  EXCLUDED.ESystemVer,
	NewBrCode =  EXCLUDED.NewBrCode,
	LastConnection =  EXCLUDED.LastConnection
`

func (q *QueriesDump) CreateBranchList(ctx context.Context, arg BranchList) error {
	_, err := q.db.ExecContext(ctx, createBranchList,
		arg.BrCode,
		arg.EbSysDate,
		arg.RunState,
		arg.OrgAddress,
		arg.TaxInfo,
		arg.DefCity,
		arg.DefProvince,
		arg.DefCountry,
		arg.DefZip,
		arg.WaivableInt,
		arg.DBVersion,
		arg.ESystemVer,
		arg.NewBrCode,
	)
	return err
}

const deleteBranchList = `-- name: DeleteBranchList :exec
DELETE FROM esystemdump.BranchList WHERE BrCode = $1
`

func (q *QueriesDump) DeleteBranchList(ctx context.Context, brCode string) error {
	_, err := q.db.ExecContext(ctx, deleteBranchList, brCode)
	return err
}

const getBranchList = `-- name: GetBranchList :one
SELECT
BrCode, EbSysDate, RunState, OrgAddress, TaxInfo, DefCity, DefProvince, DefCountry, DefZip, WaivableInt, DBVersion, ESystemVer, NewBrCode, LastConnection
FROM esystemdump.BranchList
`

func scanRowBranchList(row *sql.Row) (BranchList, error) {
	var i BranchList
	err := row.Scan(
		&i.BrCode,
		&i.EbSysDate,
		&i.RunState,
		&i.OrgAddress,
		&i.TaxInfo,
		&i.DefCity,
		&i.DefProvince,
		&i.DefCountry,
		&i.DefZip,
		&i.WaivableInt,
		&i.DBVersion,
		&i.ESystemVer,
		&i.NewBrCode,
		&i.LastConnection,
	)
	return i, err
}

func scanRowsBranchList(rows *sql.Rows) ([]BranchList, error) {
	items := []BranchList{}
	for rows.Next() {
		var i BranchList
		if err := rows.Scan(
			&i.BrCode,
			&i.EbSysDate,
			&i.RunState,
			&i.OrgAddress,
			&i.TaxInfo,
			&i.DefCity,
			&i.DefProvince,
			&i.DefCountry,
			&i.DefZip,
			&i.WaivableInt,
			&i.DBVersion,
			&i.ESystemVer,
			&i.NewBrCode,
			&i.LastConnection,
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

func (q *QueriesDump) GetBranchList(ctx context.Context, brCode string) (BranchList, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1", getBranchList)
	row := q.db.QueryRowContext(ctx, sql, brCode)
	return scanRowBranchList(row)
}

type ListBranchListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListBranchList(ctx context.Context) ([]BranchList, error) {
	rows, err := q.db.QueryContext(ctx, getBranchList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsBranchList(rows)
}

const updateBranchList = `-- name: UpdateBranchList :one
UPDATE esystemdump.BranchList SET 
  EbSysDate = $2,
  RunState = $3,
  LastConnection = CURRENT_TIMESTAMP
WHERE BrCode = $1
`

func (q *QueriesDump) UpdateBranchList(ctx context.Context, arg BranchListLight) error {
	_, err := q.db.ExecContext(ctx, updateBranchList,
		arg.BrCode,
		arg.EbSysDate,
		arg.RunState,
	)
	return err
}
