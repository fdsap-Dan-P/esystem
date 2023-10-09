package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccessRole = `-- name: CreateAccessRole: one
INSERT INTO Access_Role (
Access_Name, Description, Other_Info
) VALUES ($1, $2, $3) 
RETURNING Id, UUId, Access_Name, Description, Other_Info
`

type AccessRoleRequest struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	AccessName  string         `json:"accessName"`
	Description sql.NullString `json:"description"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccess) CreateAccessRole(ctx context.Context, arg AccessRoleRequest) (model.AccessRole, error) {
	row := q.db.QueryRowContext(ctx, createAccessRole,
		arg.AccessName,
		arg.Description,
		arg.OtherInfo,
	)
	var i model.AccessRole
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccessName,
		&i.Description,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccessRole = `-- name: DeleteAccessRole :exec
DELETE FROM Access_Role
WHERE id = $1
`

func (q *QueriesAccess) DeleteAccessRole(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccessRole, id)
	return err
}

type AccessRoleInfo struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	AccessName  string         `json:"accessName"`
	Description sql.NullString `json:"description"`
	OtherInfo   sql.NullString `json:"otherInfo"`
	ModCtr      int64          `json:"modCtr"`
	Created     sql.NullTime   `json:"created"`
	Updated     sql.NullTime   `json:"updated"`
}

const accessRoleSQL = `-- name: accessRoleSQL :one
SELECT 
Id, mr.UUId, Access_Name, Description, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Role d INNER JOIN Main_Record mr on mr.UUID = d.UUID`

func populateAccessRole(q *QueriesAccess, ctx context.Context, sql string) (AccessRoleInfo, error) {
	var i AccessRoleInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccessName,
		&i.Description,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccessRoles(q *QueriesAccess, ctx context.Context, sql string) ([]AccessRoleInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccessRoleInfo{}
	for rows.Next() {
		var i AccessRoleInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.AccessName,
			&i.Description,
			&i.OtherInfo,
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

func (q *QueriesAccess) GetAccessRole(ctx context.Context, id int64) (AccessRoleInfo, error) {
	return populateAccessRole(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", accessRoleSQL, id))
}

func (q *QueriesAccess) GetAccessRolebyUuId(ctx context.Context, uuid uuid.UUID) (AccessRoleInfo, error) {
	return populateAccessRole(q, ctx, fmt.Sprintf("%s WHERE mr.Uuid = '%v'", accessRoleSQL, uuid))
}

func (q *QueriesAccess) GetAccessRolebyName(ctx context.Context, name string) (AccessRoleInfo, error) {
	return populateAccessRole(q, ctx, fmt.Sprintf("%s WHERE Access_Name = '%v' LIMIT 1", accessRoleSQL, name))
}

type ListAccessRoleParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccess) ListAccessRole(ctx context.Context, arg ListAccessRoleParams) ([]AccessRoleInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s LIMIT %v OFFSET %v", accessRoleSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(accessRoleSQL)
	}
	return populateAccessRoles(q, ctx, sql)
}

const updateAccessRole = `-- name: UpdateAccessRole :one
UPDATE Access_Role SET 
Access_Name = $2,
Description = $3,
Other_Info = $4
WHERE id = $1
RETURNING Id, UUId, Access_Name, Description, Other_Info
`

func (q *QueriesAccess) UpdateAccessRole(ctx context.Context, arg AccessRoleRequest) (model.AccessRole, error) {
	row := q.db.QueryRowContext(ctx, updateAccessRole,
		arg.Id,
		arg.AccessName,
		arg.Description,
		arg.OtherInfo,
	)
	var i model.AccessRole
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccessName,
		&i.Description,
		&i.OtherInfo,
	)
	return i, err
}
