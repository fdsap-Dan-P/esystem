package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createUnit = `-- name: CreateUnit: one
INSERT INTO Managers (
  ManCode, Unit, AreaCode, FName, Lname, MName, VatReg, UnitAddress
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
`

type UnitRequest struct {
	UnitCode    int64          `json:"unitCode"`
	Unit        sql.NullString `json:"unit"`
	AreaCode    sql.NullInt64  `json:"areaCode"`
	FName       sql.NullString `json:"fName"`
	LName       sql.NullString `json:"lName"`
	MName       sql.NullString `json:"mName"`
	VatReg      sql.NullString `json:"vatReg"`
	UnitAddress sql.NullString `json:"unitAddress"`
}

func (q *QueriesLocal) CreateUnit(ctx context.Context, arg UnitRequest) error {
	_, err := q.db.ExecContext(ctx, createUnit,
		arg.UnitCode,
		arg.Unit,
		arg.AreaCode,
		arg.FName,
		arg.LName,
		arg.MName,
		arg.VatReg,
		arg.UnitAddress,
	)
	return err
}

const deleteUnit = `-- name: DeleteUnit :exec
DELETE FROM Managers WHERE ManCode = $1
`

func (q *QueriesLocal) DeleteUnit(ctx context.Context, UnitCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteUnit, UnitCode)
	return err
}

type UnitInfo struct {
	ModCtr      int64          `json:"modCtr"`
	BrCode      string         `json:"brCode"`
	ModAction   string         `json:"modAction"`
	UnitCode    int64          `json:"unitCode"`
	Unit        sql.NullString `json:"unit"`
	AreaCode    sql.NullInt64  `json:"areaCode"`
	FName       sql.NullString `json:"fName"`
	LName       sql.NullString `json:"lName"`
	MName       sql.NullString `json:"mName"`
	VatReg      sql.NullString `json:"vatReg"`
	UnitAddress sql.NullString `json:"unitAddress"`
}

// -- name: GetUnit :one
const getUnit = `
SELECT 
m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, ManCode, Unit, AreaCode, FName, Lname, MName, VatReg, UnitAddress
FROM OrgParms, Managers d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.ManCode
`

func scanRowUnit(row *sql.Row) (UnitInfo, error) {
	var i UnitInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.UnitCode,
		&i.Unit,
		&i.AreaCode,
		&i.FName,
		&i.LName,
		&i.MName,
		&i.VatReg,
		&i.UnitAddress,
	)
	return i, err
}

func scanRowsUnit(rows *sql.Rows) ([]UnitInfo, error) {
	items := []UnitInfo{}
	for rows.Next() {
		var i UnitInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.UnitCode,
			&i.Unit,
			&i.AreaCode,
			&i.FName,
			&i.LName,
			&i.MName,
			&i.VatReg,
			&i.UnitAddress,
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

func (q *QueriesLocal) GetUnit(ctx context.Context, UnitCode int64) (UnitInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'Managers' AND Uploaded = 0 and ManCode = $1", getUnit)
	log.Println(sql)
	row := q.db.QueryRowContext(ctx, sql, UnitCode)
	return scanRowUnit(row)
}

type ListUnitParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) UnitCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, ManCode, Unit, AreaCode, FName, 
  Lname, MName, VatReg, UnitAddress
FROM OrgParms, Managers d
`, filenamePath)
}

func (q *QueriesLocal) ListUnit(ctx context.Context) ([]UnitInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'Managers' AND Uploaded = 0`,
		getUnit)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsUnit(rows)
}

// -- name: UpdateUnit :one
const updateUnit = `
UPDATE Managers SET 
  Unit = $2,
  AreaCode = $3,
  FName = $4,
  Lname = $5,
  MName = $6,
  VatReg = $7, 
  UnitAddress = $8
WHERE ManCode = $1`

func (q *QueriesLocal) UpdateUnit(ctx context.Context, arg UnitRequest) error {
	_, err := q.db.ExecContext(ctx, updateUnit,
		arg.UnitCode,
		arg.Unit,
		arg.AreaCode,
		arg.FName,
		arg.LName,
		arg.MName,
		arg.VatReg,
		arg.UnitAddress,
	)
	return err
}
