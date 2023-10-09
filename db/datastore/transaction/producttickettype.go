package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createProductTicketType = `-- name: CreateProductTicketType: one
INSERT INTO Product_Ticket_Type(
   uuid, central_office_id, product_id, ticket_type_id, status_id, other_info )
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (UUID)
DO UPDATE SET
	uuid =  EXCLUDED.uuid,
	central_office_id =  EXCLUDED.central_office_id,
	product_id =  EXCLUDED.product_id,
	ticket_type_id =  EXCLUDED.ticket_type_id,
	status_id =  EXCLUDED.status_id,
	other_info =  EXCLUDED.other_info

RETURNING Id, uuid, central_office_id, product_id, ticket_type_id, status_id, other_info`

type ProductTicketTypeRequest struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	ProductId       int64          `json:"productId"`
	TicketTypeId    int64          `json:"ticketTypeId"`
	StatusId        int64          `json:"statusId"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateProductTicketType(ctx context.Context, arg ProductTicketTypeRequest) (model.ProductTicketType, error) {
	row := q.db.QueryRowContext(ctx, createProductTicketType,
		arg.Uuid,
		arg.CentralOfficeId,
		arg.ProductId,
		arg.TicketTypeId,
		arg.StatusId,
		arg.OtherInfo,
	)
	var i model.ProductTicketType
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.ProductId,
		&i.TicketTypeId,
		&i.StatusId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteProductTicketType = `-- name: DeleteProductTicketType :exec
DELETE FROM Product_Ticket_Type
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteProductTicketType(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProductTicketType, uuid)
	return err
}

type ProductTicketTypeInfo struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	ProductId       int64          `json:"productId"`
	TicketTypeId    int64          `json:"ticketTypeId"`
	StatusId        int64          `json:"statusId"`
	OtherInfo       sql.NullString `json:"otherInfo"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

const productTicketTypeSQL = `-- name: ProductTicketTypeSQL :one
SELECT
id, mr.UUID, central_office_id, product_id, ticket_type_id, status_id, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Product_Ticket_Type d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateProductTicketType(q *QueriesTransaction, ctx context.Context, sql string) (ProductTicketTypeInfo, error) {
	var i ProductTicketTypeInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.ProductId,
		&i.TicketTypeId,
		&i.StatusId,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateProductTicketTypes(q *QueriesTransaction, ctx context.Context, sql string) ([]ProductTicketTypeInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductTicketTypeInfo{}
	for rows.Next() {
		var i ProductTicketTypeInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.CentralOfficeId,
			&i.ProductId,
			&i.TicketTypeId,
			&i.StatusId,
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

func (q *QueriesTransaction) GetProductTicketType(ctx context.Context, id int64) (ProductTicketTypeInfo, error) {
	return populateProductTicketType(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", productTicketTypeSQL, id))
}

func (q *QueriesTransaction) GetProductTicketTypebyUuid(ctx context.Context, uuid uuid.UUID) (ProductTicketTypeInfo, error) {
	return populateProductTicketType(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", productTicketTypeSQL, uuid))
}

type ListProductTicketTypeParams struct {
	ProductId int64 `json:"productId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListProductTicketType(ctx context.Context, arg ListProductTicketTypeParams) ([]ProductTicketTypeInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			productTicketTypeSQL, arg.Limit, arg.Offset)
	} else {
		sql = productTicketTypeSQL
	}
	return populateProductTicketTypes(q, ctx, sql)
}

const updateProductTicketType = `-- name: UpdateProductTicketType :one
UPDATE Product_Ticket_Type SET 
uuid = $2,
central_office_id = $3,
product_id = $4,
ticket_type_id = $5,
status_id = $6,
other_info = $7
WHERE id = $1
RETURNING id, uuid, central_office_id, product_id, ticket_type_id, status_id, other_info
`

func (q *QueriesTransaction) UpdateProductTicketType(ctx context.Context, arg ProductTicketTypeRequest) (model.ProductTicketType, error) {
	row := q.db.QueryRowContext(ctx, updateProductTicketType,
		arg.Id,
		arg.Uuid,
		arg.CentralOfficeId,
		arg.ProductId,
		arg.TicketTypeId,
		arg.StatusId,
		arg.OtherInfo,
	)
	var i model.ProductTicketType
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.ProductId,
		&i.TicketTypeId,
		&i.StatusId,
		&i.OtherInfo,
	)
	return i, err
}
