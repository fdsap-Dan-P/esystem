package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createEmployeeEvent = `-- name: CreateEmployeeEvent: one
INSERT INTO Employee_Event (
	Employee_Id, Ticket_Id, Event_Type_Id, Office_Id, Position_Id, 
	Basic_Pay, Status_Id, Job_Grade, Job_Step, Level_Id, Employee_Type_Id, 
	Remarks, Other_Info
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
RETURNING 
  UUID, Employee_Id, Ticket_Id, Event_Type_Id, Office_Id, Position_Id, 
  Basic_Pay, Status_Id, Job_Grade, Job_Step, Level_Id, Employee_Type_Id, 
  Remarks, Other_Info
`

type EmployeeEventRequest struct {
	Uuid           uuid.UUID       `json:"uuid"`
	EmployeeId     int64           `json:"employeeId"`
	TicketId       int64           `json:"ticketId"`
	EventTypeId    int64           `json:"eventtypeId"`
	OfficeId       int64           `json:"officeId"`
	PositionId     int64           `json:"positionId"`
	BasicPay       decimal.Decimal `json:"basicPay"`
	StatusId       int64           `json:"statusId"`
	JobGrade       int16           `json:"jobGrade"`
	JobStep        int16           `json:"jobStep"`
	LevelId        sql.NullInt64   `json:"levelId"`
	EmployeeTypeId sql.NullInt64   `json:"employeeTypeId"`
	Remarks        sql.NullString  `json:"remarks"`
	OtherInfo      sql.NullString  `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateEmployeeEvent(ctx context.Context, arg EmployeeEventRequest) (model.EmployeeEvent, error) {
	row := q.db.QueryRowContext(ctx, createEmployeeEvent,
		arg.EmployeeId,
		arg.TicketId,
		arg.EventTypeId,
		arg.OfficeId,
		arg.PositionId,
		arg.BasicPay,
		arg.StatusId,
		arg.JobGrade,
		arg.JobStep,
		arg.LevelId,
		arg.EmployeeTypeId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.EmployeeEvent
	err := row.Scan(
		&i.Uuid,
		&i.EmployeeId,
		&i.TicketId,
		&i.EventTypeId,
		&i.OfficeId,
		&i.PositionId,
		&i.BasicPay,
		&i.StatusId,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.EmployeeTypeId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteEmployeeEvent = `-- name: DeleteEmployeeEvent :exec
DELETE FROM Employee_Event
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteEmployeeEvent(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEmployeeEvent, uuid)
	return err
}

type EmployeeEventInfo struct {
	Uuid           uuid.UUID       `json:"uuid"`
	EmployeeId     int64           `json:"employeeId"`
	TicketId       int64           `json:"ticketId"`
	EventTypeId    int64           `json:"eventtypeId"`
	OfficeId       int64           `json:"officeId"`
	PositionId     int64           `json:"positionId"`
	BasicPay       decimal.Decimal `json:"basicPay"`
	StatusId       int64           `json:"statusId"`
	JobGrade       int16           `json:"jobGrade"`
	JobStep        int16           `json:"jobStep"`
	LevelId        sql.NullInt64   `json:"levelId"`
	EmployeeTypeId sql.NullInt64   `json:"employeeTypeId"`
	Remarks        sql.NullString  `json:"remarks"`
	OtherInfo      sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getEmployeeEvent = `-- name: GetEmployeeEvent :one
SELECT 
	mr.UUId, Employee_Id, Ticket_Id, 
	Event_Type_Id, Office_Id, Position_Id, Basic_Pay, Status_Id, Job_Grade, 
	Job_Step, Level_Id, Employee_Type_Id, Remarks, Other_Info
	,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employee_Event d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesTransaction) GetEmployeeEvent(ctx context.Context, uuid uuid.UUID) (EmployeeEventInfo, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeEvent, uuid)
	var i EmployeeEventInfo
	err := row.Scan(
		&i.Uuid,
		&i.EmployeeId,
		&i.TicketId,
		&i.EventTypeId,
		&i.OfficeId,
		&i.PositionId,
		&i.BasicPay,
		&i.StatusId,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.EmployeeTypeId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getEmployeeEventbyUuid = `-- name: GetEmployeeEventbyUuid :one
SELECT 
	mr.UUId, Employee_Id, Ticket_Id, 
	Event_Type_Id, Office_Id, Position_Id, Basic_Pay, Status_Id, Job_Grade, 
	Job_Step, Level_Id, Employee_Type_Id, Remarks, Other_Info
	,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employee_Event d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesTransaction) GetEmployeeEventbyUuid(ctx context.Context, uuid uuid.UUID) (EmployeeEventInfo, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeEventbyUuid, uuid)
	var i EmployeeEventInfo
	err := row.Scan(
		&i.Uuid,
		&i.EmployeeId,
		&i.TicketId,
		&i.EventTypeId,
		&i.OfficeId,
		&i.PositionId,
		&i.BasicPay,
		&i.StatusId,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.EmployeeTypeId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listEmployeeEvent = `-- name: ListEmployeeEvent:many
SELECT 
	mr.UUId, Employee_Id, Ticket_Id, 
	Event_Type_Id, Office_Id, Position_Id, Basic_Pay, Status_Id, Job_Grade, 
	Job_Step, Level_Id, Employee_Type_Id, Remarks, Other_Info
	,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employee_Event d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE employee_id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListEmployeeEventParams struct {
	EmployeeId int64 `json:"employeeId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesTransaction) ListEmployeeEvent(ctx context.Context, arg ListEmployeeEventParams) ([]EmployeeEventInfo, error) {
	rows, err := q.db.QueryContext(ctx, listEmployeeEvent, arg.EmployeeId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EmployeeEventInfo{}
	for rows.Next() {
		var i EmployeeEventInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.EmployeeId,
			&i.TicketId,
			&i.EventTypeId,
			&i.OfficeId,
			&i.PositionId,
			&i.BasicPay,
			&i.StatusId,
			&i.JobGrade,
			&i.JobStep,
			&i.LevelId,
			&i.EmployeeTypeId,
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

const updateEmployeeEvent = `-- name: UpdateEmployeeEvent :one
UPDATE Employee_Event SET 
	Employee_Id = $2,
	Ticket_Id = $3,
	Event_Type_Id = $4,
	Office_Id = $5,
	Position_Id = $6,
	Basic_Pay = $7,
	Status_Id = $8,
	Job_Grade = $9,
	Job_Step = $10,
	Level_Id = $11,
	Employee_Type_Id = $12,
	Remarks = $13,
	Other_Info = $14
WHERE uuid = $1
RETURNING 
  UUID, Employee_Id, Ticket_Id, Event_Type_Id, Office_Id, Position_Id, 
  Basic_Pay, Status_Id, Job_Grade, Job_Step, Level_Id, Employee_Type_Id, 
  Remarks, Other_Info
`

func (q *QueriesTransaction) UpdateEmployeeEvent(ctx context.Context, arg EmployeeEventRequest) (model.EmployeeEvent, error) {
	row := q.db.QueryRowContext(ctx, updateEmployeeEvent,

		arg.Uuid,
		arg.EmployeeId,
		arg.TicketId,
		arg.EventTypeId,
		arg.OfficeId,
		arg.PositionId,
		arg.BasicPay,
		arg.StatusId,
		arg.JobGrade,
		arg.JobStep,
		arg.LevelId,
		arg.EmployeeTypeId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.EmployeeEvent
	err := row.Scan(
		&i.Uuid,
		&i.EmployeeId,
		&i.TicketId,
		&i.EventTypeId,
		&i.OfficeId,
		&i.PositionId,
		&i.BasicPay,
		&i.StatusId,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.EmployeeTypeId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
