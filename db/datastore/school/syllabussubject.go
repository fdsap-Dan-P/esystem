package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createSyllabusSubject = `-- name: CreateSyllabusSubject: one
INSERT INTO Syllabus_Subject(
   uuid, syllabus_id, subject_id, units, type_id, status_id, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(uuid)
DO UPDATE SET
	syllabus_id =  EXCLUDED.syllabus_id,
	subject_id =  EXCLUDED.subject_id,
	units =  EXCLUDED.units,
	type_id =  EXCLUDED.type_id,
	status_id =  EXCLUDED.status_id,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING id, uuid, syllabus_id, subject_id, units, type_id, status_id, remarks, other_info
`

type SyllabusSubjectRequest struct {
	Id         int64          `json:"id"`
	Uuid       uuid.UUID      `json:"uuid"`
	SyllabusId int64          `json:"syllabusId"`
	SubjectId  int64          `json:"subjectId"`
	Units      int64          `json:"units"`
	TypeId     int64          `json:"typeId"`
	StatusId   int64          `json:"statusId"`
	Remarks    string         `json:"remarks"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}

func (q *QueriesSchool) CreateSyllabusSubject(ctx context.Context, arg SyllabusSubjectRequest) (model.SyllabusSubject, error) {
	row := q.db.QueryRowContext(ctx, createSyllabusSubject,
		arg.Uuid,
		arg.SyllabusId,
		arg.SubjectId,
		arg.Units,
		arg.TypeId,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.SyllabusSubject
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SyllabusId,
		&i.SubjectId,
		&i.Units,
		&i.TypeId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSyllabusSubject = `-- name: DeleteSyllabusSubject :exec
DELETE FROM Syllabus_Subject
WHERE uuid = $1
`

func (q *QueriesSchool) DeleteSyllabusSubject(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSyllabusSubject, uuid)
	return err
}

type SyllabusSubjectInfo struct {
	Id         int64          `json:"id"`
	Uuid       uuid.UUID      `json:"uuid"`
	SyllabusId int64          `json:"syllabusId"`
	SubjectId  int64          `json:"subjectId"`
	Units      int64          `json:"units"`
	TypeId     int64          `json:"typeId"`
	StatusId   int64          `json:"statusId"`
	Remarks    string         `json:"remarks"`
	OtherInfo  sql.NullString `json:"otherInfo"`
	ModCtr     int64          `json:"modCtr"`
	Created    sql.NullTime   `json:"created"`
	Updated    sql.NullTime   `json:"updated"`
}

const syllabusSubjectSQL = `-- name: SyllabusSubjectSQL :one
SELECT
d.id, mr.uuid, syllabus_id, subject_id, units, type_id, status_id, remarks, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Syllabus_Subject d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateSyllabusSubject(q *QueriesSchool, ctx context.Context, sql string) (SyllabusSubjectInfo, error) {
	var i SyllabusSubjectInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SyllabusId,
		&i.SubjectId,
		&i.Units,
		&i.TypeId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateSyllabusSubjects(q *QueriesSchool, ctx context.Context, sql string) ([]SyllabusSubjectInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SyllabusSubjectInfo{}
	for rows.Next() {
		var i SyllabusSubjectInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.SyllabusId,
			&i.SubjectId,
			&i.Units,
			&i.TypeId,
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

func (q *QueriesSchool) GetSyllabusSubject(ctx context.Context, id int64) (SyllabusSubjectInfo, error) {
	return populateSyllabusSubject(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", syllabusSubjectSQL, id))
}

func (q *QueriesSchool) GetSyllabusSubjectbyUuid(ctx context.Context, uuid uuid.UUID) (SyllabusSubjectInfo, error) {
	return populateSyllabusSubject(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", syllabusSubjectSQL, uuid))
}

type ListSyllabusSubjectParams struct {
	SyllabusId int64 `json:"syllabusId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesSchool) ListSyllabusSubject(ctx context.Context, arg ListSyllabusSubjectParams) ([]SyllabusSubjectInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			syllabusSubjectSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(syllabusSubjectSQL)
	}
	return populateSyllabusSubjects(q, ctx, sql)
}

const updateSyllabusSubject = `-- name: UpdateSyllabusSubject :one
UPDATE Syllabus_Subject SET 
	uuid = $2,
	syllabus_id = $3,
	subject_id = $4,
	units = $5,
	type_id = $6,
	status_id = $7,
	remarks = $8,
	other_info = $9
WHERE id = $1
RETURNING id, uuid, syllabus_id, subject_id, units, type_id, status_id, remarks, other_info
`

func (q *QueriesSchool) UpdateSyllabusSubject(ctx context.Context, arg SyllabusSubjectRequest) (model.SyllabusSubject, error) {
	row := q.db.QueryRowContext(ctx, updateSyllabusSubject,
		arg.Id,
		arg.Uuid,
		arg.SyllabusId,
		arg.SubjectId,
		arg.Units,
		arg.TypeId,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.SyllabusSubject
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SyllabusId,
		&i.SubjectId,
		&i.Units,
		&i.TypeId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
