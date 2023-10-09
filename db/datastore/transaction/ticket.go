package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"simplebank/model"
	"simplebank/util"
	"time"

	"github.com/google/uuid"
)

const createTicket = `-- name: CreateTicket: one
INSERT INTO Ticket(
   uuid, central_office_id, ticket_type_id, ticket_date, postedby_id, 
   status_id, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(UUID)
DO UPDATE SET 
	central_office_id =  EXCLUDED.central_office_id,
	ticket_type_id =  EXCLUDED.ticket_type_id,
	ticket_date =  EXCLUDED.ticket_date,
	postedby_id =  EXCLUDED.postedby_id,
	status_id =  EXCLUDED.status_id,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING 
  id, uuid, central_office_id, ticket_type_id, ticket_date, postedby_id, 
  status_id, remarks, other_info`

type TicketRequest struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	TicketTypeId    int64          `json:"ticketTypeId"`
	TicketDate      time.Time      `json:"ticketDate"`
	PostedbyId      int64          `json:"postedbyId"`
	StatusId        int64          `json:"statusId"`
	Remarks         sql.NullString `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

type NewDailyTicketRequest struct {
	CentralOffice string         `json:"centralOffice"`
	TicketType    string         `json:"ticketType"`
	TicketDate    time.Time      `json:"ticketDate"`
	Postedby      string         `json:"postedby"`
	Status        string         `json:"status"`
	Remarks       sql.NullString `json:"remarks"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

const newTicketSQL = `-- name: TicketSQL :one
SELECT 
  t.id, t.uuid, o.id central_office_id, trntype.id ticket_type_id, t.ticket_date, u.id postedby_id, stat.id status_id, t.remarks, t.other_info
FROM Reference trnType 
INNER JOIN Office o on lower(o.Code) = lower($1)
INNER JOIN Users u on lower(u.login_name)  = lower($2)
INNER JOIN Reference stat on lower(stat.ref_type) = 'ticketstatus' 
  and lower(stat.short_name) = lower($3)
LEFT JOIN Ticket t
  on trnType.id = t.ticket_type_id 
 and o.id = t.central_office_id 
 and u.id  = t.postedby_id 
 and t.ticket_date = $4
WHERE  lower(trnType.ref_type) = 'tickettype' 
  and lower(trnType.short_name) = lower($5)
ORDER BY t.id  
LIMIT 1  
`

type NullTicket struct {
	Id              sql.NullInt64  `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId sql.NullInt64  `json:"centralOfficeId"`
	TicketTypeId    sql.NullInt64  `json:"ticketTypeId"`
	TicketDate      sql.NullTime   `json:"ticketDate"`
	PostedbyId      sql.NullInt64  `json:"postedbyId"`
	StatusId        sql.NullInt64  `json:"statusId"`
	Remarks         sql.NullString `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) NewDailyTicket(ctx context.Context, arg NewDailyTicketRequest) (model.Ticket, error) {
	var i NullTicket

	row := q.db.QueryRowContext(ctx, newTicketSQL,
		arg.CentralOffice, arg.Postedby, arg.Status, arg.TicketDate, arg.TicketType)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.TicketTypeId,
		&i.TicketDate,
		&i.PostedbyId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	dt := model.Ticket{}
	if err != nil {
		return dt, err
	}

	log.Printf("tic: %v er: %v", i, err)

	if i.Id.Valid {
		dt = model.Ticket{
			// Uuid:            util.ToUUID("a9a97b7d-26a6-4a0e-a618-8727b83a823d"),
			Id:              i.Id.Int64,
			Uuid:            i.Uuid,
			CentralOfficeId: i.CentralOfficeId.Int64,
			TicketTypeId:    i.PostedbyId.Int64,
			TicketDate:      i.TicketDate.Time,
			PostedbyId:      i.Id.Int64,
			StatusId:        i.StatusId.Int64,
			Remarks:         i.Remarks,
			OtherInfo:       i.OtherInfo,
		}
	} else {
		dt, err = q.CreateTicket(ctx, TicketRequest{
			Uuid:            util.UUID(),
			CentralOfficeId: i.CentralOfficeId.Int64,
			TicketTypeId:    i.TicketTypeId.Int64,
			TicketDate:      arg.TicketDate,
			PostedbyId:      i.PostedbyId.Int64,
			StatusId:        i.StatusId.Int64,
			Remarks:         arg.Remarks,
			OtherInfo:       arg.OtherInfo,
		})
	}
	return dt, err
}

func (q *QueriesTransaction) CreateTicket(ctx context.Context, arg TicketRequest) (model.Ticket, error) {
	row := q.db.QueryRowContext(ctx, createTicket,
		arg.Uuid,
		arg.CentralOfficeId,
		arg.TicketTypeId,
		arg.TicketDate,
		arg.PostedbyId,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Ticket
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.TicketTypeId,
		&i.TicketDate,
		&i.PostedbyId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteTicket = `-- name: DeleteTicket :exec
DELETE FROM Ticket
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicket(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicket, uuid)
	return err
}

type TicketInfo struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	TicketTypeId    int64          `json:"ticketTypeId"`
	TicketDate      time.Time      `json:"ticketDate"`
	PostedbyId      int64          `json:"postedbyId"`
	StatusId        int64          `json:"statusId"`
	Remarks         sql.NullString `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

const ticketSQL = `-- name: TicketSQL :one
SELECT
  id, mr.UUID, central_office_id, ticket_type_id, ticket_date, postedby_id, status_id, remarks, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateTicket(q *QueriesTransaction, ctx context.Context, sql string) (TicketInfo, error) {
	var i TicketInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.TicketTypeId,
		&i.TicketDate,
		&i.PostedbyId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTickets(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TicketInfo{}
	for rows.Next() {
		var i TicketInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.CentralOfficeId,
			&i.TicketTypeId,
			&i.TicketDate,
			&i.PostedbyId,
			&i.StatusId,
			&i.Remarks,
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

func (q *QueriesTransaction) GetTicket(ctx context.Context, id int64) (TicketInfo, error) {
	return populateTicket(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", ticketSQL, id))
}

func (q *QueriesTransaction) GetTicketbyUuid(ctx context.Context, uuid uuid.UUID) (TicketInfo, error) {
	return populateTicket(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", ticketSQL, uuid))
}

type ListTicketParams struct {
	TicketDate time.Time `json:"ticketDate"`
	Limit      int32     `json:"limit"`
	Offset     int32     `json:"offset"`
}

func (q *QueriesTransaction) ListTicket(ctx context.Context, arg ListTicketParams) ([]TicketInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			ticketSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(ticketSQL)
	}
	return populateTickets(q, ctx, sql)
}

const updateTicket = `-- name: UpdateTicket :one
UPDATE Ticket SET 
uuid = $2,
central_office_id = $3,
ticket_type_id = $4,
ticket_date = $5,
postedby_id = $6,
status_id = $7,
remarks = $8,
other_info = $9
WHERE id = $1
RETURNING id, uuid, central_office_id, ticket_type_id, ticket_date, postedby_id, status_id, remarks, other_info
`

func (q *QueriesTransaction) UpdateTicket(ctx context.Context, arg TicketRequest) (model.Ticket, error) {
	row := q.db.QueryRowContext(ctx, updateTicket,
		arg.Id,
		arg.Uuid,
		arg.CentralOfficeId,
		arg.TicketTypeId,
		arg.TicketDate,
		arg.PostedbyId,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Ticket
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.TicketTypeId,
		&i.TicketDate,
		&i.PostedbyId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
