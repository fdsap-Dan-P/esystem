package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createUserSpecsNumber = `-- name: CreateUserSpecsNumber: one
INSERT INTO Users_Specs_Number 
  (Users_Id, Specs_ID, Value, Value2, Measure_Id) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( Users_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Users_Id, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

type UserSpecsNumberRequest struct {
	Uuid      uuid.UUID       `json:"uuid"`
	UserId    int64           `json:"userId"`
	SpecsId   int64           `json:"specsId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

func (q *QueriesUser) CreateUserSpecsNumber(ctx context.Context, arg UserSpecsNumberRequest) (model.UserSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, createUserSpecsNumber,
		arg.UserId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.UserSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const updateUserSpecsNumber = `-- name: UpdateUserSpecsNumber :one
UPDATE Users_Specs_Number SET 
  Users_Id = $2,
  Specs_ID = $3,
  Value = $4,
  Value2 = $5,
  Measure_Id = $6
WHERE uuid = $1
RETURNING UUID, Users_Id, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

func (q *QueriesUser) UpdateUserSpecsNumber(ctx context.Context, arg UserSpecsNumberRequest) (model.UserSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, updateUserSpecsNumber,
		arg.Uuid,
		arg.UserId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.UserSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteUserSpecsNumber = `-- name: DeleteUserSpecsNumber :exec
DELETE FROM Users_Specs_Number
WHERE uuid = $1
`

func (q *QueriesUser) DeleteUserSpecsNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserSpecsNumber, uuid)
	return err
}

type UserSpecsNumberInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	UserId          int64           `json:"userId"`
	SpecsCode       string          `json:"specsCode"`
	SpecsId         int64           `json:"specsId"`
	Item            string          `json:"item"`
	ItemShortName   string          `json:"itemShortName"`
	ItemDescription string          `json:"itemDescription"`
	Value           decimal.Decimal `json:"value"`
	Value2          decimal.Decimal `json:"value2"`
	MeasureId       sql.NullInt64   `json:"measureId"`
	Measure         sql.NullString  `json:"measure"`
	MeasureUnit     sql.NullString  `json:"measureUnit"`
	ModCtr          int64           `json:"mod_ctr"`
	Created         sql.NullTime    `json:"created"`
	Updated         sql.NullTime    `json:"updated"`
}

func populateIdentitySpecNumber(q *QueriesUser, ctx context.Context, sql string) (UserSpecsNumberInfo, error) {
	var i UserSpecsNumberInfo
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
		&i.MeasureId,
		&i.Measure,
		&i.MeasureUnit,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecNumber2(q *QueriesUser, ctx context.Context, sql string) ([]UserSpecsNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []UserSpecsNumberInfo{}
	for rows.Next() {
		var i UserSpecsNumberInfo
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
			&i.MeasureId,
			&i.Measure,
			&i.MeasureUnit,

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

const userSpecsNumberSQL = `-- name: userSpecsNumberSQL
SELECT 
  mr.UUID, d.Users_Id, d.Specs_Code, d.Specs_ID, ref.Title Item, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Users_Specs_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
LEFT JOIN Reference mea on mea.ID = Measure_Id`

func (q *QueriesUser) GetUserSpecsNumber(ctx context.Context, userId int64, specsId int64) (UserSpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE d.Users_Id = %v and d.Specs_ID = %v",
		userSpecsNumberSQL, userId, specsId))
}

func (q *QueriesUser) GetUserSpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (UserSpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", userSpecsNumberSQL, uuid))
}

type ListUserSpecsNumberParams struct {
	UserId int64 `json:"UserId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUserSpecsNumber(ctx context.Context, arg ListUserSpecsNumberParams) ([]UserSpecsNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Users_Id = %v LIMIT %d OFFSET %d",
			userSpecsNumberSQL, arg.UserId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Users_Id = %v ", userSpecsNumberSQL, arg.UserId)
	}
	return populateIdentitySpecNumber2(q, ctx, sql)
}
