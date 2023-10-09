package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createUserSpecsRef = `-- name: CreateUserSpecsRef: one
INSERT INTO Users_Specs_Ref 
  (Users_Id, Specs_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Users_Id, Specs_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, Users_Id, Specs_Code, Specs_ID, Ref_Id
`

type UserSpecsRefRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	UserId  int64     `json:"userId"`
	SpecsId int64     `json:"specsId"`
	RefId   int64     `json:"refId"`
}

func (q *QueriesUser) CreateUserSpecsRef(ctx context.Context, arg UserSpecsRefRequest) (model.UserSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, createUserSpecsRef,
		arg.UserId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.UserSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const updateUserSpecsRef = `-- name: UpdateUserSpecsRef :one
UPDATE Users_Specs_Ref SET 
  Users_Id = $2,
  Specs_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, Users_Id, Specs_Code, Specs_ID, Ref_ID
`

func (q *QueriesUser) UpdateUserSpecsRef(ctx context.Context, arg UserSpecsRefRequest) (model.UserSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, updateUserSpecsRef,
		arg.Uuid,
		arg.UserId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.UserSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const deleteUserSpecsRef = `-- name: DeleteUserSpecsRef :exec
DELETE FROM Users_Specs_Ref
WHERE uuid = $1
`

func (q *QueriesUser) DeleteUserSpecsRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserSpecsRef, uuid)
	return err
}

type UserSpecsRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	UserId          int64          `json:"userId"`
	SpecsId         int64          `json:"specsId"`
	SpecsCode       string         `json:"specsCode"`
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

func populateUserSpecRef(q *QueriesUser, ctx context.Context, sql string) (UserSpecsRefInfo, error) {
	var i UserSpecsRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
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

func populateUserSpecRef2(q *QueriesUser, ctx context.Context, sql string) ([]UserSpecsRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []UserSpecsRefInfo{}
	for rows.Next() {
		var i UserSpecsRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.UserId,
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

const userSpecsRefSQL = `-- name: userSpecsRefSQL
SELECT 
  mr.UUID, d.Users_Id, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Users_Specs_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesUser) GetUserSpecsRef(ctx context.Context, userId int64, specsId int64) (UserSpecsRefInfo, error) {
	return populateUserSpecRef(q, ctx, fmt.Sprintf("%s WHERE d.Users_Id = %v and d.Specs_ID = %v",
		userSpecsRefSQL, userId, specsId))
}

func (q *QueriesUser) GetUserSpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (UserSpecsRefInfo, error) {
	return populateUserSpecRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", userSpecsRefSQL, uuid))
}

type ListUserSpecsRefParams struct {
	UserId int64 `json:"UserId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUserSpecsRef(ctx context.Context, arg ListUserSpecsRefParams) ([]UserSpecsRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Users_Id = %v LIMIT %d OFFSET %d",
			userSpecsRefSQL, arg.UserId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Users_Id = %v ", userSpecsRefSQL, arg.UserId)
	}
	return populateUserSpecRef2(q, ctx, sql)
}
