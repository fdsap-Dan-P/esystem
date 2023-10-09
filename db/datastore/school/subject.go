package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createSubject = `-- name: CreateSubject: one
INSERT INTO Subject (
	Uuid, Subject_Title, Subject_Ref_Id, Type_Id, Remarks, Other_Info
 ) VALUES ($1, $2, $3, $4, $5, $6) 
 ON CONFLICT(UUID)
 DO UPDATE SET
	subject_title =  EXCLUDED.subject_title,
	subject_ref_id =  EXCLUDED.subject_ref_id,
	type_id =  EXCLUDED.type_id,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
 RETURNING 
   Id, Uuid, Subject_Title, Subject_Ref_Id, Type_Id, Remarks, Other_Info
`

type SubjectRequest struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	SubjectTitle string         `json:"subjectTitle"`
	SubjectRefId int64          `json:"subjectRefId"`
	TypeId       sql.NullInt64  `json:"typeId"`
	Remarks      string         `json:"remarks"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

func (q *QueriesSchool) CreateSubject(ctx context.Context, arg SubjectRequest) (model.Subject, error) {
	if arg.Uuid == uuid.Nil {
		arg.Uuid, _ = uuid.NewUUID()
	}
	row := q.db.QueryRowContext(ctx, createSubject,
		arg.Uuid,
		arg.SubjectTitle,
		arg.SubjectRefId,
		arg.TypeId,
		arg.Remarks,
		arg.OtherInfo,
	)

	var i model.Subject
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SubjectTitle,
		&i.SubjectRefId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSubject = `-- name: DeleteSubject :exec
DELETE FROM Subject
WHERE id = $1
`

func (q *QueriesSchool) DeleteSubject(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSubject, id)
	return err
}

type SubjectInfo struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	SubjectTitle string         `json:"SubjectTitle"`
	SubjectRefId int64          `json:"SubjectRefId"`
	TypeId       sql.NullInt64  `json:"TypeId"`
	Remarks      string         `json:"Remarks"`
	OtherInfo    sql.NullString `json:"OtherInfo"`
	ModCtr       int64          `json:"modCtr"`
	Created      sql.NullTime   `json:"created"`
	Updated      sql.NullTime   `json:"updated"`
}

func populateSubject(q *QueriesSchool, ctx context.Context, sql string) (SubjectInfo, error) {
	var i SubjectInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SubjectTitle,
		&i.SubjectRefId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateSubjects(q *QueriesSchool, ctx context.Context, sql string) ([]SubjectInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []SubjectInfo{}
	for rows.Next() {
		var i SubjectInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.SubjectTitle,
			&i.SubjectRefId,
			&i.TypeId,
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

const subjectSQL = `-- name: subjectSQL
SELECT 
  d.Id, mr.Uuid, d.Subject_Title, d.Subject_Ref_Id, d.Type_Id, d.Remarks, d.Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Subject d 
INNER JOIN Main_Record mr on mr.Uuid = d.Uuid
`

func (q *QueriesSchool) GetSubject(ctx context.Context, id int64) (SubjectInfo, error) {
	return populateSubject(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", subjectSQL, id))
}

func (q *QueriesSchool) GetSubjectbyUuid(ctx context.Context, uuid uuid.UUID) (SubjectInfo, error) {
	return populateSubject(q, ctx, fmt.Sprintf("%s WHERE mr.Uuid = '%v'", subjectSQL, uuid))
}

type ListSubjectParams struct {
	SubjectRefId int64 `json:"subjectRefId"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *QueriesSchool) ListSubject(ctx context.Context, arg ListSubjectParams) ([]SubjectInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Subject_Ref_Id = %v LIMIT %d OFFSET %d",
			subjectSQL, arg.SubjectRefId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Subject_Ref_Id = %v ", subjectSQL, arg.SubjectRefId)
	}
	return populateSubjects(q, ctx, sql)
}

const updateSubject = `-- name: UpdateSubject :one
UPDATE Subject SET 
  Subject_Title = $2,
  Subject_Ref_Id = $3,
  Type_Id = $4,
  Remarks = $5,
  Other_Info = $6
WHERE id = $1
RETURNING 
  Id, Uuid, Subject_Title, Subject_Ref_Id, Type_Id, Remarks, Other_Info
`

func (q *QueriesSchool) UpdateSubject(ctx context.Context, arg SubjectRequest) (model.Subject, error) {
	row := q.db.QueryRowContext(ctx, updateSubject,
		arg.Id,
		arg.SubjectTitle,
		arg.SubjectRefId,
		arg.TypeId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Subject
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SubjectTitle,
		&i.SubjectRefId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
