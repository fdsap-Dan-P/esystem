package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createCustAddInfoList = `-- name: CreateCustAddInfoList: one
INSERT INTO CustAddInfoList (
	InfoCode, InfoOrder, Title, InfoType, InfoLen, InfoFormat, InputType, InfoSource
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
`

type CustAddInfoListRequest struct {
	InfoCode   int64  `json:"infoCode"`
	InfoOrder  string `json:"infoOrder"`
	Title      string `json:"title"`
	InfoType   string `json:"infoType"`
	InfoLen    int64  `json:"infoLen"`
	InfoFormat string `json:"infoFormat"`
	InputType  int64  `json:"inputType"`
	InfoSource string `json:"infoSource"`
}

func (q *QueriesLocal) CreateCustAddInfoList(ctx context.Context, arg CustAddInfoListRequest) error {
	_, err := q.db.ExecContext(ctx, createCustAddInfoList,
		arg.InfoCode,
		arg.InfoOrder,
		arg.Title,
		arg.InfoType,
		arg.InfoLen,
		arg.InfoFormat,
		arg.InputType,
		arg.InfoSource,
	)
	return err
}

const deleteCustAddInfoList = `-- name: DeleteCustAddInfoList :exec
DELETE FROM CustAddInfoList WHERE InfoCode = $1
`

func (q *QueriesLocal) DeleteCustAddInfoList(ctx context.Context, InfoCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustAddInfoList, InfoCode)
	return err
}

type CustAddInfoListInfo struct {
	ModCtr     int64  `json:"modCtr"`
	BrCode     string `json:"brCode"`
	ModAction  string `json:"modAction"`
	InfoCode   int64  `json:"infoCode"`
	InfoOrder  string `json:"infoOrder"`
	Title      string `json:"title"`
	InfoType   string `json:"infoType"`
	InfoLen    int64  `json:"infoLen"`
	InfoFormat string `json:"infoFormat"`
	InputType  int64  `json:"inputType"`
	InfoSource string `json:"infoSource"`
}

// -- name: GetCustAddInfoList :one
const getCustAddInfoList = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, InfoCode, InfoOrder, Title, InfoType, InfoLen, InfoFormat, InputType, InfoSource
FROM OrgParms, CustAddInfoList d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.InfoCode
`

func scanRowCustAddInfoList(row *sql.Row) (CustAddInfoListInfo, error) {
	var i CustAddInfoListInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.InfoCode,
		&i.InfoOrder,
		&i.Title,
		&i.InfoType,
		&i.InfoLen,
		&i.InfoFormat,
		&i.InputType,
		&i.InfoSource,
	)
	return i, err
}

func scanRowsCustAddInfoList(rows *sql.Rows) ([]CustAddInfoListInfo, error) {
	items := []CustAddInfoListInfo{}
	for rows.Next() {
		var i CustAddInfoListInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.InfoCode,
			&i.InfoOrder,
			&i.Title,
			&i.InfoType,
			&i.InfoLen,
			&i.InfoFormat,
			&i.InputType,
			&i.InfoSource,
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

func (q *QueriesLocal) GetCustAddInfoList(ctx context.Context, InfoCode int64) (CustAddInfoListInfo, error) {
	sql := fmt.Sprintf("%s WHERE  m.TableName = 'CustAddInfoList' AND Uploaded = 0 and InfoCode = $1", getCustAddInfoList)
	row := q.db.QueryRowContext(ctx, sql, InfoCode)
	return scanRowCustAddInfoList(row)
}

type ListCustAddInfoListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) CustAddInfoListCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, InfoCode, InfoOrder, Title, 
  InfoType, InfoLen, InfoFormat, InputType, InfoSource
FROM OrgParms, CustAddInfoList d
`, filenamePath)
}

func (q *QueriesLocal) ListCustAddInfoList(ctx context.Context) ([]CustAddInfoListInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'CustAddInfoList' AND Uploaded = 0`,
		getCustAddInfoList)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustAddInfoList(rows)
}

// -- name: UpdateCustAddInfoList :one
const updateCustAddInfoList = `
UPDATE CustAddInfoList SET 
	InfoOrder = $2,
	Title = $3,
	InfoType = $4,
	InfoLen = $5,
	InfoFormat = $6,
	InputType = $7,
	InfoSource = $8
WHERE InfoCode = $1`

func (q *QueriesLocal) UpdateCustAddInfoList(ctx context.Context, arg CustAddInfoListRequest) error {
	_, err := q.db.ExecContext(ctx, updateCustAddInfoList,
		arg.InfoCode,
		arg.InfoOrder,
		arg.Title,
		arg.InfoType,
		arg.InfoLen,
		arg.InfoFormat,
		arg.InputType,
		arg.InfoSource,
	)
	return err
}
