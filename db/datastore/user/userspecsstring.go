package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createUserSpecsString = `-- name: CreateUserSpecsString: one
INSERT INTO Users_Specs_String 
  (Users_ID, Specs_ID, Value) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Users_ID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value
RETURNING 
  UUID, Users_ID, Specs_Code, Specs_ID, Value
`

type UserSpecsStringRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	UserId  int64     `json:"userId"`
	SpecsId int64     `json:"specsId"`
	Value   string    `json:"value"`
}

func (q *QueriesUser) CreateUserSpecsString(ctx context.Context, arg UserSpecsStringRequest) (model.UserSpecsString, error) {
	row := q.db.QueryRowContext(ctx, createUserSpecsString,
		arg.UserId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.UserSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const updateUserSpecsString = `-- name: UpdateUserSpecsString :one
UPDATE Users_Specs_String SET 
Users_ID = $2,
  Specs_ID = $3,
  Value = $4
WHERE uuid = $1
RETURNING UUID, Users_ID, Specs_Code, Specs_ID, Value
`

func (q *QueriesUser) UpdateUserSpecsString(ctx context.Context, arg UserSpecsStringRequest) (model.UserSpecsString, error) {
	row := q.db.QueryRowContext(ctx, updateUserSpecsString,
		arg.Uuid,
		arg.UserId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.UserSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const deleteUserSpecsString = `-- name: DeleteUserSpecsString :exec
DELETE FROM Users_Specs_String
WHERE uuid = $1
`

func (q *QueriesUser) DeleteUserSpecsString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserSpecsString, uuid)
	return err
}

type UserSpecsStringInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	UserId          int64        `json:"userId"`
	SpecsCode       string       `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           string       `json:"value"`
	ModCtr          int64        `json:"mod_ctr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateUserSpecString(q *QueriesUser, ctx context.Context, sql string) (UserSpecsStringInfo, error) {
	var i UserSpecsStringInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
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
	return i, err
}

func populateUserSpecString2(q *QueriesUser, ctx context.Context, sql string) ([]UserSpecsStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []UserSpecsStringInfo{}
	for rows.Next() {
		var i UserSpecsStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.UserId,
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

const userSpecsStringSQL = `-- name: userSpecsStringSQL
SELECT 
  mr.UUID, d.Users_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name Item_Short_Name, 
  ref.Remark Item_Description, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Users_Specs_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesUser) GetUserSpecsString(ctx context.Context, userId int64, specsId int64) (UserSpecsStringInfo, error) {
	return populateUserSpecString(q, ctx, fmt.Sprintf("%s WHERE d.Users_ID = %v and d.Specs_ID = %v",
		userSpecsStringSQL, userId, specsId))
}

func (q *QueriesUser) GetUserSpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (UserSpecsStringInfo, error) {
	return populateUserSpecString(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", userSpecsStringSQL, uuid))
}

type ListUserSpecsStringParams struct {
	UserId int64 `json:"UserId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUserSpecsString(ctx context.Context, arg ListUserSpecsStringParams) ([]UserSpecsStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Users_ID = %v LIMIT %d OFFSET %d",
			userSpecsStringSQL, arg.UserId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Users_ID = %v ", userSpecsStringSQL, arg.UserId)
	}
	return populateUserSpecString2(q, ctx, sql)
}
