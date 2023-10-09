package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTrnHead = `-- name: CreateTrnHead: one
INSERT INTO Trn_Head (
  UUID, Trn_Serial, Ticket_Id, Trn_Date, Type_Id, Particular, Office_Id, 
  User_Id, Transacting_Iiid, ORNo, isFinal, isManual, Alternate_Trn, 
  Reference, Other_Info
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) 
ON CONFLICT(UUID) DO UPDATE SET
Trn_Serial = EXCLUDED.Trn_Serial, 
Ticket_Id = EXCLUDED.Ticket_Id, 
  Trn_Date = EXCLUDED.Trn_Date, 
  Type_Id = EXCLUDED.Type_Id, 
  Particular = EXCLUDED.Particular, 
  Office_Id = EXCLUDED.Office_Id, 
  User_Id = EXCLUDED.User_Id, 
  Transacting_Iiid = EXCLUDED.Transacting_Iiid, 
  ORNo = EXCLUDED.ORNo, 
  isFinal = EXCLUDED.isFinal, 
  isManual = EXCLUDED.isManual, 
  Alternate_Trn = EXCLUDED.Alternate_Trn, 
  Reference = EXCLUDED.Reference, 
  Other_Info = EXCLUDED.Other_Info
RETURNING Id, UUID, Trn_Serial, Ticket_Id, Trn_Date, Type_Id, Particular, Office_Id, 
User_Id, Transacting_Iiid, ORNo, isFinal, isManual, Alternate_Trn, 
Reference, Other_Info
`

type TrnHeadRequest struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	TrnSerial       string         `json:"trnSerial"`
	TicketId        int64          `json:"ticketId"`
	TrnDate         time.Time      `json:"trn_date"`
	TypeId          int64          `json:"typeId"`
	Particular      sql.NullString `json:"particular"`
	OfficeId        int64          `json:"officeId"`
	UserId          int64          `json:"userId"`
	TransactingIiid sql.NullInt64  `json:"transactingIiid"`
	Orno            sql.NullString `json:"orno"`
	Isfinal         sql.NullBool   `json:"isfinal"`
	Ismanual        sql.NullBool   `json:"ismanual"`
	AlternateTrn    sql.NullString `json:"alternateTrn"`
	Reference       string         `json:"reference"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

type TrnHeadFullRequest struct {
	Id              int64                       `json:"id"`
	Uuid            uuid.UUID                   `json:"uuid"`
	TrnSerial       string                      `json:"trnSerial"`
	TicketId        int64                       `json:"ticketId"`
	TrnDate         time.Time                   `json:"trn_date"`
	TypeId          int64                       `json:"typeId"`
	Particular      sql.NullString              `json:"particular"`
	OfficeId        int64                       `json:"officeId"`
	UserId          int64                       `json:"userId"`
	TransactingIiid sql.NullInt64               `json:"transactingIiid"`
	Orno            sql.NullString              `json:"orno"`
	Isfinal         sql.NullBool                `json:"isfinal"`
	Ismanual        sql.NullBool                `json:"ismanual"`
	AlternateTrn    sql.NullString              `json:"alternateTrn"`
	Reference       string                      `json:"reference"`
	OtherInfo       sql.NullString              `json:"otherInfo"`
	AccountTran     []AccountTranRequest        `json:"accountTran"`
	SpecsDate       []TrnHeadSpecsDateRequest   `json:"specsDate"`
	SpecsNumber     []TrnHeadSpecsNumberRequest `json:"SpecsNumber"`
	SpecsReference  []TrnHeadSpecsRefRequest    `json:"SpecsReference"`
	SpecsString     []TrnHeadSpecsStringRequest `json:"SpecsString"`
}

func (q *QueriesTransaction) CreateTrnHead(ctx context.Context, arg TrnHeadRequest) (model.TrnHead, error) {
	row := q.db.QueryRowContext(ctx, createTrnHead,
		arg.Uuid,
		arg.TrnSerial,
		arg.TicketId,
		arg.TrnDate,
		arg.TypeId,
		arg.Particular,
		arg.OfficeId,
		arg.UserId,
		arg.TransactingIiid,
		arg.Orno,
		arg.Isfinal,
		arg.Ismanual,
		arg.AlternateTrn,
		arg.Reference,
		arg.OtherInfo,
	)
	var i model.TrnHead
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TrnSerial,
		&i.TicketId,
		&i.TrnDate,
		&i.TypeId,
		&i.Particular,
		&i.OfficeId,
		&i.UserId,
		&i.TransactingIiid,
		&i.Orno,
		&i.Isfinal,
		&i.Ismanual,
		&i.AlternateTrn,
		&i.Reference,
		&i.OtherInfo,
	)
	return i, err
}

func (q *QueriesTransaction) CreateTrnHeadFull(ctx context.Context, arg TrnHeadRequest) (model.TrnHead, error) {
	row := q.db.QueryRowContext(ctx, createTrnHead,
		arg.Uuid,
		arg.TrnSerial,
		arg.TicketId,
		arg.TrnDate,
		arg.TypeId,
		arg.Particular,
		arg.OfficeId,
		arg.UserId,
		arg.TransactingIiid,
		arg.Orno,
		arg.Isfinal,
		arg.Ismanual,
		arg.AlternateTrn,
		arg.Reference,
		arg.OtherInfo,
	)
	var i model.TrnHead
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TrnSerial,
		&i.TicketId,
		&i.TrnDate,
		&i.TypeId,
		&i.Particular,
		&i.OfficeId,
		&i.UserId,
		&i.TransactingIiid,
		&i.Orno,
		&i.Isfinal,
		&i.Ismanual,
		&i.AlternateTrn,
		&i.Reference,
		&i.OtherInfo,
	)
	return i, err
}

const deleteTrnHead = `-- name: DeleteTrnHead :exec
DELETE FROM Trn_Head
WHERE id = $1
`

func (q *QueriesTransaction) DeleteTrnHead(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTrnHead, id)
	return err
}

type TrnHeadInfo struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	TrnSerial       string         `json:"trnSerial"`
	TicketId        int64          `json:"ticketId"`
	TrnDate         time.Time      `json:"trn_date"`
	TypeId          int64          `json:"typeId"`
	Particular      sql.NullString `json:"particular"`
	OfficeId        int64          `json:"officeId"`
	UserId          int64          `json:"userId"`
	TransactingIiid sql.NullInt64  `json:"transactingIiid"`
	Orno            sql.NullString `json:"orno"`
	Isfinal         sql.NullBool   `json:"isfinal"`
	Ismanual        sql.NullBool   `json:"ismanual"`
	AlternateTrn    sql.NullString `json:"alternateTrn"`
	Reference       string         `json:"reference"`
	OtherInfo       sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getTrnHead = `-- name: GetTrnHead :one
SELECT 
  Id, mr.UUID, Trn_Serial, Ticket_Id, Trn_Date, 
  Type_Id, Particular, Office_Id, User_Id, Transacting_Iiid, ORNo, isFinal, 
  isManual, Alternate_Trn, Reference, Other_Info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Trn_Head d 
INNER JOIN Main_Record mr on mr.UUId = d.UUId
`

func populateInventoryDetails(q *QueriesTransaction, ctx context.Context, sql string, param ...interface{}) ([]TrnHeadInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TrnHeadInfo{}
	for rows.Next() {
		var i TrnHeadInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.TrnSerial,
			&i.TicketId,
			&i.TrnDate,
			&i.TypeId,
			&i.Particular,
			&i.OfficeId,
			&i.UserId,
			&i.TransactingIiid,
			&i.Orno,
			&i.Isfinal,
			&i.Ismanual,
			&i.AlternateTrn,
			&i.Reference,
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

func (q *QueriesTransaction) GetTrnHead(ctx context.Context, id int64) (TrnHeadInfo, error) {
	script := fmt.Sprintf(`%v WHERE d.id = $1`, getTrnHead)
	log.Printf("sql: %v", script)
	items, err := populateInventoryDetails(q, ctx, script, id)
	if err != nil {
		return TrnHeadInfo{}, err
	}
	log.Printf("items: %v", items)
	if len(items) > 0 {
		return items[0], err
	} else {
		return TrnHeadInfo{}, fmt.Errorf("trnHead ID:%v not found", id)
	}
}

func (q *QueriesTransaction) GetTrnHeadbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadInfo, error) {
	script := fmt.Sprintf(`%v WHERE mr.uuid = $1`, getTrnHead)
	log.Printf("sql: %v", script)
	items, err := populateInventoryDetails(q, ctx, script, uuid)
	log.Printf("items: %v", items)
	if len(items) > 0 {
		return items[0], err
	} else {
		return TrnHeadInfo{}, fmt.Errorf("trnHead UUID:%v not found", uuid)
	}
}

type ListTrnHeadParams struct {
	TicketId int64 `json:"ticketId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTrnHead(ctx context.Context, arg ListTrnHeadParams) ([]TrnHeadInfo, error) {
	sql := fmt.Sprintf(`%v WHERE Ticket_Id = $1
	ORDER BY id LIMIT $2 OFFSET $3 `, getTrnHead)
	return populateInventoryDetails(q, ctx, sql, arg.TicketId, arg.Limit, arg.Offset)
}

const updateTrnHead = `-- name: UpdateTrnHead :one
UPDATE Trn_Head SET 
  Trn_Serial = $2,
  Ticket_Id = $3,
  Trn_Date = $4,
  Type_Id = $5,
  Particular = $6,
  Office_Id = $7,
  User_Id = $8,
  Transacting_Iiid = $9,
  ORNo = $10,
  isFinal = $11,
  isManual = $12,
  Alternate_Trn = $13,
  Reference = $14,
  Other_Info = $15
WHERE id = $1
RETURNING Id, UUID, Trn_Serial, Ticket_Id, Trn_Date, Type_Id, Particular, Office_Id, 
User_Id, Transacting_Iiid, ORNo, isFinal, isManual, Alternate_Trn, 
Reference, Other_Info
`

func (q *QueriesTransaction) UpdateTrnHead(ctx context.Context, arg TrnHeadRequest) (model.TrnHead, error) {
	row := q.db.QueryRowContext(ctx, updateTrnHead,
		arg.Id,
		arg.TrnSerial,
		arg.TicketId,
		arg.TrnDate,
		arg.TypeId,
		arg.Particular,
		arg.OfficeId,
		arg.UserId,
		arg.TransactingIiid,
		arg.Orno,
		arg.Isfinal,
		arg.Ismanual,
		arg.AlternateTrn,
		arg.Reference,
		arg.OtherInfo,
	)
	var i model.TrnHead
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TrnSerial,
		&i.TicketId,
		&i.TrnDate,
		&i.TypeId,
		&i.Particular,
		&i.OfficeId,
		&i.UserId,
		&i.TransactingIiid,
		&i.Orno,
		&i.Isfinal,
		&i.Ismanual,
		&i.AlternateTrn,
		&i.Reference,
		&i.OtherInfo,
	)
	return i, err
}
