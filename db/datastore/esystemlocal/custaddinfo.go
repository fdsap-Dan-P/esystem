package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const createCustAddInfo = `-- name: CreateCustAddInfo: one
MERGE EditedCust as T
USING (SELECT $1 CID) as S ON S.CID = T.CID
WHEN NOT MATCHED 
  THEN INSERT (CID) VALUES (S.CID);

IF EXISTS (SELECT AreaCode FROM Area WHERE AreaCode = $1)
BEGIN
  UPDATE Area SET 
    Area = $2 
  WHERE AreaCode = $1
END ELSE BEGIN
  INSERT INTO CustAddInfo (
	CID, InfoDate, InfoCode, Info, InfoValue
  ) 
  VALUES ($1, $2, $3, $4, $5);
END
DELETE EditedCust WHERE CID = $1;

`

type CustAddInfoRequest struct {
	CID       int64     `json:"cID"`
	InfoDate  time.Time `json:"infoDate"`
	InfoCode  int64     `json:"infoCode"`
	Info      string    `json:"info"`
	InfoValue int64     `json:"infoValue"`
}

func (q *QueriesLocal) CreateCustAddInfo(ctx context.Context, arg CustAddInfoRequest) error {
	_, err := q.db.ExecContext(ctx, createCustAddInfo,
		arg.CID,
		arg.InfoDate,
		arg.InfoCode,
		arg.Info,
		arg.InfoValue,
	)
	return err
}

const deleteCustAddInfo = `-- name: DeleteCustAddInfo :exec
DELETE FROM CustAddInfo WHERE CID = $1 and InfoDate = $2 and InfoCode = $3
`

func (q *QueriesLocal) DeleteCustAddInfo(ctx context.Context, cid int64, infoDate time.Time, infoCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustAddInfo, cid, infoDate, infoCode)
	return err
}

type CustAddInfoInfo struct {
	ModCtr    int64     `json:"modCtr"`
	BrCode    string    `json:"brCode"`
	ModAction string    `json:"modAction"`
	CID       int64     `json:"cID"`
	InfoDate  time.Time `json:"infoDate"`
	InfoCode  int64     `json:"infoCode"`
	Info      string    `json:"info"`
	InfoValue int64     `json:"infoValue"`
}

// -- name: GetCustAddInfo :one
const getCustAddInfo = `
SELECT 
   m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, CID, InfoDate, InfoCode, Info, InfoValue
FROM OrgParms, CustAddInfo d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.CID and m.UniqueKeyInt2 = d.InfoCode and m.UniqueKeyDate = d.InfoDate 
`

func scanRowCustAddInfo(row *sql.Row) (CustAddInfoInfo, error) {
	var i CustAddInfoInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.CID,
		&i.InfoDate,
		&i.InfoCode,
		&i.Info,
		&i.InfoValue,
	)
	return i, err
}

func scanRowsCustAddInfo(rows *sql.Rows) ([]CustAddInfoInfo, error) {
	items := []CustAddInfoInfo{}
	for rows.Next() {
		var i CustAddInfoInfo
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

func (q *QueriesLocal) GetCustAddInfo(ctx context.Context, cid int64, infoDate time.Time, infoCode int64) (CustAddInfoInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'CustAddInfo' AND Uploaded = 0 and CID = $1 and InfoDate = $2 and InfoCode = $3", getCustAddInfo)
	row := q.db.QueryRowContext(ctx, sql, cid, infoDate, infoCode)
	return scanRowCustAddInfo(row)
}

type ListCustAddInfoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) CustAddInfoCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
	0 ModCtr, OrgParms.DefBranch_Code BrCode, CID, InfoDate, 
	InfoCode, Info, InfoValue
FROM OrgParms, CustAddInfo d
`, filenamePath)
}

func (q *QueriesLocal) ListCustAddInfo(ctx context.Context) ([]CustAddInfoInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'CustAddInfo' AND Uploaded = 0`,
		getCustAddInfo)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustAddInfo(rows)
}

// -- name: UpdateCustAddInfo :one
const updateCustAddInfo = `
INSERT INTO EditedCust(CID) VALUES($1);
UPDATE CustAddInfo SET 
	Info = $4,
	InfoValue = $5
WHERE CID = $1 and InfoDate = $2 and InfoCode = $3;
DELETE EditedCust WHERE CID = $1;
`

func (q *QueriesLocal) UpdateCustAddInfo(ctx context.Context, arg CustAddInfoRequest) error {
	_, err := q.db.ExecContext(ctx, updateCustAddInfo,
		arg.CID,
		arg.InfoDate,
		arg.InfoCode,
		arg.Info,
		arg.InfoValue,
	)
	return err
}
