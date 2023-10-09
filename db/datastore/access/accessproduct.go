package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccessProduct = `-- name: CreateAccessProduct: one
INSERT INTO Access_Product 
  (Role_Id, Product_Id, Allow, Other_Info) 
VALUES ($1, $2, $3, $4) 
ON CONFLICT(Role_Id, Product_Id)
DO UPDATE SET
  Allow = Excluded.Allow,
  Other_Info = Excluded.Other_Info
RETURNING UUID, Role_Id, Product_Id, Allow, Other_Info
`

type AccessProductRequest struct {
	Uuid      uuid.UUID      `json:"uuid"`
	RoleId    int64          `json:"roleId"`
	ProductId int64          `json:"productId"`
	Allow     sql.NullBool   `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccess) CreateAccessProduct(ctx context.Context, arg AccessProductRequest) (model.AccessProduct, error) {
	row := q.db.QueryRowContext(ctx, createAccessProduct,
		arg.RoleId,
		arg.ProductId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.AccessProduct
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccessProduct = `-- name: DeleteAccessProduct :exec
DELETE FROM Access_Product
WHERE uuid = $1
`

func (q *QueriesAccess) DeleteAccessProduct(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccessProduct, uuid)
	return err
}

type AccessProductInfo struct {
	Uuid      uuid.UUID      `json:"uuid"`
	RoleId    int64          `json:"roleId"`
	ProductId int64          `json:"productId"`
	Allow     sql.NullBool   `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
	ModCtr    int64          `json:"modCtr"`
	Created   sql.NullTime   `json:"created"`
	Updated   sql.NullTime   `json:"updated"`
}

const getAccessProduct = `-- name: GetAccessProduct :one
SELECT 
mr.UUId, Role_Id, Product_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesAccess) GetAccessProduct(ctx context.Context, uuid uuid.UUID) (AccessProductInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccessProduct, uuid)
	var i AccessProductInfo
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccessProductbyUuId = `-- name: GetAccessProductbyUuId :one
SELECT 
mr.UUId, Role_Id, Product_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccess) GetAccessProductbyUuId(ctx context.Context, uuid uuid.UUID) (AccessProductInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccessProductbyUuId, uuid)
	var i AccessProductInfo
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAccessProduct = `-- name: ListAccessProduct:many
SELECT 
mr.UUId, Role_Id, Product_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Role_Id = $1
ORDER BY Product_Id
LIMIT $2
OFFSET $3
`

type ListAccessProductParams struct {
	RoleId int64 `json:"roleId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccess) ListAccessProduct(ctx context.Context, arg ListAccessProductParams) ([]AccessProductInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAccessProduct, arg.RoleId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccessProductInfo{}
	for rows.Next() {
		var i AccessProductInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.RoleId,
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

const updateAccessProduct = `-- name: UpdateAccessProduct :one
UPDATE Access_Product SET 
Role_Id = $2,
Product_Id = $3,
Allow = $4,
Other_Info = $5
WHERE uuid = $1
RETURNING UUId, Role_Id, Product_Id, Allow, Other_Info
`

func (q *QueriesAccess) UpdateAccessProduct(ctx context.Context, arg AccessProductRequest) (model.AccessProduct, error) {
	row := q.db.QueryRowContext(ctx, updateAccessProduct,
		arg.Uuid,
		arg.RoleId,
		arg.ProductId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.AccessProduct
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ProductId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}
