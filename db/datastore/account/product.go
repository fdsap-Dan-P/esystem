package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createProduct = `-- name: CreateProduct: one
INSERT INTO Product (
Code, Product_Name, Description, Normal_Balance, Isgl, 
Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6
) RETURNING Id, UUId, Code, Product_Name, Description, Normal_Balance, Isgl, 
Other_Info
`

type ProductRequest struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	Code          int64          `json:"code"`
	ProductName   string         `json:"productName"`
	Description   sql.NullString `json:"description"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateProduct(ctx context.Context, arg ProductRequest) (model.Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Code,
		arg.ProductName,
		arg.Description,
		arg.NormalBalance,
		arg.Isgl,
		arg.OtherInfo,
	)
	var i model.Product
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ProductName,
		&i.Description,
		&i.NormalBalance,
		&i.Isgl,
		&i.OtherInfo,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM Product
WHERE id = $1
`

func (q *QueriesAccount) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

type ProductInfo struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	Code          int64          `json:"code"`
	ProductName   string         `json:"productName"`
	Description   sql.NullString `json:"description"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	OtherInfo     sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getProduct = `-- name: GetProduct :one
SELECT 
Id, mr.UUId, Code, 
Product_Name, Description, Normal_Balance, Isgl, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesAccount) GetProduct(ctx context.Context, id int64) (ProductInfo, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i ProductInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ProductName,
		&i.Description,
		&i.NormalBalance,
		&i.Isgl,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getProductbyUuid = `-- name: GetProductbyUuid :one
SELECT 
Id, mr.UUId, Code, 
Product_Name, Description, Normal_Balance, Isgl, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetProductbyUuid(ctx context.Context, uuid uuid.UUID) (ProductInfo, error) {
	row := q.db.QueryRowContext(ctx, getProductbyUuid, uuid)
	var i ProductInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ProductName,
		&i.Description,
		&i.NormalBalance,
		&i.Isgl,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getProductbyName = `-- name: GetProductbyName :one
SELECT 
Id, mr.UUId, Code, 
Product_Name, Description, Normal_Balance, Isgl, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE LOWER(Product_Name) = LOWER($1) LIMIT 1
`

func (q *QueriesAccount) GetProductbyName(ctx context.Context, name string) (ProductInfo, error) {
	row := q.db.QueryRowContext(ctx, getProductbyName, name)
	var i ProductInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ProductName,
		&i.Description,
		&i.NormalBalance,
		&i.Isgl,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listProduct = `-- name: ListProduct:many
SELECT 
Id, mr.UUId, Code, 
Product_Name, Description, Normal_Balance, Isgl, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Product d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProductParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccount) ListProduct(ctx context.Context, arg ListProductParams) ([]ProductInfo, error) {
	rows, err := q.db.QueryContext(ctx, listProduct, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductInfo{}
	for rows.Next() {
		var i ProductInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Code,
			&i.ProductName,
			&i.Description,
			&i.NormalBalance,
			&i.Isgl,
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

const updateProduct = `-- name: UpdateProduct :one
UPDATE Product SET 
Code = $2,
Product_Name = $3,
Description = $4,
Normal_Balance = $5,
Isgl = $6,
Other_Info = $7
WHERE id = $1
RETURNING Id, UUId, Code, Product_Name, Description, Normal_Balance, Isgl, 
Other_Info
`

func (q *QueriesAccount) UpdateProduct(ctx context.Context, arg ProductRequest) (model.Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.Id,
		arg.Code,
		arg.ProductName,
		arg.Description,
		arg.NormalBalance,
		arg.Isgl,
		arg.OtherInfo,
	)
	var i model.Product
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ProductName,
		&i.Description,
		&i.NormalBalance,
		&i.Isgl,
		&i.OtherInfo,
	)
	return i, err
}
