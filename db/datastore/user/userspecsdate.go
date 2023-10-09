package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createUserSpecsDate = `-- name: CreateUserSpecsDate: one
INSERT INTO Users_Specs_Date 
  (Users_Id, Specs_ID, Value, Value2) 
VALUES 
  ($1, $2, $3, $4) 
ON CONFLICT( Users_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Users_Id, Specs_Code, Specs_ID, Value, Value2
`

type UserSpecsDateRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	UserId  int64     `json:"userId"`
	SpecsId int64     `json:"specsId"`
	Value   time.Time `json:"value"`
	Value2  time.Time `json:"value2"`
}

func (q *QueriesUser) CreateUserSpecsDate(ctx context.Context, arg UserSpecsDateRequest) (model.UserSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, createUserSpecsDate,
		arg.UserId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.UserSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const updateUserSpecsDate = `-- name: UpdateUserSpecsDate :one
UPDATE Users_Specs_Date SET 
    Users_Id = $2,
	Specs_ID = $3,
	Value = $4,
	Value2 = $5
WHERE uuid = $1
RETURNING UUID, Users_Id, Specs_Code, Specs_ID, Value, Value2
`

func (q *QueriesUser) UpdateUserSpecsDate(ctx context.Context, arg UserSpecsDateRequest) (model.UserSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, updateUserSpecsDate,
		arg.Uuid,
		arg.UserId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.UserSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteUserSpecsDate = `-- name: DeleteUserSpecsDate :exec
DELETE FROM Users_Specs_Date
WHERE uuid = $1
`

func (q *QueriesUser) DeleteUserSpecsDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserSpecsDate, uuid)
	return err
}

type UserSpecsDateInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	UserId          int64        `json:"userId"`
	SpecsCode       string       `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           time.Time    `json:"value"`
	Value2          time.Time    `json:"value2"`
	ModCtr          int64        `json:"modCtr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateIdentitySpecDate(q *QueriesUser, ctx context.Context, sql string) (UserSpecsDateInfo, error) {
	var i UserSpecsDateInfo
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
		&i.Value2,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecDate2(q *QueriesUser, ctx context.Context, sql string) ([]UserSpecsDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []UserSpecsDateInfo{}
	for rows.Next() {
		var i UserSpecsDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.UserId,
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

const UserSpecsDateSQL = `-- name: UserSpecsDateSQL
SELECT 
  mr.UUID, d.Users_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Users_Specs_Date d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesUser) GetUserSpecsDate(ctx context.Context, userId int64, specsId int64) (UserSpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE d.Users_ID = %v and d.Specs_ID = %v",
		UserSpecsDateSQL, userId, specsId))
}

func (q *QueriesUser) GetUserSpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (UserSpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", UserSpecsDateSQL, uuid))
}

type ListUserSpecsDateParams struct {
	UserId int64 `json:"UserId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUserSpecsDate(ctx context.Context, arg ListUserSpecsDateParams) ([]UserSpecsDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Users_ID = %v LIMIT %d OFFSET %d",
			UserSpecsDateSQL, arg.UserId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Users_ID = %v ", UserSpecsDateSQL, arg.UserId)
	}
	return populateIdentitySpecDate2(q, ctx, sql)
}
