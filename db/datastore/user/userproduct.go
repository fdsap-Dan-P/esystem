package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createUserProduct = `-- name: CreateUserProduct: one
INSERT INTO User_Product (
User_Id, Product_Id, Allow, Other_Info) 
VALUES ($1, $2, $3, $4) 
ON CONFLICT(User_ID, Product_Id) DO UPDATE SET
  Allow = EXCLUDED.Allow,
  Other_Info = EXCLUDED.Other_Info
RETURNING UUId, User_Id, Product_Id, Allow, Other_Info
`

type UserProductRequest struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	ProductId int64          `json:"productId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesUser) CreateUserProduct(ctx context.Context, arg UserProductRequest) (model.UserProduct, error) {
	row := q.db.QueryRowContext(ctx, createUserProduct,
		arg.UserId,
		arg.ProductId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.UserProduct
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}

const deleteUserProduct = `-- name: DeleteUserProduct :exec
DELETE FROM User_Product
WHERE uuid = $1
`

func (q *QueriesUser) DeleteUserProduct(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserProduct, uuid)
	return err
}

type UserProductInfo struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	ProductId int64          `json:"productId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getUserProduct = `-- name: GetUserProduct :one
SELECT 
mr.UUId, User_Id, Product_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesUser) GetUserProduct(ctx context.Context, uuid uuid.UUID) (UserProductInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserProduct, uuid)
	var i UserProductInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserProductbyUuId = `-- name: GetUserProductbyUuId :one
SELECT 
mr.UUId, User_Id, Product_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesUser) GetUserProductbyUuId(ctx context.Context, uuid uuid.UUID) (UserProductInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserProductbyUuId, uuid)
	var i UserProductInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserProductbyName = `-- name: GetUserProductbyName :one
SELECT 
mr.UUId, User_Id, Product_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Title = $1 LIMIT 1
`

func (q *QueriesUser) GetUserProductbyName(ctx context.Context, name string) (UserProductInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserProductbyName, name)
	var i UserProductInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listUserProduct = `-- name: ListUserProduct:many
SELECT 
mr.UUId, User_Id, Product_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE User_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListUserProductParams struct {
	UserId int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUserProduct(ctx context.Context, arg ListUserProductParams) ([]UserProductInfo, error) {
	rows, err := q.db.QueryContext(ctx, listUserProduct, arg.UserId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserProductInfo{}
	for rows.Next() {
		var i UserProductInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.UserId,
			&i.ProductId,
			&i.Allow,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
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

const updateUserProduct = `-- name: UpdateUserProduct :one
UPDATE User_Product SET 
User_Id = $2,
Product_Id = $3,
Allow = $4,
Other_Info = $5
WHERE uuid = $1
RETURNING UUId, User_Id, Product_Id, Allow, Other_Info
`

func (q *QueriesUser) UpdateUserProduct(ctx context.Context, arg UserProductRequest) (model.UserProduct, error) {
	row := q.db.QueryRowContext(ctx, updateUserProduct,

		arg.Uuid,
		arg.UserId,
		arg.ProductId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.UserProduct
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}
