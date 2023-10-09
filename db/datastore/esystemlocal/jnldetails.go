package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

const createJnlDetails = `-- name: CreateJnlDetails: one
INSERT INTO JnlDetails (
	Jnld_Acnt_Cd, Jnld_Jnlh_Tran, Jnld_Entry, Jnld_Db_Amt, Jnld_Cr_Amt
) 
VALUES ($1, $2, $3, $4, $5) 
`

type JnlDetailsRequest struct {
	Acc    string              `json:"acc"`
	Trn    string              `json:"trn"`
	Series sql.NullInt64       `json:"series"`
	Debit  decimal.NullDecimal `json:"debit"`
	Credit decimal.NullDecimal `json:"credit"`
}

func (q *QueriesLocal) CreateJnlDetails(ctx context.Context, arg JnlDetailsRequest) error {
	_, err := q.db.ExecContext(ctx, createJnlDetails,
		arg.Acc,
		arg.Trn,
		arg.Series,
		arg.Debit,
		arg.Credit,
	)
	return err
}

const deleteJnlDetails = `-- name: DeleteJnlDetails :exec
DELETE FROM JnlDetails WHERE Jnld_Jnlh_Tran = $1 and Jnld_Acnt_Cd = $2
`

func (q *QueriesLocal) DeleteJnlDetails(ctx context.Context, trn string, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteJnlDetails, trn, acc)
	return err
}

type JnlDetailsInfo struct {
	ModCtr    int64               `json:"modCtr"`
	BrCode    string              `json:"brCode"`
	ModAction string              `json:"modAction"`
	Acc       string              `json:"acc"`
	Trn       string              `json:"trn"`
	Series    sql.NullInt64       `json:"series"`
	Debit     decimal.NullDecimal `json:"debit"`
	Credit    decimal.NullDecimal `json:"credit"`
}

// -- name: GetJnlDetails :one
const getJnlDetails = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Jnld_Acnt_Cd, Jnld_Jnlh_Tran,
  Jnld_Entry, Jnld_Db_Amt, Jnld_Cr_Amt
  FROM OrgParms,
    (SELECT Jnld_Acnt_Cd, Jnld_Jnlh_Tran, Max(Jnld_Entry) Jnld_Entry, Sum(Jnld_Db_Amt) Jnld_Db_Amt, Sum(Jnld_Cr_Amt) Jnld_Cr_Amt
     FROM JnlDetails
     GROUP BY Jnld_Acnt_Cd, Jnld_Jnlh_Tran
     ) d
  INNER JOIN
    (SELECT Max(ModCtr) ModCtr, ModAction, UniqueKeyString1, UniqueKeyString2, min(CASE WHEN Uploaded = 1 THEN 1 ELSE 0 END) Uploaded
     FROM Modified
     WHERE TableName = 'JnlDetails'
     GROUP BY ModAction, UniqueKeyString1, UniqueKeyString2 
     ) m on m.UniqueKeyString1 = Jnld_Jnlh_Tran and m.UniqueKeyString2 = d.Jnld_Acnt_Cd
`

func scanRowJnlDetails(row *sql.Row) (JnlDetailsInfo, error) {
	var i JnlDetailsInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.Acc,
		&i.Trn,
		&i.Series,
		&i.Debit,
		&i.Credit,
	)
	return i, err
}

func scanRowsJnlDetails(rows *sql.Rows) ([]JnlDetailsInfo, error) {
	items := []JnlDetailsInfo{}
	for rows.Next() {
		var i JnlDetailsInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.Trn,
			&i.Series,
			&i.Debit,
			&i.Credit,
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

func (q *QueriesLocal) GetJnlDetails(ctx context.Context, trn string, acc string) (JnlDetailsInfo, error) {
	sql := fmt.Sprintf("%s WHERE Uploaded = 0 and Jnld_Jnlh_Tran = $1 and Jnld_Acnt_Cd = $2", getJnlDetails)
	row := q.db.QueryRowContext(ctx, sql, trn, acc)
	return scanRowJnlDetails(row)
}

type ListJnlDetailsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) JnlDetailsCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Jnld_Acnt_Cd, Jnld_Jnlh_Tran,
  Jnld_Entry, Jnld_Db_Amt, Jnld_Cr_Amt
FROM OrgParms,
 (SELECT Jnld_Acnt_Cd, Jnld_Jnlh_Tran, Max(Jnld_Entry) Jnld_Entry, Sum(Jnld_Db_Amt) Jnld_Db_Amt, Sum(Jnld_Cr_Amt) Jnld_Cr_Amt
  FROM JnlDetails
  GROUP BY Jnld_Acnt_Cd, Jnld_Jnlh_Tran
  ) d
`, filenamePath)
}

func (q *QueriesLocal) ListJnlDetails(ctx context.Context) ([]JnlDetailsInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE Uploaded = 0`,
		getJnlDetails)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsJnlDetails(rows)
}

// -- name: UpdateJnlDetails :one
const updateJnlDetails = `
UPDATE JnlDetails SET 
	Jnld_Entry = $3,
	Jnld_Db_Amt = $4,
	Jnld_Cr_Amt = $5
WHERE Jnld_Jnlh_Tran = $1 and Jnld_Acnt_Cd = $2`

func (q *QueriesLocal) UpdateJnlDetails(ctx context.Context, arg JnlDetailsRequest) error {
	_, err := q.db.ExecContext(ctx, updateJnlDetails,
		arg.Acc,
		arg.Trn,
		arg.Series,
		arg.Debit,
		arg.Credit,
	)
	return err
}
