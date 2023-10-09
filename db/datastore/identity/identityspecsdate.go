package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createIdentitySpecsDate = `-- name: CreateIdentitySpecsDate: one
INSERT INTO Identity_Specs_Date 
  (IIID, Specs_ID, Value, Value2) 
VALUES 
  ($1, $2, $3, $4) 
ON CONFLICT( IIID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, IIID, Specs_Code, Specs_ID, Value, Value2
`

type IdentitySpecsDateRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	Iiid    int64     `json:"iIID"`
	SpecsId int64     `json:"specsId"`
	Value   time.Time `json:"value"`
	Value2  time.Time `json:"value2"`
}

func (q *QueriesIdentity) CreateIdentitySpecsDate(ctx context.Context, arg IdentitySpecsDateRequest) (model.IdentitySpecsDate, error) {
	row := q.db.QueryRowContext(ctx, createIdentitySpecsDate,
		arg.Iiid,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.IdentitySpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const updateIdentitySpecsDate = `-- name: UpdateIdentitySpecsDate :one
UPDATE Identity_Specs_Date SET 
	IIID = $2,
	Specs_ID = $3,
	Value = $4,
	Value2 = $5
WHERE uuid = $1
RETURNING UUID, IIID, Specs_Code, Specs_ID, Value, Value2
`

func (q *QueriesIdentity) UpdateIdentitySpecsDate(ctx context.Context, arg IdentitySpecsDateRequest) (model.IdentitySpecsDate, error) {
	row := q.db.QueryRowContext(ctx, updateIdentitySpecsDate,
		arg.Uuid,
		arg.Iiid,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.IdentitySpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteIdentitySpecsDate = `-- name: DeleteIdentitySpecsDate :exec
DELETE FROM Identity_Specs_Date
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteIdentitySpecsDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIdentitySpecsDate, uuid)
	return err
}

type IdentitySpecsDateInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	Iiid            int64        `json:"iIID"`
	SpecsCode       int64        `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           time.Time    `json:"value"`
	Value2          time.Time    `json:"value2"`
	ModCtr          int64        `json:"mod_ctr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateIdentitySpecDate(q *QueriesIdentity, ctx context.Context, sql string) (IdentitySpecsDateInfo, error) {
	var i IdentitySpecsDateInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.Value,
		&i.Value2,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecDate2(q *QueriesIdentity, ctx context.Context, sql string) ([]IdentitySpecsDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []IdentitySpecsDateInfo{}
	for rows.Next() {
		var i IdentitySpecsDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.Value,
			&i.Value2,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
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

const identitySpecsDateSQL = `-- name: identitySpecsDateSQL
SELECT 
  mr.UUID, d.IIID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Specs_Date d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesIdentity) GetIdentitySpecsDate(ctx context.Context, iiid int64, specsId int64) (IdentitySpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE d.IIID = %v and d.Specs_ID = %v",
		identitySpecsDateSQL, iiid, specsId))
}

func (q *QueriesIdentity) GetIdentitySpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (IdentitySpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", identitySpecsDateSQL, uuid))
}

type ListIdentitySpecsDateParams struct {
	Iiid   int64 `json:"Iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListIdentitySpecsDate(ctx context.Context, arg ListIdentitySpecsDateParams) ([]IdentitySpecsDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.IIID = %v LIMIT %d OFFSET %d",
			identitySpecsDateSQL, arg.Iiid, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.IIID = %v ", identitySpecsDateSQL, arg.Iiid)
	}
	return populateIdentitySpecDate2(q, ctx, sql)
}
