package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createIdentitySpecsString = `-- name: CreateIdentitySpecsString: one
INSERT INTO Identity_Specs_String 
  (IIID, Specs_ID, Value) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( IIID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value
RETURNING 
  UUID, IIID, Specs_Code, Specs_ID, Value
`

type IdentitySpecsStringRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	Iiid    int64     `json:"iiid"`
	SpecsId int64     `json:"specsId"`
	Value   string    `json:"value"`
}

func (q *QueriesIdentity) CreateIdentitySpecsString(ctx context.Context, arg IdentitySpecsStringRequest) (model.IdentitySpecsString, error) {
	row := q.db.QueryRowContext(ctx, createIdentitySpecsString,
		arg.Iiid,
		arg.SpecsId,
		arg.Value,
	)
	var i model.IdentitySpecsString
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const updateIdentitySpecsString = `-- name: UpdateIdentitySpecsString :one
UPDATE Identity_Specs_String SET 
  IIID = $2,
  Specs_ID = $3,
  Value = $4
WHERE uuid = $1
RETURNING UUID, IIID, Specs_Code, Specs_ID, Value
`

func (q *QueriesIdentity) UpdateIdentitySpecsString(ctx context.Context, arg IdentitySpecsStringRequest) (model.IdentitySpecsString, error) {
	row := q.db.QueryRowContext(ctx, updateIdentitySpecsString,
		arg.Uuid,
		arg.Iiid,
		arg.SpecsId,
		arg.Value,
	)
	var i model.IdentitySpecsString
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const deleteIdentitySpecsString = `-- name: DeleteIdentitySpecsString :exec
DELETE FROM Identity_Specs_String
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteIdentitySpecsString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIdentitySpecsString, uuid)
	return err
}

type IdentitySpecsStringInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	Iiid            int64        `json:"iiid"`
	SpecsCode       int64        `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           string       `json:"value"`
	ModCtr          int64        `json:"mod_ctr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateIdentitySpecString(q *QueriesIdentity, ctx context.Context, sql string) (IdentitySpecsStringInfo, error) {
	var i IdentitySpecsStringInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.Value,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecString2(q *QueriesIdentity, ctx context.Context, sql string) ([]IdentitySpecsStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []IdentitySpecsStringInfo{}
	for rows.Next() {
		var i IdentitySpecsStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.Value,

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

const identitySpecsStringSQL = `-- name: identitySpecsStringSQL
SELECT 
  mr.UUID, d.Iiid, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name Item_Short_Name, 
  ref.Remark Item_Description, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Specs_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesIdentity) GetIdentitySpecsString(ctx context.Context, iiid int64, specsId int64) (IdentitySpecsStringInfo, error) {
	return populateIdentitySpecString(q, ctx, fmt.Sprintf("%s WHERE d.IIID = %v and d.Specs_ID = %v",
		identitySpecsStringSQL, iiid, specsId))
}

func (q *QueriesIdentity) GetIdentitySpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (IdentitySpecsStringInfo, error) {
	return populateIdentitySpecString(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", identitySpecsStringSQL, uuid))
}

type ListIdentitySpecsStringParams struct {
	Iiid   int64 `json:"Iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListIdentitySpecsString(ctx context.Context, arg ListIdentitySpecsStringParams) ([]IdentitySpecsStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Iiid = %v LIMIT %d OFFSET %d",
			identitySpecsStringSQL, arg.Iiid, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Iiid = %v ", identitySpecsStringSQL, arg.Iiid)
	}
	return populateIdentitySpecString2(q, ctx, sql)
}
