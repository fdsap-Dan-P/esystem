package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccessConfigRef = `-- name: CreateAccessConfigRef: one
INSERT INTO Access_Config_Ref 
  (Role_Id, Config_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Role_Id, Config_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, Role_Id, Config_Code, Config_ID, Ref_Id
`

type AccessConfigRefRequest struct {
	Uuid     uuid.UUID `json:"uuid"`
	RoleId   int64     `json:"roleId"`
	ConfigId int64     `json:"configId"`
	RefId    int64     `json:"refId"`
}

func (q *QueriesAccess) CreateAccessConfigRef(ctx context.Context, arg AccessConfigRefRequest) (model.AccessConfigRef, error) {
	row := q.db.QueryRowContext(ctx, createAccessConfigRef,
		arg.RoleId,
		arg.ConfigId,
		arg.RefId,
	)
	var i model.AccessConfigRef
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.RefId,
	)
	return i, err
}

const updateAccessConfigRef = `-- name: UpdateAccessConfigRef :one
UPDATE Access_Config_Ref SET 
  Role_Id = $2,
  Config_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, Role_Id, Config_Code, Config_ID, Ref_ID
`

func (q *QueriesAccess) UpdateAccessConfigRef(ctx context.Context, arg AccessConfigRefRequest) (model.AccessConfigRef, error) {
	row := q.db.QueryRowContext(ctx, updateAccessConfigRef,
		arg.Uuid,
		arg.RoleId,
		arg.ConfigId,
		arg.RefId,
	)
	var i model.AccessConfigRef
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.RefId,
	)
	return i, err
}

const deleteAccessConfigRef = `-- name: DeleteAccessConfigRef :exec
DELETE FROM Access_Config_Ref
WHERE uuid = $1
`

func (q *QueriesAccess) DeleteAccessConfigRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccessConfigRef, uuid)
	return err
}

type AccessConfigRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	RoleId          int64          `json:"roleId"`
	ConfigId        int64          `json:"configId"`
	ConfigCode      string         `json:"configCode"`
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

func populateaccessConfigRef(q *QueriesAccess, ctx context.Context, sql string) (AccessConfigRefInfo, error) {
	var i AccessConfigRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
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

func populateaccessConfigRef2(q *QueriesAccess, ctx context.Context, sql string) ([]AccessConfigRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccessConfigRefInfo{}
	for rows.Next() {
		var i AccessConfigRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.RoleId,
			&i.ConfigCode,
			&i.ConfigId,
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

const accessConfigRefSQL = `-- name: accessConfigRefSQL
SELECT 
  mr.UUID, d.Role_Id, d.Config_Code, d.Config_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Config_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Config_ID
`

func (q *QueriesAccess) GetAccessConfigRef(ctx context.Context, roleId int64, configId int64) (AccessConfigRefInfo, error) {
	return populateaccessConfigRef(q, ctx, fmt.Sprintf("%s WHERE d.Role_Id = %v and d.Config_ID = %v",
		accessConfigRefSQL, roleId, configId))
}

func (q *QueriesAccess) GetAccessConfigRefbyUuid(ctx context.Context, uuid uuid.UUID) (AccessConfigRefInfo, error) {
	return populateaccessConfigRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", accessConfigRefSQL, uuid))
}

type ListAccessConfigRefParams struct {
	RoleId int64 `json:"RoleId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccess) ListAccessConfigRef(ctx context.Context, arg ListAccessConfigRefParams) ([]AccessConfigRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Role_Id = %v LIMIT %d OFFSET %d",
			accessConfigRefSQL, arg.RoleId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Role_Id = %v ", accessConfigRefSQL, arg.RoleId)
	}
	return populateaccessConfigRef2(q, ctx, sql)
}
