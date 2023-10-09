package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const createSubjectEvent = `-- name: CreateSubjectEvent: one
INSERT INTO Subject_Event(
   uuid, type_id, ticket_item_id, iiid, section_subject_id, event_date, 
   grading_period, item_count, status_id, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT(UUID)
DO UPDATE SET
	type_id =  EXCLUDED.type_id,
	ticket_item_id =  EXCLUDED.ticket_item_id,
	iiid =  EXCLUDED.iiid,
	section_subject_id =  EXCLUDED.section_subject_id,
	event_date =  EXCLUDED.event_date,
	grading_period =  EXCLUDED.grading_period,
	item_count =  EXCLUDED.item_count,
	status_id =  EXCLUDED.status_id,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING 
  id, uuid, type_id, ticket_item_id, iiid, section_subject_id, event_date, 
  grading_period, item_count, status_id, remarks, other_info`

// -- (0-N/A, 1,2,3,4 Qtr, 10-Finals)
type SubjectEventRequest struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	TypeId           int64           `json:"typeId"`
	TicketItemId     int64           `json:"ticketItemId"`
	Iiid             int64           `json:"iiid"`
	SectionSubjectId int64           `json:"sectionSubjectId"`
	EventDate        time.Time       `json:"eventDate"`
	GradingPeriod    int16           `json:"gradingPeriod"`
	ItemCount        decimal.Decimal `json:"itemCount"`
	StatusId         int64           `json:"statusId"`
	Remarks          sql.NullString  `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
}

func (q *QueriesSchool) CreateSubjectEvent(ctx context.Context, arg SubjectEventRequest) (model.SubjectEvent, error) {
	row := q.db.QueryRowContext(ctx, createSubjectEvent,
		arg.Uuid,
		arg.TypeId,
		arg.TicketItemId,
		arg.Iiid,
		arg.SectionSubjectId,
		arg.EventDate,
		arg.GradingPeriod,
		arg.ItemCount,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.SubjectEvent
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.TicketItemId,
		&i.Iiid,
		&i.SectionSubjectId,
		&i.EventDate,
		&i.GradingPeriod,
		&i.ItemCount,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSubjectEvent = `-- name: DeleteSubjectEvent :exec
DELETE FROM Subject_Event
WHERE uuid = $1
`

func (q *QueriesSchool) DeleteSubjectEvent(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSubjectEvent, uuid)
	return err
}

type SubjectEventInfo struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	TypeId           int64           `json:"typeId"`
	TicketItemId     int64           `json:"ticketItemId"`
	Iiid             int64           `json:"iiid"`
	SectionSubjectId int64           `json:"sectionSubjectId"`
	EventDate        time.Time       `json:"eventDate"`
	GradingPeriod    int16           `json:"gradingPeriod"`
	ItemCount        decimal.Decimal `json:"itemCount"`
	StatusId         int64           `json:"statusId"`
	Remarks          sql.NullString  `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
	ModCtr           int64           `json:"modCtr"`
	Created          sql.NullTime    `json:"created"`
	Updated          sql.NullTime    `json:"updated"`
}

const subjectEventSQL = `-- name: SubjectEventSQL :one
SELECT
  id, mr.UUID, type_id, ticket_item_id, iiid, section_subject_id, event_date, 
  grading_period, item_count, status_id, remarks, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Subject_Event d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateSubjectEvent(q *QueriesSchool, ctx context.Context, sql string) (SubjectEventInfo, error) {
	var i SubjectEventInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.TicketItemId,
		&i.Iiid,
		&i.SectionSubjectId,
		&i.EventDate,
		&i.GradingPeriod,
		&i.ItemCount,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateSubjectEvents(q *QueriesSchool, ctx context.Context, sql string) ([]SubjectEventInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SubjectEventInfo{}
	for rows.Next() {
		var i SubjectEventInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.TypeId,
			&i.TicketItemId,
			&i.Iiid,
			&i.SectionSubjectId,
			&i.EventDate,
			&i.GradingPeriod,
			&i.ItemCount,
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

func (q *QueriesSchool) GetSubjectEvent(ctx context.Context, id int64) (SubjectEventInfo, error) {
	return populateSubjectEvent(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", subjectEventSQL, id))
}

func (q *QueriesSchool) GetSubjectEventbyUuid(ctx context.Context, uuid uuid.UUID) (SubjectEventInfo, error) {
	return populateSubjectEvent(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", subjectEventSQL, uuid))
}

type ListSubjectEventParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesSchool) ListSubjectEvent(ctx context.Context, arg ListSubjectEventParams) ([]SubjectEventInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			subjectEventSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(subjectEventSQL)
	}
	return populateSubjectEvents(q, ctx, sql)
}

const updateSubjectEvent = `-- name: UpdateSubjectEvent :one
UPDATE Subject_Event SET 
	uuid = $2,
	type_id = $3,
	ticket_item_id = $4,
	iiid = $5,
	section_subject_id = $6,
	event_date = $7,
	grading_period = $8,
	item_count = $9,
	status_id = $10,
	remarks = $11,
	other_info = $12
WHERE id = $1
RETURNING id, uuid, type_id, ticket_item_id, iiid, section_subject_id, event_date, grading_period, item_count, status_id, remarks, other_info
`

func (q *QueriesSchool) UpdateSubjectEvent(ctx context.Context, arg SubjectEventRequest) (model.SubjectEvent, error) {
	row := q.db.QueryRowContext(ctx, updateSubjectEvent,
		arg.Id,
		arg.Uuid,
		arg.TypeId,
		arg.TicketItemId,
		arg.Iiid,
		arg.SectionSubjectId,
		arg.EventDate,
		arg.GradingPeriod,
		arg.ItemCount,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.SubjectEvent
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.TicketItemId,
		&i.Iiid,
		&i.SectionSubjectId,
		&i.EventDate,
		&i.GradingPeriod,
		&i.ItemCount,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
