package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createIdentitySpecsRef = `-- name: CreateIdentitySpecsRef: one
INSERT INTO Identity_Specs_Ref 
  (IIID, Specs_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( IIID, Specs_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, IIID, Specs_Code, Specs_ID, Ref_Id
`

type IdentitySpecsRefRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	Iiid    int64     `json:"iiid"`
	SpecsId int64     `json:"specsId"`
	RefId   int64     `json:"refId"`
}

func (q *QueriesIdentity) CreateIdentitySpecsRef(ctx context.Context, arg IdentitySpecsRefRequest) (model.IdentitySpecsRef, error) {
	row := q.db.QueryRowContext(ctx, createIdentitySpecsRef,
		arg.Iiid,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.IdentitySpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const updateIdentitySpecsRef = `-- name: UpdateIdentitySpecsRef :one
UPDATE Identity_Specs_Ref SET 
  IIID = $2,
  Specs_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, IIID, Specs_Code, Specs_ID, Ref_ID
`

func (q *QueriesIdentity) UpdateIdentitySpecsRef(ctx context.Context, arg IdentitySpecsRefRequest) (model.IdentitySpecsRef, error) {
	row := q.db.QueryRowContext(ctx, updateIdentitySpecsRef,
		arg.Uuid,
		arg.Iiid,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.IdentitySpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const deleteIdentitySpecsRef = `-- name: DeleteIdentitySpecsRef :exec
DELETE FROM Identity_Specs_Ref
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteIdentitySpecsRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIdentitySpecsRef, uuid)
	return err
}

type IdentitySpecsRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	Iiid            int64          `json:"iiid"`
	SpecsCode       int64          `json:"specsCode"`
	SpecsId         int64          `json:"specsId"`
	Item            string         `json:"item"`
	ItemShortName   string         `json:"itemShortName"`
	ItemDescription string         `json:"itemDescription"`
	RefId           int64          `json:"refId"`
	MeasureId       sql.NullInt64  `json:"measureId"`
	Measure         sql.NullString `json:"measure"`
	MeasureUnit     sql.NullString `json:"measureUnit"`
	ModCtr          int64          `json:"mod_ctr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

func populateIdentitySpecRef(q *QueriesIdentity, ctx context.Context, sql string) (IdentitySpecsRefInfo, error) {
	var i IdentitySpecsRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.RefId,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecRef2(q *QueriesIdentity, ctx context.Context, sql string) ([]IdentitySpecsRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []IdentitySpecsRefInfo{}
	for rows.Next() {
		var i IdentitySpecsRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.RefId,

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

const identitySpecsRefSQL = `-- name: identitySpecsRefSQL
SELECT 
  mr.UUID, d.Iiid, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Specs_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesIdentity) GetIdentitySpecsRef(ctx context.Context, iiid int64, specsId int64) (IdentitySpecsRefInfo, error) {
	return populateIdentitySpecRef(q, ctx, fmt.Sprintf("%s WHERE d.IIID = %v and d.Specs_ID = %v",
		identitySpecsRefSQL, iiid, specsId))
}

func (q *QueriesIdentity) GetIdentitySpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (IdentitySpecsRefInfo, error) {
	return populateIdentitySpecRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", identitySpecsRefSQL, uuid))
}

type ListIdentitySpecsRefParams struct {
	Iiid   int64 `json:"Iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListIdentitySpecsRef(ctx context.Context, arg ListIdentitySpecsRefParams) ([]IdentitySpecsRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Iiid = %v LIMIT %d OFFSET %d",
			identitySpecsRefSQL, arg.Iiid, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.IIID = %v ", identitySpecsRefSQL, arg.Iiid)
	}
	return populateIdentitySpecRef2(q, ctx, sql)
}
