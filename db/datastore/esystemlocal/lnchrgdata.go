package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

const createLnChrgData = `-- name: CreateLnChrgData: one
INSERT INTO LnChrgData (
	Acc, ChrgCode, RefAcc, ChrAmnt
) 
VALUES ($1, $2, $3, $4) 
`

type LnChrgDataRequest struct {
	Acc      string          `json:"acc"`
	ChrgCode int64           `json:"chrgCode"`
	RefAcc   sql.NullString  `json:"refAcc"`
	ChrAmnt  decimal.Decimal `json:"chrAmnt"`
}

func (q *QueriesLocal) CreateLnChrgData(ctx context.Context, arg LnChrgDataRequest) error {
	_, err := q.db.ExecContext(ctx, createLnChrgData,
		arg.Acc,
		arg.ChrgCode,
		arg.RefAcc,
		arg.ChrAmnt,
	)
	return err
}

const deleteLnChrgData = `-- name: DeleteLnChrgData :exec
DELETE FROM LnChrgData WHERE  Acc = $1 and ChrgCode = $2 and isNull(RefAcc,'') = $3
`

func (q *QueriesLocal) DeleteLnChrgData(ctx context.Context, acc string, chrgCode int64, refAcc string) error {
	_, err := q.db.ExecContext(ctx, deleteLnChrgData, acc, chrgCode, refAcc)
	return err
}

type LnChrgDataInfo struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	Acc       string          `json:"acc"`
	ChrgCode  int64           `json:"chrgCode"`
	RefAcc    sql.NullString  `json:"refAcc"`
	ChrAmnt   decimal.Decimal `json:"chrAmnt"`
}

// -- name: GetLnChrgData :one
const getLnChrgData = `
SELECT 
  ModCtr, OrgParms.DefBranch_Code BrCode, ModAction, Acc, ChrgCode, RefAcc, ChrAmnt
FROM OrgParms, 
 (SELECT Acc, ChrgCode, isnull(RefAcc,'') RefAcc, sum(ChrAmnt) ChrAmnt
  FROM LnChrgData
  GROUP BY Acc, ChrgCode, isnull(RefAcc,'')
  ) d
INNER JOIN 
 (SELECT Max(ModCtr) ModCtr, ModAction, UniqueKeyInt1, UniqueKeyString1, UniqueKeyString2,
    min(CASE WHEN Uploaded = 1 THEN 1 ELSE 0 END) Uploaded
  FROM Modified
  WHERE TableName = 'LnChrgData'
  GROUP BY ModAction, UniqueKeyInt1, UniqueKeyString1, UniqueKeyString2
  ) m on m.UniqueKeyString1 = d.Acc 
  and m.UniqueKeyInt1 = d.ChrgCode and isNull(m.UniqueKeyString2,'') = isNull(d.RefAcc,'')  
`

func scanRowLnChrgData(row *sql.Row) (LnChrgDataInfo, error) {
	var i LnChrgDataInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.Acc,
		&i.ChrgCode,
		&i.RefAcc,
		&i.ChrAmnt)
	return i, err
}

func scanRowsLnChrgData(rows *sql.Rows) ([]LnChrgDataInfo, error) {
	items := []LnChrgDataInfo{}
	for rows.Next() {
		var i LnChrgDataInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.ChrgCode,
			&i.RefAcc,
			&i.ChrAmnt,
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

func (q *QueriesLocal) GetLnChrgData(ctx context.Context, acc string, code int64, refAcc string) (LnChrgDataInfo, error) {
	sql := fmt.Sprintf("%s WHERE Uploaded = 0 and Acc = $1 and ChrgCode = $2 and isNull(RefAcc,'') = $3", getLnChrgData)
	log.Printf("GetLnChrgData: %v acc:%v, code:%v, refAcc:%v", sql, acc, code, refAcc)
	row := q.db.QueryRowContext(ctx, sql, acc, code, refAcc)
	return scanRowLnChrgData(row)
}

type ListLnChrgDataParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) LnChrgDataCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Acc, ChrgCode, RefAcc, ChrAmnt
FROM OrgParms, 
 (SELECT Acc, ChrgCode, isnull(RefAcc,'') RefAcc, sum(ChrAmnt) ChrAmnt
  FROM LnChrgData
  GROUP BY Acc, ChrgCode, isnull(RefAcc,'')
  ) d
`, filenamePath)
}

func (q *QueriesLocal) ListLnChrgData(ctx context.Context) ([]LnChrgDataInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE Uploaded = 0`,
		getLnChrgData)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLnChrgData(rows)
}

// -- name: UpdateLnChrgData :one
const updateLnChrgData = `
UPDATE LnChrgData SET 
  ChrAmnt = $4
WHERE Acc = $1 and ChrgCode = $2 and isNull(RefAcc,'') = $3`

func (q *QueriesLocal) UpdateLnChrgData(ctx context.Context, arg LnChrgDataRequest) error {
	_, err := q.db.ExecContext(ctx, updateLnChrgData,
		arg.Acc,
		arg.ChrgCode,
		arg.RefAcc,
		arg.ChrAmnt,
	)
	return err
}
