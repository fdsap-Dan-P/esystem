package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createCustomerEvent = `-- name: CreateCustomerEvent: one
INSERT INTO Customer_Event (
Trn_Head_Id, Customer_Id, Type_Id, Remarks, Other_Info
) VALUES (
$1, $2, $3, $4, $5
) RETURNING UUId, Trn_Head_Id, Customer_Id, Type_Id, Remarks, Other_Info
`

type CustomerEventRequest struct {
	Uuid       uuid.UUID      `json:"uuid"`
	TrnHeadId  int64          `json:"trnHeadId"`
	CustomerId int64          `json:"customerId"`
	TypeId     int64          `json:"typeId"`
	Remarks    string         `json:"remarks"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateCustomerEvent(ctx context.Context, arg CustomerEventRequest) (model.CustomerEvent, error) {
	row := q.db.QueryRowContext(ctx, createCustomerEvent,
		arg.TrnHeadId,
		arg.CustomerId,
		arg.TypeId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.CustomerEvent
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.CustomerId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteCustomerEvent = `-- name: DeleteCustomerEvent :exec
DELETE FROM Customer_Event
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteCustomerEvent(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerEvent, uuid)
	return err
}

type CustomerEventInfo struct {
	Uuid       uuid.UUID      `json:"uuid"`
	TrnHeadId  int64          `json:"trnHeadId"`
	CustomerId int64          `json:"customerId"`
	TypeId     int64          `json:"typeId"`
	Remarks    string         `json:"remarks"`
	OtherInfo  sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getCustomerEvent = `-- name: GetCustomerEvent :one
SELECT 
mr.UUId, 
Trn_Head_Id, Customer_Id, Type_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Event d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesTransaction) GetCustomerEvent(ctx context.Context, uuid uuid.UUID) (CustomerEventInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerEvent, uuid)
	var i CustomerEventInfo
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.CustomerId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCustomerEventbyUuid = `-- name: GetCustomerEventbyUuid :one
SELECT 
mr.UUId, 
Trn_Head_Id, Customer_Id, Type_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Event d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesTransaction) GetCustomerEventbyUuid(ctx context.Context, uuid uuid.UUID) (CustomerEventInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerEventbyUuid, uuid)
	var i CustomerEventInfo
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.CustomerId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listCustomerEvent = `-- name: ListCustomerEvent:many
SELECT 
mr.UUId, 
Trn_Head_Id, Customer_Id, Type_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Event d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Customer_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListCustomerEventParams struct {
	CustomerId int64 `json:"customerId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesTransaction) ListCustomerEvent(ctx context.Context, arg ListCustomerEventParams) ([]CustomerEventInfo, error) {
	rows, err := q.db.QueryContext(ctx, listCustomerEvent, arg.CustomerId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CustomerEventInfo{}
	for rows.Next() {
		var i CustomerEventInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.CustomerId,
			&i.TypeId,
			&i.Remarks,
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

const updateCustomerEvent = `-- name: UpdateCustomerEvent :one
UPDATE Customer_Event SET 
Trn_Head_Id = $2,
Customer_Id = $3,
Type_Id = $4,
Remarks = $5,
Other_Info = $6
WHERE uuid = $1
RETURNING UUId, Trn_Head_Id, Customer_Id, Type_Id, Remarks, Other_Info
`

func (q *QueriesTransaction) UpdateCustomerEvent(ctx context.Context, arg CustomerEventRequest) (model.CustomerEvent, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerEvent,

		arg.Uuid,
		arg.TrnHeadId,
		arg.CustomerId,
		arg.TypeId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.CustomerEvent
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.CustomerId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
